package handler

import (
	"LiveDanmu/apps/gateway/video_gateway/core"
	"LiveDanmu/apps/gateway/video_gateway/core/dto"
	"LiveDanmu/apps/gateway/video_gateway/models"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/google/uuid"
)

func UploadVideoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
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
		var data models.AddVideoReq
		// 获取上下文中的claims
		claims, ok := c.Get(union_var.JWT_CONTEXT_KEY)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 类型断言
		claim, ok := claims.(*dao.MainClaims)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&data)
		if err != nil {
			resp := dto.GenFinalResponse(response.ValidateRequestFail)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 构造请求
		req := dto.GenAddVideoReq(claim.Uid, &data)
		// 调用AddVideo
		rawResp, err := core.VideoSvr.AddVideo(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
		return
	}
}

func DelVideoHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/delete/:rvid
		// 获取rvid
		rvid_ := c.Param("rvid")
		if rvid_ == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		rvid, err := strconv.ParseInt(rvid_, 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 获取上下文中的claims
		claims, ok := c.Get(union_var.JWT_CONTEXT_KEY)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 类型断言
		claim, ok := claims.(*dao.MainClaims)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 生成请求
		req := dto.GenDelVideoReq(rvid, claim.Uid, claim.Role)
		// 调用
		rawResp, err := core.VideoSvr.DelVideo(ctx, req)
		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
		return
	}
}

func JudgeAccessHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/judge/:rvid
		// 获取rvid
		rvid_ := c.Param("rvid")
		if rvid_ == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		rvid, err := strconv.ParseInt(rvid_, 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 获取上下文中的claims
		claims, ok := c.Get(union_var.JWT_CONTEXT_KEY)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 类型断言
		claim, ok := claims.(*dao.MainClaims)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
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

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
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
		claims, ok := c.Get(union_var.JWT_CONTEXT_KEY)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 类型断言
		claim, ok := claims.(*dao.MainClaims)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
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

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
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

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
		return
	}
}

func GetPreSignedUrlHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/play/:rvid
		// 获取rvid
		rvid_ := c.Param("rvid")
		if rvid_ == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		rvid, err := strconv.ParseInt(rvid_, 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
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

			c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
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

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
		return
	}
}

func GetMyVideoListHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/my?page=xxx&pageSize=xxx
		// 获取page和pageSize
		page_ := c.Query("page")
		pageSize_ := c.Query("pageSize")
		// 获取上下文中的claims
		claims, ok := c.Get(union_var.JWT_CONTEXT_KEY)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 类型断言
		claim, ok := claims.(*dao.MainClaims)
		if !ok {
			resp := dto.GenFinalResponse(response.YouDoNotHaveAccess)
			c.JSON(consts.StatusOK, resp)
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
		req := dto.GenGetMyVideoListReq(int32(page), int32(pageSize), claim.Uid)
		// 调用
		rawResp, err := core.VideoSvr.GetMyVideoList(ctx, req)

		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
		return
	}
}

func InnocentViewNumHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// video/:rvid/innocent
		// 获取rvid
		rvid_ := c.Param("rvid")
		if rvid_ == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		rvid, err := strconv.ParseInt(rvid_, 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 构造请求
		req := dto.GenInnocentViewNumReq(rvid)
		// 发起调用
		rawResp, err := core.VideoSvr.InnocentViewNum(ctx, req)

		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
		return
	}
}

func GetVideoDetailHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取rvid
		rvid_ := c.Param("rvid")
		if rvid_ == "" {
			c.JSON(consts.StatusOK, dto.GenFinalResponse[response.Response](response.EmptyRVID))
			return
		}
		rvid, err := strconv.ParseInt(rvid_, 10, 64)
		if err != nil {
			resp := dto.GenFinalResponse(response.InternalError(err))
			c.JSON(consts.StatusOK, resp)
			return
		}
		// 生成请求
		req := dto.GetVideoDetailReq(rvid)
		// 发起调用
		rawResp, err := core.VideoSvr.GetVideoDetail(ctx, req)

		if err != nil {
			resp := dto.GenFinalResponse(rawResp)
			c.JSON(consts.StatusOK, resp)
			return
		}

		c.JSON(consts.StatusOK, dto.GenFinalResponse(rawResp))
		return
	}
}

func GetNewRvidHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		rvid := uuid.New().ID()
		resp := dto.GenFinalResponse(response.FinalResponse{
			Status: 200,
			Info:   "Operation Success",
			Data:   rvid,
		})
		c.JSON(consts.StatusOK, resp)
	}
}
