package dao

import (
	"time"

	"gorm.io/gorm"
)

type DanmuData struct {
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	RVID    int64  `json:"rv_id" gorm:"column:room_id;index:idx_room_time,priority:1;not null;comment:直播间ID"`
	UserId  int64  `json:"user_id" gorm:"column:user_id;index;not null;comment:发送者用户ID"`
	Content string `json:"content" binding:"required,min=1,max=100" gorm:"column:content;type:varchar(500);not null;comment:弹幕内容"`
	Color   string `json:"color" binding:"omitempty,hexcolor" gorm:"column:color;size:20;default:'#FFFFFF';comment:弹幕颜色"`
	Ts      int64  `json:"ts" gorm:"column:ts;index:idx_room_time,priority:2;not null;comment:发送时间戳(ms)"`
}
