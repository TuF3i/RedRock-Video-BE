package handle

import (
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/core/pkg"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

func AccessTheJudge(ctx context.Context, req *videosvr.JudgeAccessReq) dto.Response {
	rvid := req.GetRvid()
	// 校验数据
	if !pkg.ValidateRVID(rvid) {
		return dto.InvalidRVID
	}
	// 调用dao层
	err := core.Dao.JudgeAccess(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}
