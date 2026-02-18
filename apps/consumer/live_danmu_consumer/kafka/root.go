package kafka

import (
	"LiveDanmu/apps/consumer/live_danmu_consumer/dao"
	"LiveDanmu/apps/consumer/live_danmu_consumer/kafka_boardcast"
	"LiveDanmu/apps/public/config/config_template"
	"context"

	"github.com/segmentio/kafka-go"
)

type ConsumerGroup struct {
	conf       *config_template.LiveDanmuConsumerConfig
	dao        *dao.Dao
	ctx        context.Context
	cancel     context.CancelFunc
	kClient    *kafka.Reader
	kBoardCast *kafka_boardcast.BoardCast
}

func GetConsumerGroup(conf *config_template.LiveDanmuConsumerConfig, dao *dao.Dao, boardCast *kafka_boardcast.BoardCast) *ConsumerGroup {
	c := ConsumerGroup{conf: conf, dao: dao, kBoardCast: boardCast}
	c.initKClient()
	return &c
}
