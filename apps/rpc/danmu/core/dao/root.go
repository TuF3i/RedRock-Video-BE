package dao

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Dao struct {
	conf *config_template.DanmuRpcConfig
	rdb  *redis.ClusterClient
	pgdb *gorm.DB
}
