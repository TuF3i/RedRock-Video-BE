package dao

import (
	"time"

	"gorm.io/gorm"
)

type GitHubUser struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`      // GitHub用户名
	Name      string `json:"name"`       // 昵称
	AvatarURL string `json:"avatar_url"` // 头像地址
	Bio       string `json:"bio"`        // 个人简介
}

type RvUser struct {
	// 通用基础字段（表主键+时间+软删除）
	ID        int64          `gorm:"primaryKey;type:bigint;autoIncrement;comment:表自增主键ID"`
	CreatedAt time.Time      `gorm:"autoCreateTime;type:datetime;comment:记录创建时间"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;type:datetime;comment:记录更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:datetime;comment:软删除标记"`

	// GitHub用户相关字段
	Uid       int64  `json:"uid" gorm:"column:github_uid;type:bigint;not null;uniqueIndex;comment:GitHub用户唯一ID"`
	Login     string `json:"login" gorm:"column:github_login;type:varchar(64);not null;comment:GitHub用户名"`
	AvatarURL string `json:"avatar_url" gorm:"column:avatar_url;type:varchar(255);comment:GitHub用户头像地址"`
	Bio       string `json:"bio" gorm:"column:bio;type:text;comment:GitHub用户个人简介"`
	Role      string `json:"role" gorm:"column:role;type:text;comment:用户"`
}

func (RvUser) TableName() string {
	return "rv_user_table"
}
