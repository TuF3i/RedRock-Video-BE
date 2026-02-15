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

func delNewDanmu(ctx context.Context, dataStruct kafkaMsg.DanmuKMsg) {
	if !core.PoolGroup.IfPoolExist(dataStruct.RVID) {
		core.PoolGroup.NewPool(ctx, dataStruct.RVID)
	}
	core.PoolGroup.BoardCastMsg(dataStruct.RVID, dto.GenRemoveDanmuWMsg(&dataStruct.Data))
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
	case kafkaMsg.DEL_LIVE_DANMU:
		delNewDanmu(ctx, dataStruct)
	}

	return nil
}

func ConsumerLoop(r *kafka.Reader) {
	for {
		// 生成TraceID
		ctx := context.Background()
		// 自动提交偏移量
		m, err := r.FetchMessage(ctx)
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
			err := r.CommitMessages(ctx, m)
			if err == nil {
				break
			}
		}
	}
}
