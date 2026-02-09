package minio

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func (r *Minio) initMinioClient() error {
	// 创建连接
	client, err := minio.New(
		r.conf.Minio.Urls[0],
		&minio.Options{
			Creds: credentials.NewStaticV4(
				r.conf.Minio.AccessKey,
				r.conf.Minio.SecretKey,
				"",
			),
			Secure: r.conf.Minio.UseSSL,
		},
	)
	if err != nil {
		return err
	}

	// 检查桶是否存在
	exists, err := client.BucketExists(context.Background(), r.conf.Minio.BlanketName)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("blanket not exists")
	}

	// 将连接写入结构体
	r.MClient = client
	return nil
}
