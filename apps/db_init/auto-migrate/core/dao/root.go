package dao

import (
	"LiveDanmu/apps/shared/config/config_template"

	"gorm.io/gorm"
)

type Dao struct {
	conf *config_template.DBInitConfig
	pgdb *gorm.DB
}

func GetDao(conf *config_template.DBInitConfig) (*Dao, error) {
	d := Dao{conf: conf}
	if err := d.initPgSQL(); err != nil {
		return nil, err
	}

	return &d, nil
}
