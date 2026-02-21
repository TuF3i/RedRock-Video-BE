package config_template

import "LiveDanmu/apps/public/logger"

type VideoGatewayConfig struct {
	PodUID string
	Hertz  HertzForVideoGateway
	Redis  RedisForVideoGateway
	Etcd   EtcdForVideoGateway
	Loki   logger.LoggerConfig
	Minio  MinioForVideoGateway
}

type HertzForVideoGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type RedisForVideoGateway struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type EtcdForVideoGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type MinioForVideoGateway struct {
	ServiceName    string
	Namespace      string
	Urls           []string
	UseSSL         bool
	AccessKey      string
	SecretKey      string
	BlanketName    string
	PicBlanketName string
}
