package config_template

import "LiveDanmu/apps/public/logger"

type DanmuGatewayConfig struct {
	PodUID string
	Hertz  HertzForDanmuGateway
	Etcd   EtcdForDanmuGateway
	Redis  RedisForDanmuGateway
	Loki   logger.LoggerConfig
	Kafka  KafkaForDanmuGateway
}

type HertzForDanmuGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type KafkaForDanmuGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type RedisForDanmuGateway struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type EtcdForDanmuGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
}
