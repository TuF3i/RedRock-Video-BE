package minio

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/minio/minio-go/v7"
)

type Minio struct {
	MClient *minio.Client
	conf    *config_template.VideoRpcConfig
}

func GetMinio(conf *config_template.VideoRpcConfig) (*Minio, error) {
	m := &Minio{conf: conf}
	err := m.initMinioClient()
	if err != nil {
		return nil, err
	}
	return m, nil
}
