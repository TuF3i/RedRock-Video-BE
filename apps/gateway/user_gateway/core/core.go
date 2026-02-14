package core

import (
	OAuth "LiveDanmu/apps/gateway/user_gateway/core/OAuth2"
	"LiveDanmu/apps/gateway/user_gateway/core/dao"
	"LiveDanmu/apps/rpc/usersvr/kitex_gen/usersvr/usersvr"

	"github.com/bwmarrin/snowflake"
)

var (
	Dao       *dao.Dao
	SnowFlake *snowflake.Node // 雪花
	OAuth2    *OAuth.OAuthCore
	UserSvr   usersvr.Client
)
