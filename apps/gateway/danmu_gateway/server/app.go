package server

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core"
	dao2 "LiveDanmu/apps/gateway/danmu_gateway/core/dao"
	"LiveDanmu/apps/gateway/danmu_gateway/core/kafka"
	"LiveDanmu/apps/gateway/danmu_gateway/core/router"
	"LiveDanmu/apps/gateway/danmu_gateway/core/websocket"
	"LiveDanmu/apps/public/config"
	logger2 "LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr/danmusvr"
	"hash/fnv"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitee.com/liumou_site/logger"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-zookeeper/resolver"
)

var k *kafka.KClient
var l *logger.LocalLogger

func onCreate() {
	l.Modular = "danmu-gateway-on-create"
	l.Info("Starting DanmuGatewayNode...")

	// 初始化配置文件
	conf, err := config.LoadDanmuGatewayConfig()
	if err != nil {
		l.Error("Load Configuration Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化etcd
	discovery, err := resolver.NewZookeeperResolver(conf.Etcd.Urls, 10*time.Second)
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

	//初始化DanmuSvr
	svr, err := danmusvr.NewClient(
		union_var.DANMU_SVR,
		client.WithResolver(discovery),
		client.WithRPCTimeout(5*time.Second),
	)
	if err != nil {
		l.Error("Init DanmuSvr Error: %v", err.Error())
		os.Exit(1)
	}
	core.DanmuSvr = svr

	// 初始化WS连接池组
	group := websocket.NewPoolGroup()
	core.PoolGroup = group

	// 启动Kafka监听携程
	kClient := kafka.GetKClient(conf)
	kClient.StartConsume()
	k = kClient

	// 初始化日志
	llog, err := logger2.GetLogger(conf.Loki)
	if err != nil {
		l.Error("Init LokiLogger Error: %v", err.Error())
		os.Exit(1)
	}
	core.Logger = llog

	// 启动HertzAPI网关
	router.HertzApi(conf)

	l.Info("Starting DanmuGatewayNode Successfully!")
}

func onDestroy() {
	l.Modular = "danmu-gateway-on-destroy"
	l.Info("Shutdown DanmuGatewayNode...")
	// 中止消费携程
	err := k.StopConsume()
	if err != nil {
		l.Error("Stop Consumer Error: %v", err.Error())
	}

	// 关闭PoolGroup
	for k, _ := range core.PoolGroup.Pools {
		core.PoolGroup.CancelPool(k)
	}

	// 关闭Logger
	err = core.Logger.SyncClean()
	if err != nil {
		l.Error("Stop LokiLogger Error: %v", err.Error())
	}

	l.Info("Shutdown DanmuGatewayNode Successfully!")
}

func RunDanmuGateway() {
	// 初始化局部日志
	l = logger.NewLogger(1)

	// 初始化监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	onCreate()

	<-quit

	onDestroy()
}
