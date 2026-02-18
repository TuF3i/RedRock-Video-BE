package server

import (
	"LiveDanmu/apps/gateway/video_gateway/core"
	dao2 "LiveDanmu/apps/gateway/video_gateway/core/dao"
	minio2 "LiveDanmu/apps/gateway/video_gateway/core/minio"
	"LiveDanmu/apps/gateway/video_gateway/core/router"
	"LiveDanmu/apps/public/config"
	logger2 "LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr/videosvr"
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
	l.Modular = "video-gateway-on-create"
	l.Info("Starting VideoGatewayNode...")

	// 初始化配置文件
	conf, err := config.LoadVideoGatewayConfig()
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

	//VideoSvr
	svr, err := videosvr.NewClient(
		union_var.VIDEO_SVR,
		client.WithResolver(discovery),
		client.WithHostPorts(""),
		client.WithRPCTimeout(5*time.Second),
	)
	if err != nil {
		l.Error("Init DanmuSvr Error: %v", err.Error())
		os.Exit(1)
	}
	core.VideoSvr = svr

	// 初始化Minio
	minio, err := minio2.GetMinio(conf)
	if err != nil {
		l.Error("Init Minio Error: %v", err.Error())
		os.Exit(1)
	}
	core.Minio = minio

	// 初始化日志
	llog, err := logger2.GetLogger(conf.Loki)
	if err != nil {
		l.Error("Init LokiLogger Error: %v", err.Error())
		os.Exit(1)
	}
	core.Logger = llog

	// 启动HertzAPI网关
	router.HertzApi(conf)

	l.Info("Starting VideoGatewayNode Successfully!")
}

func onDestroy() {
	l.Modular = "video-gateway-on-destroy"
	l.Info("Shutdown VideoGatewayNode...")

	// 关闭Logger
	err := core.Logger.SyncClean()
	if err != nil {
		l.Error("Stop LokiLogger Error: %v", err.Error())
	}

	l.Info("Shutdown VideoGatewayNode Successfully!")
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
