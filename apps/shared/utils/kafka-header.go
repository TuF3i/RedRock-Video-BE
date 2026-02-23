package utils

import "github.com/segmentio/kafka-go"

func GetHeaderValue(msg kafka.Message, key string) string {
	for _, header := range msg.Headers {
		if header.Key == key {
			return string(header.Value)
		}
	}
	return ""
}
