package dao

import (
	"LiveDanmu/apps/public/dto"
	publicDao "LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/utils"
	"context"
	"errors"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
)

const ON_CONTINUE = 1000

func (r *Dao) getHotDanmuR(ctx context.Context, vid int64) ([]publicDao.DanmuData, dto.Response) {
	var results []publicDao.DanmuData
	// 生成对应的RedisKey
	keyForHotDanmu := utils.GenHotDanmuKey(vid)
	// 读取弹幕
	rawJsonList, err := r.rdb.LRange(ctx, keyForHotDanmu, 0, -1).Result()
	if err != nil {
		return nil, dto.ServerInternalError(err)
	}
	// 空切片直接Return
	if len(rawJsonList) == 0 {
		return []publicDao.DanmuData{}, dto.OperationSuccess
	}
	// 解析Json
	for _, data := range rawJsonList {
		var DanmuData publicDao.DanmuData
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(data), &DanmuData); err != nil {
			// TODO Logger
			continue
		}
		results = append(results, DanmuData)
	}

	return results, dto.OperationSuccess
}

func (r *Dao) getFullDanmuR(ctx context.Context, vid int64) ([]publicDao.DanmuData, dto.Response) {
	var results []publicDao.DanmuData
	// 生成对应的RedisKey
	keyForFullDanmu := utils.GenFullDanmuKey(vid)
	// 读取弹幕
	rawJsonList, err := r.rdb.LRange(ctx, keyForFullDanmu, 0, -1).Result()
	if err != nil {
		return nil, dto.ServerInternalError(err)
	}
	// 空切片直接Return
	if len(rawJsonList) == 0 {
		return []publicDao.DanmuData{}, dto.OperationSuccess
	}
	// 解析Json
	for _, data := range rawJsonList {
		var DanmuData publicDao.DanmuData
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(data), &DanmuData); err != nil {
			// TODO Logger
			continue
		}
		results = append(results, DanmuData)
	}

	return results, dto.OperationSuccess
}

func (r *Dao) setHotDanmuR(ctx context.Context, vid int64, data []publicDao.DanmuData) dto.Response {
	danmuBytes := make([]interface{}, 0, len(data))
	// 生成对应的RedisKey
	keyForHotDanmu := utils.GenHotDanmuKey(vid)
	// 填充数据
	for _, danmu := range data {
		danmuString, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(danmu)
		if err != nil {
			// TODO Logger
			continue
		}
		danmuBytes = append(danmuBytes, danmuString)
	}
	// 用管道写redis
	pipe := r.rdb.Pipeline()
	pipe.LPush(ctx, keyForHotDanmu, danmuBytes...)
	pipe.Expire(ctx, keyForHotDanmu, 24*time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}

func (r *Dao) setFullDanmuR(ctx context.Context, vid int64, data []publicDao.DanmuData) dto.Response {
	danmuBytes := make([]interface{}, 0, len(data))
	// 生成对应的RedisKey
	keyForFullDanmu := utils.GenFullDanmuKey(vid)
	// 填充数据
	for _, danmu := range data {
		danmuString, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(danmu)
		if err != nil {
			// TODO Logger
			continue
		}
		danmuBytes = append(danmuBytes, danmuString)
	}
	// 用管道写redis
	pipe := r.rdb.Pipeline()
	pipe.LPush(ctx, keyForFullDanmu, danmuBytes...)
	pipe.Expire(ctx, keyForFullDanmu, 24*time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}

func (r *Dao) incrementHotR(ctx context.Context, vid int64) dto.Response {
	// 生成redis的键
	keyForHotDanmuCounter := utils.GenHotDanmuCounterKey(vid)
	keyForHotDanmu := utils.GenHotDanmuKey(vid)
	// 获取计数器值
	counter, err := r.rdb.Get(ctx, keyForHotDanmuCounter).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			counter = 0 // 显式初始化
		} else {
			return dto.ServerInternalError(err)
		}
	}
	// 如果值到达续期线就续期
	if counter >= ON_CONTINUE {
		pipe := r.rdb.Pipeline()
		pipe.Expire(ctx, keyForHotDanmu, 24*time.Hour)
		pipe.Set(ctx, keyForHotDanmuCounter, 0, 0)
		_, err := pipe.Exec(ctx)
		if err != nil {
			return dto.ServerInternalError(err)
		}
		return dto.OperationSuccess
	}
	// 递增值
	err = r.rdb.Incr(ctx, keyForHotDanmuCounter).Err()
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}

func (r *Dao) incrementFullR(ctx context.Context, vid int64) dto.Response {
	// 生成redis的键
	keyForFullDanmuCounter := utils.GenFullDanmuCounterKey(vid)
	keyForFullDanmu := utils.GenFullDanmuKey(vid)
	// 获取计数器值
	counter, err := r.rdb.Get(ctx, keyForFullDanmuCounter).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			counter = 0 // 显式初始化
		} else {
			return dto.ServerInternalError(err)
		}
	}
	// 如果值到达续期线就续期
	if counter >= ON_CONTINUE {
		pipe := r.rdb.Pipeline()
		pipe.Expire(ctx, keyForFullDanmu, 24*time.Hour)
		pipe.Set(ctx, keyForFullDanmuCounter, 0, 0)
		_, err := pipe.Exec(ctx)
		if err != nil {
			return dto.ServerInternalError(err)
		}
		return dto.OperationSuccess
	}
	// 递增值
	err = r.rdb.Incr(ctx, keyForFullDanmuCounter).Err()
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}
