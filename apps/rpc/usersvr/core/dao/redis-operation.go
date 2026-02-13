package dao

import (
	"context"
	"time"
)

func (r *Dao) setNewValue(ctx context.Context, key string, value string, expireTime time.Duration) error {
	err := r.rdb.Set(ctx, key, value, expireTime).Err()
	return err
}

func (r *Dao) getKeyValue(ctx context.Context, key string) (string, error) {
	data, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return data, nil
}

func (r *Dao) ifKeyExist(ctx context.Context, key string) (bool, error) {
	count, err := r.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
