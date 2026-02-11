package core

import (
	"LiveDanmu/apps/gateway/video_gateway/core/dao"
	"LiveDanmu/apps/gateway/video_gateway/core/minio"
	"LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr/videosvr"

	"github.com/bwmarrin/snowflake"
)

var (
	Logger    *logger.NewLogger
	SnowFlake *snowflake.Node // 雪花
	Dao       *dao.Dao
	Minio     *minio.Minio
	VideoSvr  videosvr.Client
)
