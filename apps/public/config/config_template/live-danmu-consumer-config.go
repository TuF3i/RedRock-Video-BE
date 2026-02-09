package config_template

type LiveDanmuConsumerConfig struct {
	PodUID  string
	GroupID string
	KafKa   KafkaForLiveDanmuConsumer
	PgSQL   PostgresForLiveDanmuConsumer
	Loki    LokiConfigForLiveDanmuConsumer
}

type KafkaForLiveDanmuConsumer struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type PostgresForLiveDanmuConsumer struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}

type LokiConfigForLiveDanmuConsumer struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}
