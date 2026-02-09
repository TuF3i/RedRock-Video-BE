package minio

import (
	"LiveDanmu/apps/public/union_var"
	"context"

	"github.com/minio/minio-go/v7"
)

func (r *Minio) GetSignedUrl(ctx context.Context, minioKey string) (string, error) {
	// 获取预签名URL
	url, err := r.MClient.PresignedGetObject(
		ctx,
		r.conf.Minio.BlanketName,
		minioKey,
		union_var.MINIO_EXPIRE_TIME, // key有效期
		nil,
	)
	if err != nil {
		return "", err
	}
	// 返回url
	return url.String(), nil
}

func (r *Minio) CheckIfFaceExist(ctx context.Context, minioKey string) (bool, error) {
	// 检查文件是否存在
	_, err := r.MClient.StatObject(
		ctx,
		r.conf.Minio.PicBlanketName,
		minioKey,
		minio.StatObjectOptions{},
	)
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		if errResponse.Code == "NoSuchKey" || errResponse.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Minio) CheckIfVideoExist(ctx context.Context, minioKey string) (bool, error) {
	// 检查文件是否存在
	_, err := r.MClient.StatObject(
		ctx,
		r.conf.Minio.BlanketName,
		minioKey,
		minio.StatObjectOptions{},
	)
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		if errResponse.Code == "NoSuchKey" || errResponse.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Minio) DelVideo(ctx context.Context, minioKey string) error {
	// 删除对象
	err := r.MClient.RemoveObject(
		ctx,
		r.conf.Minio.BlanketName,
		minioKey,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *Minio) DelFace(ctx context.Context, minioKey string) error {
	// 删除对象
	err := r.MClient.RemoveObject(
		ctx,
		r.conf.Minio.PicBlanketName,
		minioKey,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}
