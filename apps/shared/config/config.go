package config

import (
	"LiveDanmu/apps/shared/config/config_reader"
	"LiveDanmu/apps/shared/config/config_template"
	"LiveDanmu/apps/shared/config/parse_string"
	"os"
)

func LoadDanmuGatewayConfig() (*config_template.DanmuGatewayConfig, error) {
	conf, err := config_reader.DanmuGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.Kafka.Urls = KAFKA_CLUSTER_ADDRS
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	conf.Kafka.Urls = parse_string.GetAddrs(conf.Kafka.Hosts)

	return conf, nil
}

func LoadDanmuRpcConfig() (*config_template.DanmuRpcConfig, error) {
	conf, err := config_reader.DanmuRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.KafKa.Urls = KAFKA_CLUSTER_ADDRS
		conf.PgSQL.Host = PGSQL_ADDR
		conf.PgSQL.Port = PGSQL_PORT
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.KafKa.Urls = parse_string.GetAddrs(conf.KafKa.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadLiveDanmuConsumerConfig() (*config_template.LiveDanmuConsumerConfig, error) {
	conf, err := config_reader.LiveDanmuConsumerConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.KafKa.Urls = KAFKA_CLUSTER_ADDRS
		conf.PgSQL.Host = PGSQL_ADDR
		conf.PgSQL.Port = PGSQL_PORT
		return conf, nil
	}

	conf.KafKa.Urls = parse_string.GetAddrs(conf.KafKa.Hosts)

	return conf, nil
}

func LoadVideoDanmuConsumerConfig() (*config_template.VideoDanmuConsumerConfig, error) {
	conf, err := config_reader.VideoDanmuConsumerConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.KafKa.Urls = KAFKA_CLUSTER_ADDRS
		conf.PgSQL.Host = PGSQL_ADDR
		conf.PgSQL.Port = PGSQL_PORT
		return conf, nil
	}

	conf.KafKa.Urls = parse_string.GetAddrs(conf.KafKa.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadVideoGatewayConfig() (*config_template.VideoGatewayConfig, error) {
	conf, err := config_reader.VideoGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.Minio.Host = MINIO_HOST
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadVideoRpcConfig() (*config_template.VideoRpcConfig, error) {
	conf, err := config_reader.VideoRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.Minio.Host = MINIO_HOST
		conf.PgSQL.Host = PGSQL_ADDR
		conf.PgSQL.Port = PGSQL_PORT
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadUserGatewayConfig() (*config_template.UserGatewayConfig, error) {
	conf, err := config_reader.UserGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadUserRpcConfig() (*config_template.UserRpcConfig, error) {
	conf, err := config_reader.UserRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.PgSQL.Host = PGSQL_ADDR
		conf.PgSQL.Port = PGSQL_PORT
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadLiveGatewayConfig() (*config_template.LiveGatewayConfig, error) {
	conf, err := config_reader.LiveGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadLiveRpcConfig() (*config_template.LiveRpcConfig, error) {
	conf, err := config_reader.LiveRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Registry.Urls = ZOOKEEPER_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.PgSQL.Host = PGSQL_ADDR
		conf.PgSQL.Port = PGSQL_PORT
		conf.Kafka.Urls = KAFKA_CLUSTER_ADDRS
		return conf, nil
	}

	conf.Registry.Urls = parse_string.GetAddrs(conf.Registry.Hosts)

	conf.Redis.Urls = parse_string.GetAddrs(conf.Redis.Hosts)

	return conf, nil
}

func LoadDBInitConfig() (*config_template.DBInitConfig, error) {
	conf, err := config_reader.DBInitConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.PgSQL.Host = PGSQL_ADDR
		conf.PgSQL.Port = PGSQL_PORT
		return conf, nil
	}

	return conf, nil
}
