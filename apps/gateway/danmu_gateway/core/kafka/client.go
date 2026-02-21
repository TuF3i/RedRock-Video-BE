package kafka

import (
	kafka2 "LiveDanmu/apps/public/models/kafka"
	"time"

	"github.com/segmentio/kafka-go"
)

func (r *KClient) initKafkaClient() {
	// 连接拨号器
	dialer := &kafka.Dialer{
		ClientID: r.conf.PodUID,
		Timeout:  10 * time.Second,
	}

	r.consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     r.conf.Kafka.Urls,
		GroupID:     r.conf.PodUID, // 使用不同的GroupID已达到广播的效果
		Dialer:      dialer,
		Topic:       kafka2.LIVE_DANMU_BOARDCAST_TOPIC,
		StartOffset: kafka.LastOffset,
	})
}
