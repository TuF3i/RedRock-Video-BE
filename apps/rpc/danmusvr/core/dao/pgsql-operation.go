package dao

import (
	publicDao "LiveDanmu/apps/public/models/dao"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *Dao) getFullDanmuP(ctx context.Context, vid int64) ([]*publicDao.DanmuData, error) {
	var results []*publicDao.DanmuData
	err := r.pgdb.Where("rv_id = ?", vid).Order("ts DESC").Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *Dao) checkIfDanmuExistOnPgSQL(tx *gorm.DB, danID int64) (bool, error) {
	var dest publicDao.DanmuData
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

func (r *Dao) delVideoDanmu(Tx *gorm.DB, danID int64) error {
	var dest publicDao.DanmuData
	if err := Tx.
		Where("dan_id = ?", danID).
		Delete(&dest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (r *Dao) getVideoDanmuDetail(tx *gorm.DB, danID int64) (*publicDao.DanmuData, error) {
	var data *publicDao.DanmuData
	err := tx.Where("dan_id = ?", danID).First(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Dao) getUserInfo(uid int64) (publicDao.DUserInfo, error) {
	var data publicDao.RvUser
	err := r.pgdb.Where("uid = ?", uid).Select("uid", "avatar_url", "github_login").Find(data).Error
	if err != nil {
		return publicDao.DUserInfo{
			Uid:       0,
			UserName:  "",
			AvatarURL: "",
		}, err
	}

	return publicDao.DUserInfo{
		Uid:       data.Uid,
		UserName:  data.Login,
		AvatarURL: data.AvatarURL,
	}, nil
}
