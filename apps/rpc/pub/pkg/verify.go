package pkg

import (
	"regexp"
	"unicode/utf8"
)

var colorRe = regexp.MustCompile(`(?i)^#([0-9a-f]{3}|[0-9a-f]{6})$`)

// 校验房间ID
func ValidateRoomID(roomID int64) bool {
	if roomID < 1 || roomID > 9999999999 {
		return false
	}
	return true
}

// 校验用户ID
func ValidateUserID(userID int64) bool {
	if userID < 1 {
		return false
	}
	return true
}

// 校验颜色
func ValidateColor(color string) bool {
	if !colorRe.MatchString(color) {
		return false
	}
	return true
}

// 校验弹幕内容
func ValidateContent(content string) bool {
	if utf8.RuneCountInString(content) == 0 || utf8.RuneCountInString(content) > 60 {
		return false
	}
	for _, r := range content {
		if r == '\n' || r == '\r' || r == '\t' {
			return false
		}
	}
	return true
}
