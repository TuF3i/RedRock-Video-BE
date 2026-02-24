package config_template

import "LiveDanmu/apps/shared/logger"

type VideoGatewayConfig struct {
	ContainerName string
	Hertz         HertzForVideoGateway
	Redis         RedisForVideoGateway
	Registry      RegistryForVideoGateway
	Loki          logger.LoggerConfig
	Minio         MinioForVideoGateway
}

type HertzForVideoGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type RedisForVideoGateway struct {
	Password string
	Hosts    string
	Urls     []string
}

type RegistryForVideoGateway struct {
	Hosts string
	Urls  []string
}

type MinioForVideoGateway struct {
	Host           string
	UseSSL         bool
	AccessKey      string
	SecretKey      string
	BlanketName    string
	PicBlanketName string
}
