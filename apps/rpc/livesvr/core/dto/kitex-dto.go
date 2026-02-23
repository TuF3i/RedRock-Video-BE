package dto

import (
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
)

// DtoResp 接口
type DtoResp interface {
	GetStatus() int64
	GetInfo() string
}

// KitexResp KitexResponse泛型接口
type KitexResp interface {
	*livesvr.GetLiveInfoResp | *livesvr.GetLiveListResp | *livesvr.StartLiveResp | *livesvr.StopLiveResp | *livesvr.GetMyLiveListResp
}

// GenKitexResp 生成Kitex响应
func GenKitexResp[T KitexResp](resp DtoResp, data interface{}) T {
	// 声明T的实例
	var res T
	// 类型推断
	switch v := any(res).(type) {
	case *livesvr.GetLiveInfoResp:
		v = new(livesvr.GetLiveInfoResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*livesvr.LiveDetail)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *livesvr.GetLiveListResp:
		v = new(livesvr.GetLiveListResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*livesvr.GetLiveListData)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *livesvr.StartLiveResp:
		v = new(livesvr.StartLiveResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*livesvr.LiveDetail)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *livesvr.StopLiveResp:
		v = new(livesvr.StopLiveResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *livesvr.GetMyLiveListResp:
		v = new(livesvr.GetMyLiveListResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*livesvr.GetMyLiveListData)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	}
	return res
}
