package kafka

import (
	"LiveDanmu/apps/consumer/video_danmu_consumer/dao"
	"LiveDanmu/apps/public/config/config_template"
	"context"

	"github.com/segmentio/kafka-go"
)

type ConsumerGroup struct {
	conf    *config_template.VideoDanmuConsumerConfig
	ctx     context.Context
	cancel  context.CancelFunc
	dao     *dao.Dao
	kClient *kafka.Reader
}

func GetConsumerGroup(conf *config_template.VideoDanmuConsumerConfig, dao *dao.Dao) *ConsumerGroup {
	c := ConsumerGroup{conf: conf, dao: dao}
	c.initKClient()
	return &c
}
