package dao

import (
	"time"

	"gorm.io/gorm"
)

type LUserInfo struct {
	Uid       int64  `json:"uid" gorm:"-"`
	UserName  string `json:"user_name" gorm:"-"`
	AvatarURL string `json:"avatar_url" gorm:"-"`
}

type LiveInfo struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index;comment:删除时间（软删除）"`

	RVID             int64  `gorm:"column:rv_id;type:bigint;not null;index:idx_rv_id;comment:直播间唯一ID"`
	OwerId           int64  `gorm:"column:ower_id;type:bigint;not null;index:idx_ower_id;comment:主播用户ID"`
	Title            string `gorm:"column:title;type:varchar(64);not null;index:idx_title;comment:直播名称"`
	StreamName       string `gorm:"column:stream_name;type:varchar(64);not null;uniqueIndex:uk_stream_name;comment:推流名称（如 live/test123）"`
	UpstreamPassword string `gorm:"column:upstream_password;type:varchar(64);not null;comment:推流密码（建议加密存储）"`

	User LUserInfo `json:"user"`
}

// TableName 显式指定表名（可选，若结构体名复数形式与表名不一致时推荐）
func (l LiveInfo) TableName() string {
	return "live_info_table"
}
