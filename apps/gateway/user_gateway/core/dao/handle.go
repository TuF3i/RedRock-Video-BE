package dao

import (
	"LiveDanmu/apps/public/utils"
	"context"
)

func (r *Dao) CheckIfAccessTokenExist(ctx context.Context, token string) (bool, error) {
	keyForAccessToken := utils.GenAccessTokenKey(token)
	ok, err := r.checkKeyExistence(ctx, keyForAccessToken)
	return ok, err
}

func (r *Dao) CheckIfRefreshTokenExist(ctx context.Context, token string) (bool, error) {
	keyForRefreshToken := utils.GenRefreshTokenKey(token)
	ok, err := r.checkKeyExistence(ctx, keyForRefreshToken)
	return ok, err
}
