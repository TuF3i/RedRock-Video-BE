package core

import (
	"LiveDanmu/apps/gateway/user_gateway/core/dao"

	"github.com/bwmarrin/snowflake"
)

var (
	Dao       *dao.Dao
	SnowFlake *snowflake.Node // 雪花
)
