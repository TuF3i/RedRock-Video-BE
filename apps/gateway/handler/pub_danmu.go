package handler

import (
	"LiveDanmu/apps/gateway"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func PublishDanMuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 生成TraceID
		traceID := gateway.SnowFlake.Generate().String()
		ctx = context.WithValue(ctx, gateway.TRACE_ID_KEY, traceID)
		// 调用pub_danmu

	}
}
