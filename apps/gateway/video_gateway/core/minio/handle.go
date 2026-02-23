package minio

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func (r *Minio) UploadFile(ctx context.Context, minioKey string, file *multipart.FileHeader) error {
	// 打开文件
	f, err := file.Open()
	if err != nil {
		return err
	}
	// 上传到私有桶
	_, err = r.MClient.PutObject(
		ctx,
		r.conf.Minio.BlanketName,
		minioKey,
		f,
		file.Size,
		minio.PutObjectOptions{ContentType: "video/mp4"},
	)
	if err != nil {
		return err
	}

	_ = f.Close()
	return nil
}

func (r *Minio) UploadFaceFile(ctx context.Context, minioKey string, file *multipart.FileHeader) error {
	// 打开文件
	f, err := file.Open()
	if err != nil {
		return err
	}
	// 上传到私有桶
	_, err = r.MClient.PutObject(
		ctx,
		r.conf.Minio.PicBlanketName,
		minioKey,
		f,
		file.Size,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}

	_ = f.Close()
	return nil
}
