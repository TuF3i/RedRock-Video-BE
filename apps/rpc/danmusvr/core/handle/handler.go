package handle

import (
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"errors"
)

// DanmuSvrImpl implements the last handler interface defined in the IDL.
type DanmuSvrImpl struct{}

// PubDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubDanmu(ctx context.Context, req *danmusvr.PubReq) (resp *danmusvr.PubResp, err error) {
	rawResp := PubVideoDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenFinalRespForPubDanMu(rawResp), rawResp
	}
	return dto.GenFinalRespForPubDanMu(rawResp), nil
}

// PubLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubLiveDanmu(ctx context.Context, req *danmusvr.PubLiveReq) (resp *danmusvr.PubLiveResp, err error) {
	rawResp := PubLiveDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenFinalRespForPubLiveDanMu(rawResp), rawResp
	}
	return dto.GenFinalRespForPubLiveDanMu(rawResp), nil
}

// GetDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetDanmu(ctx context.Context, req *danmusvr.GetReq) (resp *danmusvr.GetResp, err error) {
	data, rawResp := GetFullDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenFinalRespForGetDanMu(rawResp, data), rawResp
	}
	return dto.GenFinalRespForGetDanMu(rawResp, data), nil
}

// GetTop implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetTop(ctx context.Context, req *danmusvr.GetTopReq) (resp *danmusvr.GetTopResp, err error) {
	data, rawResp := GetHotDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenFinalRespForGetHotDanMu(rawResp, data), rawResp
	}
	return dto.GenFinalRespForGetHotDanMu(rawResp, data), nil
}

// DelLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) DelLiveDanmu(ctx context.Context, req *danmusvr.DelLiveReq) (resp *danmusvr.DelLiveResp, err error) {
	rawResp := DelLiveDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenFinalRespForDelLiveDanMu(rawResp), rawResp
	}
	return dto.GenFinalRespForDelLiveDanMu(rawResp), nil
}

// DelDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) DelDanmu(ctx context.Context, req *danmusvr.DelReq) (resp *danmusvr.DelResp, err error) {
	rawResp := DelVideoDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenFinalRespForDelVideoDanMu(rawResp), rawResp
	}
	return dto.GenFinalRespForDelVideoDanMu(rawResp), nil
}
