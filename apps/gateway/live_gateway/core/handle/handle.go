package handle

import (
	"LiveDanmu/apps/gateway/live_gateway/core"
	"LiveDanmu/apps/gateway/live_gateway/core/dto"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GetLiveInfoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 获取rvid
		rvid_ := c.Query("rvid")
		rvid := utils.RVIDDecoder(rvid_)
		// 构造请求
		req := dto.GenGetLiveInfoReq(rvid, claim.Uid)
		// 发起调用
		rawResp, err := core.LiveSvr.GetLiveInfo(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func GetLiveListHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
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
		// 构造请求
		req := dto.GenGetLiveListReq(int32(page), int32(pageSize))
		// 发起调用
		rawResp, err := core.LiveSvr.GetLiveList(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func StartLiveHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// TODO 获取title
		title :=
		// 构造请求
		req := dto.GenStartLiveReq()
	}
}

func StopLiveHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}

func SRSAuthHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}
