package dto

import "LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"

func GenLoginReq(data *usersvr.RvUserInfo) *usersvr.LoginReq {
	return &usersvr.LoginReq{
		UserInfo: data,
	}
}

func GenRefreshAccessTokenReq(refreshToken string) *usersvr.RefreshReq {
	return &usersvr.RefreshReq{RefreshToken: refreshToken}
}

func GenGetUserInfoReq(uid int64) *usersvr.GetUserInfoReq {
	return &usersvr.GetUserInfoReq{Uid: uid}
}

func GenSetAdminRoleReq(uid int64) *usersvr.SetAdminRoleReq {
	return &usersvr.SetAdminRoleReq{Uid: uid}
}

func GenLogoutReq(uid int64) *usersvr.LogoutReq {
	return &usersvr.LogoutReq{Uid: uid}
}

func GenGetAdminerReq(page int32, pageSize int32) *usersvr.GetAdminerReq {
	return &usersvr.GetAdminerReq{
		Page:     page,
		PageSize: pageSize,
	}
}

func GenGetUsersReq(page int32, pageSize int32) *usersvr.GetUsersReq {
	return &usersvr.GetUsersReq{
		Page:     page,
		PageSize: pageSize,
	}
}
