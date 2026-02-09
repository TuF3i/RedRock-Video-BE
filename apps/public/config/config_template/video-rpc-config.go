package config_template

type VideoRpcConfig struct {
	PodUID string
	Redis  RedisForVideoRpc
	PgSQL  PostgresForVideoRpc
	Loki   LokiConfigForVideoRpc
	Etcd   EtcdForVideoRpc
	Minio  MinioForVideoRpc
}

type RedisForVideoRpc struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type LokiConfigForVideoRpc struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}

type EtcdForVideoRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type MinioForVideoRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
	UseSSL      bool
	AccessKey   string
	SecretKey   string
	BlanketName string
}

type PostgresForVideoRpc struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}
