package kafka

import (
	"LiveDanmu/apps/consumer/video_danmu_consumer"
	"LiveDanmu/apps/shared/union_var"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (r *ConsumerGroup) initKClient() {
	dialer := &kafka.Dialer{
		ClientID: r.conf.ContainerName,
		Timeout:  10 * time.Second,
	}

	r.kClient = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     r.conf.KafKa.Urls,
		GroupID:     r.conf.GroupID,
		Topic:       union_var.VIDEO_DANMU_PUB_TOPIC,
		Dialer:      dialer,
		StartOffset: kafka.FirstOffset,
	})

	video_danmu_consumer.Logger.INFO("Kafka Reader initialized",
		zap.Strings("brokers", r.conf.KafKa.Urls),
		zap.String("group_id", r.conf.GroupID),
		zap.String("topic", union_var.VIDEO_DANMU_PUB_TOPIC),
		zap.String("client_id", r.conf.ContainerName),
	)
}
