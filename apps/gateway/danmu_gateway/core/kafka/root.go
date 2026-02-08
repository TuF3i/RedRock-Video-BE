package kafka

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/segmentio/kafka-go"
)

type KClient struct {
	conf     *config_template.DanmuGatewayConfig
	consumer *kafka.Reader
}

func GetKClient(conf *config_template.DanmuGatewayConfig) *KClient {
	k := &KClient{conf: conf}
	k.initKafkaClient()
	return k
}
