package dao

import (
	"LiveDanmu/apps/public/config/config_template"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Dao struct {
	conf          *config_template.VideoRpcConfig
	rdb           *redis.ClusterClient
	pgdb          *gorm.DB
	isSyncRunning atomic.Bool
	userSyncPool  map[int64]*atomic.Bool
}

func GetDao(conf *config_template.VideoRpcConfig) (*Dao, error) {
	d := Dao{conf: conf, isSyncRunning: atomic.Bool{}, userSyncPool: make(map[int64]*atomic.Bool)}
	if err := d.initPgSQL(); err != nil {
		return nil, err
	}
	if err := d.initRedisClient(); err != nil {
		return nil, err
	}

	return &d, nil
}
