package config_template

type LiveGatewayConfig struct {
	PodUID string
	Hertz  HertzForLiveGateway
	Etcd   EtcdForLiveGateway
	Loki   LokiConfigForLiveGateway
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

type LokiConfigForLiveGateway struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}
