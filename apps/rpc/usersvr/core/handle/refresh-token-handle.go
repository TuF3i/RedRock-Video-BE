package handle

import (
	"LiveDanmu/apps/public/jwt"
	"LiveDanmu/apps/rpc/usersvr/core"
	"LiveDanmu/apps/rpc/usersvr/core/dto"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr"
	"context"
)

func GetRefreshToken(ctx context.Context, req *usersvr.RefreshReq) (dto.Response, *string) {
	// 读取refreshToken
	refreshToken := req.GetRefreshToken()

	// 解析RefreshToken
	claims, err := jwt.VerifyRefreshToken(refreshToken)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 获取claim字段
	uid := claims.Uid
	role := claims.Role

	// 在redis内验证refreshToken
	ok, err := core.Dao.VerifyRefreshToken(ctx, uid, refreshToken)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	if !ok {
		return dto.InvalidRefreshToken, nil
	}

	// 生成新的AccessToken
	accessToken, err := jwt.GenerateAccessToken(uid, role)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	// 将新accessToken写入Redis
	err = core.Dao.SetNewAccessToken(ctx, uid, accessToken)
	if err != nil {
		return dto.ServerInternalError(err), nil
	}

	return dto.OperationSuccess, &accessToken
}
