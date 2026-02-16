package handle

import (
	"LiveDanmu/apps/rpc/livesvr/core"
	"LiveDanmu/apps/rpc/livesvr/core/dto"
	"LiveDanmu/apps/rpc/livesvr/core/pkg"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"
)

func StopLive(ctx context.Context, req *livesvr.StopLiveReq) dto.Response {
	// 获取参数
	rvid := req.GetRvid()
	uid := req.GetUid()
	// 校验字段
	if !pkg.ValidateRVID(rvid) {
		return dto.InvalidRVID
	}
	if !pkg.ValidateUID(uid) {
		return dto.InvalidUID
	}
	// 检查是否存在
	ok, err := core.Dao.CheckIfExist(ctx, rvid)
	if !ok {
		return dto.LiveNotExist
	}
	// 获取数据
	data, err := core.Dao.GetLiveInfo(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err)
	}
	// 检查权限
	if data.OwerId != uid {
		return dto.NoPermission
	}
	// 关闭直播
	err = core.Dao.StopLive(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err)
	}
	return dto.OperationSuccess
}
