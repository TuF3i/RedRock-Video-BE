package dao

import (
	"LiveDanmu/apps/public/models/dao"

	"gorm.io/gorm"
)

func (r *Dao) newRecord(tx *gorm.DB, data *dao.LiveInfo) error {
	err := tx.Create(data).Error
	return err
}

func (r *Dao) delRecord(tx *gorm.DB, rvid int64) error {
	err := tx.Where("rv_id = ?", rvid).Delete(&dao.LiveInfo{}).Error
	return err
}

func (r *Dao) getRecords(tx *gorm.DB, page int32, pageSize int32) ([]*dao.LiveInfo, int64, error) {
	var dataSet []*dao.LiveInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&dao.LiveInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	return dataSet, total, nil
}

func (r *Dao) getRecordDetail(tx *gorm.DB, rvid int64) (*dao.LiveInfo, error) {
	data := &dao.LiveInfo{}
	err := tx.Where("rv_id = ?", rvid).Find(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Dao) getRecordDetailForUsers(tx *gorm.DB, uid int64) ([]*dao.LiveInfo, int64, error) {
	var dataSet []*dao.LiveInfo
	var total int64

	err := tx.Model(&dao.LiveInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Where("ower_id = ?", uid).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	return dataSet, total, nil
}

func (r *Dao) ifRecordExist(tx *gorm.DB, rvid int64) (bool, error) {
	var count int64
	err := tx.Model(&dao.LiveInfo{}).Where("rv_id = ?", rvid).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *Dao) getUserInfo(uid int64) (dao.LUserInfo, error) {
	var data dao.RvUser
	err := r.pgdb.Where("github_uid = ?", uid).Select("github_uid", "github_login", "avatar_url").First(&data).Error
	if err != nil {
		return dao.LUserInfo{
			Uid:       0,
			UserName:  "",
			AvatarURL: "",
		}, err
	}

	return dao.LUserInfo{
		Uid:       data.Uid,
		UserName:  data.Login,
		AvatarURL: data.AvatarURL,
	}, nil
}
