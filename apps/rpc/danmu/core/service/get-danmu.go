package service

import (
	"LiveDanmu/apps/rpc/danmu/core"
	"LiveDanmu/apps/rpc/danmu/kitex_gen/danmusvr"
	"context"
)

func GetHotDanmu(ctx context.Context, req *danmusvr.GetTopReq) *danmusvr.GetTopResp {
	core.Dao.ReadHotDanmu(ctx, req.)
}
