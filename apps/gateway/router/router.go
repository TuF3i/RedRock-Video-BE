package router

import (
	"LiveDanmu/apps/gateway"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/hertz-contrib/monitor-prometheus"
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

func HertzApi() {
	// 构造Url
	url := fmt.Sprintf("%v:%v", gateway.Config.Hertz.IPAddr, gateway.Config.Hertz.Port)
	// 创建服务核心
	h = server.Default(server.WithHostPorts(url))
	// 初始化路由
	initRouter(h)
	// 设置日志内核
	hlog.SetLogger(gateway.Logger)
	// 启动Hertz引擎
	go func() { h.Spin() }()
}

func initRouter(h *server.Hertz) {}
