package logger

import (
	"go.uber.org/zap"
)

type NewLogger struct {
	Logger   *zap.Logger
	LokiHook *lokiHook
}

func GetLogger(conf LokiConfig) (*NewLogger, error) {
	l := NewLogger{}
	l.Logger, l.LokiHook = initZapWithLoki(conf)
	return &l, nil
}
func (r *NewLogger) SyncClean() error {
	err := r.Logger.Sync() // 刷新Zap缓冲区
	r.LokiHook.Close()     // 关闭Loki Hook，等待所有日志推送完成
	if err != nil {
		return err
	}
	return nil
}

func (r *NewLogger) INFO(msg string, fields ...zap.Field) {
	r.Logger.Info(msg, fields...)
}

func (r *NewLogger) WARN(msg string, fields ...zap.Field) {
	r.Logger.Warn(msg, fields...)
}
func (r *NewLogger) DEBUG(msg string, fields ...zap.Field) {
	r.Logger.Debug(msg, fields...)
}

func (r *NewLogger) PANIC(msg string, fields ...zap.Field) {
	r.Logger.Panic(msg, fields...)
}

func (r *NewLogger) FATAL(msg string, field ...zap.Field) {
	r.Logger.Fatal(msg, field...)
}
