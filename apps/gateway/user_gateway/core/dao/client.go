package dao

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
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
