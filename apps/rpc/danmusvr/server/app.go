package app

import (
	"LiveDanmu/apps/public/config"
	logger2 "LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/public/logger/adapter"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/danmusvr/core"
	dao2 "LiveDanmu/apps/rpc/danmusvr/core/dao"
	"LiveDanmu/apps/rpc/danmusvr/core/handle"
	"LiveDanmu/apps/rpc/danmusvr/core/kafka"
	"LiveDanmu/apps/rpc/danmusvr/core/middleware"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr/danmusvr"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitee.com/liumou_site/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var l *logger.LocalLogger
var svr server.Server

func onCreate() {
	l.Modular = "danmu-svr-on-create"
	l.Info("Starting DanmuSvrNode...")

	// 初始化配置文件
	conf, err := config.LoadDanmuRpcConfig()
	if err != nil {
		l.Error("Load Configuration Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化etcd
	registry, err := etcd.NewEtcdRegistry(conf.Etcd.Urls, etcd.WithDialTimeoutOpt(5*time.Second))
	if err != nil {
		l.Error("Init Etcd Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化Dao
	dao, err := dao2.GetDao(conf)
	if err != nil {
		l.Error("Init Dao Error: %v", err.Error())
		os.Exit(1)
	}
	core.Dao = dao

	// 初始化kafka
	kp := kafka.GetKClient(conf)
	core.KClient = kp

	// 初始化日志
	llog, err := logger2.GetLogger(conf.Loki)
	if err != nil {
		l.Error("Init LokiLogger Error: %v", err.Error())
		os.Exit(1)
	}
	core.Logger = llog

	// 向注册中心注册服务
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		l.Error("Resolve TCPAddr Error: %v", err.Error())
		os.Exit(1)
	}

	svr = danmusvr.NewServer(
		new(handle.DanmuSvrImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: union_var.DANMU_SVR,
		}),
		server.WithRegistry(registry),
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware.PreInit),
		server.WithMiddleware(middleware.DanmuPoolReleaseMiddleware),
		server.WithLogger(adapter.NewKitexLokiLogger(core.Logger.Logger)),
	)

	// 启动服务
	err = svr.Run()
	if err != nil {
		l.Error("Run DanmuSvr Error: %v", err.Error())
		os.Exit(1)
	}

	l.Info("Starting DanmuSvrNode Successfully!")
}

func onDestroy() {
	l.Modular = "danmu-svr-on-destroy"
	l.Info("Stop DanmuSvrNode...")

	// 停止微服务
	err := svr.Stop()
	if err != nil {
		l.Error("Stop DanmuSvr Error: %v", err.Error())
	}

	// 停止kafka生产者
	err = core.KClient.StopProducer()
	if err != nil {
		l.Error("Stop KProducer Error: %v", err.Error())
	}

	// 关闭Logger
	err = core.Logger.SyncClean()
	if err != nil {
		l.Error("Stop LokiLogger Error: %v", err.Error())
	}

	l.Info("Shutdown DanmuSvrNode Successfully!")
}

func RunDanmuSvr() {
	// 初始化局部日志
	l = logger.NewLogger(1)

	// 初始化监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	onCreate()

	<-quit

	onDestroy()
}
