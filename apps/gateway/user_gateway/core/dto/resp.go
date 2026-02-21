package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
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
	if reflect.ValueOf(resp).IsNil() {
		return response.FinalResponse{
			Status: 0,
			Info:   "nil response",
			Data:   nil,
		}
	}
	// 组装
	switch v := any(resp).(type) {
	// 微服务响应
	case *usersvr.LoginResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.RefreshResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.GetUserInfoResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.SetAdminRoleResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *usersvr.GetAdminerResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.GetUsersResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.LogoutResp:
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
