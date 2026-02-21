package handle

import (
	"LiveDanmu/apps/rpc/danmusvr/core"
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/core/pkg"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"fmt"
)

func PubVideoDanmu(ctx context.Context, req *danmusvr.PubVideoReq) dto.Response {
	// 提取弹幕数据
	danmuData := req.DanmuMsg
	// 校验RoomID
	if !pkg.ValidateRoomID(danmuData.Rvid) {
		return dto.InvalidRoomID
	}
	// 校验UserID
	if !pkg.ValidateUserID(danmuData.Uid) {
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
	err := core.KClient.SendVideoDanmuMsg(ctx, danmuData)
	fmt.Printf("PubVideoDanmu Called \n")
	if err != nil {
		return dto.ServerInternalError(err)
	}
	return dto.OperationSuccess
}

func PubLiveDanmu(ctx context.Context, req *danmusvr.PubLiveReq) dto.Response {
	// 提取弹幕数据
	danmuData := req.DanmuMsg
	// 校验RoomID
	if !pkg.ValidateRoomID(danmuData.Rvid) {
		return dto.InvalidRoomID
	}
	// 校验UserID
	if !pkg.ValidateUserID(danmuData.Uid) {
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
	err := core.KClient.SendLiveDanmuMsg(ctx, danmuData)
	if err != nil {
		return dto.ServerInternalError(err)
	}
	return dto.OperationSuccess
}
