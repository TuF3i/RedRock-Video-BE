package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/utils"
	"context"
	"errors"
	"strconv"
)

func (r *Dao) StartLive(ctx context.Context, data *dao.LiveInfo) error {
	// 启动数据库事物
	tx := r.pgdb.Begin()
	// 检查pgsql中是否存在记录
	ok, err := r.ifRecordExist(tx, data.RVID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if ok {
		tx.Rollback()
		return nil
	}
	// 向pgsql中写入值
	err = r.newRecord(tx, data)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 向redis中写入
	key := utils.GenLiveListKey()
	err = r.newField(ctx, key, strconv.FormatInt(data.RVID, 10), data)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *Dao) StopLive(ctx context.Context, rvid int64) error {
	// 启动数据库事物
	tx := r.pgdb.Begin()
	// 删除redis的整个key
	key := utils.GenLiveListKey()
	err := r.delKey(ctx, key)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 检查pgsql中是否存在记录
	ok, err := r.ifRecordExist(tx, rvid)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !ok {
		tx.Rollback()
		return nil
	}
	// 从pgsql中删除
	err = r.delRecord(tx, rvid)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *Dao) GetLiveList(ctx context.Context, page int32, pageSize int32) ([]*dao.LiveInfo, int64, error) {
	// 启动数据库事物
	tx := r.pgdb.Begin()
	key := utils.GenLiveListKey()
	// 从redis中读取缓存
	data, total, err := r.getFields(ctx, key, page, pageSize)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	if data != nil {
		tx.Rollback()
		return data, total, nil
	}

	// 从pgsql里读数据
	data, total, err = r.getRecords(tx, page, pageSize)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	// 异步更新
	if !r.isSyncRunning.Load() {
		go func(data []*dao.LiveInfo, r *Dao) {
			r.isSyncRunning.Store(true)
			ctx := context.Background()
			for _, v := range data {
				_ = r.newField(ctx, key, strconv.FormatInt(v.RVID, 10), v)
			}
			r.isSyncRunning.Store(false)
		}(data, r)
	}

	tx.Commit()
	return data, total, nil
}

func (r *Dao) GetLiveInfo(ctx context.Context, rvid int64) (*dao.LiveInfo, error) {
	// 启动数据库事物
	tx := r.pgdb.Begin()
	key := utils.GenLiveListKey()
	// 从redis读取缓存
	data, err := r.getFieldDetail(ctx, key, strconv.FormatInt(rvid, 10))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if data != nil {
		tx.Rollback()
		return data, nil
	}
	// 检查pgsql中是否存在结果
	ok, err := r.ifRecordExist(tx, rvid)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if !ok {
		tx.Rollback()
		return nil, errors.New("live not exists")
	}
	// 从pgsql里读取数据
	data, err = r.getRecordDetail(tx, rvid)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 向redis内写入数据
	_ = r.newField(ctx, key, strconv.FormatInt(rvid, 10), data)

	tx.Commit()
	return data, nil
}

func (r *Dao) CheckIfExist(ctx context.Context, rvid int64) (bool, error) {
	//开启数据库事务
	tx := r.pgdb.Begin()
	defer tx.Commit()
	// 检查记录是否存在
	ok, err := r.ifRecordExist(tx, rvid)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	return true, nil
}
