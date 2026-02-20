package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
)

type KitexReqs interface {
	*livesvr.GetLiveInfoReq | *livesvr.GetLiveListReq | *livesvr.StartLiveReq | *livesvr.StopLiveReq
}

type KitexResps interface {
	*livesvr.GetLiveInfoResp | *livesvr.GetLiveListResp | *livesvr.StartLiveResp | *livesvr.StopLiveResp | *livesvr.GetMyLiveListResp | response.Response
}

type Kresp interface {
	GetStatus() int64
	GetInfo() string
}

var (
	// 编译期校验：实现Kresp
	_ Kresp = (*livesvr.GetLiveInfoResp)(nil)
	_ Kresp = (*livesvr.GetLiveListResp)(nil)
	_ Kresp = (*livesvr.StartLiveResp)(nil)
	_ Kresp = (*livesvr.StopLiveResp)(nil)
	_ Kresp = response.Response{}
)
