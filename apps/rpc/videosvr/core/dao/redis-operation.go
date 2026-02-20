package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"
	"errors"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
)

const (
	ON_REFRESH = 500 // 500播放就刷新
)

func (r *Dao) checkIfKeyExist(ctx context.Context, key string) (bool, error) {
	count, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *Dao) setNewValue(ctx context.Context, key string, value string) error {
	err := r.rdb.Set(ctx, key, value, union_var.MINIO_EXPIRE_TIME).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Dao) getExpireLast(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := r.rdb.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

func (r *Dao) getValueInKey(ctx context.Context, key string) (string, error) {
	value, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (r *Dao) newField(ctx context.Context, key string, field string, data *dao.VideoInfo) error {
	// 序列化
	raw, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		return err
	}
	// 存如Redis
	err = r.rdb.HSet(ctx, key, field, raw).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Dao) delKey(ctx context.Context, key string) error {
	err := r.rdb.Del(ctx, key).Err()
	return err
}

func (r *Dao) getFields(ctx context.Context, key string, page int32, pageSize int32) ([]*dao.VideoInfo, int64, error) {
	var dataSet []*dao.VideoInfo

	rawList, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, 0, err
	}

	if len(rawList) == 0 {
		return nil, 0, nil
	}

	var rawValues []string
	for _, v := range rawList {
		rawValues = append(rawValues, v)
	}

	total := len(rawList)
	offset := (page - 1) * pageSize
	end := offset + pageSize

	for i := offset; i < end; i++ {
		videoInfo := &dao.VideoInfo{}
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(rawValues[i]), videoInfo); err != nil {
			continue
		}
		dataSet = append(dataSet, videoInfo)
	}

	return dataSet, int64(total), nil
}

func (r *Dao) getFieldDetail(ctx context.Context, key string, field string) (*dao.VideoInfo, error) {
	data := &dao.VideoInfo{}

	raw, err := r.rdb.HGet(ctx, key, field).Result()
	if err != nil {
		return nil, err
	}

	err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(raw), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Dao) incrementHotR(ctx context.Context, rvid int64) bool {
	// 生成redis的键
	keyForVideoFieldCounter := utils.GenVideoFieldCounterKey(rvid)
	// 获取计数器值
	counter, err := r.rdb.Get(ctx, keyForVideoFieldCounter).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			counter = 0 // 显式初始化
		} else {
			return false
		}
	}
	// 如果值到达续期线就续期
	if counter >= ON_REFRESH {
		r.rdb.Set(ctx, keyForVideoFieldCounter, 0, 0)
		return true
	}
	// 递增值
	err = r.rdb.Incr(ctx, keyForVideoFieldCounter).Err()
	if err != nil {
		return false
	}

	return false
}
