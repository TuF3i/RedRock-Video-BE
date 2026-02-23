package core

import (
	"LiveDanmu/apps/shared/logger"
	"LiveDanmu/apps/rpc/usersvr/core/dao"
)

var (
	Dao    *dao.Dao
	Logger *logger.NewLogger
)
