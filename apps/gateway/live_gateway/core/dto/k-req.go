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

func GenStopLiveReq(rvid int64, uid int64) *livesvr.StopLiveReq {
	return &livesvr.StopLiveReq{
		Rvid: rvid,
		Uid:  uid,
	}
}

func GenSRSAuthReq(rvid int64, key string) *livesvr.SRSAuthReq {
	return &livesvr.SRSAuthReq{
		Rvid:     rvid,
		Password: key,
	}
}
