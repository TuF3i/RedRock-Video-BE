package adapter

import (
	"LiveDanmu/apps/shared/union_var"
	"context"
	"io"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.uber.org/zap"
)

type HertzLokiLogger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// 确保实现了所有接口
var _ hlog.FullLogger = (*HertzLokiLogger)(nil)

// NewHertzZapLogger 创建实例
func NewHertzZapLogger(lokiLogger *zap.Logger) *HertzLokiLogger {
	return &HertzLokiLogger{
		logger: lokiLogger,
		sugar:  lokiLogger.Sugar(),
	}
}

// 基础接口
func (l *HertzLokiLogger) Trace(v ...interface{}) {
	l.sugar.Debug(v...)
}

func (l *HertzLokiLogger) Debug(v ...interface{}) {
	l.sugar.Debug(v...)
}

func (l *HertzLokiLogger) Info(v ...interface{}) {
	l.sugar.Info(v...)
}

func (l *HertzLokiLogger) Notice(v ...interface{}) {
	l.sugar.Info(v...)
}

func (l *HertzLokiLogger) Warn(v ...interface{}) {
	l.sugar.Warn(v...)
}

func (l *HertzLokiLogger) Error(v ...interface{}) {
	l.sugar.Error(v...)
}

func (l *HertzLokiLogger) Fatal(v ...interface{}) {
	l.sugar.Fatal(v...)
}

// 格式化接口
func (l *HertzLokiLogger) Tracef(format string, v ...interface{}) {
	l.sugar.Debugf(format, v...)
}

func (l *HertzLokiLogger) Debugf(format string, v ...interface{}) {
	l.sugar.Debugf(format, v...)
}

func (l *HertzLokiLogger) Infof(format string, v ...interface{}) {
	l.sugar.Infof(format, v...)
}

func (l *HertzLokiLogger) Noticef(format string, v ...interface{}) {
	l.sugar.Infof(format, v...)
}

func (l *HertzLokiLogger) Warnf(format string, v ...interface{}) {
	l.sugar.Warnf(format, v...)
}

func (l *HertzLokiLogger) Errorf(format string, v ...interface{}) {
	l.sugar.Errorf(format, v...)
}

func (l *HertzLokiLogger) Fatalf(format string, v ...interface{}) {
	l.sugar.Fatalf(format, v...)
}

// 带ctx的链路追踪接口
func (l *HertzLokiLogger) ctxSugar(ctx context.Context) *zap.SugaredLogger {
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

func (l *HertzLokiLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Debugf(format, v...)
}

func (l *HertzLokiLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Debugf(format, v...)
}

func (l *HertzLokiLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Infof(format, v...)
}

func (l *HertzLokiLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Infof(format, v...)
}

func (l *HertzLokiLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Warnf(format, v...)
}

func (l *HertzLokiLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Errorf(format, v...)
}

func (l *HertzLokiLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.ctxSugar(ctx).Fatalf(format, v...)
}

// SetLevel 接口
func (l *HertzLokiLogger) SetLevel(level hlog.Level) {
	// 全局核心设置过了，这里忽略
}

func (l *HertzLokiLogger) SetOutput(output io.Writer) {
	// 忽略自定义输出
}

func (l *HertzLokiLogger) Sync() {
	_ = l.logger.Sync()
}
