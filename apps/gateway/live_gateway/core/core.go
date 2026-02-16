package core

import (
	"LiveDanmu/apps/gateway/live_gateway/core/dao"
	"LiveDanmu/apps/rpc/livesvr/kitex_gen/livesvr/livesvr"

	"github.com/bwmarrin/snowflake"
)

var (
	Dao       *dao.Dao
	SnowFlake *snowflake.Node // 雪花
	LiveSvr   livesvr.Client
)
