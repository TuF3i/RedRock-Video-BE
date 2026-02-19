package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"context"
)

func (r *Dao) InsertDanmuIntoDBs(ctx context.Context, data *dao.DanmuData) error {
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

	// 检查redis里是否有相关记录
	ok, err = r.checkIfDanmuExistOnRedis(ctx, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 如果没有就跳过
	if !ok {
		tx.Commit()
		return nil
	}

	// 向redis里插入数据
	err = r.insertDanmuToRedis(ctx, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
