package config_template

type VideoDanmuConsumerConfig struct {
	PodUID  string
	GroupID string
	KafKa   KafkaForVideoDanmuConsumer
	PgSQL   PostgresForVideoDanmuConsumer
	Redis   RedisForVideoDanmuConsumer
}

type KafkaForVideoDanmuConsumer struct {
	ServiceName string
	Namespace   string
	Urls        []string
}

type PostgresForVideoDanmuConsumer struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}

type RedisForVideoDanmuConsumer struct {
	Password    string
	ServiceName string
	Namespace   string
	Urls        []string
}
