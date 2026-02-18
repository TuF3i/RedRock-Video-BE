package core

import (
	"LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/rpc/usersvr/core/dao"
)

var (
	Dao    *dao.Dao
	Logger *logger.NewLogger
)
