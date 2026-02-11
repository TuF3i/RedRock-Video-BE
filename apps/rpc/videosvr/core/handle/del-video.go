package handle

import (
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/core/pkg"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

func DelVideo(ctx context.Context, req *videosvr.DelVideoReq) dto.Response {
	rvid := req.Rvid
	uid := req.Uid
	role := req.Role
	// 校验数据
	if !pkg.ValidateRVID(rvid) {
		return dto.InvalidRVID
	}
	if !pkg.ValidateAuthorID(uid) {
		return dto.InvalidAuthorID
	}
	// 从Dao层获取视频信息
	vInfo, err := core.Dao.GetVideoInfo(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err)
	}
	// 校验数据
	if uid != vInfo.AuthorID && role != union_var.JWT_ROLE_ADMIN {
		return dto.NoPermission
	}
	// 调用Dao层
	tx, err := core.Dao.DelVideoRecord(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	// 从minio中删除视频
	ok, err := core.Minio.CheckIfVideoExist(ctx, utils.RVIDEncoder(rvid))
	if err != nil {
		tx.Rollback()
		return dto.ServerInternalError(err)
	}
	if ok {
		err := core.Minio.DelVideo(ctx, utils.RVIDEncoder(rvid))
		if err != nil {
			tx.Rollback()
			return dto.ServerInternalError(err)
		}
	}

	// 从minio中删除封面
	ok, err = core.Minio.CheckIfFaceExist(ctx, utils.RVIDEncoder(rvid))
	if err != nil {
		tx.Rollback()
		return dto.ServerInternalError(err)
	}
	if ok {
		err := core.Minio.DelFace(ctx, utils.RVIDEncoder(rvid))
		if err != nil {
			tx.Rollback()
			return dto.ServerInternalError(err)
		}
	}

	tx.Commit()
	return dto.OperationSuccess
}
