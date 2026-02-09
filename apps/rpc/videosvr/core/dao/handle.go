package dao

import (
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"
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
