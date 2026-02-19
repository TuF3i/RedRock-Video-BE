package handler

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core"
	"LiveDanmu/apps/gateway/danmu_gateway/core/dto"
	"LiveDanmu/apps/gateway/danmu_gateway/core/models"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/public/union_var"
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/websocket"
)

// websocket设置
var upgrader = websocket.HertzUpgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	}, // 跨域
}

func PubVideoDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData models.VideoDanmuReq
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(response.ValidateRequestFail))
			return
		}
		// 填充danmuData
		danmuData.UID = claim.Uid
		// 转换结构体
		pubReq := dto.GenPubReq(danmuData)
		// 调用PubDanmu微服务
		resp, err := core.DanmuSvr.PubVideoDanmu(ctx, pubReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
	}
}

func PubLiveDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData models.LiveDanmuReq
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(response.ValidateRequestFail))
			return
		}
		// 填充danmuData
		danmuData.UID = claim.Uid
		// 转换结构体
		pubReq := dto.GenPubLiveReq(danmuData)
		// 调用PubDanmu微服务
		resp, err := core.DanmuSvr.PubLiveDanmu(ctx, pubReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
	}
}

func GetHotDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/hot/:rvid
		// 从路由中提取rvid
		rvid := c.Param("rvid")
		if rvid == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(response.EmptyRVID))
			return
		}
		// 将string转为int64
		num, err := strconv.ParseInt(rvid, 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 生成GetTopReq
		getTopReq := dto.GenGetTopReq(num)
		// 调用GetTop
		resp, err := core.DanmuSvr.GetTop(ctx, getTopReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
	}
}

func GetFullDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/full/:rvid
		// 从路由中提取rvid
		rvid_ := c.Param("rvid")
		if rvid_ == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(response.EmptyRVID))
			return
		}
		// 将string转为int64
		rvid, err := strconv.ParseInt(rvid_, 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 生成GetTopReq
		getReq := dto.GenGetDanmuReq(rvid)
		// 调用GetDanmu
		resp, err := core.DanmuSvr.GetDanmu(ctx, getReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
			return
		}
		finalResp := dto.GenFinalResponse(resp)
		c.JSON(consts.StatusOK, finalResp)
	}
}

func DelDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 从路由中提取rvid
		rvid_ := c.Param("rvid")
		if rvid_ == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(response.EmptyRVID))
			return
		}
		// 将string转为int64
		rvid, err := strconv.ParseInt(rvid_, 10, 64)
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
			return
		}
		// 转换结构体
		delReq := dto.GenDelReq(rvid, claim.Uid)
		// 调用PubDanmu微服务
		resp, err := core.DanmuSvr.DelDanmu(ctx, delReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
			return
		}

		finalResp := dto.GenFinalResponse(resp)
		c.JSON(consts.StatusOK, finalResp)
	}
}

func LiveDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/live/:rvid
		// 将连接升级成ws
		err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
			// 从路由中提取rvid
			rvid_ := c.Param("rvid")
			if rvid_ == "" {
				c.JSON(consts.StatusOK, dto.GenFinalResponse(response.EmptyRVID))
				return
			}
			// 将string转为int64
			rvid, err := strconv.ParseInt(rvid_, 10, 64)
			if err != nil {
				resp := response.InternalError(err)
				c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
				return
			}
			// 在连接池内新建连接
			err = core.PoolGroup.AddConnToGroup(rvid, conn)
			if err != nil {
				resp := response.InternalError(err)
				c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
				return
			}
		})
		// 是否升级成功
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse(resp))
			return
		}
	}
}
