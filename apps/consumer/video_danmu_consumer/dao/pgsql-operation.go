package dao

import (
	"LiveDanmu/apps/shared/models"
	"errors"

	"gorm.io/gorm"
)

func (r *Dao) checkIfDanmuExistOnPgSQL(Tx *gorm.DB, danID int64) (bool, error) {
	var dest models.DanmuData
	if err := Tx.
		Where("dan_id = ?", danID).
		First(&dest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Dao) insertDanmuToPgSQL(Tx *gorm.DB, data *models.DanmuData) error {
	if err := Tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
