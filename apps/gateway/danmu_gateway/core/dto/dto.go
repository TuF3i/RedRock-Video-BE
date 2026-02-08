package dto

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core/models"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
)

func GenPubReq(raw dao.DanmuData) *danmusvr.PubReq {
	return &danmusvr.PubReq{
		DanmuMsg: &danmusvr.DanmuMsg{
			RoomId:  raw.RVID,
			UserId:  raw.UserId,
			Content: raw.Content,
			Color:   raw.Color,
			Ts:      raw.Ts,
		}}
}

func GenPubLiveReq(raw dao.DanmuData) *danmusvr.PubLiveReq {
	return &danmusvr.PubLiveReq{
		DanmuMsg: &danmusvr.DanmuMsg{
			RoomId:  raw.RVID,
			UserId:  raw.UserId,
			Content: raw.Content,
			Color:   raw.Color,
			Ts:      raw.Ts,
		}}
}

func GenGetTopReq(rvid int64) *danmusvr.GetTopReq {
	return &danmusvr.GetTopReq{BV: rvid}
}

func GenFinalResponseForGetTopReq(raw *danmusvr.GetTopResp) response.FinalResponse {
	return response.FinalResponse{
		Status: uint(raw.Status),
		Info:   raw.Info,
		Data:   raw.Data,
	}
}

func GenGetDanmuReq(rvid int64) *danmusvr.GetReq {
	return &danmusvr.GetReq{BV: rvid}
}

func GenFinalResponseForGetDanmuReq(raw *danmusvr.GetResp) response.FinalResponse {
	return response.FinalResponse{
		Status: uint(raw.Status),
		Info:   raw.Info,
		Data:   raw.Data,
	}
}

func GenAddDanmuWMsg(raw *dao.DanmuData) models.WebsocketMsg {
	return models.WebsocketMsg{
		Msg:      "danmu.add",
		Data:     raw,
		MataData: make(map[string]interface{}),
	}
}

func GenRemoveDanmuWMsg(raw *dao.DanmuData) models.WebsocketMsg {
	return models.WebsocketMsg{
		Msg:      "danmu.del",
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

func GenDelLiveReq(raw dao.DanmuData) *danmusvr.DelLiveReq {
	return &danmusvr.DelLiveReq{
		DanmuMsg: &danmusvr.DanmuMsg{
			RoomId:  raw.RVID,
			UserId:  raw.UserId,
			Content: raw.Content,
			Color:   raw.Color,
			Ts:      raw.Ts,
		}}
}

func GenDelReq(raw dao.DanmuData) *danmusvr.DelReq {
	return &danmusvr.DelReq{
		DanmuMsg: &danmusvr.DanmuMsg{
			RoomId:  raw.RVID,
			UserId:  raw.UserId,
			Content: raw.Content,
			Color:   raw.Color,
			Ts:      raw.Ts,
		}}
}
