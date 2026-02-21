package kafka_boardcast

import (
	kafkaCfg "LiveDanmu/apps/public/models/kafka"
	"time"

	"github.com/segmentio/kafka-go"
)

func (b *BoardCast) initKClient() {
	// 连接拨号器
	dialer := &kafka.Transport{
		ClientID:    b.conf.PodUID,
		DialTimeout: 10 * time.Second,
	}
	// 弹幕广播器
	b.kClient = &kafka.Writer{
		Addr:                   kafka.TCP(b.conf.KafKa.Urls...),
		Topic:                  kafkaCfg.LIVE_DANMU_BOARDCAST_TOPIC,
		MaxAttempts:            3, // 重试次数
		BatchSize:              1,
		BatchTimeout:           100 * time.Millisecond, // 超时时间
		RequiredAcks:           1,
		Async:                  false,
		AllowAutoTopicCreation: true,
		Transport:              dialer,
	}
}
