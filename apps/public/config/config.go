package config

import (
	"LiveDanmu/apps/public/config/config_reader"
	"LiveDanmu/apps/public/config/config_template"
	"LiveDanmu/apps/public/config/dns_lookup"
	"os"
)

func LoadDanmuGatewayConfig() (*config_template.DanmuGatewayConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.DanmuGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.Kafka.Urls = KAFKA_CLUSTER_ADDRS

		return conf, nil
	}

	// 服务发现
	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Kafka.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Kafka.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Kafka.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Kafka.ServiceName, conf.Kafka.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Kafka.Urls = addrList
	}

	return conf, nil
}

func LoadDanmuRpcConfig() (*config_template.DanmuRpcConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.DanmuRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.KafKa.Urls = KAFKA_CLUSTER_ADDRS
		conf.PgSQL.Urls = []string{"127.0.0.1:5432"}

		return conf, nil
	}

	// 服务发现
	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.KafKa.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.KafKa.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.KafKa.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.KafKa.ServiceName, conf.KafKa.Namespace)
		if err != nil {
			return nil, err
		}
		conf.KafKa.Urls = addrList
	}

	if conf.PgSQL.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.PgSQL.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}

func LoadLiveDanmuConsumerConfig() (*config_template.LiveDanmuConsumerConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.LiveDanmuConsumerConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.KafKa.Urls = KAFKA_CLUSTER_ADDRS
		conf.PgSQL.Urls = []string{"127.0.0.1:5432"}

		return conf, nil
	}

	// 服务发现
	if conf.KafKa.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.KafKa.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.KafKa.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.KafKa.ServiceName, conf.KafKa.Namespace)
		if err != nil {
			return nil, err
		}
		conf.KafKa.Urls = addrList
	}

	if conf.PgSQL.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.PgSQL.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}

func LoadVideoDanmuConsumerConfig() (*config_template.VideoDanmuConsumerConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.VideoDanmuConsumerConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.KafKa.Urls = KAFKA_CLUSTER_ADDRS
		conf.PgSQL.Urls = []string{"127.0.0.1:5432"}

		return conf, nil
	}

	// 服务发现
	if conf.KafKa.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.KafKa.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.KafKa.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.KafKa.ServiceName, conf.KafKa.Namespace)
		if err != nil {
			return nil, err
		}
		conf.KafKa.Urls = addrList
	}

	if conf.PgSQL.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.PgSQL.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}

func LoadVideoGatewayConfig() (*config_template.VideoGatewayConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.VideoGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.Minio.Urls = MINIO_ADDRS

		return conf, nil
	}

	// 服务发现
	if conf.Minio.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Minio.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Minio.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Minio.ServiceName, conf.Minio.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Minio.Urls = addrList
	}

	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}

func LoadVideoRpcConfig() (*config_template.VideoRpcConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.VideoRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.Minio.Urls = MINIO_ADDRS
		conf.PgSQL.Urls = []string{"127.0.0.1:5432"}

		return conf, nil
	}

	// 服务发现
	if conf.Minio.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Minio.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Minio.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Minio.ServiceName, conf.Minio.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Minio.Urls = addrList
	}

	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	if conf.PgSQL.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.PgSQL.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	}

	return conf, nil
}

func LoadUserGatewayConfig() (*config_template.UserGatewayConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.UserGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS

		return conf, nil
	}

	// 服务发现
	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}

func LoadUserRpcConfig() (*config_template.UserRpcConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.UserRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.PgSQL.Urls = []string{"127.0.0.1:5432"}

		return conf, nil
	}

	// 服务发现
	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.PgSQL.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.PgSQL.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}

func LoadLiveGatewayConfig() (*config_template.LiveGatewayConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.LiveGatewayConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS

		return conf, nil
	}

	// 服务发现
	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}

func LoadLiveRpcConfig() (*config_template.LiveRpcConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.LiveRpcConfigLoader()
	if err != nil {
		return nil, err
	}

	if os.Getenv("RV_DEBUG") != "" {
		conf.Etcd.Urls = ETCD_ADDRS
		conf.Loki.LokiAddr = LOKI_ADDRS
		conf.Redis.Urls = REDIS_CLUSTER_ADDRS
		conf.PgSQL.Urls = []string{"127.0.0.1:5432"}

		return conf, nil
	}

	// 服务发现
	if conf.Etcd.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Etcd.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Etcd.Urls = addrList
	}

	if conf.PgSQL.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.PgSQL.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
		if err != nil {
			return nil, err
		}
		conf.PgSQL.Urls = addrList
	}

	if conf.Redis.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Redis.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Redis.Urls = addrList
	}

	if conf.Loki.Namespace == "" {
		addrList, err := dns_lookup.ServiceDiscoveryOverDocker(conf.Loki.ServiceName)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	} else {
		addrList, err := dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
		if err != nil {
			return nil, err
		}
		conf.Loki.LokiAddr = addrList
	}

	return conf, nil
}
