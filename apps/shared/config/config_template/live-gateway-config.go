package config_template

import "LiveDanmu/apps/shared/logger"

type LiveGatewayConfig struct {
	ContainerName string
	Hertz         HertzForLiveGateway
	Registry      RegistryForLiveGateway
	Loki          logger.LoggerConfig
	Redis         RedisForLiveGateway
}

type HertzForLiveGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type RegistryForLiveGateway struct {
	Hosts string
	Urls  []string
}

type RedisForLiveGateway struct {
	Password string
	Hosts    string
	Urls     []string
}
