package kafka

import (
	"LiveDanmu/apps/consumer/video_danmu_consumer"
	kafka2 "LiveDanmu/apps/public/models/kafka"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (r *ConsumerGroup) initKClient() {
	dialer := &kafka.Dialer{
		ClientID: r.conf.PodUID,
		Timeout:  10 * time.Second,
	}

	r.kClient = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     r.conf.KafKa.Urls,
		GroupID:     r.conf.GroupID,
		Topic:       kafka2.VIDEO_DANMU_PUB_TOPIC,
		Dialer:      dialer,
		StartOffset: kafka.FirstOffset,

		//MinBytes:        1,
		//MaxBytes:        10e6,
		//MaxWait:         10 * time.Millisecond,
		//ReadLagInterval: -1,
	})

	video_danmu_consumer.Logger.INFO("Kafka Reader initialized",
		zap.Strings("brokers", r.conf.KafKa.Urls),
		zap.String("group_id", r.conf.GroupID),
		zap.String("topic", kafka2.VIDEO_DANMU_PUB_TOPIC),
		zap.String("client_id", r.conf.PodUID),
	)
}
