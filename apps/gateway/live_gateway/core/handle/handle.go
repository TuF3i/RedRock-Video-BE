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
		data := new(dto.HStartLiveRequest)
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 获取标题
		err := c.BindAndValidate(data)
		if err != nil {
			rawResp := response.InternalError(err)
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 构造请求
		req := dto.GenStartLiveReq(claim.Uid, data.Title)
		// 发起调用
		rawResp, err := core.LiveSvr.StartLive(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func StopLiveHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 获取rvid
		rvid_ := c.Query("rvid")
		rvid := utils.RVIDDecoder(rvid_)
		// 构造请求
		req := dto.GenStopLiveReq(rvid, claim.Uid)
		// 发起请求
		rawResp, err := core.LiveSvr.StopLive(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}
		resp := dto.GenFinalResponse(rawResp)
		c.JSON(consts.StatusOK, resp)
	}
}

func SRSAuthHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取字段
		streamName := c.Query("stream")
		key := c.Query("key")
		rvid := utils.RVIDDecoder(streamName)
		// 构造请求
		req := dto.GenSRSAuthReq(rvid, key)
		// 发起调用
		resp, _ := core.LiveSvr.SRSAuth(ctx, req)

		c.JSON(consts.StatusOK, map[string]int{"code": int(resp.Ok)})
	}
}
