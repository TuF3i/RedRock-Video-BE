package main

import (
	videosvr "LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
)

// VideoSvrImpl implements the last service interface defined in the IDL.
type VideoSvrImpl struct{}

// AddVideo implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) AddVideo(ctx context.Context, req *videosvr.AddVideoReq) (resp *videosvr.AddVideoResp, err error) {
	// TODO: Your code here...
	return
}

// DelVideo implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) DelVideo(ctx context.Context, req *videosvr.DelVideoReq) (resp *videosvr.DelVideoResp, err error) {
	// TODO: Your code here...
	return
}

// JudgeAccess implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) JudgeAccess(ctx context.Context, req *videosvr.JudgeAccessReq) (resp *videosvr.JudgeAccessResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideoList implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) GetVideoList(ctx context.Context, req *videosvr.GetPreSignedUrlReq) (resp *videosvr.GetVideoListResp, err error) {
	// TODO: Your code here...
	return
}

// GetPreSignedUrl implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) GetPreSignedUrl(ctx context.Context, req *videosvr.GetPreSignedUrlReq) (resp *videosvr.GetPreSignedUrlResp, err error) {
	// TODO: Your code here...
	return
}
