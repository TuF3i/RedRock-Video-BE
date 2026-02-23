package config_template

type DBInitConfig struct {
	PgSQL PostgresForDBInit
}

type PostgresForDBInit struct {
	User        string
	Password    string
	DBName      string
	ServiceName string
	Namespace   string
	Urls        []string
}
