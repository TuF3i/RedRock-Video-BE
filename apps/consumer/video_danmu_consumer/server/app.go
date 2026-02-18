package server

import (
	"LiveDanmu/apps/consumer/video_danmu_consumer"
	dao2 "LiveDanmu/apps/consumer/video_danmu_consumer/dao"
	"LiveDanmu/apps/consumer/video_danmu_consumer/kafka"
	"LiveDanmu/apps/public/config"
	logger2 "LiveDanmu/apps/public/logger"
	"os"
	"os/signal"
	"syscall"

	"gitee.com/liumou_site/logger"
)

var kc *kafka.ConsumerGroup
var l *logger.LocalLogger

func onCreate() {
	l.Modular = "video-danmu-consumer-on-create"
	l.Info("Starting VideoDanmuConsumerNode...")

	// 初始化配置文件
	conf, err := config.LoadVideoDanmuConsumerConfig()
	if err != nil {
		l.Error("Load Configuration Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化Dao
	dao, err := dao2.GetDao(conf)
	if err != nil {
		l.Error("Init Dao Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化日志
	llog, err := logger2.GetLogger(conf.Loki)
	if err != nil {
		l.Error("Init LokiLogger Error: %v", err.Error())
		os.Exit(1)
	}
	video_danmu_consumer.Logger = llog

	// 初始化消费者组
	kc = kafka.GetConsumerGroup(conf, dao)

	// 启动消费者携程
	kc.StartConsume()

	l.Info("Starting VideoDanmuConsumerNode Successfully!")
}

func onDestroy() {
	l.Modular = "video-danmu-consumer-on-destroy"
	l.Info("Stop VideoDanmuConsumerNode...")

	// 停止消费者携程
	err := kc.StopConsume()
	if err != nil {
		l.Error("Stop Consumer Error: %v", err.Error())
	}

	// 关闭Logger
	err = video_danmu_consumer.Logger.SyncClean()
	if err != nil {
		l.Error("Stop LokiLogger Error: %v", err.Error())
	}

	l.Info("Stop VideoDanmuConsumerNode Successfully!")
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
