package kafka

import (
	"LiveDanmu/apps/rpc/danmu/dto"
	"LiveDanmu/apps/rpc/danmu/kitex_gen/danmusvr"
	"context"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

func (r *KClient) produceDanmuKMsg(ctx context.Context, data *danmusvr.DanmuMsg, writer *kafka.Writer) dto.Response {
	// 序列化Json
	msg, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(*data)
	if err != nil {
		return dto.ServerInternalError(err)
	}
	// 组装弹幕消息
	kmsg := kafka.Message{
		Key:   []byte(strconv.FormatInt(data.RoomId, 10)),
		Value: msg,
		Headers: []kafka.Header{
			{Key: "version", Value: []byte("1.0")},
		},
	}
	// 发送消息
	err = writer.WriteMessages(ctx, kmsg)
	if err != nil {
		return dto.ServerInternalError(err)
	}

	return dto.OperationSuccess
}

func (r *KClient) SendVideoDanmuMsg(ctx context.Context, msg *danmusvr.DanmuMsg) dto.Response {
	resp := r.produceDanmuKMsg(ctx, msg, r.videoDanmuWriter)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}

func (r *KClient) SendLiveDanmuMsg(ctx context.Context, msg *danmusvr.DanmuMsg) dto.Response {
	resp := r.produceDanmuKMsg(ctx, msg, r.liveDanmuWriter)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	return dto.OperationSuccess
}
