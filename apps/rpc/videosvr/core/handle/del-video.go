package handle

import (
	"LiveDanmu/apps/public/utils"
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/core/pkg"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

func DelVideo(ctx context.Context, req *videosvr.DelVideoReq) dto.Response {
	data := req.Rvid
	// 校验数据
	if !pkg.ValidateRVID(data) {
		return dto.InvalidRVID
	}
	//
	// 调用Dao层
	tx, err := core.Dao.DelVideoRecord(ctx, data)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	// 从minio中删除视频
	ok, err := core.Minio.CheckIfVideoExist(ctx, utils.RVIDEncoder(data))
	if err != nil {
		tx.Rollback()
		return dto.ServerInternalError(err)
	}
	if ok {
		err := core.Minio.DelVideo(ctx, utils.RVIDEncoder(data))
		if err != nil {
			tx.Rollback()
			return dto.ServerInternalError(err)
		}
	}

	// 从minio中删除封面
	ok, err = core.Minio.CheckIfFaceExist(ctx, utils.RVIDEncoder(data))
	if err != nil {
		tx.Rollback()
		return dto.ServerInternalError(err)
	}
	if ok {
		err := core.Minio.DelFace(ctx, utils.RVIDEncoder(data))
		if err != nil {
			tx.Rollback()
			return dto.ServerInternalError(err)
		}
	}

	tx.Commit()
	return dto.OperationSuccess
}
