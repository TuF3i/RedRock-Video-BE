package handle

import (
	"LiveDanmu/apps/rpc/danmusvr/core"
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/core/pkg"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
)

func DelVideoDanmu(ctx context.Context, req *danmusvr.DelReq) dto.Response {
	// 提取弹幕数据
	danID := req.DanId
	uid := req.Uid
	// 校验danID
	if !pkg.ValidateDanID(danID) {
		return dto.InvalidDanID
	}
	// 校验uid
	if !pkg.ValidateUserID(uid) {
		return dto.InvalidUserID
	}
	// 检查弹幕是否存在
	ok, err := core.Dao.IfDanmuExist(danID)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	if !ok {
		return dto.DanmuNotExist
	}

	// 获取弹幕信息
	data, err := core.Dao.GetDanmuDetail(danID)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	// 校验权限
	if uid != data.UserId {
		return dto.NoPermission
	}

	// 发送kafka消息
	err = core.Dao.DelVideoDanmu(ctx, danID)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}
