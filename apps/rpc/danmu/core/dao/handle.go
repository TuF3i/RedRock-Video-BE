package dao

import (
	"LiveDanmu/apps/public/dto"
	"LiveDanmu/apps/public/models/dao"
	"context"
	"errors"
)

// Compare cy的小工具,防止切片越界
func Compare(max int, data []dao.DanmuData) []dao.DanmuData {
	length := len(data)
	if length >= max {
		return data[:999]
	}
	return data
}

func (r *Dao) ReadHotDanmu(ctx context.Context, vid int64) ([]dao.DanmuData, dto.Response) {
	// 从redis读数据
	data, resp := r.getHotDanmuR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}
	// redis里没有就穿透到pgsql
	if len(data) == 0 {
		// 从pgsql里拉数据
		data, resp := r.getFullDanmuP(ctx, vid)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}
		// 向redis里写入hotDanmu
		resp = r.setHotDanmuR(ctx, vid, Compare(1000, data))
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}
		// 向redis里写入fullDanmu
		resp = r.setFullDanmuR(ctx, vid, data)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}

		return data, dto.OperationSuccess
	}
	// 计数器递增
	resp = r.incrementHotR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}

	return data, dto.OperationSuccess
}

func (r *Dao) ReadFullDanmu(ctx context.Context, vid int64) ([]dao.DanmuData, dto.Response) {
	// 从redis拉数据
	data, resp := r.getFullDanmuR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}
	// 如果mysql里没就走pgsql
	if len(data) == 0 {
		// 从pgsql里拉数据
		data, resp := r.getFullDanmuP(ctx, vid)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}
		// 向redis里写入fullDanmu
		resp = r.setFullDanmuR(ctx, vid, data)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}

		return data, dto.OperationSuccess
	}
	// 计数器递增
	resp = r.incrementFullR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}

	return data, dto.OperationSuccess
}
