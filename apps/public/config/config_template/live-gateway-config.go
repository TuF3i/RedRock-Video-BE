package config_template

import "LiveDanmu/apps/public/logger"

type LiveGatewayConfig struct {
	PodUID string
	Hertz  HertzForLiveGateway
	Etcd   EtcdForLiveGateway
	Loki   logger.LokiConfig
	Redis  RedisForLiveGateway
}

type HertzForLiveGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type EtcdForLiveGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type RedisForLiveGateway struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}
