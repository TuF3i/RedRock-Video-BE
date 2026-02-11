package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
)

type KitexReqs interface {
	*videosvr.AddVideoReq | *videosvr.DelVideoReq | *videosvr.JudgeAccessReq | *videosvr.GetVideoListReq | *videosvr.GetPreSignedUrlReq
}

type KitexResps interface {
	*videosvr.AddVideoResp | *videosvr.DelVideoResp | *videosvr.JudgeAccessResp | *videosvr.GetVideoListResp | *videosvr.GetPreSignedUrlResp | *videosvr.GetJudgeListResp | response.Response
}

type Kresp interface {
	GetStatus() int64
	GetInfo() string
}

var (
	// 编译期校验：实现Kresp
	_ Kresp = (*videosvr.AddVideoResp)(nil)
	_ Kresp = (*videosvr.DelVideoResp)(nil)
	_ Kresp = (*videosvr.JudgeAccessResp)(nil)
	_ Kresp = (*videosvr.GetVideoListResp)(nil)
	_ Kresp = (*videosvr.GetPreSignedUrlResp)(nil)
	_ Kresp = (*videosvr.GetJudgeListResp)(nil)
	_ Kresp = response.Response{}
)
