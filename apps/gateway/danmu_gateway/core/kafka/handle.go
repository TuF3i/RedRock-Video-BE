package kafka

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core"
	"LiveDanmu/apps/gateway/danmu_gateway/core/dto"
	kafkaMsg "LiveDanmu/apps/public/models/kafka"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

const RETRY_COUNT = 5

func boardCastNewDanmu(ctx context.Context, dataStruct kafkaMsg.DanmuKMsg) {
	if !core.PoolGroup.IfPoolExist(dataStruct.RVID) {
		core.PoolGroup.NewPool(ctx, dataStruct.RVID)
	}
	core.PoolGroup.BoardCastMsg(dataStruct.RVID, dto.GenAddDanmuWMsg(&dataStruct.Data))
}

func closeLive(dataStruct kafkaMsg.DanmuKMsg) {
	if !core.PoolGroup.IfPoolExist(dataStruct.RVID) {
		return
	}
	core.PoolGroup.BoardCastMsg(dataStruct.RVID, dto.GenLiveOffWMsg())
	core.PoolGroup.CancelPool(dataStruct.RVID)
}

func process(ctx context.Context, m kafka.Message) error {
	var dataStruct kafkaMsg.DanmuKMsg
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(m.Value, dataStruct)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, union_var.TRACE_ID_KEY, utils.GetHeaderValue(m, union_var.TRACE_ID_KEY))

	// 提取操作
	switch dataStruct.OP {
	case kafkaMsg.CLOSE_LIVE:
		closeLive(dataStruct)
	case kafkaMsg.PUB_LIVE_DANMU:
		boardCastNewDanmu(ctx, dataStruct)
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
