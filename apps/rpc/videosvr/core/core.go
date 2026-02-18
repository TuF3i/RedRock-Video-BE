package core

import (
	"LiveDanmu/apps/public/logger"
	"LiveDanmu/apps/rpc/videosvr/core/dao"
	"LiveDanmu/apps/rpc/videosvr/core/minio"
	"LiveDanmu/apps/rpc/videosvr/kitex_gen/videosvr/videosvr"
)

var (
	Dao      *dao.Dao
	Minio    *minio.Minio
	VideoSvr videosvr.Client
	Logger   *logger.NewLogger
)
