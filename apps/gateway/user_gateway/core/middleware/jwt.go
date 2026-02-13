package middleware

import (
	"LiveDanmu/apps/gateway/user_gateway/core"
	"LiveDanmu/apps/public/jwt"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/public/union_var"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func JWTMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头提取Authorization-Header
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(consts.StatusOK, response.EmptyJWTString)
			c.Abort()
		}
		// 提取AccessToken
		token := jwt.StripBearer(authHeader)
		// 验证并解析JWT
		claims, err := jwt.VerifyAccessToken(token)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			c.Abort()
		}
		// 在Redis中校验Token
		ok, err := core.Dao.CheckIfAccessTokenExist(ctx, claims.Uid, token)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			c.Abort()
		}
		// 判断Token是否正确
		if !ok {
			c.JSON(consts.StatusOK, response.JWTNotRegisteredInRedis)
			c.Abort()
		}
		// 将claims写入上下文
		c.Set(union_var.JWT_CONTEXT_KEY, claims)
		c.Next(ctx)

		return
	}
}
