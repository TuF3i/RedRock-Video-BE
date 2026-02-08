package kafka

import (
	kafka2 "LiveDanmu/apps/public/models/kafka"
	"time"

	"github.com/segmentio/kafka-go"
)

func (r *ConsumerGroup) initKClient() {
	// 连接拨号器
	dialer := &kafka.Dialer{
		ClientID: r.conf.PodUID,
		Timeout:  10 * time.Second,
	}

	r.kClient = kafka.NewReader(kafka.ReaderConfig{
		Brokers: r.conf.KafKa.Urls,
		GroupID: r.conf.GroupID,
		Dialer:  dialer,
		Topic:   kafka2.VIDEO_DANMU_PUB_TOPIC,

		MinBytes:        1,                     // 立即返回，不等待批次填满
		MaxBytes:        10e6,                  // 10MB 上限
		MaxWait:         10 * time.Millisecond, // 最长等待10ms
		ReadLagInterval: -1,                    // 禁用延迟统计，减少开销
	})
}
