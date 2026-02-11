package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

func convertAtoB(dataSet *dao.VideoInfo) *videosvr.VideoInfo {
	return &videosvr.VideoInfo{
		Rvid:        dataSet.RVID,
		FaceUrl:     dataSet.FaceUrl,
		MinioKey:    dataSet.MinioKey,
		Title:       dataSet.Title,
		Description: dataSet.Description,
		ViewNum:     dataSet.ViewNum,
		UseFace:     dataSet.UseFace,
		InJudge:     dataSet.InJudge,
		AuthorId:    dataSet.AuthorID,
		AuthorName:  dataSet.AuthorName,
	}
}

func batchDaoToKitex(dataSet []*dao.VideoInfo) []*videosvr.VideoInfo {
	res := make([]*videosvr.VideoInfo, len(dataSet))
	for k, v := range dataSet {
		res[k] = convertAtoB(v)
	}
	return res
}

func GetVideoList(ctx context.Context, req *videosvr.GetVideoListReq) (dto.Response, *videosvr.GetVideoListData) {
	// 调用dao层
	dataSet, total, err := core.Dao.GetVideoList(ctx, req.Page, req.PageSize)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 组装结构体
	data := &videosvr.GetVideoListData{
		Total:  total,
		Videos: batchDaoToKitex(dataSet),
	}

	return dto.OperationSuccess, data
}

func GetJudgeList(ctx context.Context, req *videosvr.GetJudgeListReq) (dto.Response, *videosvr.GetVideoListData) {
	// 调用dao层
	dataSet, total, err := core.Dao.GetJudgingVideoList(ctx, req.Page, req.PageSize)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 组装结构体
	data := &videosvr.GetVideoListData{
		Total:  total,
		Videos: batchDaoToKitex(dataSet),
	}

	return dto.OperationSuccess, data
}

func GetPreSignedUrl(ctx context.Context, req *videosvr.GetPreSignedUrlReq) (dto.Response, string) {
	// 在数据库中查询视频信息
	vInfo, err := core.Dao.GetVideoInfo(ctx, req.GetRvid())
	if err != nil {
		return dto.ServerInternalError(err), ""
	}
	// 权限判断
	if vInfo.InJudge {
		if req.GetRole() != union_var.JWT_ROLE_ADMIN && req.GetUid() != vInfo.AuthorID {
			return dto.NoPermission, ""
		}
	}
	// 游客单独处理
	if req.GetRole() == union_var.JWT_ROLE_GUEST {
		url, err := core.Minio.GetSignedUrl(ctx, utils.RVIDEncoder(req.Rvid))
		if err != nil {
			return dto.ServerInternalError(err), ""
		}

		return dto.OperationSuccess, url
	}
	// 查询Url
	ok, err := core.Dao.IfNeedToGenNewPreSignedUrl(ctx, req.GetUid(), req.GetRvid())
	if err != nil {
		return dto.ServerInternalError(err), ""
	}
	// 需要续期
	if ok {
		url, err := core.Minio.GetSignedUrl(ctx, utils.RVIDEncoder(req.Rvid))
		if err != nil {
			return dto.ServerInternalError(err), ""
		}
		err = core.Dao.SetPreSignedUrlToRedis(ctx, url, req.GetUid(), req.GetRvid())
		if err != nil {
			return dto.ServerInternalError(err), ""
		}

		return dto.OperationSuccess, url
	}
	// 不需续期
	url, err := core.Dao.GetPreSignedUrlFromRedis(ctx, req.GetUid(), req.GetRvid())
	if err != nil {
		return dto.ServerInternalError(err), ""
	}

	return dto.OperationSuccess, url
}
