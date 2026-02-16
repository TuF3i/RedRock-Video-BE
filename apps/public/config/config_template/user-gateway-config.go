package config_template

import "LiveDanmu/apps/public/logger"

type UserGatewayConfig struct {
	PodUID string
	Hertz  HertzForUserGateway
	Redis  RedisForUserGateway
	Etcd   EtcdForUserGateway
	Loki   logger.LokiConfig
	OAuth  OAuthConfForUserGateway
}

type HertzForUserGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type RedisForUserGateway struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type EtcdForUserGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type OAuthConfForUserGateway struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}
