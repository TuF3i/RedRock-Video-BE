package dto

import "LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"

func GenGetUserInfoReq(uid int64) *usersvr.GetUserInfoReq {
	return &usersvr.GetUserInfoReq{Uid: uid}
}
