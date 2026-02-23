package middleware

import (
	"LiveDanmu/apps/gateway/live_gateway/core"
	response2 "LiveDanmu/apps/gateway/response"
	"LiveDanmu/apps/shared/jwt"
	"LiveDanmu/apps/shared/union_var"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func JWTMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头提取Authorization-Header
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(consts.StatusOK, response2.EmptyJWTString)
			c.Abort()
			return
		}
		// 提取AccessToken
		token := jwt.StripBearer(authHeader)
		// 验证并解析JWT
		claims, err := jwt.VerifyAccessToken(token)
		if err != nil {
			c.JSON(consts.StatusOK, response2.InternalError(err))
			c.Abort()
			return
		}
		// 在Redis中校验Token
		ok, err := core.Dao.CheckIfAccessTokenExist(ctx, claims.Uid, token)
		if err != nil {
			c.JSON(consts.StatusOK, response2.InternalError(err))
			c.Abort()
			return
		}
		// 判断Token是否正确
		if !ok {
			c.JSON(consts.StatusOK, response2.JWTNotRegisteredInRedis)
			c.Abort()
			return
		}
		// 将claims写入上下文
		c.Set(union_var.JWT_CONTEXT_KEY, claims)
		c.Next(ctx)

		return
	}
}
