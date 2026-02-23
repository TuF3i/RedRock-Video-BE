package dao

import (
	"LiveDanmu/apps/shared/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *Dao) getFullDanmuP(ctx context.Context, vid int64) ([]*models.DanmuData, error) {
	var results []*models.DanmuData
	err := r.pgdb.Where("rvid = ?", vid).Order("ts DESC").Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *Dao) checkIfDanmuExistOnPgSQL(tx *gorm.DB, danID int64) (bool, error) {
	var dest models.DanmuData
	if err := tx.
		Where("dan_id = ?", danID).
		First(&dest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Dao) delVideoDanmu(tx *gorm.DB, danID int64) error {
	var dest models.DanmuData
	if err := tx.
		Where("dan_id = ?", danID).
		Delete(&dest).Error; err != nil {
		return err
	}
	return nil
}

func (r *Dao) getVideoDanmuDetail(tx *gorm.DB, danID int64) (*models.DanmuData, error) {
	var data models.DanmuData
	err := tx.Where("dan_id = ?", danID).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Dao) getUserInfo(uid int64) (models.DUserInfo, error) {
	var data models.RvUser
	err := r.pgdb.Where("github_uid = ?", uid).Select("github_uid", "avatar_url", "github_login").Find(&data).Error
	if err != nil {
		return models.DUserInfo{
			Uid:       0,
			UserName:  "",
			AvatarURL: "",
		}, err
	}

	return models.DUserInfo{
		Uid:       data.Uid,
		UserName:  data.Login,
		AvatarURL: data.AvatarURL,
	}, nil
}

func (r *Dao) getUserInfoBatch(uids []int64) (map[int64]models.DUserInfo, error) {
	if len(uids) == 0 {
		return make(map[int64]models.DUserInfo), nil
	}

	var users []models.RvUser
	err := r.pgdb.Where("github_uid IN ?", uids).Select("github_uid", "avatar_url", "github_login").Find(&users).Error
	if err != nil {
		return nil, err
	}

	result := make(map[int64]models.DUserInfo, len(users))
	for _, user := range users {
		result[user.Uid] = models.DUserInfo{
			Uid:       user.Uid,
			UserName:  user.Login,
			AvatarURL: user.AvatarURL,
		}
	}

	return result, nil
}
