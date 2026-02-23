package dao

import (
	"LiveDanmu/apps/shared/models"
	"context"
)

// Compare cy的小工具,防止切片越界
func Compare(max int, data []*models.DanmuData) []*models.DanmuData {
	length := len(data)
	if length >= max {
		return data[:max]
	}
	return data
}

func (r *Dao) ReadHotDanmu(ctx context.Context, vid int64) ([]*models.DanmuData, error) {
	// 从redis读数据
	data, err := r.getHotDanmuR(ctx, vid)
	if err != nil {
		return nil, err
	}
	// redis里没有就穿透到pgsql
	if len(data) == 0 {
		// 从pgsql里拉数据
		data, err := r.getFullDanmuP(ctx, vid)
		if err != nil {
			return nil, err
		}
		// 批量获取用户信息，避免N+1查询
		userIDs := make([]int64, len(data))
		for i, val := range data {
			userIDs[i] = val.UserId
		}
		userMap, _ := r.getUserInfoBatch(userIDs)
		for _, val := range data {
			if user, ok := userMap[val.UserId]; ok {
				val.User = user
			}
		}
		// 向redis里写入hotDanmu
		err = r.setHotDanmuR(ctx, vid, Compare(1000, data))
		if err != nil {
			return nil, err
		}

		// 向redis里写入fullDanmu
		err = r.setFullDanmuR(ctx, vid, data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
	// 计数器递增
	err = r.incrementHotR(ctx, vid)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Dao) ReadFullDanmu(ctx context.Context, vid int64) ([]*models.DanmuData, error) {
	// 从redis拉数据
	data, err := r.getFullDanmuR(ctx, vid)
	if err != nil {
		return nil, err
	}
	// 如果redis里没就走pgsql
	if len(data) == 0 {
		// 从pgsql里拉数据
		data, err := r.getFullDanmuP(ctx, vid)
		if err != nil {
			return nil, err
		}
		// 批量获取用户信息，避免N+1查询
		userIDs := make([]int64, len(data))
		for i, val := range data {
			userIDs[i] = val.UserId
		}
		userMap, _ := r.getUserInfoBatch(userIDs)
		for _, val := range data {
			if user, ok := userMap[val.UserId]; ok {
				val.User = user
			}
		}
		// 向redis里写入fullDanmu
		err = r.setFullDanmuR(ctx, vid, data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
	// 计数器递增
	err = r.incrementFullR(ctx, vid)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Dao) DelVideoDanmu(ctx context.Context, danID int64) error {
	// 检查弹幕是否存在
	tx := r.pgdb.Begin()
	ok, err := r.checkIfDanmuExistOnPgSQL(tx, danID)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 字段不存在直接返回
	if !ok {
		tx.Commit()
		return nil
	}

	// 删除弹幕
	err = r.delVideoDanmu(tx, danID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	// 从redis中删除整个key，下次访问时自动补位
	err = r.delDanmuInRedis(ctx, danID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Dao) IfDanmuExist(danID int64) (bool, error) {
	tx := r.pgdb.Begin()
	ok, err := r.checkIfDanmuExistOnPgSQL(tx, danID)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	// 字段不存在直接返回
	if !ok {
		tx.Commit()
		return false, nil
	}

	tx.Commit()
	return true, nil
}

func (r *Dao) GetDanmuDetail(danID int64) (*models.DanmuData, error) {
	tx := r.pgdb.Begin()
	defer tx.Rollback()

	data, err := r.getVideoDanmuDetail(tx, danID)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return data, nil
}
