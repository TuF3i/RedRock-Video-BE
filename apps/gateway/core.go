package gateway

import (
	"LiveDanmu/apps/gateway/models"

	"github.com/bwmarrin/snowflake"
	hertzzap "github.com/hertz-contrib/logger/zap"
)

const (
	TRACE_ID_KEY = "trace_id"
)

var (
	Logger    *hertzzap.Logger // 日志组件
	Config    *models.Config
	SnowFlake *snowflake.Node
)
