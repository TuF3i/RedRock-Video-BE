package dto

import (
	response2 "LiveDanmu/apps/gateway/response"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"reflect"
)

func genFinalResp(resp Kresp, data interface{}) response2.FinalResponse {
	return response2.FinalResponse{
		Status: resp.GetStatus(),
		Info:   resp.GetInfo(),
		Data:   data,
	}
}

func GenFinalResponse[T KitexResps](resp T) response2.FinalResponse {
	// 空指针检查
	t := reflect.ValueOf(resp)
	if t.Kind() == reflect.Ptr && t.IsNil() {
		return response2.FinalResponse{
			Status: 0,
			Info:   "nil response",
			Data:   nil,
		}
	}
	// 组装
	switch v := any(resp).(type) {
	// 微服务响应
	case *videosvr.AddVideoResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.DelVideoResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.GetJudgeListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.JudgeAccessResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.GetVideoListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.GetPreSignedUrlResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.InnocentViewNumResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.GetMyVideoListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.GetVideoDetailResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case response2.Response:
		return genFinalResp(v, nil)
	case response2.FinalResponse:
		return v
	}
	// 兜底
	return response2.FinalResponse{}
}
