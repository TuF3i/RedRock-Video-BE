package dao

import (
	"time"

	"gorm.io/gorm"
)

type DUserInfo struct {
	Uid       int64  `json:"uid" gorm:"-"`
	UserName  string `json:"user_name" gorm:"-"`
	AvatarURL string `json:"avatar_url" gorm:"-"`
}

type DanmuData struct {
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	DanID   int64  `json:"dan_id" gorm:"column:dan_id;index:idx_dan_id,priority:1;not null;comment:弹幕ID"`
	RVID    int64  `json:"rv_id" gorm:"column:rvid;index:idx_rvid,priority:1;not null;comment:RVID"`
	UserId  int64  `json:"user_id" gorm:"column:user_id;index;not null;comment:发送者用户ID"`
	Content string `json:"content" binding:"required,min=1,max=100" gorm:"column:content;type:varchar(500);not null;comment:弹幕内容"`
	Color   string `json:"color" binding:"omitempty,hexcolor" gorm:"column:color;size:20;default:'#FFFFFF';comment:弹幕颜色"`
	Ts      int64  `json:"ts" gorm:"column:ts;index:idx_room_time,priority:2;not null;comment:发送时间戳(ms)"`

	User DUserInfo `json:"user"`
}

func (DanmuData) TableName() string {
	return "danmu_data_table"
}
