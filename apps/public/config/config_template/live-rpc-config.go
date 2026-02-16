package config_template

type LiveRpcConfig struct {
	PodUID string
	Etcd   EtcdForLiveRpc
	PgSQL  PostgresForLiveRpc
	Redis  RedisForLiveRpc
	Kafka  KafkaForLiveRpc
	Loki   LokiConfigForLiveRpc
}

type EtcdForLiveRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type PostgresForLiveRpc struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}

type RedisForLiveRpc struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type LokiConfigForLiveRpc struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}

type KafkaForLiveRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}
