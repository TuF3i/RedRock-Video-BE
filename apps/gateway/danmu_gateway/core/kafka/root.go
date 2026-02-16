package kafka

import (
	"LiveDanmu/apps/public/config/config_template"
	"context"

	"github.com/segmentio/kafka-go"
)

type KClient struct {
	conf     *config_template.DanmuGatewayConfig
	consumer *kafka.Reader
	ctx      context.Context
	cancel   context.CancelFunc
}

func GetKClient(conf *config_template.DanmuGatewayConfig) *KClient {
	k := &KClient{conf: conf}
	k.initKafkaClient()
	return k
}
