package dto

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core/models"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"

	"github.com/google/uuid"
)

// GenPubReq Pub
func GenPubReq(raw models.VideoDanmuReq) *danmusvr.PubVideoReq {
	return &danmusvr.PubVideoReq{
		DanmuMsg: &danmusvr.PubDanmuData{
			DanId:     int64(uuid.New().ID()),
			Rvid:      raw.RVID,
			Uid:       raw.UID,
			Content:   raw.Content,
			Color:     raw.Color,
			TimeStamp: raw.Ts,
		}}
}

// GenPubLiveReq PubLive
func GenPubLiveReq(raw models.LiveDanmuReq) *danmusvr.PubLiveReq {
	return &danmusvr.PubLiveReq{
		DanmuMsg: &danmusvr.PubDanmuData{
			DanId:     int64(uuid.New().ID()),
			Rvid:      raw.RVID,
			Uid:       raw.UID,
			Content:   raw.Content,
			Color:     raw.Color,
			TimeStamp: raw.Ts,
		}}
}

// GenGetTopReq GetTop
func GenGetTopReq(rvid int64) *danmusvr.GetTopReq {
	return &danmusvr.GetTopReq{Rvid: rvid}
}

// GenGetDanmuReq GetDanmu
func GenGetDanmuReq(rvid int64) *danmusvr.GetFullReq {
	return &danmusvr.GetFullReq{Rvid: rvid}
}

func GenAddDanmuWMsg(raw *dao.DanmuData) models.WebsocketMsg {
	return models.WebsocketMsg{
		Msg:      "danmu.add",
		Data:     raw,
		MataData: make(map[string]interface{}),
	}
}

func GenLiveOffWMsg() models.WebsocketMsg {
	return models.WebsocketMsg{
		Msg:      "live.off",
		Data:     nil,
		MataData: make(map[string]interface{}),
	}
}

func GenDelReq(danID int64, uid int64) *danmusvr.DelReq {
	return &danmusvr.DelReq{DanId: danID, Uid: uid}
}
