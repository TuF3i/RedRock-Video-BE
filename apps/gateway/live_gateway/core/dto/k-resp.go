package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
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
	case *livesvr.GetLiveInfoResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *livesvr.GetLiveListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *livesvr.StartLiveResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *livesvr.StopLiveResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case response.Response:
		return genFinalResp(v, nil)
	case *livesvr.GetMyLiveListResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	}
	// 兜底
	return response.FinalResponse{}
}
