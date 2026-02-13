package dao

import (
	"LiveDanmu/apps/public/models/dao"

	"gorm.io/gorm"
)

func (r *Dao) ifRecordExist(tx *gorm.DB, uid int64) (bool, error) {
	var count int64
	err := tx.Model(&dao.RvUser{}).Where("uid = ?", uid).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *Dao) getRecordDetail(tx *gorm.DB, uid int64) (*dao.RvUser, error) {
	data := new(dao.RvUser)
	err := tx.Where("uid = ?", uid).First(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Dao) newRecord(tx *gorm.DB, data *dao.RvUser) error {
	return tx.Create(data).Error
}
