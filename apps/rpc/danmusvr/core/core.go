package core

import (
	"LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/rpc/danmusvr/core/dao"
	"LiveDanmu/apps/rpc/danmusvr/core/kafka"
)

var (
	KClient *kafka.KClient
	Logger  *logger.NewLogger
	Dao     *dao.Dao
)
