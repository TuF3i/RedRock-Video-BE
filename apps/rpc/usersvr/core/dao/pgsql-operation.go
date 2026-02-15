package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/union_var"
	"strconv"

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
	if strconv.FormatInt(data.Uid, 10) == r.conf.AdminId {
		data.Role = union_var.JWT_ROLE_ADMIN
	} else {
		data.Role = union_var.JWT_ROLE_USER
	}
	return tx.Create(data).Error
}

func (r *Dao) getRecordDetails(tx *gorm.DB, page int32, pageSize int32, role string) ([]*dao.RvUser, int64, error) {
	var dataSet []*dao.RvUser
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&dao.RvUser{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Where("role = ?", role).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	return dataSet, total, nil
}

func (r *Dao) setColumnValue(tx *gorm.DB, uid int64, column string, value interface{}) error {
	err := tx.Model(&dao.RvUser{}).Where("uid = ?", uid).Update(column, value).Error
	if err != nil {
		return err
	}

	return nil
}
