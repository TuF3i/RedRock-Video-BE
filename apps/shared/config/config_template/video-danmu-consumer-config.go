package config_template

import "LiveDanmu/apps/shared/logger"

type VideoDanmuConsumerConfig struct {
	PodUID  string
	GroupID string
	KafKa   KafkaForVideoDanmuConsumer
	PgSQL   PostgresForVideoDanmuConsumer
	Redis   RedisForVideoDanmuConsumer
	Loki    logger.LoggerConfig
}

type KafkaForVideoDanmuConsumer struct {
	Hosts string
	Urls  []string
}

type PostgresForVideoDanmuConsumer struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

type RedisForVideoDanmuConsumer struct {
	Password string
	Hosts    string
	Urls     []string
}
