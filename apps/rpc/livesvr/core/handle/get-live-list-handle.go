package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/livesvr/core"
	"LiveDanmu/apps/rpc/livesvr/core/dto"
	"LiveDanmu/apps/rpc/livesvr/core/pkg"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"
)

func convertDao2LiveListInfo(raw *dao.LiveInfo) *livesvr.LiveListInfo {
	return &livesvr.LiveListInfo{
		Rvid:       raw.RVID,
		Title:      raw.Title,
		StreamName: raw.StreamName,
		UserInfo: &livesvr.UserInfo{
			Uid:       raw.User.Uid,
			UserName:  raw.User.UserName,
			AvatarUrl: raw.User.AvatarURL,
		},
	}
}

func batchA2B(raw []*dao.LiveInfo) []*livesvr.LiveListInfo {
	res := make([]*livesvr.LiveListInfo, 0, len(raw))
	for i, v := range raw {
		res[i] = convertDao2LiveListInfo(v)
	}

	return res
}

func batchC2D(raw []*dao.LiveInfo) []*livesvr.LiveDetail {
	res := make([]*livesvr.LiveDetail, 0, len(raw))
	for i, v := range raw {
		res[i] = convertDao2LiveDetail(v)
	}

	return res
}

func GetLiveList(ctx context.Context, req *livesvr.GetLiveListReq) (dto.Response, *livesvr.GetLiveListData) {
	page := req.GetPage()
	pageSize := req.GetPageSize()
	// 从数据库读取数据
	raw, total, err := core.Dao.GetLiveList(ctx, page, pageSize)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}
	// 转换结构体
	data := batchA2B(raw)
	// 组装响应
	respData := &livesvr.GetLiveListData{
		Total: total,
		Lives: data,
	}

	return dto.OperationSuccess, respData
}

func GetMyLiveList(ctx context.Context, req *livesvr.GetMyLiveListReq) (dto.Response, *livesvr.GetMyLiveListData) {
	uid := req.GetUid()
	// 校验UID
	if !pkg.ValidateUID(uid) {
		return dto.InvalidUID, nil
	}
	// 读数据库
	raw, total, err := core.Dao.GetUserLiveList(ctx, uid)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 转换结构体
	data := batchC2D(raw)
	//组装响应
	respData := &livesvr.GetMyLiveListData{
		Total: total,
		Lives: data,
	}

	return dto.OperationSuccess, respData
}
