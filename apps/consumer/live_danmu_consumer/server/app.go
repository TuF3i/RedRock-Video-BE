package server

import (
	"LiveDanmu/apps/consumer/live_danmu_consumer"
	dao2 "LiveDanmu/apps/consumer/live_danmu_consumer/dao"
	"LiveDanmu/apps/consumer/live_danmu_consumer/kafka"
	"LiveDanmu/apps/consumer/live_danmu_consumer/kafka_boardcast"
	"LiveDanmu/apps/public/config"
	logger2 "LiveDanmu/apps/public/logger"
	"os"
	"os/signal"
	"syscall"

	"gitee.com/liumou_site/logger"
)

var l *logger.LocalLogger
var kc *kafka.ConsumerGroup
var kb *kafka_boardcast.BoardCast

func onCreate() {
	l.Modular = "live-danmu-consumer-on-create"
	l.Info("Starting LiveDanmuConsumerNode...")

	// 初始化配置文件
	conf, err := config.LoadLiveDanmuConsumerConfig()
	if err != nil {
		l.Error("Load Configuration Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化日志
	llog, err := logger2.GetLogger(conf.Loki)
	if err != nil {
		l.Error("Init LokiLogger Error: %v", err.Error())
		os.Exit(1)
	}
	live_danmu_consumer.Logger = llog

	// 初始化dao
	dao, err := dao2.GetDao(conf)
	if err != nil {
		l.Error("Init Dao Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化kafka广播
	kb = kafka_boardcast.GetBoardCast(conf)

	// 初始化消费者组
	kc = kafka.GetConsumerGroup(conf, dao, kb)

	// 启动消费者组
	kc.StartConsume()

	l.Info("Starting LiveDanmuConsumerNode Successfully!")
}

func onDestroy() {
	l.Modular = "live-danmu-consumer-on-destroy"
	l.Info("Stop LiveDanmuConsumerNode...")

	// 关闭消费者携程
	err := kc.StopConsume()
	if err != nil {
		l.Error("Stop Consumer Error: %v", err.Error())
	}

	// 关闭广播
	err = kb.StopBoardCast()
	if err != nil {
		l.Error("Stop BoardCast Error: %v", err.Error())
	}

	// 关闭Logger
	err = live_danmu_consumer.Logger.SyncClean()
	if err != nil {
		l.Error("Stop LokiLogger Error: %v", err.Error())
	}

	l.Info("Shutdown LiveDanmuConsumerNode Successfully!")
}

func RunLiveDanmuConsumer() {
	// 初始化局部日志
	l = logger.NewLogger(1)

	// 初始化监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	onCreate()

	<-quit

	onDestroy()
}
