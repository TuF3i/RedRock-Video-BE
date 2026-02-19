package dto

import (
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
)

//// GenFinalRespForPubDanMu 发送弹幕最终响应
//func GenFinalRespForPubDanMu(resp Response) *danmusvr.PubResp {
//	return &danmusvr.PubResp{
//		Status: int64(resp.Status),
//		Info:   resp.Info,
//	}
//}
//
//// GenFinalRespForGetDanMu 获取所有弹幕最终响应
//func GenFinalRespForGetDanMu(resp Response, data []*danmusvr.DanmuMsg) *danmusvr.GetResp {
//	return &danmusvr.GetResp{
//		Status: int64(resp.Status),
//		Info:   resp.Info,
//		Data:   data,
//	}
//}
//
//// GenFinalRespForGetHotDanMu 获取前1000条弹幕最终响应
//func GenFinalRespForGetHotDanMu(resp Response, data []*danmusvr.DanmuMsg) *danmusvr.GetTopResp {
//	return &danmusvr.GetTopResp{
//		Status: int64(resp.Status),
//		Info:   resp.Info,
//		Data:   data,
//	}
//}
//
//// GenFinalRespForPubLiveDanMu 发送直播弹幕
//func GenFinalRespForPubLiveDanMu(resp Response) *danmusvr.PubLiveResp {
//	return &danmusvr.PubLiveResp{
//		Status: int64(resp.Status),
//		Info:   resp.Info,
//	}
//}
//
//// GenFinalRespForDelLiveDanMu 删除直播弹幕
//func GenFinalRespForDelLiveDanMu(resp Response) *danmusvr.DelLiveResp {
//	return &danmusvr.DelLiveResp{
//		Status: int64(resp.Status),
//		Info:   resp.Info,
//	}
//}
//
//// GenFinalRespForDelVideoDanMu 删除视频弹幕
//func GenFinalRespForDelVideoDanMu(resp Response) *danmusvr.DelResp {
//	return &danmusvr.DelResp{
//		Status: int64(resp.Status),
//		Info:   resp.Info,
//	}
//}

// DtoResp 接口
type DtoResp interface {
	GetStatus() int64
	GetInfo() string
}

// KitexResp KitexResponse泛型接口
type KitexResp interface {
	*danmusvr.PubVideoResp | *danmusvr.PubLiveResp | *danmusvr.GetFullResp | *danmusvr.GetTopResp | *danmusvr.DelLiveResp | *danmusvr.DelResp
}

// GenKitexResp 生成Kitex响应
func GenKitexResp[T KitexResp](resp DtoResp, data interface{}) T {
	// 声明T的实例
	var res T
	// 类型推断
	switch v := any(res).(type) {
	case *danmusvr.PubVideoResp:
		v = new(danmusvr.PubVideoResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *danmusvr.PubLiveResp:
		v = new(danmusvr.PubLiveResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *danmusvr.GetFullResp:
		v = new(danmusvr.GetFullResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.([]*danmusvr.GetDanmuData)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *danmusvr.GetTopResp:
		v = new(danmusvr.GetTopResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.([]*danmusvr.GetDanmuData)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *danmusvr.DelResp:
		v = new(danmusvr.DelResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	}
	return res
}
