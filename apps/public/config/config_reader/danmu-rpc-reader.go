package config_reader

import (
	"LiveDanmu/apps/public/config/config_template"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func setDefaultForDanmuRpc(v *viper.Viper) {
	v.SetDefault(config_template.DANMU_RPC_ETCD_SERVICENAME, "etcd")
	v.SetDefault(config_template.DANMU_RPC_ETCD_NAMESPACE, "")
	v.SetDefault(config_template.DANMU_RPC_KAFKA_SERVICENAME, "kafka")
	v.SetDefault(config_template.DANMU_RPC_KAFKA_NAMESPACE, "")
	v.SetDefault(config_template.DANMU_RPC_PGSQL_SERVICENAME, "pgpool")
	v.SetDefault(config_template.DANMU_RPC_PGSQL_NAMESPACE, "")
	v.SetDefault(config_template.DANMU_RPC_PGSQL_USER, "root")
	v.SetDefault(config_template.DANMU_RPC_PGSQL_PASSWORD, "")
	v.SetDefault(config_template.DANMU_RPC_PGSQL_DBNAME, "rvideo")
	v.SetDefault(config_template.DANMU_RPC_REDIS_SERVICENAME, "redis")
	v.SetDefault(config_template.DANMU_RPC_REDIS_NAMESPACE, "")
	v.SetDefault(config_template.DANMU_RPC_REDIS_PASSWORD, "")
	v.SetDefault(config_template.DANMU_RPC_POD_UID, uuid.New().String())
	v.SetDefault(config_template.DANMU_RPC_LOKI_NAMESPACE, "")
	v.SetDefault(config_template.DANMU_RPC_LOKI_SERVICENAME, "loki")
	v.SetDefault(config_template.DANMU_RPC_LOKI_SERVICE, "DANMU_RPC")
	v.SetDefault(config_template.DANMU_RPC_LOKI_LEVEL, "INFO")
	v.SetDefault(config_template.DANMU_RPC_LOKI_ENV, "proc")
}

func DanmuRpcConfigLoader() (*config_template.DanmuRpcConfig, error) {
	// 初始化结构体指针
	conf := new(config_template.DanmuRpcConfig)
	// 初始化Viper
	v := viper.New()
	//这样环境变量需要以 DANMU_RPC_ 开头，如 DANMU_RPC_HERTZ_LISTENADDR
	v.SetEnvPrefix("DANMU_RPC")
	// 加载默认配置
	setDefaultForDanmuRpc(v)
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
