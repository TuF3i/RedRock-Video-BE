package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/utils"
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/core/pkg"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
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
	// 校验字段
	if !pkg.ValidateRVID(data.Rvid) {
		return dto.InvalidRVID
	}
	if !pkg.ValidateUid(data.Uid) {
		return dto.InvalidUid
	}
	if !pkg.ValidateTitle(data.Title) {
		return dto.InvalidTitle
	}
	if !pkg.ValidateDescription(data.Description) {
		return dto.InvalidDescription
	}

	// 转换结构体
	videoData := genVideoInfoS(data)
	// 调用dao层
	err := core.Dao.NewVideoRecord(ctx, videoData)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}
