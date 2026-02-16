package dto

import "LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"

func GenGetLiveInfoReq(rvid int64, uid int64) *livesvr.GetLiveInfoReq {
	return &livesvr.GetLiveInfoReq{
		Rvid: rvid,
		Uid:  uid,
	}
}

func GenGetLiveListReq(page int32, pageSize int32) *livesvr.GetLiveListReq {
	return &livesvr.GetLiveListReq{
		Page:     page,
		PageSize: pageSize,
	}
}

func GenStartLiveReq(owerID int64, title string) *livesvr.StartLiveReq {
	return &livesvr.StartLiveReq{
		OwerId: owerID,
		Title:  title,
	}
}
