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
		Addr:         kafka.TCP(r.conf.KafKa.Urls...),
		Topic:        kafkaCfg.LIVE_DANMU_PUB_TOPIC,
		Balancer:     &RoomPartitioner{},
		MaxAttempts:  3, // 重试次数
		BatchSize:    1,
		BatchTimeout: 100 * time.Millisecond, // 超时时间
		RequiredAcks: 1,
		Async:        false,
		//Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
		Transport:              dialer,
	}

	// 视频弹幕生产者
	r.videoDanmuWriter = &kafka.Writer{
		Addr:  kafka.TCP(r.conf.KafKa.Urls...),
		Topic: kafkaCfg.VIDEO_DANMU_PUB_TOPIC,
		//Balancer:               &RoomPartitioner{},
		MaxAttempts:  3, // 重试次数
		BatchSize:    1,
		BatchTimeout: 5 * time.Millisecond, // 超时时间
		RequiredAcks: 1,
		Async:        false,
		//Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
		Transport:              dialer,
		WriteTimeout:           5 * time.Second, // 新增写超时，避免快速失败
		ReadTimeout:            5 * time.Second, // 新增读超时
	}

	// 集群控制器
	for _, addr := range r.conf.KafKa.Urls { // 故障轮询
		conn, err := kafka.Dial("tcp", addr)
		if err != nil {
			continue
		}
		r.utilClient = conn
	}

	// 网关广播控制器（删弹幕用的）
	r.boardCastController = &kafka.Writer{
		Addr:         kafka.TCP(r.conf.KafKa.Urls...),
		Topic:        kafkaCfg.LIVE_DANMU_BOARDCAST_TOPIC,
		MaxAttempts:  3, // 重试次数
		BatchSize:    1,
		BatchTimeout: 1 * time.Millisecond, // 超时时间
		RequiredAcks: 1,
		Async:        false,
		//Compression:            kafka.Snappy,
		AllowAutoTopicCreation: true,
		Transport:              dialer,
	}
}
