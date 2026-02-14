package dto

import "LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"

func GenLoginReq(data *usersvr.RvUserInfo) *usersvr.LoginReq {
	return &usersvr.LoginReq{
		UserInfo: data,
	}
}
