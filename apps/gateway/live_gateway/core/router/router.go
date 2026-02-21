package router

import (
	"LiveDanmu/apps/gateway/live_gateway/core"
	"LiveDanmu/apps/gateway/live_gateway/core/handle"
	"LiveDanmu/apps/gateway/live_gateway/core/middleware"
	"LiveDanmu/apps/public/config/config_template"
	"LiveDanmu/apps/public/logger/adapter"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/hertz-contrib/monitor-prometheus"
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

func HertzApi(conf *config_template.LiveGatewayConfig) {
	// 构造Url
	url := fmt.Sprintf("%v:%v", conf.Hertz.ListenAddr, conf.Hertz.ListenPort)
	monitorUrl := fmt.Sprintf("%v:%v", conf.Hertz.ListenAddr, conf.Hertz.MonitoringPort)
	// 创建服务核心
	h = server.Default(server.WithHostPorts(url), server.WithTracer(prometheus.NewServerTracer(monitorUrl, "/hertz")))
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
	g := h.Group("/live")
	{
		g.GET("/info", middleware.JWTMiddleware(), handle.GetLiveInfoHandleFunc())
		g.GET("/list", handle.GetLiveListHandleFunc())
		g.GET("/list/my", middleware.JWTMiddleware(), handle.GetMyLiveListHandleFunc())
		g.POST("/start", middleware.JWTMiddleware(), handle.StartLiveHandleFunc()) //
		g.GET("/stop", middleware.JWTMiddleware(), handle.StopLiveHandleFunc())    //
		g.POST("/srs/auth", handle.SRSAuthHandleFunc())
	}
}
