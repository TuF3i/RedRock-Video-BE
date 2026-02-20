package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"LiveDanmu/apps/rpc/videosvr/core"
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	"LiveDanmu/apps/rpc/videosvr/core/pkg"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

func convertAtoB(data *dao.VideoInfo) *videosvr.VideoDetail {
	return &videosvr.VideoDetail{
		Rvid: data.RVID,
		MataInfo: &videosvr.VideoMataInfo{
			FaceKey:     data.FaceUrl,
			MinioKey:    data.MinioKey,
			Title:       data.Title,
			Description: data.Description,
			ViewNum:     data.ViewNum,
		},
		UserInfo: &videosvr.VideoUserInfo{
			AuthorId:   data.User.Uid,
			AuthorName: data.User.UserName,
			AvatarUrl:  data.User.AvatarURL,
		},
		InJudge: data.InJudge,
	}
}

func convertC2D(data *dao.VideoInfo) *videosvr.VideoListData {
	return &videosvr.VideoListData{
		Rvid:    data.RVID,
		Title:   data.Title,
		FaceKey: data.FaceUrl,
		UserInfo: &videosvr.VideoUserInfo{
			AuthorId:   data.User.Uid,
			AuthorName: data.User.UserName,
			AvatarUrl:  data.User.AvatarURL,
		},
		InJudge: data.InJudge,
	}
}

func batchDaoToVideoDetail(dataSet []*dao.VideoInfo) []*videosvr.VideoDetail {
	res := make([]*videosvr.VideoDetail, len(dataSet))
	for k, v := range dataSet {
		res[k] = convertAtoB(v)
	}
	return res
}

func batchDaoToVideoListData(dataSet []*dao.VideoInfo) []*videosvr.VideoListData {
	res := make([]*videosvr.VideoListData, len(dataSet))
	for k, v := range dataSet {
		res[k] = convertC2D(v)
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
		Videos: batchDaoToVideoListData(dataSet),
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
		Videos: batchDaoToVideoListData(dataSet),
	}

	return dto.OperationSuccess, data
}

func GetVideoDetail(ctx context.Context, req *videosvr.GetVideoDetailReq) (dto.Response, *videosvr.VideoDetail) {
	rvid := req.GetRvid()
	// 校验rvid
	if !pkg.ValidateRVID(rvid) {
		return dto.InvalidRVID, nil
	}
	// 获取视频信息
	raw, err := core.Dao.GetVideoInfo(ctx, rvid)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 转换结构体
	data := convertAtoB(raw)

	return dto.OperationSuccess, data
}

func GetMyVideoList(ctx context.Context, req *videosvr.GetMyVideoListReq) (dto.Response, *videosvr.GetVideoListData) {
	page := req.GetPage()
	pageSize := req.GetPageSize()
	uid := req.GetUid()
	// 调用Dao
	dataSet, total, err := core.Dao.GetUserVideoList(ctx, page, pageSize, uid)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 组装结构体
	data := &videosvr.GetVideoListData{
		Total:  total,
		Videos: batchDaoToVideoListData(dataSet),
	}

	return dto.OperationSuccess, data
}

func GetPreSignedUrl(ctx context.Context, req *videosvr.GetPreSignedUrlReq) (dto.Response, *string) {
	// 在数据库中查询视频信息
	vInfo, err := core.Dao.GetVideoInfo(ctx, req.GetRvid())
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 权限判断
	if vInfo.InJudge {
		if req.GetRole() != union_var.JWT_ROLE_ADMIN && req.GetUid() != vInfo.UID {
			return dto.NoPermission, nil
		}
	}
	// 游客单独处理
	if req.GetRole() == union_var.JWT_ROLE_GUEST {
		url, err := core.Minio.GetSignedUrl(ctx, utils.RVIDEncoder(req.Rvid))
		if err != nil {
			return dto.ServerInternalError(err), nil
		}

		return dto.OperationSuccess, &url
	}
	// 查询Url
	ok, err := core.Dao.IfNeedToGenNewPreSignedUrl(ctx, req.GetUid(), req.GetRvid())
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 需要续期
	if ok {
		url, err := core.Minio.GetSignedUrl(ctx, utils.RVIDEncoder(req.Rvid))
		if err != nil {
			return dto.ServerInternalError(err), nil
		}
		err = core.Dao.SetPreSignedUrlToRedis(ctx, url, req.GetUid(), req.GetRvid())
		if err != nil {
			return dto.ServerInternalError(err), nil
		}

		return dto.OperationSuccess, &url
	}
	// 不需续期
	url, err := core.Dao.GetPreSignedUrlFromRedis(ctx, req.GetUid(), req.GetRvid())
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	return dto.OperationSuccess, &url
}
