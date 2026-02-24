package parse_string

import (
	"strings"
)

func GetAddrs(hosts string) []string {
	return strings.Split(hosts, ",")
}
