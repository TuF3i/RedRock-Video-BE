package handle

import (
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

func InnocentViewNum(ctx context.Context, req *videosvr.InnocentViewNumReq) dto.Response {
	rvid := req.GetRvid()
	// 调用dao
	err := core.Dao.InnocentViewNum(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}
