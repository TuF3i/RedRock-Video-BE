package dao

import (
	"LiveDanmu/apps/shared/models"
)

func (r *Dao) AutoMigrate() error {
	err := r.pgdb.AutoMigrate(
		&models.RvUser{},
		&models.LiveInfo{},
		&models.VideoInfo{},
		&models.DanmuData{},
	)
	if err != nil {
		return err
	}

	return nil
}
