package dao

import (
	"time"

	"gorm.io/gorm"
)

type VideoInfo struct {
	// gorm段
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	// 数据段
	RVID        int64  `gorm:"column:rvid;primaryKey;autoIncrement" json:"rvid"`
	FaceUrl     string `gorm:"column:face_url;size:512" json:"face_url"`
	MinioKey    string `gorm:"column:minio_key;size:512" json:"minio_key"`
	Title       string `gorm:"column:title;size:255" json:"title"`
	Description string `gorm:"column:description;type:text" json:"description"`
	ViewNum     int64  `gorm:"column:view_num;default:0" json:"view_num"`
	// 属性段
	UseFace bool `gorm:"column:use_face;default:false" json:"use_face"`
	InJudge bool `gorm:"column:in_judge;default:true" json:"in_judge"`
}
