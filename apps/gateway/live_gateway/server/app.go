package server

import (
	"LiveDanmu/apps/gateway/live_gateway/core"
	dao2 "LiveDanmu/apps/gateway/live_gateway/core/dao"
	"LiveDanmu/apps/gateway/live_gateway/core/router"
	"LiveDanmu/apps/public/config"
	logger2 "LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr/livesvr"
	"hash/fnv"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitee.com/liumou_site/logger"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var l *logger.LocalLogger

func onCreate() {
	l.Modular = "live-gateway-on-create"
	l.Info("Starting LiveGatewayNode...")

	// 初始化配置文件
	conf, err := config.LoadLiveGatewayConfig()
	if err != nil {
		l.Error("Load Configuration Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化etcd
	discovery, err := etcd.NewEtcdResolver(conf.Etcd.Urls)
	if err != nil {
		l.Error("Init Etcd Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化snowflake
	hash := fnv.New64a()
	_, err = hash.Write([]byte(conf.PodUID))
	if err != nil {
		l.Error("Init SnowFlake Error: %v", err.Error())
		os.Exit(1)
	}
	// 取模 1024，确保节点 ID 在 0-1023 范围内
	nodeID := int64(hash.Sum64() % 1024)
	snowFlake, err := snowflake.NewNode(nodeID)
	if err != nil {
		l.Error("Init SnowFlake Error: %v", err.Error())
		os.Exit(1)
	}
	core.SnowFlake = snowFlake

	// 初始化Dao
	dao, err := dao2.GetDao(conf)
	if err != nil {
		l.Error("Init Dao Error: %v", err.Error())
		os.Exit(1)
	}
	core.Dao = dao

	// 初始化LiveSvr
	svr, err := livesvr.NewClient(
		union_var.LIVE_SVR,
		client.WithResolver(discovery),
		client.WithHostPorts(""),
		client.WithRPCTimeout(5*time.Second),
	)
	if err != nil {
		l.Error("Init DanmuSvr Error: %v", err.Error())
		os.Exit(1)
	}
	core.LiveSvr = svr

	// 初始化日志
	llog, err := logger2.GetLogger(conf.Loki)
	if err != nil {
		l.Error("Init LokiLogger Error: %v", err.Error())
		os.Exit(1)
	}
	core.Logger = llog

	// 启动HertzAPI网关
	router.HertzApi(conf)

	l.Info("Starting LiveGatewayNode Successfully!")

}

func onDestroy() {
	l.Modular = "live-gateway-on-destroy"
	l.Info("Shutdown LiveGatewayNode...")

	// 关闭Logger
	err := core.Logger.SyncClean()
	if err != nil {
		l.Error("Stop LokiLogger Error: %v", err.Error())
	}

	l.Info("Shutdown LiveGatewayNode Successfully!")

}

func RunLiveGateway() {
	// 初始化局部日志
	l = logger.NewLogger(1)

	// 初始化监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	onCreate()

	<-quit

	onDestroy()
}
