package handle

import (
	"LiveDanmu/apps/rpc/usersvr/core"
	"LiveDanmu/apps/rpc/usersvr/core/dto"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"context"
)

func SetAdminRole(ctx context.Context, req *usersvr.SetAdminRoleReq) dto.Response {
	uid := req.GetUid()
	// 检查用户是否存在
	ok, err := core.Dao.IfUserExist(uid)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	if !ok {
		return dto.NoUserExist
	}

	// 将用户改为管理员
	err = core.Dao.SetAdminRole(ctx, uid)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}
