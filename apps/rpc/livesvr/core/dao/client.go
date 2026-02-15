package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (r *Dao) initRedisClient() error {
	// 创建集群连接
	rdb := redis.NewClusterClient(
		&redis.ClusterOptions{
			// 基础配置
			Addrs:        r.conf.Redis.Urls,
			Password:     r.conf.Redis.Password,
			MaxRedirects: 3, // 最大重定向次数

			// 连接池配置
			PoolSize:     10, // 每个节点的连接池大小
			MinIdleConns: 5,  // 最小空闲连接
			MaxRetries:   3,  // 失败重试次数
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
		},
	)
	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return err
	}

	r.rdb = rdb
	return nil
}

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
