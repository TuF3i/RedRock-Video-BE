package dto

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
)

func GenAddVideoReq(data *dao.VideoInfo) *videosvr.AddVideoReq {
	return &videosvr.AddVideoReq{VideoInfo: &videosvr.VideoInfo{
		Rvid:        data.RVID,
		FaceUrl:     data.FaceUrl,
		MinioKey:    data.MinioKey,
		Title:       data.Title,
		Description: data.Description,
		ViewNum:     0,
		UseFace:     data.UseFace,
		InJudge:     true,
		AuthorId:    data.AuthorID,
		AuthorName:  "",
	}}
}

func GenDelVideoReq(rvid int64, uid int64, role string) *videosvr.DelVideoReq {
	return &videosvr.DelVideoReq{Rvid: rvid, Uid: uid, Role: role}
}

func GenJudgeAccessReq(rvid int64) *videosvr.JudgeAccessReq {
	return &videosvr.JudgeAccessReq{Rvid: rvid}
}

func GenGetVideoListReq(page int32, pageSize int32) *videosvr.GetVideoListReq {
	return &videosvr.GetVideoListReq{
		Page:     page,
		PageSize: pageSize,
	}
}

func GenGetJudgeListReq(page int32, pageSize int32) *videosvr.GetJudgeListReq {
	return &videosvr.GetJudgeListReq{
		Page:     page,
		PageSize: pageSize,
	}
}

func GenGetPreSignedUrlReq(rvid int64, uid int64, role string) *videosvr.GetPreSignedUrlReq {
	return &videosvr.GetPreSignedUrlReq{
		Rvid: rvid,
		Uid:  uid,
		Role: role,
	}
}
