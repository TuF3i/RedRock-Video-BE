package config_reader

import (
	"LiveDanmu/apps/shared/config/config_template"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func setDefaultForUserRpc(v *viper.Viper) {
	v.SetDefault(config_template.USER_RPC_REGISTRY_HOSTS, "zookeeper:2181")
	v.SetDefault(config_template.USER_RPC_PGSQL_HOST, "pgpool")
	v.SetDefault(config_template.USER_RPC_PGSQL_PORT, "5432")
	v.SetDefault(config_template.USER_RPC_PGSQL_USER, "root")
	v.SetDefault(config_template.USER_RPC_PGSQL_PASSWORD, "")
	v.SetDefault(config_template.USER_RPC_PGSQL_DBNAME, "rvideo")
	v.SetDefault(config_template.USER_RPC_REDIS_HOSTS, "redis-1:6379,redis-2:6379,redis-3:6379")
	v.SetDefault(config_template.USER_RPC_REDIS_PASSWORD, "")
	v.SetDefault(config_template.USER_RPC_POD_UID, uuid.New().String())
	v.SetDefault(config_template.USER_RPC_LOKI_SERVICE, "USER_RPC")
	v.SetDefault(config_template.USER_RPC_LOKI_LEVEL, "INFO")
	v.SetDefault(config_template.USER_RPC_LOKI_ENV, "proc")
	v.SetDefault(config_template.USER_RPC_ADMIN, "")
}

func UserRpcConfigLoader() (*config_template.UserRpcConfig, error) {
	// 初始化结构体指针
	conf := new(config_template.UserRpcConfig)
	// 初始化Viper
	v := viper.New()
	//这样环境变量需要以 USER_RPC_ 开头，如 USER_RPC_ETCD_SERVICENAME
	v.SetEnvPrefix("USER_RPC")
	// 加载默认配置
	setDefaultForUserRpc(v)
	// 加载环境变量
	v.AutomaticEnv()
	//设置键名转换器（将环境变量中的 _ 映射到结构体的嵌套字段）
	//例如：USER_RPC_ETCD_SERVICE_NAME -> Etcd.ServiceName
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 解析配置到结构体
	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
