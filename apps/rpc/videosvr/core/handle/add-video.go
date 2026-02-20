package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/utils"
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/core/pkg"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"

	"go.uber.org/zap"
)

func genVideoInfoS(data *videosvr.AddVideoData) *dao.VideoInfo {
	return &dao.VideoInfo{
		RVID:        data.Rvid,
		UID:         data.Uid,
		FaceUrl:     utils.RVIDEncoder(data.Rvid),
		MinioKey:    utils.RVIDEncoder(data.Rvid),
		Title:       data.Title,
		Description: data.Description,
		ViewNum:     0,
		InJudge:     true,
	}
}

func AddVideo(ctx context.Context, req *videosvr.AddVideoReq) dto.Response {
	data := req.AddVideoData
	rvid := data.Rvid
	uid := data.Uid

	core.Logger.INFO("AddVideo start", zap.Int64("rvid", rvid), zap.Int64("uid", uid), zap.String("title", data.Title))

	if !pkg.ValidateRVID(rvid) {
		core.Logger.WARN("AddVideo invalid rvid", zap.Int64("rvid", rvid))
		return dto.InvalidRVID
	}
	if !pkg.ValidateUid(uid) {
		core.Logger.WARN("AddVideo invalid uid", zap.Int64("uid", uid))
		return dto.InvalidUid
	}
	if !pkg.ValidateTitle(data.Title) {
		core.Logger.WARN("AddVideo invalid title", zap.Int64("rvid", rvid), zap.String("title", data.Title))
		return dto.InvalidTitle
	}
	if !pkg.ValidateDescription(data.Description) {
		core.Logger.WARN("AddVideo invalid description", zap.Int64("rvid", rvid))
		return dto.InvalidDescription
	}

	videoData := genVideoInfoS(data)
	err := core.Dao.NewVideoRecord(ctx, videoData)
	if err != nil {
		core.Logger.WARN("AddVideo failed", zap.Int64("rvid", rvid), zap.Int64("uid", uid), zap.Error(err))
		return dto.ServerInternalError(err)
	}

	core.Logger.INFO("AddVideo success", zap.Int64("rvid", rvid), zap.Int64("uid", uid))
	return dto.OperationSuccess
}
