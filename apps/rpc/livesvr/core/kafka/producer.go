package kafka

import (
	"LiveDanmu/apps/public/models/dao"
	KMsg "LiveDanmu/apps/public/models/kafka"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/livesvr/core/dto"
	"context"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

func (r *KClient) genDanmuKMsg(rvid int64) KMsg.DanmuKMsg {
	// 结构体转换
	return KMsg.DanmuKMsg{
		RVID: rvid,
		OP:   KMsg.CLOSE_LIVE,
		Data: dao.DanmuData{},
	}
}

func (r *KClient) produceDanmuKMsg(ctx context.Context, rvid int64, writer *kafka.Writer) dto.Response {
	// 生成KMsg
	source := r.genDanmuKMsg(rvid)
	// 序列化Json
	msg, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(source)
	if err != nil {
		return dto.ServerInternalError(err)
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
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}

func (r *KClient) SendLiveOffMsg(ctx context.Context, rvid int64) dto.Response {
	resp := r.produceDanmuKMsg(ctx, rvid, r.boardCastController)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}
