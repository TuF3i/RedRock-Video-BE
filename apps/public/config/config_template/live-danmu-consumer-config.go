package config_template

type LiveDanmuConsumerConfig struct {
	PodUID  string
	GroupID string
	KafKa   KafkaForLiveDanmuConsumer
	PgSQL   PostgresForLiveDanmuConsumer
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
