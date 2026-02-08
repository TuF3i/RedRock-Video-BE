package kafka

import (
	"LiveDanmu/apps/consumer/video_danmu_consumer"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"

	kafkaMsg "LiveDanmu/apps/public/models/kafka"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

const RETRY_COUNT = 5

func process(ctx context.Context, m kafka.Message) error {
	var dataStruct kafkaMsg.DanmuKMsg
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(m.Value, dataStruct)
	if err != nil {
		return err
	}
	
	ctx = context.WithValue(ctx, union_var.TRACE_ID_KEY, utils.GetHeaderValue(m, union_var.TRACE_ID_KEY))

	err = video_danmu_consumer.Dao.InsertDanmuIntoDBs(ctx, &dataStruct.Data)
	if err != nil {
		return err
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
