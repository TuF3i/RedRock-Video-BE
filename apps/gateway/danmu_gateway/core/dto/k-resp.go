package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
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
	v := any(resp)
	if v == nil {
		return response.FinalResponse{
			Status: 0,
			Info:   "nil response",
			Data:   nil,
		}
	}
	// 组装
	switch v := any(resp).(type) {
	// 微服务响应
	case *danmusvr.PubVideoResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *danmusvr.PubLiveResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *danmusvr.GetFullResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *danmusvr.GetTopResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *danmusvr.DelLiveResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *danmusvr.DelResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case response.Response:
		return genFinalResp(v, nil)
	}
	// 兜底
	return response.FinalResponse{}
}
