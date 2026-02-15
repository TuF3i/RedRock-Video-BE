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

func batchRvUser2RvUserInfo(raw []*dao.RvUser) []*usersvr.RvUserInfo {
	dataSet := make([]*usersvr.RvUserInfo, 0, len(raw))
	for i, v := range raw {
		dataSet[i] = convertRvUser2RvUserInfo(v)
	}
	return dataSet
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

func GetUserList(ctx context.Context, req *usersvr.GetUsersReq) (dto.Response, *usersvr.GetUserListData) {
	page := req.GetPage()
	pageSize := req.GetPageSize()
	// 获取数据
	rawData, total, err := core.Dao.GetUserList(ctx, page, pageSize)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 转换结构体
	users := batchRvUser2RvUserInfo(rawData)

	// 组装GetUserListData
	data := &usersvr.GetUserListData{
		Total: total,
		Users: users,
	}

	return dto.OperationSuccess, data
}

func GetAdminList(ctx context.Context, req *usersvr.GetAdminerReq) (dto.Response, *usersvr.GetUserListData) {
	page := req.GetPage()
	pageSize := req.GetPageSize()
	// 获取数据
	rawData, total, err := core.Dao.GetAdminerList(ctx, page, pageSize)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 转换结构体
	users := batchRvUser2RvUserInfo(rawData)

	// 组装GetUserListData
	data := &usersvr.GetUserListData{
		Total: total,
		Users: users,
	}

	return dto.OperationSuccess, data
}
