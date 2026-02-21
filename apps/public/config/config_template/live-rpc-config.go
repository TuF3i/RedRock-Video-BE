package config_template

import "LiveDanmu/apps/public/logger"

type LiveRpcConfig struct {
	PodUID string
	Etcd   EtcdForLiveRpc
	PgSQL  PostgresForLiveRpc
	Redis  RedisForLiveRpc
	Kafka  KafkaForLiveRpc
	Loki   logger.LoggerConfig
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

type KafkaForLiveRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}
