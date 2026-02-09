package handler

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core"
	"LiveDanmu/apps/gateway/danmu_gateway/core/dto"
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

func PubDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData dao.DanmuData
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, response.ValidateRequestFail)
			return
		}
		// 填充danmuData
		danmuData.UserId = claim.Uid
		// 转换结构体
		pubReq := dto.GenPubReq(danmuData)
		// 调用PubDanmu微服务
		resp, err := core.DanmuSvr.PubDanmu(ctx, pubReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponseForPubReq(resp))
			return
		}

		c.JSON(consts.StatusOK, response.OperationSuccess)
	}
}

func PubLiveDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData dao.DanmuData
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, response.ValidateRequestFail)
			return
		}
		// 填充danmuData
		danmuData.UserId = claim.Uid
		// 转换结构体
		pubReq := dto.GenPubLiveReq(danmuData)
		// 调用PubDanmu微服务
		resp, err := core.DanmuSvr.PubLiveDanmu(ctx, pubReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponseForPubLive(resp))
			return
		}

		c.JSON(consts.StatusOK, response.OperationSuccess)
	}
}

func GetHotDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/hot/:rvid
		// 从路由中提取rvid
		rvid := c.Param("rvid")
		if rvid == "" {
			c.JSON(consts.StatusOK, response.EmptyRVID)
			return
		}
		// 将string转为int64
		num, err := strconv.ParseInt(rvid, 10, 64)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}
		// 生成GetTopReq
		getTopReq := dto.GenGetTopReq(num)
		// 调用GetTop
		resp, err := core.DanmuSvr.GetTop(ctx, getTopReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponseForGetTopReq(resp))
			return
		}
		finalResp := dto.GenFinalResponseForGetTopReq(resp)
		c.JSON(consts.StatusOK, finalResp)
	}
}

func GetFullDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/full/:rvid
		// 从路由中提取rvid
		rvid := c.Param("rvid")
		if rvid == "" {
			c.JSON(consts.StatusOK, response.EmptyRVID)
			return
		}
		// 将string转为int64
		num, err := strconv.ParseInt(rvid, 10, 64)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}
		// 生成GetTopReq
		getReq := dto.GenGetDanmuReq(num)
		// 调用GetDanmu
		resp, err := core.DanmuSvr.GetDanmu(ctx, getReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponseForGetDanmuReq(resp))
			return
		}
		finalResp := dto.GenFinalResponseForGetDanmuReq(resp)
		c.JSON(consts.StatusOK, finalResp)
	}
}

func DelLiveDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData dao.DanmuData
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, response.ValidateRequestFail)
			return
		}
		// 校验权限
		if claim.Role != union_var.JWT_ROLE_ADMIN || claim.Uid != danmuData.UserId {
			c.JSON(consts.StatusOK, response.YouDoNotHaveAccess)
			return
		}
		// 转换结构体
		delLiveReq := dto.GenDelLiveReq(danmuData)
		// 调用PubDanmu微服务
		resp, err := core.DanmuSvr.DelLiveDanmu(ctx, delLiveReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponseForDelLiveReq(resp))
			return
		}

		c.JSON(consts.StatusOK, response.OperationSuccess)
	}
}

func DelDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData dao.DanmuData
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, response.ValidateRequestFail)
			return
		}
		// 校验权限
		if claim.Role != union_var.JWT_ROLE_ADMIN || claim.Uid != danmuData.UserId {
			c.JSON(consts.StatusOK, response.YouDoNotHaveAccess)
			return
		}
		// 转换结构体
		delReq := dto.GenDelReq(danmuData)
		// 调用PubDanmu微服务
		resp, err := core.DanmuSvr.DelDanmu(ctx, delReq)
		if err != nil {
			c.JSON(consts.StatusOK, dto.GenFinalResponseForDelReq(resp))
			return
		}

		c.JSON(consts.StatusOK, response.OperationSuccess)
	}
}

func LiveDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/live/:rvid
		// 将连接升级成ws
		err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
			// 从路由中提取rvid
			rvid := c.Param("rvid")
			if rvid == "" {
				c.JSON(consts.StatusOK, response.EmptyRVID)
				return
			}
			// 将string转为int64
			num, err := strconv.ParseInt(rvid, 10, 64)
			if err != nil {
				c.JSON(consts.StatusOK, response.InternalError(err))
				return
			}
			// 在连接池内新建连接
			err = core.PoolGroup.AddConnToGroup(num, conn)
			if err != nil {
				c.JSON(consts.StatusOK, response.InternalError(err))
				return
			}
		})
		// 是否升级成功
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}
	}
}
