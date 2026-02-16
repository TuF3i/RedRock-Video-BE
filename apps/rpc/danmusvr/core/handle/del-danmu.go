package handle

import (
	"LiveDanmu/apps/rpc/danmusvr/core"
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/core/pkg"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"errors"
)

func DelVideoDanmu(ctx context.Context, req *danmusvr.DelReq) dto.Response {
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
	resp := core.Dao.DelVideoDanmu(ctx, danmuData)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}
