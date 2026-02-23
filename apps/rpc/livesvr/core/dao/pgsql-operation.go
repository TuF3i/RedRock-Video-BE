package dao

import (
	"LiveDanmu/apps/shared/models"

	"gorm.io/gorm"
)

func (r *Dao) newRecord(tx *gorm.DB, data *models.LiveInfo) error {
	err := tx.Create(data).Error
	return err
}

func (r *Dao) delRecord(tx *gorm.DB, rvid int64) error {
	err := tx.Where("rv_id = ?", rvid).Delete(&models.LiveInfo{}).Error
	return err
}

func (r *Dao) getRecords(tx *gorm.DB, page int32, pageSize int32) ([]*models.LiveInfo, int64, error) {
	var dataSet []*models.LiveInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&models.LiveInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	return dataSet, total, nil
}

func (r *Dao) getRecordDetail(tx *gorm.DB, rvid int64) (*models.LiveInfo, error) {
	var data models.LiveInfo
	err := tx.Where("rv_id = ?", rvid).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Dao) getRecordDetailForUsers(tx *gorm.DB, uid int64) ([]*models.LiveInfo, int64, error) {
	var dataSet []*models.LiveInfo
	var total int64

	err := tx.Model(&models.LiveInfo{}).Count(&total).Error
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
	err := tx.Model(&models.LiveInfo{}).Where("rv_id = ?", rvid).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *Dao) getUserInfo(uid int64) (models.LUserInfo, error) {
	var data models.RvUser
	err := r.pgdb.Where("github_uid = ?", uid).Select("github_uid", "github_login", "avatar_url").First(&data).Error
	if err != nil {
		return models.LUserInfo{
			Uid:       0,
			UserName:  "",
			AvatarURL: "",
		}, err
	}

	return models.LUserInfo{
		Uid:       data.Uid,
		UserName:  data.Login,
		AvatarURL: data.AvatarURL,
	}, nil
}

func (r *Dao) getUserInfoBatch(uids []int64) (map[int64]models.LUserInfo, error) {
	if len(uids) == 0 {
		return make(map[int64]models.LUserInfo), nil
	}

	var users []models.RvUser
	err := r.pgdb.Where("github_uid IN ?", uids).Select("github_uid", "github_login", "avatar_url").Find(&users).Error
	if err != nil {
		return nil, err
	}

	result := make(map[int64]models.LUserInfo, len(users))
	for _, user := range users {
		result[user.Uid] = models.LUserInfo{
			Uid:       user.Uid,
			UserName:  user.Login,
			AvatarURL: user.AvatarURL,
		}
	}

	return result, nil
}
