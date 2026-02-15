package dto

import (
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
)

// DtoResp 接口
type DtoResp interface {
	GetStatus() int64
	GetInfo() string
}

// KitexResp KitexResponse泛型接口
type KitexResp interface {
	*usersvr.LoginResp | *usersvr.RefreshResp | *usersvr.GetUserInfoResp | *usersvr.SetAdminRoleResp | *usersvr.GetAdminerResp | *usersvr.LogoutResp | *usersvr.GetUsersResp
}

// GenKitexResp 生成Kitex响应
func GenKitexResp[T KitexResp](resp DtoResp, data interface{}) T {
	// 声明T的实例
	var res T
	// 类型推断
	switch v := any(res).(type) {
	case *usersvr.LogoutResp:
		v = new(usersvr.LogoutResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *usersvr.RefreshResp:
		v = new(usersvr.RefreshResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*string)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *usersvr.GetUserInfoResp:
		v = new(usersvr.GetUserInfoResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*usersvr.RvUserInfo)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *usersvr.SetAdminRoleResp:
		v = new(usersvr.SetAdminRoleResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		res = any(v).(T)
	case *usersvr.GetAdminerResp:
		v = new(usersvr.GetAdminerResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*usersvr.GetUserListData)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *usersvr.LoginResp:
		v = new(usersvr.LoginResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*usersvr.LoginData)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	case *usersvr.GetUsersResp:
		v = new(usersvr.GetUsersResp)
		v.SetStatus(resp.GetStatus())
		v.SetInfo(resp.GetInfo())
		// 类型断言
		val, ok := data.(*usersvr.GetUserListData)
		if ok {
			v.SetData(val)
		}
		res = any(v).(T)
	}
	return res
}
