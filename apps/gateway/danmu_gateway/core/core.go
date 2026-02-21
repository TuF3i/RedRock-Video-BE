package core

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core/dao"
	"LiveDanmu/apps/gateway/danmu_gateway/core/websocket"
	"LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr/danmusvr"

	"github.com/bwmarrin/snowflake"
)

var (
	Dao       *dao.Dao
	SnowFlake *snowflake.Node // 雪花
	DanmuSvr  danmusvr.Client
	PoolGroup *websocket.WebSocketManager
	Logger    *logger.NewLogger
)
