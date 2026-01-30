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

	// 直播弹幕生产者
	r.liveDanmuWriter = &kafka.Writer{
		Addr:                   kafka.TCP(r.conf.KafKa.Urls...),
		Topic:                  kafkaCfg.HOT_DANMU_PUB_TOPIC,
		Balancer:               &RoomPartitioner{},
		MaxAttempts:            1, // 重试次数
		BatchSize:              1,
		BatchTimeout:           1 * time.Millisecond, // 超时时间
		RequiredAcks:           1,
		Async:                  false,
		Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
		Transport:              dialer,
	}

	// 视频弹幕生产者
	r.videoDanmuWriter = &kafka.Writer{
		Addr:                   kafka.TCP(r.conf.KafKa.Urls...),
		Topic:                  kafkaCfg.VIDEO_DANMU_PUB_TOPIC,
		Balancer:               &RoomPartitioner{},
		MaxAttempts:            5, // 重试次数
		BatchSize:              15,
		BatchTimeout:           5 * time.Millisecond, // 超时时间
		RequiredAcks:           1,
		Async:                  false,
		Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
		Transport:              dialer,
	}

	for _, addr := range r.conf.KafKa.Urls { // 故障轮询
		conn, err := kafka.Dial("tcp", addr)
		if err != nil {
			continue
		}
		r.utilClient = conn
	}
}
