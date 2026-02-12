package config_reader

import (
	"LiveDanmu/apps/public/config/config_template"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func setDefaultForLiveGateway(v *viper.Viper) {
	v.SetDefault(config_template.LIVE_GATEWAY_HERTZ_LISTENADDR, "0.0.0.0")
	v.SetDefault(config_template.LIVE_GATEWAY_HERTZ_LISTENPORT, "8080")
	v.SetDefault(config_template.LIVE_GATEWAY_HERTZ_MONITORINGPORT, "8081")
	v.SetDefault(config_template.LIVE_GATEWAY_ETCD_SERVICENAME, "etcd")
	v.SetDefault(config_template.LIVE_GATEWAY_ETCD_NAMESPACE, "middleware")
	v.SetDefault(config_template.LIVE_GATEWAY_LOKI_NAMESPACE, "middleware")
	v.SetDefault(config_template.LIVE_GATEWAY_LOKI_SERVICENAME, "loki")
	v.SetDefault(config_template.LIVE_GATEWAY_LOKI_SERVICE, "LIVE_GATEWAY")
	v.SetDefault(config_template.LIVE_GATEWAY_LOKI_ENV, "proc")
	v.SetDefault(config_template.LIVE_GATEWAY_LOKI_LEVEL, "INFO")
	v.SetDefault(config_template.LIVE_GATEWAY_POD_UID, uuid.New().String())
}

func LiveGatewayConfigLoader() (*config_template.LiveGatewayConfig, error) {
	// 初始化结构体指针
	conf := new(config_template.LiveGatewayConfig)
	// 初始化Viper
	v := viper.New()
	//这样环境变量需要以 LIVE_GATEWAY_ 开头，如 LIVE_GATEWAY_HERTZ_LISTENADDR
	v.SetEnvPrefix("LIVE_GATEWAY")
	// 加载默认配置
	setDefaultForLiveGateway(v)
	// 加载环境变量
	v.AutomaticEnv()
	//设置键名转换器（将环境变量中的 _ 映射到结构体的嵌套字段）
	//例如：LIVE_GATEWAY_HERTZ_LISTEN_ADDR -> Hertz.ListenAddr
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 解析配置到结构体
	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
