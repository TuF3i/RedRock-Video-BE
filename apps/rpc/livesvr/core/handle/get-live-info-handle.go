package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/livesvr/core"
	"LiveDanmu/apps/rpc/livesvr/core/dto"
	"LiveDanmu/apps/rpc/livesvr/core/pkg"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"
)

func convertDao2LiveDetail(raw *dao.LiveInfo) *livesvr.LiveDetail {
	return &livesvr.LiveDetail{
		Rvid:             raw.RVID,
		OwerId:           raw.OwerId,
		Title:            raw.Title,
		StreamName:       raw.StreamName,
		UpstreamPassword: raw.UpstreamPassword,
	}
}

func GetLiveInfo(ctx context.Context, req *livesvr.GetLiveInfoReq) (dto.Response, *livesvr.LiveDetail) {
	// 获取参数
	rvid := req.GetRvid()
	uid := req.GetUid()
	// 校验字段
	if !pkg.ValidateRVID(rvid) {
		return dto.InvalidRVID, nil
	}
	if !pkg.ValidateUID(uid) {
		return dto.InvalidUID, nil
	}
	// 检查是否存在
	ok, err := core.Dao.CheckIfExist(ctx, rvid)
	if !ok {
		return dto.LiveNotExist, nil
	}
	// 获取数据
	data, err := core.Dao.GetLiveInfo(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 检查权限
	if data.OwerId != uid {
		return dto.NoPermission, nil
	}

	// 转换结构体
	d := convertDao2LiveDetail(data)

	return dto.OperationSuccess, d
}
