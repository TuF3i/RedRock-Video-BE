package dto

import (
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
)

type KitexReqs interface {
	*usersvr.LoginReq | *usersvr.RefreshReq | *usersvr.GetUserInfoReq | *usersvr.SetAdminRoleReq | *usersvr.LogoutReq
}

type KitexResps interface {
	*usersvr.LoginResp | *usersvr.RefreshResp | *usersvr.GetUserInfoResp | *usersvr.SetAdminRoleResp | *usersvr.GetAdminerResp | *usersvr.GetUsersResp | *usersvr.LogoutResp | response.Response
}

type Kresp interface {
	GetStatus() int64
	GetInfo() string
}

var (
	// 编译期校验：实现Kresp
	_ Kresp = (*usersvr.LoginResp)(nil)
	_ Kresp = (*usersvr.RefreshResp)(nil)
	_ Kresp = (*usersvr.GetUserInfoResp)(nil)
	_ Kresp = (*usersvr.SetAdminRoleResp)(nil)
	_ Kresp = (*usersvr.GetAdminerResp)(nil)
	_ Kresp = (*usersvr.GetUsersResp)(nil)
	_ Kresp = (*usersvr.LogoutResp)(nil)
	_ Kresp = response.Response{}
)
