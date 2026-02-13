package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/usersvr/core"
	"LiveDanmu/apps/rpc/usersvr/core/dto"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"context"
)

func convertRvUser2RvUserInfo(raw *dao.RvUser) *usersvr.RvUserInfo {
	return &usersvr.RvUserInfo{
		Uid:       raw.Uid,
		UserName:  raw.Login,
		AvatarUrl: raw.AvatarURL,
		Bio:       raw.Bio,
		Role:      raw.Role,
	}
}

func GetUserInfo(ctx context.Context, req *usersvr.GetUserInfoReq) (dto.Response, *usersvr.RvUserInfo) {
	// 获取UID
	uid := req.GetUid()
	// 查询用户是否存在
	ok, err := core.Dao.IfUserExist(uid)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 用户不存在
	if !ok {
		return dto.NoUserExist, nil
	}
	// 查询信息
	data, err := core.Dao.GetUserInfo(uid)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 转换结构体
	uInfo := convertRvUser2RvUserInfo(data)

	return dto.OperationSuccess, uInfo
}
