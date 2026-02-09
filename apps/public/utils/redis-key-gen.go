package utils

import (
	"fmt"
)

func GenHotDanmuKey(vid int64) string {
	return fmt.Sprintf("dm:data:hot:%v", vid)
}

func GenFullDanmuKey(vid int64) string {
	return fmt.Sprintf("dm:data:full:%v", vid)
}

func GenHotDanmuCounterKey(vid int64) string {
	return fmt.Sprintf("dm:counter:hot:%v", vid)
}

func GenFullDanmuCounterKey(vid int64) string {
	return fmt.Sprintf("dm:counter:full:%v", vid)
}

func GenAccessTokenKey(token string) string {
	return fmt.Sprintf("auth:token:access:%v", token)
}

func GenRefreshTokenKey(token string) string {
	return fmt.Sprintf("auth:token:refresh:%v", token)
}

func GenPreSignedUrlKey(uid int64, rvid int64) string {
	return fmt.Sprintf("minio:presigenedurl:%v:%v", uid, rvid)
}
