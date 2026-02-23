package dao

import (
	"LiveDanmu/apps/shared/config/config_template"

	"gorm.io/gorm"
)

type Dao struct {
	conf *config_template.LiveDanmuConsumerConfig
	pgdb *gorm.DB
}

func GetDao(conf *config_template.LiveDanmuConsumerConfig) (*Dao, error) {
	d := Dao{conf: conf}
	if err := d.initPgSQL(); err != nil {
		return nil, err
	}

	return &d, nil
}
