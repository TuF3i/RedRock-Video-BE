package config_template

import "LiveDanmu/apps/shared/logger"

type UserRpcConfig struct {
	PodUID   string
	AdminId  string
	Registry RegistryForUserRpc
	PgSQL    PostgresForUserRpc
	Redis    RedisForUserRpc
	Loki     logger.LoggerConfig
}

type RegistryForUserRpc struct {
	Hosts string
	Urls  []string
}

type PostgresForUserRpc struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

type RedisForUserRpc struct {
	Password string
	Hosts    string
	Urls     []string
}
