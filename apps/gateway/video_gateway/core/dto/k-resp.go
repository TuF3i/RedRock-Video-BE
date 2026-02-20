package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
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
	if resp == nil {
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
	case response.Response:
		return genFinalResp(v, nil)
	case response.FinalResponse:
		return v
	}
	// 兜底
	return response.FinalResponse{}
}
