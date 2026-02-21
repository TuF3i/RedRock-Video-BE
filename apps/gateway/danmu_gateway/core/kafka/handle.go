package kafka

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core"
	ws "LiveDanmu/apps/gateway/danmu_gateway/core/websocket"
	kafkaMsg "LiveDanmu/apps/public/models/kafka"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const RETRY_COUNT = 5

//type DanmuKMsg struct {
//	RVID int64
//	OP   string
//	Data dao.DanmuData
//}

func process(ctx context.Context, m kafka.Message) error {
	var dataStruct kafkaMsg.DanmuKMsg
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(m.Value, &dataStruct)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, union_var.TRACE_ID_KEY, utils.GetHeaderValue(m, union_var.TRACE_ID_KEY))

	raw, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(dataStruct.Data)
	if err != nil {
		return err
	}

	core.Logger.INFO("PushDanmaku", zap.Any("Data", dataStruct.Data))

	err = ws.SendDanmaku(dataStruct.RVID, raw)
	if err != nil {
		return err
	}

	return nil
}

func (r *KClient) consumerLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 生成TraceID
			ctx := context.Background()
			// 手动提交偏移量
			m, err := r.consumer.FetchMessage(ctx)
			if err != nil {
				continue
			}
			// 处理消息
			for i := 0; i < RETRY_COUNT; i++ {
				err := process(ctx, m)
				if err == nil {
					break
				}
			}
			for i := 0; i < RETRY_COUNT; i++ {
				err := r.consumer.CommitMessages(ctx, m)
				if err == nil {
					break
				}
			}
		}
	}
}

func (r *KClient) StartConsume() {
	r.ctx, r.cancel = context.WithCancel(context.Background())

	go r.consumerLoop(r.ctx)
}

func (r *KClient) StopConsume() error {
	r.cancel()
	return r.consumer.Close()
}
