package kafka

import (
	"LiveDanmu/apps/shared/union_var"
	"time"

	"github.com/segmentio/kafka-go"
)

func (r *ConsumerGroup) initKClient() {
	// 连接拨号器
	dialer := &kafka.Dialer{
		ClientID: r.conf.ContainerName,
		Timeout:  10 * time.Second,
	}

	r.kClient = kafka.NewReader(kafka.ReaderConfig{
		Brokers: r.conf.KafKa.Urls,
		GroupID: r.conf.GroupID,
		Dialer:  dialer,
		Topic:   union_var.LIVE_DANMU_PUB_TOPIC,

		StartOffset: kafka.LastOffset,
	})
}
