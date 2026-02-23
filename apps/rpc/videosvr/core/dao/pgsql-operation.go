package dao

import (
	"LiveDanmu/apps/shared/models"
	"errors"

	"gorm.io/gorm"
)

func (r *Dao) checkIfRecordExist(tx *gorm.DB, rvid int64) (bool, error) {
	err := tx.Where("rvid = ?", rvid).First(&models.VideoInfo{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Dao) creatNewRecord(tx *gorm.DB, data *models.VideoInfo) error {
	err := tx.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Dao) delARecord(tx *gorm.DB, rvid int64) error {
	err := tx.Where("rvid = ?", rvid).Delete(&models.VideoInfo{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Dao) getDetailOfARecord(tx *gorm.DB, rvid int64) (*models.VideoInfo, error) {
	var data models.VideoInfo
	err := tx.Where("rvid = ?", rvid).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Dao) getRecordList(tx *gorm.DB, page int32, pageSize int32) ([]*models.VideoInfo, int64, error) {
	var dataSet []*models.VideoInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&models.VideoInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Where("in_judge = ?", false).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	if len(dataSet) > 0 {
		uids := make([]int64, len(dataSet))
		for i, val := range dataSet {
			uids[i] = val.UID
		}
		userMap, _ := r.getUserInfoBatch(uids)
		for _, val := range dataSet {
			if user, ok := userMap[val.UID]; ok {
				val.User = user
			}
		}
	}

	return dataSet, total, nil
}

func (r *Dao) getUserInfo(uid int64) (models.VUserInfo, error) {
	var data models.RvUser
	err := r.pgdb.Where("github_uid = ?", uid).Select("github_uid", "github_login", "avatar_url").First(&data).Error
	if err != nil {
		return models.VUserInfo{
			Uid:       0,
			UserName:  "",
			AvatarURL: "",
		}, err
	}

	return models.VUserInfo{
		Uid:       data.Uid,
		UserName:  data.Login,
		AvatarURL: data.AvatarURL,
	}, nil
}

func (r *Dao) getUserInfoBatch(uids []int64) (map[int64]models.VUserInfo, error) {
	if len(uids) == 0 {
		return make(map[int64]models.VUserInfo), nil
	}

	var users []models.RvUser
	err := r.pgdb.Where("github_uid IN ?", uids).Select("github_uid", "github_login", "avatar_url").Find(&users).Error
	if err != nil {
		return nil, err
	}

	result := make(map[int64]models.VUserInfo, len(users))
	for _, user := range users {
		result[user.Uid] = models.VUserInfo{
			Uid:       user.Uid,
			UserName:  user.Login,
			AvatarURL: user.AvatarURL,
		}
	}

	return result, nil
}

func (r *Dao) getJudgingRecordList(tx *gorm.DB, page int32, pageSize int32) ([]*models.VideoInfo, int64, error) {
	var dataSet []*models.VideoInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&models.VideoInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Where("in_judge = ?", true).Select("rvid", "title", "face_url", "in_judge").Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	return dataSet, total, nil
}

func (r *Dao) getUserRecordList(tx *gorm.DB, page int32, pageSize int32, uid int64) ([]*models.VideoInfo, int64, error) {
	var dataSet []*models.VideoInfo
	var total int64

	offset := (page - 1) * pageSize
	err := tx.Model(&models.VideoInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int(offset)).Limit(int(pageSize)).Where("uid = ?", uid).Find(&dataSet).Error
	if err != nil {
		return nil, 0, err
	}

	if len(dataSet) > 0 {
		uids := make([]int64, len(dataSet))
		for i, val := range dataSet {
			uids[i] = val.UID
		}
		userMap, _ := r.getUserInfoBatch(uids)
		for _, val := range dataSet {
			if user, ok := userMap[val.UID]; ok {
				val.User = user
			}
		}
	}

	return dataSet, total, nil
}

func (r *Dao) setRecordColumn(tx *gorm.DB, rvid int64, column string, value interface{}) error {
	err := tx.Model(&models.VideoInfo{}).Where("rvid = ?", rvid).Update(column, value).Error
	if err != nil {
		return err
	}
	return nil
}
