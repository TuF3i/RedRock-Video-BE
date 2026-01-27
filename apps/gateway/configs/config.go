package configs

import (
	"LiveDanmu/apps/gateway/models"
	"strings"

	"github.com/spf13/viper"
)

func ViperConfig() (*models.Config, error) {
	// 初始化配置文件结构体
	conf := new(models.Config)

	// 设置默认值
	viper.SetDefault("hertz.ipaddr", "0.0.0.0")
	viper.SetDefault("hertz.port", "8080")

	// 读环境变量
	viper.SetEnvPrefix("APP") // 环境变量前缀，如 APP_HERTZ_IPADDR
	viper.AutomaticEnv()

	// 处理嵌套结构：将 . 替换为 _，使 HERTZ_IPADDR 能匹配 hertz.ipaddr
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 反序列化
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
