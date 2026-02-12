package config_template

type UserGatewayConfig struct {
	PodUID string
	Hertz  HertzForUserGateway
	Redis  RedisForUserGateway
	Etcd   EtcdForUserGateway
	Loki   LokiConfigForUserGateway
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

type LokiConfigForUserGateway struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}

type OAuthConfForUserGateway struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}
