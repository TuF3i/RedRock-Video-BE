package kafka

import (
	kafkaCfg "LiveDanmu/apps/public/models/kafka"
	"time"

	"github.com/segmentio/kafka-go"
)

func (r *KClient) initKafkaClient() {
	// 连接拨号器
	dialer := &kafka.Transport{
		ClientID:    r.conf.PodUID,
		DialTimeout: 10 * time.Second,
	}

	// 网关广播控制器（删弹幕用的）
	r.boardCastController = &kafka.Writer{
		Addr:                   kafka.TCP(r.conf.Kafka.Urls...),
		Topic:                  kafkaCfg.LIVE_DANMU_BOARDCAST_TOPIC,
		MaxAttempts:            1, // 重试次数
		BatchSize:              1,
		BatchTimeout:           1 * time.Millisecond, // 超时时间
		RequiredAcks:           1,
		Async:                  false,
		Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
		Transport:              dialer,
	}
}
