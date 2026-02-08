package config

import (
	"LiveDanmu/apps/public/config/config_reader"
	"LiveDanmu/apps/public/config/config_template"
	"LiveDanmu/apps/public/config/dns_lookup"
)

func LoadDanmuGatewayConfig() (*config_template.DanmuGatewayConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.DanmuGatewayConfigLoader()
	if err != nil {
		return nil, err
	}
	// 服务发现
	addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Etcd.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Loki.LokiAddr = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Redis.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Kafka.ServiceName, conf.Kafka.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Kafka.Urls = addrList

	return conf, nil
}

func LoadDanmuRpcConfig() (*config_template.DanmuRpcConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.DanmuRpcConfigLoader()
	if err != nil {
		return nil, err
	}
	// 服务发现
	addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Etcd.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.KafKa.ServiceName, conf.KafKa.Namespace)
	if err != nil {
		return nil, err
	}
	conf.KafKa.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
	if err != nil {
		return nil, err
	}
	conf.PgSQL.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Redis.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Loki.LokiAddr = addrList

	return conf, nil
}

func LoadLiveDanmuConsumerConfig() (*config_template.LiveDanmuConsumerConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.LiveDanmuConsumerConfigLoader()
	if err != nil {
		return nil, err
	}
	// 服务发现
	addrList, err := dns_lookup.ServiceDiscovery(conf.KafKa.ServiceName, conf.KafKa.Namespace)
	if err != nil {
		return nil, err
	}
	conf.KafKa.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
	if err != nil {
		return nil, err
	}
	conf.PgSQL.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Loki.LokiAddr = addrList

	return conf, nil
}

func LoadVideoDanmuConsumerConfig() (*config_template.VideoDanmuConsumerConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.VideoDanmuConsumerConfigLoader()
	if err != nil {
		return nil, err
	}
	// 服务发现
	addrList, err := dns_lookup.ServiceDiscovery(conf.KafKa.ServiceName, conf.KafKa.Namespace)
	if err != nil {
		return nil, err
	}
	conf.KafKa.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.PgSQL.ServiceName, conf.PgSQL.Namespace)
	if err != nil {
		return nil, err
	}
	conf.PgSQL.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Redis.ServiceName, conf.Redis.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Redis.Urls = addrList

	addrList, err = dns_lookup.ServiceDiscovery(conf.Loki.ServiceName, conf.Loki.Namespace)
	if err != nil {
		return nil, err
	}
	conf.Loki.LokiAddr = addrList

	return conf, nil
}
