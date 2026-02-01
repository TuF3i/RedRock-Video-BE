package service

import (
	"LiveDanmu/apps/public/dto"
	"LiveDanmu/apps/rpc/danmu/core"
	"LiveDanmu/apps/rpc/danmu/core/pkg"
	"LiveDanmu/apps/rpc/danmu/kitex_gen/danmusvr"
	"context"
	"errors"
)

func PubVideoDanmu(ctx context.Context, req *danmusvr.PubReq) dto.Response {
	// 提取弹幕数据
	danmuData := req.DanmuMsg
	// 校验RoomID
	if !pkg.ValidateRoomID(danmuData.RoomId) {
		return dto.InvalidRoomID
	}
	// 校验UserID
	if !pkg.ValidateUserID(danmuData.UserId) {
		return dto.InvalidUserID
	}
	// 校验Color
	if !pkg.ValidateColor(danmuData.Color) {
		return dto.InvalidColor
	}
	// 校验Content
	if !pkg.ValidateContent(danmuData.Content) {
		return dto.InvalidContent
	}
	// 发送kafka消息
	resp := core.KClient.SendVideoDanmuMsg(ctx, danmuData)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}

func PubLiveDanmu(ctx context.Context, req *danmusvr.PubReq) dto.Response {
	// 提取弹幕数据
	danmuData := req.DanmuMsg
	// 校验RoomID
	if !pkg.ValidateRoomID(danmuData.RoomId) {
		return dto.InvalidRoomID
	}
	// 校验UserID
	if !pkg.ValidateUserID(danmuData.UserId) {
		return dto.InvalidUserID
	}
	// 校验Color
	if !pkg.ValidateColor(danmuData.Color) {
		return dto.InvalidColor
	}
	// 校验Content
	if !pkg.ValidateContent(danmuData.Content) {
		return dto.InvalidContent
	}
	// 发送kafka消息
	resp := core.KClient.SendLiveDanmuMsg(ctx, danmuData)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}
