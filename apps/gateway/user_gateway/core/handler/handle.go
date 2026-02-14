package handler

import (
	"LiveDanmu/apps/gateway/user_gateway/core"
	"LiveDanmu/apps/gateway/user_gateway/core/dto"
	"LiveDanmu/apps/gateway/user_gateway/core/pkg"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func UserLoginHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 重定向到GitHub
		core.OAuth2.AuthHandle(c)
	}
}

func CallbackHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 处理Github回调
		user := core.OAuth2.CallbackHandler(ctx, c)
		// 结构体转换
		rvUserInfo := pkg.ConvertGitHubUser2RvUserInfo(user)
		// 构造请求
		req := dto.GenLoginReq(rvUserInfo)
		// 发起调用
		rawResp, err := core.UserSvr.UserLogin(ctx, req)
		// 生成最终响应
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

	}
}

func RefreshAccessTokenHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}

func GetUserInfoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}

func SetAdminRoleHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}

func GetAdminerHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}

func GetUsersHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}

func LogoutHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}
