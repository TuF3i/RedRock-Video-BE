package kafka

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/segmentio/kafka-go"
)

type KClient struct {
	liveDanmuWriter  *kafka.Writer
	videoDanmuWriter *kafka.Writer
	utilClient       *kafka.Conn
	conf             *config_template.DanmuRpcConfig
}

func GetKClient(conf *config_template.DanmuRpcConfig) *KClient {
	k := &KClient{conf: conf}
	k.initKafkaClient()
	return k
}
