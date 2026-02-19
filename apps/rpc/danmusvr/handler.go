package main

import (
	danmusvr "LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
)

// DanmuSvrImpl implements the last service interface defined in the IDL.
type DanmuSvrImpl struct{}

// PubVideoDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubVideoDanmu(ctx context.Context, req *danmusvr.PubVideoReq) (resp *danmusvr.PubVideoResp, err error) {
	// TODO: Your code here...
	return
}

// PubLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubLiveDanmu(ctx context.Context, req *danmusvr.PubLiveReq) (resp *danmusvr.PubLiveResp, err error) {
	// TODO: Your code here...
	return
}

// GetDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetDanmu(ctx context.Context, req *danmusvr.GetFullReq) (resp *danmusvr.GetFullResp, err error) {
	// TODO: Your code here...
	return
}

// GetTop implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetTop(ctx context.Context, req *danmusvr.GetTopReq) (resp *danmusvr.GetTopResp, err error) {
	// TODO: Your code here...
	return
}

// DelLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) DelLiveDanmu(ctx context.Context, req *danmusvr.DelLiveReq) (resp *danmusvr.DelLiveResp, err error) {
	// TODO: Your code here...
	return
}

// DelDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) DelDanmu(ctx context.Context, req *danmusvr.DelReq) (resp *danmusvr.DelResp, err error) {
	// TODO: Your code here...
	return
}
