package dao

import (
	"LiveDanmu/apps/public/union_var"
	"context"
	"time"
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
