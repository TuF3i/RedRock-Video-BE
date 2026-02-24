package kafka

import (
	"LiveDanmu/apps/shared/union_var"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func getGroupID(ContainerName string) string {
	if ContainerName == "default-container-name" {
		return uuid.New().String()
	}

	return ContainerName
}

func (r *KClient) initKafkaClient() {
	// 连接拨号器
	dialer := &kafka.Dialer{
		ClientID: r.conf.ContainerName,
		Timeout:  10 * time.Second,
	}

	r.consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     r.conf.Kafka.Urls,
		GroupID:     getGroupID(r.conf.ContainerName), // 使用不同的GroupID已达到广播的效果
		Dialer:      dialer,
		Topic:       union_var.LIVE_DANMU_BOARDCAST_TOPIC,
		StartOffset: kafka.LastOffset,
	})
}
