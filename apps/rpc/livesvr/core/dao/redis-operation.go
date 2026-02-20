package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"context"

	jsoniter "github.com/json-iterator/go"
)

func (r *Dao) newField(ctx context.Context, key string, field string, data *dao.LiveInfo) error {
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

func (r *Dao) getAllFields(ctx context.Context, key string) ([]*dao.LiveInfo, error) {
	var dataSet []*dao.LiveInfo
	rawList, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	for _, v := range rawList {
		liveInfo := &dao.LiveInfo{}
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(v), liveInfo); err != nil {
			continue
		}
		dataSet = append(dataSet, liveInfo)
	}

	return dataSet, nil
}

func (r *Dao) getFields(ctx context.Context, key string, page int32, pageSize int32) ([]*dao.LiveInfo, int64, error) {
	var dataSet []*dao.LiveInfo

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
		liveInfo := &dao.LiveInfo{}
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(rawValues[i]), liveInfo); err != nil {
			continue
		}
		dataSet = append(dataSet, liveInfo)
	}

	return dataSet, int64(total), nil
}

func (r *Dao) getFieldDetail(ctx context.Context, key string, field string) (*dao.LiveInfo, error) {
	data := &dao.LiveInfo{}

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
