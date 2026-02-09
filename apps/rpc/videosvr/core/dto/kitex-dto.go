package dto

import "LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"

// DtoResp 接口
type DtoResp interface {
	GetStatus()
	GetInfo()
}

// KitexResponse泛型接口
type KitexResp interface {
	*videosvr.AddVideoResp | *videosvr.DelVideoResp | *videosvr.JudgeAccessResp | *videosvr.GetVideoListResp | *videosvr.GetPreSignedUrlResp
}
