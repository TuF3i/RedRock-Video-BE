package handle

import (
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"errors"
)

// DanmuSvrImpl implements the last handler interface defined in the IDL.
type DanmuSvrImpl struct{}

// PubVideoDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubVideoDanmu(ctx context.Context, req *danmusvr.PubVideoReq) (resp *danmusvr.PubVideoResp, err error) {
	rawResp := PubVideoDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenKitexResp[*danmusvr.PubVideoResp](rawResp, nil), rawResp
	}
	return dto.GenKitexResp[*danmusvr.PubVideoResp](rawResp, nil), nil
}

// PubLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubLiveDanmu(ctx context.Context, req *danmusvr.PubLiveReq) (resp *danmusvr.PubLiveResp, err error) {
	rawResp := PubLiveDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenKitexResp[*danmusvr.PubLiveResp](rawResp, nil), rawResp
	}
	return dto.GenKitexResp[*danmusvr.PubLiveResp](rawResp, nil), nil
}

// GetDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetDanmu(ctx context.Context, req *danmusvr.GetFullReq) (resp *danmusvr.GetFullResp, err error) {
	data, rawResp := GetFullDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenKitexResp[*danmusvr.GetFullResp](rawResp, data), rawResp
	}
	return dto.GenKitexResp[*danmusvr.GetFullResp](rawResp, data), nil
}

// GetTop implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetTop(ctx context.Context, req *danmusvr.GetTopReq) (resp *danmusvr.GetTopResp, err error) {
	data, rawResp := GetHotDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenKitexResp[*danmusvr.GetTopResp](rawResp, data), rawResp
	}
	return dto.GenKitexResp[*danmusvr.GetTopResp](rawResp, data), nil
}

// DelLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) DelLiveDanmu(ctx context.Context, req *danmusvr.DelLiveReq) (resp *danmusvr.DelLiveResp, err error) {
	// 废弃
	return
}

// DelDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) DelDanmu(ctx context.Context, req *danmusvr.DelReq) (resp *danmusvr.DelResp, err error) {
	rawResp := DelVideoDanmu(ctx, req)
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return dto.GenKitexResp[*danmusvr.DelResp](rawResp, nil), rawResp
	}
	return dto.GenKitexResp[*danmusvr.DelResp](rawResp, nil), nil
}
