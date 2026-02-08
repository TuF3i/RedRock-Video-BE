package config_template

type DanmuGatewayConfig struct {
	PodUID string
	Hertz  HertzForDanmuGateway
	Etcd   EtcdForDanmuGateway
	Redis  RedisForDanmuGateway
	Loki   LokiConfigForDanmuGateway
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

type LokiConfigForDanmuGateway struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"` // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service"`   // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`       // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`     // 日志级别，如debug/info/error
}
