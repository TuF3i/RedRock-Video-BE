package middleware

import (
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/danmusvr/core/handle"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func DanmuPoolReleaseMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// 执行后续业务逻辑
		err = next(ctx, req, resp)
		// Top型类型断言
		if dmResp, ok := resp.(*danmusvr.GetTopResp); ok {
			// 释放内存
			if dmResp.Data != nil {
				handle.ReleaseDanmuMsg(dmResp.Data)
			}
		}
		// Full型类型断言
		if dmResp, ok := resp.(*danmusvr.GetFullResp); ok {
			// 释放内存
			if dmResp.Data != nil {
				handle.ReleaseDanmuMsg(dmResp.Data)
			}
		}
		// 其他直接return nil
		return err
	}
}

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
