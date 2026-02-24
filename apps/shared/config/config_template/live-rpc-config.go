package config_template

import "LiveDanmu/apps/shared/logger"

type LiveRpcConfig struct {
	ContainerName string
	Registry      RegistryForLiveRpc
	PgSQL         PostgresForLiveRpc
	Redis         RedisForLiveRpc
	Kafka         KafkaForLiveRpc
	Loki          logger.LoggerConfig
}

type RegistryForLiveRpc struct {
	Hosts string
	Urls  []string
}

type PostgresForLiveRpc struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

type RedisForLiveRpc struct {
	Password string
	Hosts    string
	Urls     []string
}

type KafkaForLiveRpc struct {
	Hosts string
	Urls  []string
}
