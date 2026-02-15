package core

import (
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr/usersvr"
	"LiveDanmu/apps/rpc/videosvr/core/dao"
	"LiveDanmu/apps/rpc/videosvr/core/minio"
)

var (
	Dao     *dao.Dao
	Minio   *minio.Minio
	UserSvr usersvr.Client
)
