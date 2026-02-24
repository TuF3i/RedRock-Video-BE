package config_template

import "LiveDanmu/apps/shared/logger"

type UserGatewayConfig struct {
	PodUID   string
	Hertz    HertzForUserGateway
	Redis    RedisForUserGateway
	Registry RegistryForUserGateway
	Loki     logger.LoggerConfig
	Oauth    OAuthConfForUserGateway
}

type HertzForUserGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type RedisForUserGateway struct {
	Password string
	Hosts    string
	Urls     []string
}

type RegistryForUserGateway struct {
	Hosts string
	Urls  []string
}

type OAuthConfForUserGateway struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}
