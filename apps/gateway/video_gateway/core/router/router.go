package router

import (
	"LiveDanmu/apps/gateway/video_gateway/core"
	"LiveDanmu/apps/gateway/video_gateway/core/handler"
	"LiveDanmu/apps/gateway/video_gateway/core/middleware"
	"LiveDanmu/apps/public/config/config_template"
	"LiveDanmu/apps/public/logger/adapter"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	prometheus "github.com/hertz-contrib/monitor-prometheus"
)

var h *server.Hertz

func HertzShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.Shutdown(ctx); err != nil { // 会触发优雅停服
		return err
	}
	return nil
}

func HertzApi(conf *config_template.VideoGatewayConfig) {
	// 构造Url
	url := fmt.Sprintf("%v:%v", conf.Hertz.ListenAddr, conf.Hertz.ListenPort)
	monitorUrl := fmt.Sprintf("%v:%v", conf.Hertz.ListenAddr, conf.Hertz.MonitoringPort)
	// 创建服务核心
	h = server.Default(server.WithHostPorts(url), server.WithMaxRequestBodySize(1024*1024*1024), server.WithTracer(prometheus.NewServerTracer(monitorUrl, "/hertz")))
	// 注册TraceID生成中间件
	h.Use(middleware.TraceIDMiddleware())
	// 初始化路由
	initRouter(h)
	// 设置日志内核
	hlog.SetLogger(adapter.NewHertzZapLogger(core.Logger.Logger))
	// 启动Hertz引擎
	go func() { h.Spin() }()
}

func initRouter(h *server.Hertz) {
	g := h.Group("/video")
	{
		// 获取视频列表
		g.GET("/list", handler.GetVideoListHandleFunc())
		// 获取我的视频列表
		g.GET("/list/my", middleware.JWTMiddleware(), handler.GetMyVideoListHandleFunc())
		// 获取新的RVID
		g.GET("/new/rvid", middleware.JWTMiddleware(), handler.GetNewRvidHandleFunc())
		// 发布视频
		g.POST("/new", middleware.JWTMiddleware(), handler.AddVideoHandleFunc())
		// 增加观看数
		g.PATCH("/:rvid/innocent", handler.InnocentViewNumHandleFunc())
		// 查看视频详情
		g.GET("/:rvid/detail", handler.GetVideoDetailHandleFunc())
		// 删除视频
		g.DELETE("/:rvid", middleware.JWTMiddleware(), handler.DelVideoHandleFunc())

		// 审核子分组
		judgeGroup := g.Group("/judge")
		judgeGroup.Use(middleware.JWTMiddleware())
		{
			// 获取审核列表
			judgeGroup.GET("/list", handler.GetJudgeListHandleFunc())
			// 审核通过
			judgeGroup.PATCH("/:rvid", handler.JudgeAccessHandleFunc())
		}

		// 获取预签名URL
		g.GET("/:rvid/play-url", middleware.LooseJWTMiddleware(), handler.GetPreSignedUrlHandleFunc())

		// 上传子分组
		uploadGroup := g.Group("/:rvid/upload")
		uploadGroup.Use(middleware.JWTMiddleware())
		{
			// 上传视频文件
			uploadGroup.POST("/video", handler.UploadVideoHandleFunc())
			// 上传封面文件
			uploadGroup.POST("/cover", handler.UploadFaceHandleFunc())
		}
	}

}
