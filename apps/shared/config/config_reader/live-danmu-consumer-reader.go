package config_reader

import (
	"LiveDanmu/apps/shared/config/config_template"
	"strings"

	"github.com/spf13/viper"
)

func setDefaultForLiveDanmuConsumer(v *viper.Viper) {
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_KAFKA_HOSTS, "kafka:9092")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_PGSQL_HOST, "pgpool")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_PGSQL_PORT, "5432")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_PGSQL_USER, "root")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_PGSQL_PASSWORD, "")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_PGSQL_DBNAME, "rvideo")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_GROUPID, "live-danmu-consumer-group-union")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_CONTAINERNAME, "default-container-name")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_LOKI_SERVICE, "LIVE_DANMU_CONSUMER")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_LOKI_LEVEL, "INFO")
	v.SetDefault(config_template.LIVE_DANMU_CONSUMER_LOKI_ENV, "proc")
}

func LiveDanmuConsumerConfigLoader() (*config_template.LiveDanmuConsumerConfig, error) {
	// 初始化结构体指针
	conf := new(config_template.LiveDanmuConsumerConfig)
	// 初始化Viper
	v := viper.New()
	//这样环境变量需要以 DANMU_GATEWAY_ 开头，如 DANMU_GATEWAY_HERTZ_LISTENADDR
	v.SetEnvPrefix("LIVE_DANMU_CONSUMER")
	// 加载默认配置
	setDefaultForLiveDanmuConsumer(v)
	// 加载环境变量
	v.AutomaticEnv()
	//设置键名转换器（将环境变量中的 _ 映射到结构体的嵌套字段）
	//例如：DANMU_HERTZ_LISTEN_ADDR -> Hertz.ListenAddr
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 解析配置到结构体
	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
