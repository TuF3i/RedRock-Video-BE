package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
)

type KitexReqs interface {
	*danmusvr.PubVideoReq | *danmusvr.PubLiveReq | *danmusvr.GetFullReq | *danmusvr.GetTopReq | *danmusvr.DelLiveReq | *danmusvr.DelReq
}

type KitexResps interface {
	*danmusvr.PubVideoResp | *danmusvr.PubLiveResp | *danmusvr.GetFullResp | *danmusvr.GetTopResp | *danmusvr.DelLiveResp | *danmusvr.DelResp | response.Response
}

type Kresp interface {
	GetStatus() int64
	GetInfo() string
}

var (
	// 编译期校验：实现Kresp
	_ Kresp = (*danmusvr.PubVideoResp)(nil)
	_ Kresp = (*danmusvr.PubLiveResp)(nil)
	_ Kresp = (*danmusvr.GetFullResp)(nil)
	_ Kresp = (*danmusvr.GetTopResp)(nil)
	_ Kresp = (*danmusvr.DelLiveResp)(nil)
	_ Kresp = (*danmusvr.DelResp)(nil)
	_ Kresp = response.Response{}
)
