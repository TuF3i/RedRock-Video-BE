package config_reader

import (
	"LiveDanmu/apps/shared/config/config_template"
	"strings"

	"github.com/spf13/viper"
)

func setDefaultForDBInit(v *viper.Viper) {
	v.SetDefault(config_template.DB_INIT_PGSQL_HOST, "pgsql")
	v.SetDefault(config_template.DB_INIT_PGSQL_PORT, "5432")
	v.SetDefault(config_template.DB_INIT_PGSQL_USER, "root")
	v.SetDefault(config_template.DB_INIT_PGSQL_PASSWORD, "")
	v.SetDefault(config_template.DB_INIT_PGSQL_DBNAME, "rvideo")
}

func DBInitConfigLoader() (*config_template.DBInitConfig, error) {
	// 初始化结构体指针
	conf := new(config_template.DBInitConfig)
	// 初始化Viper
	v := viper.New()
	//这样环境变量需要以 DANMU_GATEWAY_ 开头，如 DANMU_GATEWAY_HERTZ_LISTENADDR
	v.SetEnvPrefix("DB_INIT")
	// 加载默认配置
	setDefaultForDBInit(v)
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
