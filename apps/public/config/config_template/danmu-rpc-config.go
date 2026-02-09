package config_template

type DanmuRpcConfig struct {
	PodUID string
	Etcd   EtcdForDanmuRpc
	KafKa  KafkaForDanmuRpc
	PgSQL  PostgresForDanmuRpc
	Redis  RedisForDanmuRpc
	Loki   LokiConfigForDanmuRpc
}

type EtcdForDanmuRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type KafkaForDanmuRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type PostgresForDanmuRpc struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}

type RedisForDanmuRpc struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type LokiConfigForDanmuRpc struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}
