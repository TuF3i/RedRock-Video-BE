package kafka

import (
	KMsg "LiveDanmu/apps/shared/models"
	"LiveDanmu/apps/shared/union_var"
	"context"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

func (r *KClient) genDanmuKMsg(rvid int64) KMsg.DanmuKMsg {
	// 结构体转换
	return KMsg.DanmuKMsg{
		RVID: rvid,
		OP:   union_var.CLOSE_LIVE,
		Data: KMsg.DanmuData{},
	}
}

func (r *KClient) produceDanmuKMsg(ctx context.Context, rvid int64, writer *kafka.Writer) error {
	// 生成KMsg
	source := r.genDanmuKMsg(rvid)
	// 序列化Json
	msg, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(source)
	if err != nil {
		return err
	}
	// 组装弹幕消息
	kmsg := kafka.Message{
		Key:   []byte(strconv.FormatInt(source.RVID, 10)),
		Value: msg,
		Headers: []kafka.Header{
			{Key: "version", Value: []byte("1.0")},
			{Key: union_var.TRACE_ID_KEY, Value: []byte(ctx.Value(union_var.TRACE_ID_KEY).(string))},
		},
	}
	// 发送消息
	err = writer.WriteMessages(ctx, kmsg)
	if err != nil {
		return err
	}

	return nil
}

func (r *KClient) SendLiveOffMsg(ctx context.Context, rvid int64) error {
	return r.produceDanmuKMsg(ctx, rvid, r.boardCastController)
}
