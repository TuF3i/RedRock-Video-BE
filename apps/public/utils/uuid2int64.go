package utils

import (
	"hash/fnv"
)

func UUIDToInt64(u string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(u))
	return int64(h.Sum64())
}
