package core

import (
	"LiveDanmu/apps/rpc/videosvr/core/dao"
	"LiveDanmu/apps/rpc/videosvr/core/minio"
)

var (
	Dao   *dao.Dao
	Minio *minio.Minio
)
