package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/utils"
	"LiveDanmu/apps/rpc/livesvr/core"
	"LiveDanmu/apps/rpc/livesvr/core/dto"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"

	"github.com/google/uuid"
)

func StartLive(ctx context.Context, req *livesvr.StartLiveReq) (dto.Response, *livesvr.LiveDetail) {
	// 获取字段值
	title := req.GetTitle()
	owerID := req.GetOwerId()
	// 生成rvid和密码
	rvid := int64(uuid.New().ID())
	upstreamPassword := uuid.New().String()[:8]
	streamName := utils.RVIDEncoder(rvid)
	// 组装结构体
	data := &dao.LiveInfo{
		RVID:             rvid,
		OwerId:           owerID,
		Title:            title,
		StreamName:       streamName,
		UpstreamPassword: upstreamPassword,
	}
	// 写入数据库
	err := core.Dao.StartLive(ctx, data)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 转换结构体
	respData := convertDao2LiveDetail(data)

	return dto.OperationSuccess, respData
}
