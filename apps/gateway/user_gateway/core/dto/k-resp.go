package dto

import (
	response2 "LiveDanmu/apps/gateway/response"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
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
	case *usersvr.LoginResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.RefreshResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.GetUserInfoResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.SetAdminRoleResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case *usersvr.GetAdminerResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.GetUsersResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, v.GetData())
		}
		return genFinalResp(v, nil)
	case *usersvr.LogoutResp:
		if v.GetStatus() == 0 {
			return genFinalResp(response2.OperationSuccess, nil)
		}
		return genFinalResp(v, nil)
	case response2.Response:
		return genFinalResp(v, nil)
	}
	// 兜底
	return response2.FinalResponse{}
}
