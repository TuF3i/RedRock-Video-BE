package core

import (
	"LiveDanmu/apps/rpc/danmusvr/core/dao"
	"LiveDanmu/apps/rpc/danmusvr/core/kafka"
	"LiveDanmu/apps/shared/logger"
)

var (
	KClient *kafka.KClient
	Logger  *logger.NewLogger
	Dao     *dao.Dao
)
