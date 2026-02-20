package dao

import (
	"time"

	"gorm.io/gorm"
)

type VUserInfo struct {
	Uid       int64  `json:"uid" gorm:"-"`
	UserName  string `json:"user_name" gorm:"-"`
	AvatarURL string `json:"avatar_url" gorm:"-"`
}

type VideoInfo struct {
	// gorm段
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	// 数据段
	RVID        int64  `gorm:"column:rvid" json:"rvid"`
	FaceUrl     string `gorm:"column:face_url;size:512" json:"face_url"`
	MinioKey    string `gorm:"column:minio_key;size:512" json:"minio_key"`
	Title       string `gorm:"column:title;size:255" json:"title"`
	Description string `gorm:"column:description;type:text" json:"description"`
	ViewNum     int64  `gorm:"column:view_num;default:0" json:"view_num"`
	// 属性段
	InJudge bool `gorm:"column:in_judge;default:true" json:"-"`
	// 用户段
	UID int64 `gorm:"column:uid" json:"uid"`

	User VUserInfo `json:"user"`
}

func (VideoInfo) TableName() string {
	return "video_info_table"
}
