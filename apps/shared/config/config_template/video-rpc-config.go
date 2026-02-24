package config_template

import "LiveDanmu/apps/shared/logger"

type VideoRpcConfig struct {
	ContainerName string
	Redis         RedisForVideoRpc
	PgSQL         PostgresForVideoRpc
	Loki          logger.LoggerConfig
	Registry      RegistryForVideoRpc
	Minio         MinioForVideoRpc
}

type RedisForVideoRpc struct {
	Password string
	Hosts    string
	Urls     []string
}

type RegistryForVideoRpc struct {
	Hosts string
	Urls  []string
}

type MinioForVideoRpc struct {
	Host           string
	UseSSL         bool
	AccessKey      string
	SecretKey      string
	BlanketName    string
	PicBlanketName string
}

type PostgresForVideoRpc struct {
	User     string
	Password string
	DBName   string
	Port     string
	Host     string
}
