package danmu

import (
	"LiveDanmu/apps/public/dto"
	"LiveDanmu/apps/rpc/danmu/core"
	danmusvr "LiveDanmu/apps/rpc/danmu/kitex_gen/danmusvr"
	"context"
)

// DanmuSvrImpl implements the last service interface defined in the IDL.
type DanmuSvrImpl struct{}

// PubDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubDanmu(ctx context.Context, req *danmusvr.PubReq) (resp *danmusvr.PubResp, err error) {
	rawResp := core.KClient.SendVideoDanmuMsg(ctx, req.DanmuMsg)
	return dto.GenFinalRespForPubDanMu(rawResp), nil
}

// PubLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubLiveDanmu(ctx context.Context, req *danmusvr.PubLiveReq) (resp *danmusvr.PubLiveResp, err error) {
	rawResp := core.KClient.SendLiveDanmuMsg(ctx, req.DanmuMsg)
	return dto.GenFinalRespForPubLiveDanMu(rawResp), nil
}

// GetDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetDanmu(ctx context.Context, req *danmusvr.GetReq) (resp *danmusvr.GetResp, err error) {
	// TODO: Your code here...
	return
}

// GetTop implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetTop(ctx context.Context, req *danmusvr.GetTopReq) (resp *danmusvr.GetTopResp, err error) {
	// TODO: Your code here...
	return
}
