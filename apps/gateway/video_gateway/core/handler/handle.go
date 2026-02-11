package handler

import (
	"LiveDanmu/apps/gateway/video_gateway/core"
	"LiveDanmu/apps/gateway/video_gateway/core/dto"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func UploadVideoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/upload/:rvid
		// 获取rvid
		RawRvid := c.Param("rvid")
		if RawRvid == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		// 解码rvid
		rvid := utils.RVIDDecoder(RawRvid)
		// 获取文件句柄
		file, err := c.FormFile("file")
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}
		// 调用minio
		err = core.Minio.UploadFile(ctx, utils.RVIDEncoder(rvid), file)
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.OperationSuccess))
	}
}

func UploadFaceHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/upload/face/:rvid
		// 获取rvid
		RawRvid := c.Param("rvid")
		if RawRvid == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		// 解码rvid
		rvid := utils.RVIDDecoder(RawRvid)
		// 获取文件句柄
		file, err := c.FormFile("file")
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}
		// 调用minio
		err = core.Minio.UploadFaceFile(ctx, utils.RVIDEncoder(rvid), file)
		if err != nil {
			resp := response.InternalError(err)
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](resp))
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.OperationSuccess))
	}
}

func AddVideoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 视频数据
		var data dao.VideoInfo
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&data)
		if err != nil {
			resp := dto.GenFinalResponse(response.ValidateRequestFail)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 填充结构体
		data.AuthorID = claim.Uid
		// 构造请求
		req := dto.GenAddVideoReq(&data)
		// 调用AddVideo
		rawResp, err := core.VideoSvr.AddVideo(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(response.OperationSuccess))
		return
	}
}

func DelVideoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/delete/:rvid
		// 获取rvid
		RawRvid := c.Param("rvid")
		if RawRvid == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		// 解码rvid
		rvid := utils.RVIDDecoder(RawRvid)
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 生成请求
		req := dto.GenDelVideoReq(rvid, claim.Uid, claim.Role)
		// 调用
		rawResp, err := core.VideoSvr.DelVideo(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(response.OperationSuccess))
		return
	}
}

func JudgeAccessHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/judge/:rvid
		// 获取rvid
		RawRvid := c.Param("rvid")
		if RawRvid == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		// 解码rvid
		rvid := utils.RVIDDecoder(RawRvid)
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 权限判断
		if claim.Role != union_var.JWT_ROLE_ADMIN {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.YouDoNotHaveAccess))
			return
		}
		// 生成请求
		req := dto.GenJudgeAccessReq(rvid)
		// 调用
		rawResp, err := core.VideoSvr.JudgeAccess(ctx, req)

		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(response.OperationSuccess))
		return
	}
}

func GetJudgeListHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/judge?page=xxx&pageSize=xxx
		// 获取page和pageSize
		page_ := c.Query("page")
		pageSize_ := c.Query("pageSize")
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 权限判断
		if claim.Role != union_var.JWT_ROLE_ADMIN {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.YouDoNotHaveAccess))
			return
		}
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
		// 生成请求
		req := dto.GenGetJudgeListReq(int32(page), int32(pageSize))
		// 调用
		rawResp, err := core.VideoSvr.GetJudgeList(ctx, req)

		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(response.OperationSuccess))
		return
	}
}

func GetVideoListHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/judge?page=xxx&pageSize=xxx
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
		// 生成请求
		req := dto.GenGetVideoListReq(int32(page), int32(pageSize))
		// 调用
		rawResp, err := core.VideoSvr.GetVideoList(ctx, req)

		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(response.OperationSuccess))
		return
	}
}

func GetPreSignedUrlHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/play/:rvid
		// 获取rvid
		RawRvid := c.Param("rvid")
		if RawRvid == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		// 解码rvid
		rvid := utils.RVIDDecoder(RawRvid)
		// 获取上下文中的claims
		claims, ok := c.Get(union_var.JWT_CONTEXT_KEY)
		if !ok {
			// 生成请求
			req := dto.GenGetPreSignedUrlReq(rvid, 1, union_var.JWT_ROLE_GUEST)
			// 调用
			rawResp, err := core.VideoSvr.GetPreSignedUrl(ctx, req)

			if err != nil {
				resp := dto.GenFinalResponse(rawResp)
				c.JSON(consts.StatusOK, resp)
				return
			}

			c.JSON(consts.StatusOK, dto.GenFinalResponse(response.OperationSuccess))
			return
		}
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 生成请求
		req := dto.GenGetPreSignedUrlReq(rvid, claim.Uid, claim.Role)
		// 调用
		rawResp, err := core.VideoSvr.GetPreSignedUrl(ctx, req)

		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(response.OperationSuccess))
		return
	}
}
