package handle

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/danmusvr/core"
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"sync"
)

// 可复用指针内存池，避免重复申请内存
var danmuMsgPool = sync.Pool{New: func() interface{} { return &danmusvr.GetDanmuData{} }}

// 生成DanmuMsg
func genGetMsg(data []*dao.DanmuData) []*danmusvr.GetDanmuData {
	results := make([]*danmusvr.GetDanmuData, 0, len(data))
	for _, dm := range data {
		// 从内存池里取出一个对象
		result := danmuMsgPool.Get().(*danmusvr.GetDanmuData)
		// 覆写内存字段
		*result = danmusvr.GetDanmuData{
			DanId:     dm.DanID,
			Rvid:      dm.RVID,
			Content:   dm.Content,
			Color:     dm.Color,
			TimeStamp: dm.Ts,
			UserInfo: &danmusvr.UserInfo{
				Uid:       dm.User.Uid,
				UserName:  dm.User.UserName,
				AvatarUrl: dm.User.AvatarURL,
			},
		}
		results = append(results, result)
	}

	return results
}

// ReleaseDanmuMsg 使用完释放
func ReleaseDanmuMsg(ori []*danmusvr.GetDanmuData) {
	for _, msg := range ori {
		danmuMsgPool.Put(msg)
	}
	ori = nil
}

func GetHotDanmu(ctx context.Context, req *danmusvr.GetTopReq) ([]*danmusvr.GetDanmuData, dto.Response) {
	// 从数据库读弹幕
	dm, err := core.Dao.ReadHotDanmu(ctx, req.Rvid)
	if err != nil {
		return nil, dto.ServerInternalError(err)
	}
	// 转换结构体类型
	data := genGetMsg(dm)

	return data, dto.OperationSuccess
}

func GetFullDanmu(ctx context.Context, req *danmusvr.GetFullReq) ([]*danmusvr.GetDanmuData, dto.Response) {
	// 从数据库读弹幕
	dm, err := core.Dao.ReadFullDanmu(ctx, req.Rvid)
	if err != nil {
		return nil, dto.ServerInternalError(err)
	}
	// 转换结构体类型
	data := genGetMsg(dm)

	return data, dto.OperationSuccess
}
