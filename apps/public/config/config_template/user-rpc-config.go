package config_template

import "LiveDanmu/apps/public/logger"

type UserRpcConfig struct {
	PodUID  string
	AdminId string
	Etcd    EtcdForUserRpc
	PgSQL   PostgresForUserRpc
	Redis   RedisForUserRpc
	Loki    logger.LokiConfig
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
