package dto

import "LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"

// DtoResp 接口
type DtoResp interface {
	GetStatus() int64
	GetInfo() string
}

// KitexResp KitexResponse泛型接口
type KitexResp interface {
	*videosvr.AddVideoResp | *videosvr.DelVideoResp | *videosvr.JudgeAccessResp | *videosvr.GetVideoListResp | *videosvr.GetPreSignedUrlResp | *videosvr.GetJudgeListResp
}

// GenKitexResp 生成Kitex响应
func GenKitexResp[T KitexResp](resp DtoResp, data interface{}) T {
	// 声明T的实例
	var res T
	// 类型推断
	switch v := any(res).(type) {
	case *videosvr.AddVideoResp:
		v = new(videosvr.AddVideoResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *videosvr.DelVideoResp:
		v = new(videosvr.DelVideoResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *videosvr.JudgeAccessResp:
		v = new(videosvr.JudgeAccessResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *videosvr.GetVideoListResp:
		v = new(videosvr.GetVideoListResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		v.SetData(data.(*videosvr.GetVideoListData))
		res = any(v).(T)
	case *videosvr.GetPreSignedUrlResp:
		v = new(videosvr.GetPreSignedUrlResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		v.SetData(data.(string))
		res = any(v).(T)
	case *videosvr.GetJudgeListResp:
		v = new(videosvr.GetJudgeListResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		v.SetData(data.(*videosvr.GetVideoListData))
		res = any(v).(T)
	}
	return res
}
