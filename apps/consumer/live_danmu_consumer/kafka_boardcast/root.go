package kafka_boardcast

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/segmentio/kafka-go"
)

type BoardCast struct {
	conf    *config_template.LiveDanmuConsumerConfig
	kClient *kafka.Writer
}

func GetBoardCast(conf *config_template.LiveDanmuConsumerConfig) *BoardCast {
	b := BoardCast{conf: conf}
	b.initKClient()
	return &b
}

func (r *BoardCast) StopBoardCast() error {
	return r.kClient.Close()
}
