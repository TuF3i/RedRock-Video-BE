package dto

import (
	"LiveDanmu/apps/gateway/video_gateway/models"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
)

func GenAddVideoReq(uid int64, data *models.AddVideoReq) *videosvr.AddVideoReq {
	return &videosvr.AddVideoReq{AddVideoData: &videosvr.AddVideoData{
		Rvid:        data.Rvid,
		Uid:         uid,
		Title:       data.Title,
		Description: data.Description,
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

func GenGetMyVideoListReq(page int32, pageSize int32, uid int64) *videosvr.GetMyVideoListReq {
	return &videosvr.GetMyVideoListReq{
		Page:     page,
		PageSize: pageSize,
		Uid:      uid,
	}
}

func GenInnocentViewNumReq(rvid int64) *videosvr.InnocentViewNumReq {
	return &videosvr.InnocentViewNumReq{Rvid: rvid}
}
