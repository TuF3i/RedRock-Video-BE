package config_template

import "LiveDanmu/apps/public/logger"

type LiveDanmuConsumerConfig struct {
	PodUID  string
	GroupID string
	KafKa   KafkaForLiveDanmuConsumer
	PgSQL   PostgresForLiveDanmuConsumer
	Loki    logger.LokiConfig
}

type KafkaForLiveDanmuConsumer struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type PostgresForLiveDanmuConsumer struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}
