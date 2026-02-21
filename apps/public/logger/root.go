package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig 日志配置项
type LoggerConfig struct {
	Service string // 服务名
	Env     string // 环境
	Level   string // 日志级别
}

// ------------ Zap 日志初始化核心方法 ------------
// initZap 初始化Zap日志器
func initZap(config LoggerConfig) *zap.Logger {
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
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// 3. 构建Zap Core（编码器+输出器+日志级别）
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		level,
	)

	// 4. 创建Logger
	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.ErrorOutput(zapcore.AddSync(os.Stderr)),
	)

	// 5. 替换Zap全局日志器
	zap.ReplaceGlobals(logger)

	return logger
}
