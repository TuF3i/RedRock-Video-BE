package app

import (
	"LiveDanmu/apps/public/config"
	logger2 "LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/public/logger/adapter"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/usersvr/core"
	dao2 "LiveDanmu/apps/rpc/usersvr/core/dao"
	"LiveDanmu/apps/rpc/usersvr/core/handle"
	"LiveDanmu/apps/rpc/usersvr/core/middleware"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr/usersvr"
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

//func main() {
//	svr := usersvr.NewServer(new(handle.UserSvrImpl))
//
//	err := svr.Run()
//
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

var l *logger.LocalLogger
var svr server.Server

func onCreate() {
	l.Modular = "user-svr-on-create"
	l.Info("Starting UserSvrNode...")

	// 初始化配置文件
	conf, err := config.LoadUserRpcConfig()
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

	svr = usersvr.NewServer(
		new(handle.UserSvrImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: union_var.USER_SVR,
		}),
		server.WithRegistry(registry),
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware.PreInit),
		server.WithLogger(adapter.NewKitexLokiLogger(core.Logger.Logger)),
	)

	// 启动服务
	err = svr.Run()
	if err != nil {
		l.Error("Run DanmuSvr Error: %v", err.Error())
		os.Exit(1)
	}

	l.Info("Starting UserSvrNode Successfully!")
}

func onDestroy() {
	l.Modular = "user-svr-on-create"
	l.Info("Stop UserSvrNode...")

	// 停止微服务
	err := svr.Stop()
	if err != nil {
		l.Error("Stop DanmuSvr Error: %v", err.Error())
	}

	// 关闭Logger
	err = core.Logger.SyncClean()
	if err != nil {
		l.Error("Stop LokiLogger Error: %v", err.Error())
	}

	l.Info("Shutdown UserSvrNode Successfully!")
}

func RunUserSvr() {
	// 初始化局部日志
	l = logger.NewLogger(1)

	// 初始化监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	onCreate()

	<-quit

	onDestroy()
}
