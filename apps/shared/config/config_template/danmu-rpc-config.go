package config_template

import "LiveDanmu/apps/shared/logger"

type DanmuRpcConfig struct {
	ContainerName string
	Registry      RegistryForDanmuRpc
	KafKa         KafkaForDanmuRpc
	PgSQL         PostgresForDanmuRpc
	Redis         RedisForDanmuRpc
	Loki          logger.LoggerConfig
}

type RegistryForDanmuRpc struct {
	Urls  []string
	Hosts string
}

type KafkaForDanmuRpc struct {
	Hosts string
	Urls  []string
}

type PostgresForDanmuRpc struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

type RedisForDanmuRpc struct {
	Password string
	Urls     []string
	Hosts    string
}
