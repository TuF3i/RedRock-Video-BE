package handle

import (
	"LiveDanmu/apps/rpc/videosvr/core/dto"
	videosvr "LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr"
	"context"
	"errors"
)

// VideoSvrImpl implements the last service_tag.yaml interface defined in the IDL.
type VideoSvrImpl struct{}

// AddVideo implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) AddVideo(ctx context.Context, req *videosvr.AddVideoReq) (resp *videosvr.AddVideoResp, err error) {
	// 调用方法
	rawResp := AddVideo(ctx, req)
	resp = dto.GenKitexResp[*videosvr.AddVideoResp](rawResp, nil)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// DelVideo implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) DelVideo(ctx context.Context, req *videosvr.DelVideoReq) (resp *videosvr.DelVideoResp, err error) {
	// 调用方法
	rawResp := DelVideo(ctx, req)
	resp = dto.GenKitexResp[*videosvr.DelVideoResp](rawResp, nil)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// JudgeAccess implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) JudgeAccess(ctx context.Context, req *videosvr.JudgeAccessReq) (resp *videosvr.JudgeAccessResp, err error) {
	// 调用方法
	rawResp := AccessTheJudge(ctx, req)
	resp = dto.GenKitexResp[*videosvr.JudgeAccessResp](rawResp, nil)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// GetVideoList implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) GetVideoList(ctx context.Context, req *videosvr.GetVideoListReq) (resp *videosvr.GetVideoListResp, err error) {
	// 调用方法
	rawResp, data := GetVideoList(ctx, req)
	resp = dto.GenKitexResp[*videosvr.GetVideoListResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil

}

// GetPreSignedUrl implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) GetPreSignedUrl(ctx context.Context, req *videosvr.GetPreSignedUrlReq) (resp *videosvr.GetPreSignedUrlResp, err error) {
	// 调用方法
	rawResp, data := GetPreSignedUrl(ctx, req)
	resp = dto.GenKitexResp[*videosvr.GetPreSignedUrlResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}

// GetJudgeList implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) GetJudgeList(ctx context.Context, req *videosvr.GetJudgeListReq) (resp *videosvr.GetJudgeListResp, err error) {
	// 调用方法
	rawResp, data := GetJudgeList(ctx, req)
	resp = dto.GenKitexResp[*videosvr.GetJudgeListResp](rawResp, data)
	// 判断响应
	if !errors.Is(rawResp, dto.OperationSuccess) {
		return resp, rawResp
	}

	return resp, nil
}
