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
