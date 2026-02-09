package pkg

import (
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// 1. 校验 RVID：必须是大于 0 的整数
func ValidateRVID(rvid int64) bool {
	return rvid > 0
}

// 2. 校验 FaceUrl：非空、长度≤512、符合http/https URL格式
func ValidateFaceUrl(faceUrl string) bool {
	trimmed := strings.TrimSpace(faceUrl)
	// 非空校验
	if trimmed == "" {
		return false
	}
	// 长度校验（≤512字符）
	if utf8.RuneCountInString(trimmed) > 512 {
		return false
	}
	// URL格式校验
	parsedUrl, err := url.ParseRequestURI(trimmed)
	if err != nil || (parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https") {
		return false
	}
	return true
}

// 3. 校验 MinioKey：非空、长度≤512、仅含合法字符（字母/数字/下划线/斜杠/点/短横线）
func ValidateMinioKey(minioKey string) bool {
	trimmed := strings.TrimSpace(minioKey)
	if trimmed == "" {
		return false
	}
	if utf8.RuneCountInString(trimmed) > 512 {
		return false
	}
	// 合法字符校验
	validKeyRegex := regexp.MustCompile(`^[a-zA-Z0-9_\./-]+$`)
	return validKeyRegex.MatchString(trimmed)
}

// 4. 校验 Title：非空（去空白后）、长度≤255
func ValidateTitle(title string) bool {
	trimmed := strings.TrimSpace(title)
	if trimmed == "" {
		return false
	}
	return utf8.RuneCountInString(trimmed) <= 255
}

// 5. 校验 Description：可选字段，仅校验长度≤10000（为空也返回true）
func ValidateDescription(description string) bool {
	return utf8.RuneCountInString(description) <= 10000
}

// 9. 校验 AuthorID：必须是大于 0 的整数
func ValidateAuthorID(authorID int64) bool {
	return authorID > 0
}

// 10. 校验 AuthorName：非空（去空白后）、长度≤255
func ValidateAuthorName(authorName string) bool {
	trimmed := strings.TrimSpace(authorName)
	if trimmed == "" {
		return false
	}
	return utf8.RuneCountInString(trimmed) <= 255
}
