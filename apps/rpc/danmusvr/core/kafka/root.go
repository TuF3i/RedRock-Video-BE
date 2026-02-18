package kafka

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/segmentio/kafka-go"
)

type KClient struct {
	liveDanmuWriter     *kafka.Writer
	videoDanmuWriter    *kafka.Writer
	boardCastController *kafka.Writer
	utilClient          *kafka.Conn
	conf                *config_template.DanmuRpcConfig
}

func GetKClient(conf *config_template.DanmuRpcConfig) *KClient {
	k := &KClient{conf: conf}
	k.initKafkaClient()
	return k
}

func (r *KClient) StopProducer() error {
	if err := r.liveDanmuWriter.Close(); err != nil {
		return err
	}

	if err := r.videoDanmuWriter.Close(); err != nil {
		return err
	}

	if err := r.boardCastController.Close(); err != nil {
		return err
	}

	if err := r.utilClient.Close(); err != nil {
		return err
	}

	return nil
}
