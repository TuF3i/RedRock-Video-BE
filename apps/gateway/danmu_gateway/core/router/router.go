package router

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core"
	"LiveDanmu/apps/gateway/danmu_gateway/core/handler"
	"LiveDanmu/apps/gateway/danmu_gateway/core/middleware"
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

func HertzApi(conf *config_template.DanmuGatewayConfig) {
	// 构造Url
	url := fmt.Sprintf("%v:%v", conf.Hertz.ListenAddr, conf.Hertz.ListenPort)
	// 创建服务核心
	h = server.Default(server.WithHostPorts(url), server.WithTracer(prometheus.NewServerTracer(conf.Hertz.MonitoringPort, "/hertz")))
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
	g := h.Group("/danmu")
	{
		// 发布视频弹幕
		g.POST("/video", middleware.JWTMiddleware(), handler.PubDanmuHandleFunc())
		// 发布直播弹幕
		g.POST("/live", middleware.JWTMiddleware(), handler.PubLiveDanmuHandleFunc())
		// 删除直播弹幕
		g.DELETE("/live", middleware.JWTMiddleware(), handler.DelDanmuHandleFunc())
		// 建立直播实时ws
		g.GET("/live/:rvid", handler.LiveDanmuHandleFunc())
		// 获取首屏视频弹幕
		g.GET("/hot/:rvid", handler.GetHotDanmuHandleFunc())
		// 获取全量视频弹幕
		g.GET("/full/:rvid", handler.GetFullDanmuHandleFunc())
	}

}
