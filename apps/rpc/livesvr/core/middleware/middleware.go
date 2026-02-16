package middleware

import (
	"LiveDanmu/apps/public/union_var"
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func PreInit(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// 将MetaData中写入Context
		if traceID, ok := metainfo.GetPersistentValue(ctx, union_var.TRACE_ID_KEY); ok {
			ctx = context.WithValue(ctx, union_var.TRACE_ID_KEY, traceID)
		} else {
			ctx = context.WithValue(ctx, union_var.TRACE_ID_KEY, "")
		}

		// 继续执行后续逻辑
		err = next(ctx, req, resp)
		return err
	}
}
