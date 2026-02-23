package dao

import (
	"LiveDanmu/apps/shared/models"
	"context"
)

func (r *Dao) InsertDanmuIntoDBs(ctx context.Context, data *models.DanmuData) error {
	// 开启数据库事务
	tx := r.pgdb.Begin()

	// 检查Pg中有没有相关记录
	ok, err := r.checkIfDanmuExistOnPgSQL(tx, data.DanID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 如果有直接跳过
	if ok {
		return nil
	}

	// 向Pg中插入记录
	err = r.insertDanmuToPgSQL(tx, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
