package dao

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (r *Dao) initPgSQL() error {
	// 创建连接URL
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s "+
			"sslmode=disable "+
			"connect_timeout=5 "+
			"target_session_attrs=read-write", // 确保连接主库写数据

		r.conf.PgSQL.Urls[0], // K8s 内部访问
		5432,
		r.conf.PgSQL.User,
		r.conf.PgSQL.Password,
		r.conf.PgSQL.DBName,
	)
	// 连接pgpool
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger:      logger.Default.LogMode(logger.Info),
			PrepareStmt: false, // 关键：禁用 prepared statement（节点切换后失效）
		},
	)
	if err != nil {
		return err
	}

	// 配置连接池（应对节点切换）
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // 短连接，快速释放故障节点连接

	r.pgdb = db
	return nil
}
