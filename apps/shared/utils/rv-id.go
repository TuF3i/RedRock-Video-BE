package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func RVIDEncoder(raw int64) string {
	return fmt.Sprintf("Rv%v", raw)
}

func RVIDDecoder(raw string) int64 {
	numStr := strings.TrimPrefix(raw, "Rv")
	num, _ := strconv.ParseInt(numStr, 10, 64)
	return num
}
