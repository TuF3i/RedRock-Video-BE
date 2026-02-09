package dto

import "LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"

var (
	OperationSuccess = Response{Status: 40200, Info: "Operation Success"}
	InvalidRoomID    = Response{Status: 40001, Info: "Invalid RoomID"}
	InvalidUserID    = Response{Status: 40002, Info: "Invalid UserID"}
	InvalidColor     = Response{Status: 40003, Info: "Invalid Color"}
	InvalidContent   = Response{Status: 40004, Info: "Invalid Content"}
)

// Response 业务层错误封装
type Response struct {
	Status uint   `json:"status"`
	Info   string `json:"info"`
}

func (r Response) Error() string {
	return r.Info
}

// ServerInternalError 服务器内部错误封装
func ServerInternalError(err error) Response {
	return Response{
		Status: 500,
		Info:   err.Error(),
	}
}

// GenFinalRespForPubDanMu 发送弹幕最终响应
func GenFinalRespForPubDanMu(resp Response) *danmusvr.PubResp {
	return &danmusvr.PubResp{
		Status: int64(resp.Status),
		Info:   resp.Info,
	}
}

// GenFinalRespForGetDanMu 获取所有弹幕最终响应
func GenFinalRespForGetDanMu(resp Response, data []*danmusvr.DanmuMsg) *danmusvr.GetResp {
	return &danmusvr.GetResp{
		Status: int64(resp.Status),
		Info:   resp.Info,
		Data:   data,
	}
}

// GenFinalRespForGetHotDanMu 获取前1000条弹幕最终响应
func GenFinalRespForGetHotDanMu(resp Response, data []*danmusvr.DanmuMsg) *danmusvr.GetTopResp {
	return &danmusvr.GetTopResp{
		Status: int64(resp.Status),
		Info:   resp.Info,
		Data:   data,
	}
}

// GenFinalRespForPubLiveDanMu 发送直播弹幕
func GenFinalRespForPubLiveDanMu(resp Response) *danmusvr.PubLiveResp {
	return &danmusvr.PubLiveResp{
		Status: int64(resp.Status),
		Info:   resp.Info,
	}
}

// GenFinalRespForDelLiveDanMu 删除直播弹幕
func GenFinalRespForDelLiveDanMu(resp Response) *danmusvr.DelLiveResp {
	return &danmusvr.DelLiveResp{
		Status: int64(resp.Status),
		Info:   resp.Info,
	}
}

// GenFinalRespForDelVideoDanMu 删除视频弹幕
func GenFinalRespForDelVideoDanMu(resp Response) *danmusvr.DelResp {
	return &danmusvr.DelResp{
		Status: int64(resp.Status),
		Info:   resp.Info,
	}
}
