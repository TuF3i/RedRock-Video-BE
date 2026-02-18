package core

import (
	"LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/rpc/livesvr/core/dao"
	"LiveDanmu/apps/rpc/livesvr/core/kafka"
)

var (
	Dao    *dao.Dao
	Kafka  *kafka.KClient
	Logger *logger.NewLogger
)
