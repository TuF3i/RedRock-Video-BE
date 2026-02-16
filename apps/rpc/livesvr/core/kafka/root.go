package kafka

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/segmentio/kafka-go"
)

type KClient struct {
	boardCastController *kafka.Writer
	conf                *config_template.LiveRpcConfig
}

func GetKClient(conf *config_template.LiveRpcConfig) *KClient {
	k := &KClient{conf: conf}
	k.initKafkaClient()
	return k
}
