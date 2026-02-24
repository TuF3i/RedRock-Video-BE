package config_template

type DBInitConfig struct {
	PgSQL PostgresForDBInit
}

type PostgresForDBInit struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}
