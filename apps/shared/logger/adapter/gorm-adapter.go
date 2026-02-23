package adapter

import (
	"LiveDanmu/apps/shared/union_var"
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type LokiGormLogger struct {
	lokiLogger    *zap.Logger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// NewLokiGormLogger 注册新日志核心
func NewLokiGormLogger(lokiLogger *zap.Logger) *LokiGormLogger {
	return &LokiGormLogger{
		lokiLogger:    lokiLogger,
		LogLevel:      logger.Info,
		SlowThreshold: 500 * time.Millisecond,
	}
}

// LogMode 设置日志级别
func (r *LokiGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *r
	newLogger.LogLevel = level
	return &newLogger
}

func (r *LokiGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if r.LogLevel >= logger.Info {
		r.lokiLogger.Sugar().Infof(msg, data...)
	}
}

func (r *LokiGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if r.LogLevel >= logger.Info {
		r.lokiLogger.Sugar().Warnf(msg, data...)
	}
}

func (r *LokiGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if r.LogLevel >= logger.Info {
		r.lokiLogger.Sugar().Errorf(msg, data...)
	}
}

func (r *LokiGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if r.LogLevel <= logger.Silent {
		return
	}

	getTraceIDFromContext := func(ctx context.Context) string {
		// 从自定义 context value 获取
		if traceID, ok := ctx.Value(union_var.TRACE_ID_KEY).(string); ok {
			return traceID
		}

		return ""
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	fields := []zap.Field{
		zap.String(union_var.TRACE_ID_KEY, getTraceIDFromContext(ctx)),
		zap.Duration("elapsed", elapsed),
		zap.Int64("rows", rows),
		zap.String("sql", sql),
	}

	switch {
	case err != nil && r.LogLevel >= logger.Error:
		r.lokiLogger.Error("gorm trace error", append(fields, zap.Error(err))...)
	case elapsed > r.SlowThreshold && r.SlowThreshold != 0 && r.LogLevel >= logger.Warn:
		r.lokiLogger.Warn("gorm slow query", fields...)
	case r.LogLevel >= logger.Info:
		r.lokiLogger.Info("gorm query", fields...)
	}
}
