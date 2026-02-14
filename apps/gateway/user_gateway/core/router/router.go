package router

import (
	"LiveDanmu/apps/gateway/user_gateway/core"
	"LiveDanmu/apps/gateway/user_gateway/core/handler"
	"LiveDanmu/apps/gateway/user_gateway/core/middleware"
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
	g := h.Group("/user")
	{
		// OAuth2认证组
		authGroup := g.Group("/auth")
		{
			authGroup.GET("", handler.UserLoginHandleFunc())
			authGroup.GET("/callback", handler.CallbackHandleFunc())
			authGroup.GET("/logout", handler.LogoutHandleFunc())
		}

		// 刷新AccessToken
		g.GET("/refresh", middleware.JWTRefreshMiddleware(), handler.RefreshAccessTokenHandleFunc())

		// 用户信息组
		infoGroup := g.Group("/info", middleware.JWTMiddleware())
		{
			infoGroup.GET("/user", handler.GetUserInfoHandleFunc())
			infoGroup.GET("/users", handler.GetUsersHandleFunc())
			infoGroup.GET("/adminer", handler.GetAdminerHandleFunc())
		}

		// 给用户设置Admin权限
		g.GET("/set/adminer", middleware.JWTRefreshMiddleware(), handler.SetAdminRoleHandleFunc())
	}
}
