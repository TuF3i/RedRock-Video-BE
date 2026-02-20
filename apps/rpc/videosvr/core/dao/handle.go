package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"
	"errors"
	"strconv"
	"sync/atomic"
	"time"

	"gorm.io/gorm"
)

func (r *Dao) SetPreSignedUrlToRedis(ctx context.Context, url string, uid int64, rvid int64) error {
	key := utils.GenPreSignedUrlKey(uid, rvid)
	err := r.setNewValue(ctx, key, url)
	if err != nil {
		return err
	}
	return nil
}

func (r *Dao) IfNeedToGenNewPreSignedUrl(ctx context.Context, uid int64, rvid int64) (bool, error) {
	key := utils.GenPreSignedUrlKey(uid, rvid)
	ttl, err := r.getExpireLast(ctx, key)
	if err != nil {
		return false, err
	}
	if ttl < union_var.MINIO_ON_CONTINUE_TIME {
		return true, nil
	}
	return false, nil
}

func (r *Dao) GetPreSignedUrlFromRedis(ctx context.Context, uid int64, rvid int64) (string, error) {
	key := utils.GenPreSignedUrlKey(uid, rvid)
	url, err := r.getValueInKey(ctx, key)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (r *Dao) NewVideoRecord(ctx context.Context, data *dao.VideoInfo) error {
	tx := r.pgdb.Begin()
	key := utils.GenVideoListKey()
	keyForUserVideoList := utils.GenUserVideoListKey(data.UID)

	ok, err := r.checkIfRecordExist(tx, data.RVID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if ok {
		tx.Rollback()
		return errors.New("record already exists")
	}
	err = r.creatNewRecord(tx, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	_ = r.delKey(ctx, key)

	_ = r.delKey(ctx, keyForUserVideoList)

	tx.Commit()
	return nil
}

func (r *Dao) DelVideoRecord(ctx context.Context, rvid int64, uid int64) (*gorm.DB, error) {
	tx := r.pgdb.Begin()
	key := utils.GenVideoListKey()
	keyForUserVideoList := utils.GenUserVideoListKey(uid)

	ok, err := r.checkIfRecordExist(tx, rvid)
	if err != nil {
		tx.Rollback()
		return tx, err
	}
	if !ok {
		tx.Rollback()
		return tx, errors.New("record not exists")
	}
	err = r.delARecord(tx, rvid)
	if err != nil {
		tx.Rollback()
		return tx, err
	}

	_ = r.delKey(ctx, key)
	_ = r.delKey(ctx, keyForUserVideoList)

	//tx.Commit()
	return tx, nil
}

func (r *Dao) GetVideoInfo(ctx context.Context, rvid int64) (*dao.VideoInfo, error) {
	tx := r.pgdb.Begin()
	key := utils.GenVideoListKey()

	data, err := r.getFieldDetail(ctx, key, strconv.FormatInt(rvid, 10))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if data != nil {
		tx.Rollback()
		return data, nil
	}

	ok, err := r.checkIfRecordExist(tx, rvid)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if !ok {
		tx.Rollback()
		return nil, errors.New("record not exists")
	}
	data, err = r.getDetailOfARecord(tx, rvid)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_ = r.newField(ctx, key, strconv.FormatInt(rvid, 10), data)

	tx.Commit()
	return data, nil
}

func (r *Dao) GetVideoList(ctx context.Context, page int32, pageSize int32) ([]*dao.VideoInfo, int64, error) {
	tx := r.pgdb.Begin()
	key := utils.GenVideoListKey()

	dataSet, total, err := r.getFields(ctx, key, page, pageSize)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	if dataSet != nil {
		tx.Rollback()
		return dataSet, total, nil
	}

	dataSet, total, err = r.getRecordList(tx, page, pageSize)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	// 异步更新
	if !r.isSyncRunning.Load() {
		go func(data []*dao.VideoInfo, r *Dao) {
			r.isSyncRunning.Store(true)
			ctx := context.Background()
			for _, v := range data {
				_ = r.newField(ctx, key, strconv.FormatInt(v.RVID, 10), v)
			}
			r.rdb.Expire(ctx, key, 24*time.Hour)
			r.isSyncRunning.Store(false)
		}(dataSet, r)
	}

	tx.Commit()
	return dataSet, total, nil
}

func (r *Dao) GetUserVideoList(ctx context.Context, page int32, pageSize int32, uid int64) ([]*dao.VideoInfo, int64, error) {
	tx := r.pgdb.Begin()
	key := utils.GenUserVideoListKey(uid)

	dataSet, total, err := r.getFields(ctx, key, page, pageSize)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	if dataSet != nil {
		tx.Rollback()
		return dataSet, total, nil
	}

	dataSet, total, err = r.getUserRecordList(tx, page, pageSize, uid)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	signal, ok := r.userSyncPool[uid]
	if !ok {
		signal = &atomic.Bool{}
		r.userSyncPool[uid] = signal
	}

	// 异步更新
	if !signal.Load() {
		go func(data []*dao.VideoInfo, r *Dao, signal *atomic.Bool) {
			signal.Store(true)
			ctx := context.Background()
			for _, v := range data {
				_ = r.newField(ctx, key, strconv.FormatInt(v.RVID, 10), v)
			}
			r.rdb.Expire(ctx, key, 24*time.Hour)
			signal.Store(false)
		}(dataSet, r, signal)
	}

	tx.Commit()
	return dataSet, total, nil
}

func (r *Dao) GetJudgingVideoList(ctx context.Context, page int32, pageSize int32) ([]*dao.VideoInfo, int64, error) {
	tx := r.pgdb.Begin()
	dataSet, total, err := r.getJudgingRecordList(tx, page, pageSize)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	tx.Commit()
	return dataSet, total, nil
}

func (r *Dao) JudgeAccess(ctx context.Context, rvid int64) error {
	tx := r.pgdb.Begin()
	key := utils.GenVideoListKey()
	err := r.setRecordColumn(tx, rvid, "in_judge", false)
	if err != nil {
		tx.Rollback()
		return err
	}

	_ = r.delKey(ctx, key)

	tx.Commit()
	return nil
}

func (r *Dao) InnocentViewNum(ctx context.Context, rvid int64) error {
	// 从redis读取数据
	data, err := r.getFieldDetail(ctx, utils.GenVideoListKey(), strconv.FormatInt(rvid, 10))
	if err != nil {
		return err
	}
	// 播放递增
	data.ViewNum++
	// 写用户视频列表
	err = r.newField(ctx, utils.GenUserVideoListKey(data.UID), strconv.FormatInt(rvid, 10), data)
	// 写总视频列表
	err = r.newField(ctx, utils.GenVideoListKey(), strconv.FormatInt(rvid, 10), data)
	// 判断是否要写入pgsql
	if r.incrementHotR(ctx, rvid) {
		tx := r.pgdb.Begin()
		err := r.setRecordColumn(tx, rvid, "view_num", data.ViewNum)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
