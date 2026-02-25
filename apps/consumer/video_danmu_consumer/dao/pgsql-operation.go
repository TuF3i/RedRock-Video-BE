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

func (r *Dao) getUserInfoFromPgSQL(Tx *gorm.DB, uid int64) (*models.DUserInfo, error) {
	var raw models.RvUser
	if err := Tx.Where("github_uid = ?", uid).Select("github_uid", "avatar_url", "github_login").First(&raw).Error; err != nil {
		return nil, err
	}

	data := models.DUserInfo{
		Uid:       raw.Uid,
		UserName:  raw.Login,
		AvatarURL: raw.AvatarURL,
	}

	return &data, nil
}
