package logger

import (
	"go.uber.org/zap"
)

type NewLogger struct {
	Logger *zap.Logger
}

func GetLogger(conf LoggerConfig) (*NewLogger, error) {
	l := NewLogger{}
	l.Logger = initZap(conf)
	return &l, nil
}

func (r *NewLogger) SyncClean() error {
	err := r.Logger.Sync()
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
