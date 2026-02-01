package core

import (
	"LiveDanmu/apps/rpc/danmu/core/dao"
	"LiveDanmu/apps/rpc/danmu/core/kafka"
)

var (
	KClient *kafka.KClient
	Dao     *dao.Dao
)
