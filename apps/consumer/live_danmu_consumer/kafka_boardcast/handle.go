package kafka_boardcast

import (
	kafkaMsg "LiveDanmu/apps/shared/models"
	"LiveDanmu/apps/shared/union_var"
	"context"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

func (b *BoardCast) SendDanmuMsg(ctx context.Context, data *kafkaMsg.DanmuKMsg) error {
	// 序列化Json
	msg, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		return err
	}
	// 组装弹幕消息
	kmsg := kafka.Message{
		Key:   []byte(strconv.FormatInt(data.RVID, 10)),
		Value: msg,
		Headers: []kafka.Header{
			{Key: "version", Value: []byte("1.0")},
			{Key: union_var.TRACE_ID_KEY, Value: []byte(ctx.Value(union_var.TRACE_ID_KEY).(string))},
		},
	}
	// 发送消息
	err = b.kClient.WriteMessages(ctx, kmsg)
	if err != nil {
		return err
	}

	return nil
}
