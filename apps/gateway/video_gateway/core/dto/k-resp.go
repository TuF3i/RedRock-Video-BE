package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"reflect"
)

func genFinalResp(resp Kresp, data interface{}) response.FinalResponse {
	return response.FinalResponse{
		Status: resp.GetStatus(),
		Info:   resp.GetInfo(),
		Data:   data,
	}
}

func GenFinalResponse[T KitexResps](resp T) response.FinalResponse {
	// 空指针检查
	t := reflect.ValueOf(resp)
	if t.Kind() == reflect.Ptr && t.IsNil() {
		return response.FinalResponse{
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
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.DelVideoResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.GetJudgeListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.JudgeAccessResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.GetVideoListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.GetPreSignedUrlResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.InnocentViewNumResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *videosvr.GetMyVideoListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *videosvr.GetVideoDetailResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case response.Response:
		return genFinalResp(v, nil)
	case response.FinalResponse:
		return v
	}
	// 兜底
	return response.FinalResponse{}
}
