package dao

import (
	"LiveDanmu/apps/shared/models"
	"context"

	jsoniter "github.com/json-iterator/go"
)

func (r *Dao) newField(ctx context.Context, key string, field string, data *models.LiveInfo) error {
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

func (r *Dao) getAllFields(ctx context.Context, key string) ([]*models.LiveInfo, error) {
	rawList, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if len(rawList) == 0 {
		return nil, nil
	}

	// 预分配切片容量
	dataSet := make([]*models.LiveInfo, 0, len(rawList))
	for _, v := range rawList {
		liveInfo := &models.LiveInfo{}
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(v), liveInfo); err != nil {
			continue
		}
		dataSet = append(dataSet, liveInfo)
	}

	return dataSet, nil
}

// TODO 抄代码
func (r *Dao) getFields(ctx context.Context, key string, page int32, pageSize int32) ([]*models.LiveInfo, int64, error) {
	rawList, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, 0, err
	}

	if len(rawList) == 0 {
		return nil, 0, nil
	}

	// 将 map 的值转换为切片
	rawValues := make([]string, 0, len(rawList))
	for _, v := range rawList {
		rawValues = append(rawValues, v)
	}

	total := len(rawValues)

	// ✅ 添加分页边界检查
	offset := int((page - 1) * pageSize)
	if offset < 0 || offset >= total {
		return nil, int64(total), nil
	}

	end := offset + int(pageSize)
	if end > total {
		end = total
	}

	// 预分配切片容量
	dataSet := make([]*models.LiveInfo, 0, end-offset)
	for i := offset; i < end; i++ {
		liveInfo := &models.LiveInfo{}
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(rawValues[i]), liveInfo); err != nil {
			continue
		}
		dataSet = append(dataSet, liveInfo)
	}

	return dataSet, int64(total), nil
}

func (r *Dao) getFieldDetail(ctx context.Context, key string, field string) (*models.LiveInfo, error) {
	data := &models.LiveInfo{}

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
