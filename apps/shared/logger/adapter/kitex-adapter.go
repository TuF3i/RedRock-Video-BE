package adapter

import (
	"LiveDanmu/apps/shared/union_var"
	"context"
	"io"

	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
)

type KitexLokiLogger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	level  klog.Level
}

var _ klog.FullLogger = (*KitexLokiLogger)(nil)

func NewKitexLokiLogger(lokiLogger *zap.Logger) *KitexLokiLogger {
	return &KitexLokiLogger{
		logger: lokiLogger,
		sugar:  lokiLogger.Sugar(),
		level:  klog.LevelInfo,
	}
}

// 基础日志
func (l *KitexLokiLogger) Trace(v ...interface{})                  { l.sugar.Debug(v...) }
func (l *KitexLokiLogger) Debug(v ...interface{})                  { l.sugar.Debug(v...) }
func (l *KitexLokiLogger) Info(v ...interface{})                   { l.sugar.Info(v...) }
func (l *KitexLokiLogger) Notice(v ...interface{})                 { l.sugar.Info(v...) }
func (l *KitexLokiLogger) Warn(v ...interface{})                   { l.sugar.Warn(v...) }
func (l *KitexLokiLogger) Error(v ...interface{})                  { l.sugar.Error(v...) }
func (l *KitexLokiLogger) Fatal(v ...interface{})                  { l.sugar.Fatal(v...) }
func (l *KitexLokiLogger) Tracef(format string, v ...interface{})  { l.sugar.Debugf(format, v...) }
func (l *KitexLokiLogger) Debugf(format string, v ...interface{})  { l.sugar.Debugf(format, v...) }
func (l *KitexLokiLogger) Infof(format string, v ...interface{})   { l.sugar.Infof(format, v...) }
func (l *KitexLokiLogger) Noticef(format string, v ...interface{}) { l.sugar.Infof(format, v...) }
func (l *KitexLokiLogger) Warnf(format string, v ...interface{})   { l.sugar.Warnf(format, v...) }
func (l *KitexLokiLogger) Errorf(format string, v ...interface{})  { l.sugar.Errorf(format, v...) }
func (l *KitexLokiLogger) Fatalf(format string, v ...interface{})  { l.sugar.Fatalf(format, v...) }

// ctxSugar
func (l *KitexLokiLogger) ctxSugar(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return l.sugar
	}

	// 闭包函数
	getTraceIDFromContext := func(ctx context.Context) string {
		// 从自定义 context value 获取
		if traceID, ok := ctx.Value(union_var.TRACE_ID_KEY).(string); ok {
			return traceID
		}

		return ""
	}

	traceID := getTraceIDFromContext(ctx)
	if traceID == "" {
		return l.sugar
	}

	return l.sugar.With(union_var.TRACE_ID_KEY, traceID)
}

// 带上下文的日志
func (l *KitexLokiLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Debugf(format, v...)
}

func (l *KitexLokiLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Debugf(format, v...)
}

func (l *KitexLokiLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Infof(format, v...)
}

func (l *KitexLokiLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Infof(format, v...)
}

func (l *KitexLokiLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Warnf(format, v...)
}

func (l *KitexLokiLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Errorf(format, v...)
}

func (l *KitexLokiLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Fatalf(format, v...)
}

func (l *KitexLokiLogger) SetLevel(level klog.Level) {
	// 级别控制由核心接管，这里不做处理
}

func (l *KitexLokiLogger) SetOutput(output io.Writer) {
	// 不做自定义输出
}

func (l *KitexLokiLogger) Level() klog.Level {
	return l.level
}

func (l *KitexLokiLogger) Sync() {
	_ = l.logger.Sync()
}
