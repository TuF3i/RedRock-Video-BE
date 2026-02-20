package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"errors"

	"gorm.io/gorm"
)

func (r *Dao) checkIfRecordExist(tx *gorm.DB, rvid int64) (bool, error) {
	err := tx.Where("rvid = ?", rvid).First(&dao.VideoInfo{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Dao) creatNewRecord(tx *gorm.DB, data *dao.VideoInfo) error {
	err := tx.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Dao) delARecord(tx *gorm.DB, rvid int64) error {
	err := tx.Where("rvid = ?", rvid).Delete(&dao.VideoInfo{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Dao) getDetailOfARecord(tx *gorm.DB, rvid int64) (*dao.VideoInfo, error) {
	data := new(dao.VideoInfo)
	err := tx.Where("rvid = ?", rvid).First(data).Error
	if err != nil {
		return nil, err
	}

	data.User, _ = r.getUserInfo(data.UID)

	return data, nil
}

func (r *Dao) getRecordList(tx *gorm.DB, page int32, pageSize int32) ([]*dao.VideoInfo, int64, error) {
	var dataSet []*dao.VideoInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&dao.VideoInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Where("in_judge = ?", false).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	for _, val := range dataSet {
		val.User, _ = r.getUserInfo(val.UID)
	}

	return dataSet, total, nil
}

func (r *Dao) getUserInfo(uid int64) (dao.VUserInfo, error) {
	var data dao.RvUser
	err := r.pgdb.Where("github_uid = ?", uid).Select("github_uid", "github_login", "avatar_url").First(&data).Error
	if err != nil {
		return dao.VUserInfo{
			Uid:       0,
			UserName:  "",
			AvatarURL: "",
		}, err
	}

	return dao.VUserInfo{
		Uid:       data.Uid,
		UserName:  data.Login,
		AvatarURL: data.AvatarURL,
	}, nil
}

func (r *Dao) getJudgingRecordList(tx *gorm.DB, page int32, pageSize int32) ([]*dao.VideoInfo, int64, error) {
	var dataSet []*dao.VideoInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&dao.VideoInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Where("in_judge = ?", true).Select("rvid", "title", "face_url", "in_judge").Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	for _, val := range dataSet {
		val.User, _ = r.getUserInfo(val.UID)
	}

	return dataSet, total, nil
}

func (r *Dao) getUserRecordList(tx *gorm.DB, page int32, pageSize int32, uid int64) ([]*dao.VideoInfo, int64, error) {
	var dataSet []*dao.VideoInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&dao.VideoInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Where("uid = ?", uid).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	for _, val := range dataSet {
		val.User, _ = r.getUserInfo(val.UID)
	}

	return dataSet, total, nil
}

func (r *Dao) setRecordColumn(tx *gorm.DB, rvid int64, column string, value interface{}) error {
	err := tx.Model(&dao.VideoInfo{}).Where("rvid = ?", rvid).Update(column, value).Error
	if err != nil {
		return err
	}
	return nil
}
