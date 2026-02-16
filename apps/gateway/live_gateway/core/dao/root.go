package dao

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/redis/go-redis/v9"
)

type Dao struct {
	conf *config_template.LiveGatewayConfig
	rdb  *redis.ClusterClient
}

func GetDao(conf *config_template.LiveGatewayConfig) (*Dao, error) {
	d := Dao{conf: conf}
	if err := d.initRedisClient(); err != nil {
		return nil, err
	}

	return &d, nil
}
