package config_template

type UserRpcConfig struct {
	PodUID  string
	AdminId string
	Etcd    EtcdForUserRpc
	PgSQL   PostgresForUserRpc
	Redis   RedisForUserRpc
	Loki    LokiConfigForUserRpc
}

type EtcdForUserRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type PostgresForUserRpc struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}

type RedisForUserRpc struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type LokiConfigForUserRpc struct {
	ServiceName string
	Namespace   string
	LokiAddr    []string `mapstructure:"loki_addr"`        // Loki地址，如http://127.0.0.1:3100
	Service     string   `mapstructure:"service_tag.yaml"` // 服务名，作为Loki标签
	Env         string   `mapstructure:"env"`              // 环境，如dev/test/prod，作为Loki标签
	Level       string   `mapstructure:"level"`            // 日志级别，如debug/info/error
}
