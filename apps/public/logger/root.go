package logger

import (
	"net/http"
	"sync"

	"go.uber.org/zap/zapcore"
)

// ------------ 全局常量配置（生产可抽离到配置中心） ------------
const (
	lokiChanSize  = 1000 // 日志缓冲通道大小，根据业务QPS调整
	lokiWorkerNum = 3    // 异步推送协程数，建议3-5个
)

// LokiConfig ------------ Loki 相关结构体定义 ------------
// LokiConfig Loki配置项，适配viper/nacos等配置中心（mapstructure标签）
type LokiConfig struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}

// LokiPushRequest Loki /loki/api/v1/push 接口标准请求体
type LokiPushRequest struct {
	Streams []LokiStream `json:"streams"`
}

// LokiStream Loki 流结构，包含标签和日志条目
type LokiStream struct {
	Labels  string     `json:"stream"` // 标签，必须是JSON字符串格式
	Entries [][]string `json:"values"` // 日志条目，格式：[[时间戳(RFC3339Nano), 日志内容], ...]
}

// ------------ Zap Hook 核心结构体 ------------
// lokiHook 实现zap.Hook接口，生产级异步钩子
type lokiHook struct {
	config  LokiConfig          // 配置信息
	client  *http.Client        // 全局复用HTTP客户端，避免频繁创建连接
	logChan chan *zapcore.Entry // 日志缓冲通道，异步解耦
	wg      sync.WaitGroup      // 等待组，保证优雅退出时所有日志推送完成
}
