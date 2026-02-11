package dto

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core/models"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
)

// Pub
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

func GenFinalResponseForPubReq(raw *danmusvr.PubResp) response.FinalResponse {
	return response.FinalResponse{
		Status: raw.Status,
		Info:   raw.Info,
		Data:   nil,
	}
}

// PubLive
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

func GenFinalResponseForPubLive(raw *danmusvr.PubLiveResp) response.FinalResponse {
	return response.FinalResponse{
		Status: raw.Status,
		Info:   raw.Info,
		Data:   nil,
	}
}

// GetTop
func GenGetTopReq(rvid int64) *danmusvr.GetTopReq {
	return &danmusvr.GetTopReq{BV: rvid}
}

func GenFinalResponseForGetTopReq(raw *danmusvr.GetTopResp) response.FinalResponse {
	return response.FinalResponse{
		Status: raw.Status,
		Info:   raw.Info,
		Data:   raw.Data,
	}
}

// GetDAnmu
func GenGetDanmuReq(rvid int64) *danmusvr.GetReq {
	return &danmusvr.GetReq{BV: rvid}
}

func GenFinalResponseForGetDanmuReq(raw *danmusvr.GetResp) response.FinalResponse {
	return response.FinalResponse{
		Status: raw.Status,
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

// DelLive
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

func GenFinalResponseForDelLiveReq(raw *danmusvr.DelLiveResp) response.FinalResponse {
	return response.FinalResponse{
		Status: raw.Status,
		Info:   raw.Info,
		Data:   nil,
	}
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

func GenFinalResponseForDelReq(raw *danmusvr.DelResp) response.FinalResponse {
	return response.FinalResponse{
		Status: raw.Status,
		Info:   raw.Info,
		Data:   nil,
	}
}
