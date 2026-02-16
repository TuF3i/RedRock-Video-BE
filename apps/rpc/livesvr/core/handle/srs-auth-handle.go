package handle

import (
	"LiveDanmu/apps/rpc/livesvr/core"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"
)

func SRSAuth(ctx context.Context, req *livesvr.SRSAuthReq) *livesvr.SRSAuthResp {
	rvid := req.GetRvid()
	password := req.GetPassword()

	ok, err := core.Dao.CheckIfExist(ctx, rvid)
	if err != nil || !ok {
		return &livesvr.SRSAuthResp{Ok: 1}
	}

	data, err := core.Dao.GetLiveInfo(ctx, rvid)
	if err != nil {
		return &livesvr.SRSAuthResp{Ok: 1}
	}

	if data.UpstreamPassword != password {
		return &livesvr.SRSAuthResp{Ok: 1}
	}

	return &livesvr.SRSAuthResp{Ok: 0}
}
