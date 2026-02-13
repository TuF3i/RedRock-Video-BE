package dao

import (
	"LiveDanmu/apps/public/utils"
	"context"
)

func (r *Dao) CheckIfAccessTokenExist(ctx context.Context, uid int64, token string) (bool, error) {
	keyForAccessToken := utils.GenAccessTokenKey(uid)
	// 检查Token是否存在
	ok, err := r.checkKeyExistence(ctx, keyForAccessToken)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	// 检查值
	val, err := r.getKeyValue(ctx, keyForAccessToken)
	if err != nil {
		return false, err
	}
	if val != token {
		return false, nil
	}

	return true, nil
}

func (r *Dao) CheckIfRefreshTokenExist(ctx context.Context, uid int64, token string) (bool, error) {
	keyForRefreshToken := utils.GenRefreshTokenKey(uid)
	// 检查Token是否存在
	ok, err := r.checkKeyExistence(ctx, keyForRefreshToken)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	// 检查值
	val, err := r.getKeyValue(ctx, keyForRefreshToken)
	if err != nil {
		return false, err
	}
	if val != token {
		return false, nil
	}

	return true, nil
}
