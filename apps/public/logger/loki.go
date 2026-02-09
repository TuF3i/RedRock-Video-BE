package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ------------ Loki Hook 核心方法 ------------
// NewLokiHook 创建生产级Loki Hook，初始化通道和工作协程
func NewLokiHook(config LokiConfig) *lokiHook {
	// 生产级HTTP客户端配置：连接池、超时、保活
	transport := &http.Transport{
		MaxIdleConns:        10,               // 最大空闲连接数
		IdleConnTimeout:     30 * time.Second, // 空闲连接超时时间
		DisableKeepAlives:   false,            // 开启连接保活
		TLSHandshakeTimeout: 5 * time.Second,  // TLS握手超时
	}
	client := &http.Client{
		Timeout:   3 * time.Second, // 请求超时时间，避免Loki不可达阻塞
		Transport: transport,
	}

	hook := &lokiHook{
		config:  config,
		client:  client,
		logChan: make(chan *zapcore.Entry, lokiChanSize), // 带缓冲通道，削峰填谷
	}

	// 启动多协程工作池
	hook.wg.Add(lokiWorkerNum)
	for i := 0; i < lokiWorkerNum; i++ {
		go hook.worker(i) // 传入协程ID，方便调试
	}

	return hook
}

// Fire 实现zap.Hook接口，日志触发时执行：异步放入通道，立即返回不阻塞业务
func (h *lokiHook) Fire(entry zapcore.Entry) error {
	select {
	case h.logChan <- &entry: // 通道有空闲，放入日志
	default: // 通道满时丢弃日志，避免阻塞业务，可根据需求改为落盘本地
		zap.L().Warn("loki log channel is full, discard log",
			zap.String("service_tag.yaml", h.config.Service),
			zap.String("msg", entry.Message[:min(100, len(entry.Message))])) // 截断超长消息
	}
	return nil
}

// worker 工作协程，循环从通道取日志推送到Loki
func (h *lokiHook) worker(id int) {
	defer h.wg.Done()
	zap.L().Info("loki push worker start", zap.Int("worker_id", id))
	// 通道关闭时，for range会自动退出
	for entry := range h.logChan {
		if err := h.pushToLoki(entry); err != nil {
			zap.L().Error("push log to loki failed",
				zap.Int("worker_id", id),
				zap.Error(err),
				zap.String("msg", entry.Message[:min(100, len(entry.Message))]))
		}
	}
	zap.L().Info("loki push worker exit", zap.Int("worker_id", id))
}

// pushToLoki 实际推送逻辑：构造Loki请求体并发送HTTP请求
func (h *lokiHook) pushToLoki(entry *zapcore.Entry) error {
	// 1. 构造Loki低基数标签（禁止用trace_id/order_id等高基数字段）
	labels := map[string]string{
		"service_tag.yaml": h.config.Service,
		"env":              h.config.Env,
		"level":            entry.Level.String(),
	}
	labelsJson, err := json.Marshal(labels)
	if err != nil {
		return fmt.Errorf("marshal labels failed: %w", err)
	}

	// 2. 构造结构化日志内容，整合zap原生字段+自定义Field
	logContent := make(map[string]interface{})
	logContent["time"] = entry.Time.Format(time.RFC3339)
	logContent["msg"] = entry.Message
	logContent["caller"] = entry.Caller.String() // 代码调用位置：文件:行号
	if entry.Stack != "" {
		logContent["stack"] = entry.Stack // 堆栈信息，仅错误日志有
	}

	// 3. 构造Loki标准请求体
	pushReq := LokiPushRequest{
		Streams: []LokiStream{
			{
				Labels: string(labelsJson),
				Entries: [][]string{
					{entry.Time.UTC().Format(time.RFC3339Nano), marshalLogContent(logContent)},
				},
			},
		},
	}
	reqBody, _ := json.Marshal(pushReq) // 忽略序列化错误，避免阻塞业务

	// 4. 发送POST请求到Loki Push API
	lokiUrl := "http://" + h.config.LokiAddr[0] + ":3100" + "/loki/api/v1/push"
	resp, err := h.client.Post(lokiUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	// 5. 校验Loki响应（成功返回204 No Content）
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("loki response error, status: %s, url: %s", resp.Status, lokiUrl)
	}
	return nil
}

// Close 优雅关闭：关闭通道并等待所有工作协程完成
func (h *lokiHook) Close() {
	close(h.logChan) // 关闭通道，工作协程会自动退出for range
	h.wg.Wait()      // 等待所有未推送的日志处理完成
	zap.L().Info("all loki push worker done, hook closed")
}

// ------------ 工具函数 ------------
// marshalLogContent 安全序列化日志内容，避免JSON序列化panic
func marshalLogContent(content map[string]interface{}) string {
	data, err := json.Marshal(content)
	if err != nil {
		return fmt.Sprintf(`{"msg":"marshal log content failed", "error":"%s", "original_msg":"%v"}`, err.Error(), content["msg"])
	}
	return string(data)
}

// min 取最小值，避免字符串截断时索引越界
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ------------ Zap 日志初始化核心方法 ------------
// InitZapWithLoki 初始化Zap日志器并集成Loki Hook，返回logger和hook（用于优雅关闭）
func initZapWithLoki(config LokiConfig) (*zap.Logger, *lokiHook) {
	// 1. 解析日志级别，默认InfoLevel
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		zap.L().Warn("parse log level failed, use default info level", zap.Error(err))
		level = zapcore.InfoLevel
	}

	// 2. 生产级JSON编码器配置（控制台输出结构化日志）
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 级别大写：DEBUG/INFO/ERROR
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 时间格式：ISO8601
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 调用者：文件:行号（短路径）
		EncodeDuration: zapcore.SecondsDurationEncoder, // 耗时秒数
	}

	// 3. 构建Zap Core（编码器+输出器+日志级别）
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                   // JSON格式输出，便于采集
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 输出到控制台，生产可追加文件
		level,
	)

	// 4. 创建Loki Hook并集成到Zap
	lokiHook := NewLokiHook(config)
	logger := zap.New(core,
		zap.AddCaller(),                             // 开启调用者信息
		zap.AddStacktrace(zapcore.ErrorLevel),       // 错误级别打印堆栈
		zap.Hooks(lokiHook.Fire),                    // 集成Loki Hook
		zap.ErrorOutput(zapcore.AddSync(os.Stderr)), // 错误日志单独输出到标准错误
	)

	// 5. 替换Zap全局日志器，方便全局使用zap.L()
	zap.ReplaceGlobals(logger)

	return logger, lokiHook
}

// ------------ 主函数测试 ------------
//func main() {
//	// 1. 配置Loki（生产可通过viper从yaml/nacos读取）
//	lokiConfig := LokiConfig{
//		LokiAddr: []string{"http://127.0.0.1:3100"}, // 本地Loki地址，确保3100端口可访问
//		Service:  "pay-service_tag.yaml",                     // 你的微服务名
//		Env:      "prod",                            // 运行环境
//		Level:    "debug",                           // 日志级别
//	}
//
//	// 2. 初始化Zap+Loki
//	logger, lokiHook := initZapWithLoki(lokiConfig)
//	// 3. 优雅退出：刷新缓冲区+关闭Loki Hook+等待推送完成
//	defer func() {
//		_ = logger.Sync() // 刷新Zap缓冲区
//		lokiHook.Close()  // 关闭Loki Hook，等待所有日志推送完成
//		fmt.Println("program exit gracefully")
//	}()
//
//	// 4. 测试推送不同级别日志（带自定义Field）
//	logger.Debug("用户发起支付请求",
//		zap.String("trace_id", "OTEL-20260204-001"),
//		zap.Int64("user_id", 100001),
//		zap.String("pay_way", "wechat"),
//	)
//
//	logger.Info("支付成功",
//		zap.String("trace_id", "OTEL-20260204-002"),
//		zap.Int64("order_id", 888888),
//		zap.Float64("amount", 199.00),
//		zap.Bool("success", true),
//		zap.Duration("cost", 120*time.Millisecond),
//	)
//
//	logger.Error("支付失败",
//		zap.String("trace_id", "OTEL-20260204-003"),
//		zap.Int64("user_id", 100001),
//		zap.Int64("order_id", 888889),
//		zap.Error(fmt.Errorf("微信支付接口超时，code:408, timeout:5s")),
//	)
//
//	// 测试批量日志（模拟高并发场景）
//	for i := 0; i < 10; i++ {
//		logger.Info("批量支付通知",
//			zap.String("trace_id", fmt.Sprintf("OTEL-20260204-%03d", i+10)),
//			zap.Int("batch_no", i+1),
//			zap.Int("count", 50),
//		)
//	}
//
//}
