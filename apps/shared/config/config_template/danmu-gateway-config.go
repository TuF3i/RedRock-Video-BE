package config_template

import "LiveDanmu/apps/shared/logger"

type DanmuGatewayConfig struct {
	ContainerName string
	Hertz         HertzForDanmuGateway
	Registry      RegistryForDanmuGateway
	Redis         RedisForDanmuGateway
	Loki          logger.LoggerConfig
	Kafka         KafkaForDanmuGateway
}

type HertzForDanmuGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type KafkaForDanmuGateway struct {
	Urls  []string
	Hosts string
}

type RedisForDanmuGateway struct {
	Password string
	Urls     []string
	Hosts    string
}

type RegistryForDanmuGateway struct {
	Urls  []string
	Hosts string
}
