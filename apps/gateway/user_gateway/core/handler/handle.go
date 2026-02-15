package handler

import (
	"LiveDanmu/apps/gateway/user_gateway/core"
	"LiveDanmu/apps/gateway/user_gateway/core/dto"
	"LiveDanmu/apps/gateway/user_gateway/core/pkg"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/public/union_var"
	"context"
	"strconv"

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
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func RefreshAccessTokenHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的refreshToken
		token, _ := c.Get(union_var.JWT_REFRESH_KEY)
		// 类型断言
		refreshToken := token.(string)
		// 生成请求
		req := dto.GenRefreshAccessTokenReq(refreshToken)
		// 发起调用
		rawResp, err := core.UserSvr.RefreshToken(ctx, req)
		// 生成最终响应
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func GetUserInfoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 生成请求
		req := dto.GenGetUserInfoReq(claim.Uid)
		// 发起调用
		rawResp, err := core.UserSvr.GetUserInfo(ctx, req)
		// 生成最终响应
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func SetAdminRoleHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 权限判断
		if claim.Role != union_var.JWT_ROLE_ADMIN {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 获取uid
		// set?uid=xxx
		uid, err := strconv.ParseInt(c.Query("uid"), 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 生成请求
		req := dto.GenSetAdminRoleReq(uid)
		// 发起调用
		rawResp, err := core.UserSvr.SetAdminRole(ctx, req)
		// 生成最终响应
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func GetAdminerHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 权限判断
		if claim.Role != union_var.JWT_ROLE_ADMIN {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 获取page和pageSize
		page_ := c.Query("page")
		pageSize_ := c.Query("pageSize")
		// 类型转换
		page, err := strconv.Atoi(page_)
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}
		pageSize, err := strconv.Atoi(pageSize_)
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}
		// 构造req
		req := dto.GenGetAdminerReq(int32(page), int32(pageSize))
		// 发起调用
		rawResp, err := core.UserSvr.GetAdminer(ctx, req)
		// 生成最终响应
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func GetUsersHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 权限判断
		if claim.Role != union_var.JWT_ROLE_ADMIN {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 获取page和pageSize
		page_ := c.Query("page")
		pageSize_ := c.Query("pageSize")
		// 类型转换
		page, err := strconv.Atoi(page_)
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}
		pageSize, err := strconv.Atoi(pageSize_)
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}
		// 构造req
		req := dto.GenGetUsersReq(int32(page), int32(pageSize))
		// 发起调用
		rawResp, err := core.UserSvr.GetUsers(ctx, req)
		// 生成最终响应
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func LogoutHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 生成请求
		req := dto.GenLogoutReq(claim.Uid)
		// 发起调用
		rawResp, err := core.UserSvr.Logout(ctx, req)
		// 生成最终响应
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}
