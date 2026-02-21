package config_reader

import (
	"LiveDanmu/apps/public/config/config_template"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func setDefaultForVideoDanmuConsumer(v *viper.Viper) {
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_KAFKA_SERVICENAME, "kafka")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_KAFKA_NAMESPACE, "")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_PGSQL_SERVICENAME, "pgpool")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_PGSQL_NAMESPACE, "")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_PGSQL_USER, "root")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_PGSQL_PASSWORD, "eeelcgkklo12l13l17gg")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_PGSQL_DBNAME, "rvideo")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_REDIS_SERVICENAME, "redis")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_REDIS_NAMESPACE, "")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_REDIS_PASSWORD, "eeelcgkklo12l13l17gg")
	v.SetDefault(config_template.VIEDO_DANMU_CONSUMER_GROUPID, "video-danmu-consumer-group-union")
	v.SetDefault(config_template.VIDEO_DANMU_CONSUMER_POD_UID, uuid.New().String())
	v.SetDefault(config_template.VIEDO_DANMU_CONSUMER_LOKI_NAMESPACE, "")
	v.SetDefault(config_template.VIEDO_DANMU_CONSUMER_LOKI_SERVICENAME, "loki")
	v.SetDefault(config_template.VIEDO_DANMU_CONSUMER_LOKI_SERVICE, "VIEDO_DANMU_CONSUMER")
	v.SetDefault(config_template.VIEDO_DANMU_CONSUMER_LOKI_LEVEL, "INFO")
	v.SetDefault(config_template.VIEDO_DANMU_CONSUMER_LOKI_ENV, "proc")
}

func VideoDanmuConsumerConfigLoader() (*config_template.VideoDanmuConsumerConfig, error) {
	// 初始化结构体指针
	conf := new(config_template.VideoDanmuConsumerConfig)
	// 初始化Viper
	v := viper.New()
	//这样环境变量需要以 DANMU_GATEWAY_ 开头，如 DANMU_GATEWAY_HERTZ_LISTENADDR
	v.SetEnvPrefix("VIDEO_DANMU_CONSUMER")
	// 加载默认配置
	setDefaultForVideoDanmuConsumer(v)
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
