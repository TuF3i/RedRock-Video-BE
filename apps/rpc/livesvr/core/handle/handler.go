package handle

import (
	"LiveDanmu/apps/rpc/livesvr/core/dto"
	livesvr "LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"
	"errors"
)

// LiveSvrImpl implements the last service interface defined in the IDL.
type LiveSvrImpl struct{}

// GetLiveInfo implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) GetLiveInfo(ctx context.Context, req *livesvr.GetLiveInfoReq) (resp *livesvr.GetLiveInfoResp, err error) {
	// 调用
	rawResp, data := GetLiveInfo(ctx, req)
	// 装换上响应
	resp = dto.GenKitexResp[*livesvr.GetLiveInfoResp](rawResp, data)
	// 判断错误
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// GetLiveList implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) GetLiveList(ctx context.Context, req *livesvr.GetLiveListReq) (resp *livesvr.GetLiveListResp, err error) {
	// 调用
	rawResp, data := GetLiveList(ctx, req)
	// 装换上响应
	resp = dto.GenKitexResp[*livesvr.GetLiveListResp](rawResp, data)
	// 判断错误
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// StartLive implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) StartLive(ctx context.Context, req *livesvr.StartLiveReq) (resp *livesvr.StartLiveResp, err error) {
	// 调用
	rawResp, data := StartLive(ctx, req)
	// 装换上响应
	resp = dto.GenKitexResp[*livesvr.StartLiveResp](rawResp, data)
	// 判断错误
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// StopLive implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) StopLive(ctx context.Context, req *livesvr.StopLiveReq) (resp *livesvr.StopLiveResp, err error) {
	// 调用
	rawResp := StopLive(ctx, req)
	// 装换上响应
	resp = dto.GenKitexResp[*livesvr.StopLiveResp](rawResp, nil)
	// 判断错误
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// SRSAuth implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) SRSAuth(ctx context.Context, req *livesvr.SRSAuthReq) (resp *livesvr.SRSAuthResp, err error) {
	resp = SRSAuth(ctx, req)
	return resp, nil
}

func (s *LiveSvrImpl) GetMyLiveList(ctx context.Context, req *livesvr.GetMyLiveListReq) (resp *livesvr.GetMyLiveListResp, err error) {
	// 调用
	rawResp, data := GetMyLiveList(ctx, req)
	// 装换上响应
	resp = dto.GenKitexResp[*livesvr.GetMyLiveListResp](rawResp, data)
	// 判断错误
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}
