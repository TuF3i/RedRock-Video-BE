package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/livesvr/core"
	"LiveDanmu/apps/rpc/livesvr/core/dto"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr"
	"context"
)

func batchA2B(raw []*dao.LiveInfo) []*livesvr.LiveInfo {
	res := make([]*livesvr.LiveInfo, 0, len(raw))
	for i, v := range raw {
		res[i] = convertDao2Livesvr(v)
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
