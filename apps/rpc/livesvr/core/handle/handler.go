package handle

import (
	livesvr "LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"
)

// LiveSvrImpl implements the last service interface defined in the IDL.
type LiveSvrImpl struct{}

// GetLiveInfo implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) GetLiveInfo(ctx context.Context, req *livesvr.GetLiveInfoReq) (resp *livesvr.GetLiveInfoResp, err error) {
	// TODO: Your code here...
	return
}

// GetLiveList implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) GetLiveList(ctx context.Context) (resp *livesvr.GetLiveListResp, err error) {
	// TODO: Your code here...
	return
}

// StartLive implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) StartLive(ctx context.Context, req *livesvr.StartLiveReq) (resp *livesvr.StartLiveResp, err error) {
	// TODO: Your code here...
	return
}

// StopLive implements the LiveSvrImpl interface.
func (s *LiveSvrImpl) StopLive(ctx context.Context, req *livesvr.StopLiveReq) (resp *livesvr.StopLiveResp, err error) {
	// TODO: Your code here...
	return
}
