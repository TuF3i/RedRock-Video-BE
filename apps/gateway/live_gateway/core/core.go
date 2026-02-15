package core

import (
	"LiveDanmu/apps/gateway/live_gateway/core/dao"

	"github.com/bwmarrin/snowflake"
)

var (
	Dao       *dao.Dao
	SnowFlake *snowflake.Node // 雪花
)
