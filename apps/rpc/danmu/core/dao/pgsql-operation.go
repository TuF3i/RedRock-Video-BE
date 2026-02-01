package dao

import (
	"LiveDanmu/apps/public/dto"
	publicDao "LiveDanmu/apps/public/models/dao"
	"context"
)

func (r *Dao) getFullDanmuP(ctx context.Context, vid int64) ([]publicDao.DanmuData, dto.Response) {
	var results []publicDao.DanmuData
	err := r.pgdb.Where("rv_id = ?", vid).Order("ts DESC").Find(&results).Error
	if err != nil {
		return []publicDao.DanmuData{}, dto.ServerInternalError(err)
	}

	return results, dto.OperationSuccess
}
