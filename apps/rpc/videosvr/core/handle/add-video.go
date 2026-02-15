package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/core/pkg"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

func genVideoInfoS(data *videosvr.VideoInfo) *dao.VideoInfo {
	return &dao.VideoInfo{
		RVID:        data.Rvid,
		FaceUrl:     data.FaceUrl,
		MinioKey:    data.MinioKey,
		Title:       data.Title,
		Description: data.Description,
		UseFace:     data.UseFace,
		AuthorID:    data.AuthorId,
		AuthorName:  data.AuthorName,

		// 设置字段默认值
		InJudge: true,
		ViewNum: 0,
	}
}

func AddVideo(ctx context.Context, req *videosvr.AddVideoReq) dto.Response {
	data := req.VideoInfo
	// 校验字段
	if !pkg.ValidateRVID(data.Rvid) {
		return dto.InvalidRVID
	}
	if !pkg.ValidateFaceUrl(data.FaceUrl) {
		return dto.InvalidFaceUrl
	}
	if !pkg.ValidateMinioKey(data.MinioKey) {
		return dto.InvalidMinioKey
	}
	if !pkg.ValidateDescription(data.Description) {
		return dto.InvalidDescription
	}
	if !pkg.ValidateAuthorID(data.AuthorId) {
		return dto.InvalidAuthorID
	}

	// 从UserSvr获取用户名
	req_ := dto.GenGetUserInfoReq(data.AuthorId)
	resp, err := core.UserSvr.GetUserInfo(ctx, req_)
	if err != nil {
		return dto.ServerInternalError(err)
	}
	if pkg.ValidateAuthorName(data.AuthorName) {
		return dto.InvalidAuthorName
	}
	data.AuthorName = resp.GetData().GetUserName()

	// 转换结构体
	videoData := genVideoInfoS(data)
	// 调用dao层
	err = core.Dao.NewVideoRecord(ctx, videoData)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}
