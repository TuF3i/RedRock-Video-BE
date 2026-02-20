package dao

import (
	"LiveDanmu/apps/public/jwt"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"
)

func (r *Dao) AddUser(data *dao.RvUser) error {
	uid := data.Uid
	tx := r.pgdb.Begin()
	// 查询用户是否存在
	ok, err := r.ifRecordExist(tx, uid)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 用户存在直接返回
	if ok {
		tx.Rollback()
		return nil
	}
	// 不存在直接创建
	err = r.newRecord(tx, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *Dao) IfUserExist(uid int64) (bool, error) {
	tx := r.pgdb.Begin()
	defer tx.Rollback()
	// 查询用户是否存在
	ok, err := r.ifRecordExist(tx, uid)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	tx.Commit()

	return true, nil
}

func (r *Dao) GetUserInfo(uid int64) (*dao.RvUser, error) {
	tx := r.pgdb.Begin()
	defer tx.Rollback()
	// 获取用户信息
	data, err := r.getRecordDetail(tx, uid)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return data, nil
}

func (r *Dao) SetNewAccessToken(ctx context.Context, uid int64, token string) error {
	key := utils.GenAccessTokenKey(uid)
	err := r.setNewValue(ctx, key, token, jwt.GetAccessTokenExpireTime())
	if err != nil {
		return err
	}

	return nil
}

func (r *Dao) SetNewRefreshToken(ctx context.Context, uid int64, token string) error {
	key := utils.GenRefreshTokenKey(uid)
	err := r.setNewValue(ctx, key, token, jwt.GetRefreshTokenExpireTime())
	if err != nil {
		return err
	}

	return nil
}

func (r *Dao) VerifyRefreshToken(ctx context.Context, uid int64, token string) (bool, error) {
	key := utils.GenRefreshTokenKey(uid)
	data, err := r.getKeyValue(ctx, key)
	if err != nil {
		return false, err
	}

	if data == token {
		return true, nil
	}

	return false, nil
}

func (r *Dao) SetAdminRole(ctx context.Context, uid int64) error {
	tx := r.pgdb.Begin()
	err := r.setColumnValue(tx, uid, "role", union_var.JWT_ROLE_ADMIN)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *Dao) GetAdminerList(ctx context.Context, page int32, pageSize int32) ([]*dao.RvUser, int64, error) {
	tx := r.pgdb.Begin()
	defer tx.Rollback()
	dataSet, total, err := r.getRecordDetails(tx, page, pageSize, union_var.JWT_ROLE_ADMIN)
	if err != nil {
		return nil, 0, err
	}

	tx.Commit()

	return dataSet, total, nil
}

func (r *Dao) GetUserList(ctx context.Context, page int32, pageSize int32) ([]*dao.RvUser, int64, error) {
	tx := r.pgdb.Begin()
	defer tx.Rollback()
	dataSet, total, err := r.getRecordDetails(tx, page, pageSize, union_var.JWT_ROLE_USER)
	if err != nil {
		return nil, 0, err
	}

	tx.Commit()

	return dataSet, total, nil
}

func (r *Dao) Logout(ctx context.Context, uid int64) error {
	keyForAccessToken := utils.GenAccessTokenKey(uid)
	keyForRefreshToken := utils.GenRefreshTokenKey(uid)

	ok, err := r.ifKeyExist(ctx, keyForAccessToken)
	if err != nil {
		return err
	}
	if ok {
		if err := r.delKeyValue(ctx, keyForAccessToken); err != nil {
			return err
		}
	}

	ok, err = r.ifKeyExist(ctx, keyForRefreshToken)
	if err != nil {
		return err
	}
	if ok {
		if err := r.delKeyValue(ctx, keyForRefreshToken); err != nil {
			return err
		}
	}

	return nil
}
