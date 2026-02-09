package config_template

type VideoGatewayConfig struct {
	PodUID string
	Hertz  HertzForVideoGateway
	Redis  RedisForVideoGateway
	Etcd   EtcdForVideoGateway
	Loki   LokiConfigForVideoGateway
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

type LokiConfigForVideoGateway struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}

type EtcdForVideoGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type MinioForVideoGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
	UseSSL      bool
	AccessKey   string
	SecretKey   string
	BlanketName string
}
