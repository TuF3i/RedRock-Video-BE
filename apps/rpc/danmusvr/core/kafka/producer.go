package kafka

import (
	"LiveDanmu/apps/public/models/dao"
	KMsg "LiveDanmu/apps/public/models/kafka"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

func (r *KClient) genDanmuKMsg(msg *danmusvr.PubDanmuData) KMsg.DanmuKMsg {
	// 结构体转换
	return KMsg.DanmuKMsg{
		RVID: msg.Rvid,
		OP:   KMsg.PUB_LIVE_DANMU,
		Data: dao.DanmuData{
			DanID:   msg.DanId,
			RVID:    msg.Rvid,
			UserId:  msg.Uid,
			Content: msg.Content,
			Color:   msg.Color,
			Ts:      msg.TimeStamp,
		},
	}
}

func (r *KClient) produceDanmuKMsg(ctx context.Context, data *danmusvr.PubDanmuData, writer *kafka.Writer) dto.Response {
	// 生成KMsg
	source := r.genDanmuKMsg(data)
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

func (r *KClient) SendVideoDanmuMsg(ctx context.Context, msg *danmusvr.PubDanmuData) dto.Response {
	resp := r.produceDanmuKMsg(ctx, msg, r.videoDanmuWriter)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}

func (r *KClient) SendLiveDanmuMsg(ctx context.Context, msg *danmusvr.PubDanmuData) dto.Response {
	resp := r.produceDanmuKMsg(ctx, msg, r.liveDanmuWriter)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}
