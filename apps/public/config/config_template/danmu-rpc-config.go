package config_template

type DanmuRpcConfig struct {
	PodUID string
	Etcd   EtcdForDanmuRpc
	KafKa  KafkaForDanmuRpc
	PgSQL  PostgresForDanmuRpc
	Redis  RedisForDanmuRpc
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
