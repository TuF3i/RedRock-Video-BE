package config_template

import "LiveDanmu/apps/public/logger"

type VideoRpcConfig struct {
	PodUID string
	Redis  RedisForVideoRpc
	PgSQL  PostgresForVideoRpc
	Loki   logger.LoggerConfig
	Etcd   EtcdForVideoRpc
	Minio  MinioForVideoRpc
}

type RedisForVideoRpc struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}

type EtcdForVideoRpc struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type MinioForVideoRpc struct {
	ServiceName    string
	Namespace      string
	Urls           []string
	UseSSL         bool
	AccessKey      string
	SecretKey      string
	BlanketName    string
	PicBlanketName string
}

type PostgresForVideoRpc struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}
