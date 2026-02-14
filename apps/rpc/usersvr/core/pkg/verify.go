package pkg

import (
	"net/url"
	"regexp"
	"strings"
)

// 1. 校验ID字段：GitHub用户ID为正整数
func ValidateGitHubUserID(id int64) bool {
	return id > 0
}

// 2. 校验Login字段：符合GitHub用户名规则
// 规则：3-39个字符，仅包含字母、数字、连字符(-)，不能以连字符开头/结尾
var loginRegex = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,37}[a-zA-Z0-9])?$`)

func ValidateGitHubUserLogin(login string) bool {
	if len(login) > 39 {
		return false
	}
	return loginRegex.MatchString(login)
}

// 4. 校验AvatarURL字段：非空且为有效的URL格式
func ValidateGitHubUserAvatarURL(avatarURL string) bool {
	if avatarURL == "" {
		return false
	}
	// 解析URL验证格式合法性
	_, err := url.ParseRequestURI(avatarURL)
	if err != nil {
		return false
	}
	// 额外确保是http/https协议（GitHub头像均为该协议）
	u, err := url.Parse(avatarURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return false
	}
	return true
}

// 5. 校验Bio字段：可选字段（可空），非空时长度不超过160字符（GitHub官方限制）
func ValidateGitHubUserBio(bio string) bool {
	// 空值合法
	if bio == "" {
		return true
	}
	// 非空时不超过160字符
	return len(strings.TrimSpace(bio)) <= 160
}
