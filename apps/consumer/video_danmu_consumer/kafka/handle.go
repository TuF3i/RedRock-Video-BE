package kafka

import (
	"LiveDanmu/apps/consumer/video_danmu_consumer"
	"LiveDanmu/apps/consumer/video_danmu_consumer/dao"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"

	kafkaMsg "LiveDanmu/apps/public/models/kafka"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const RETRY_COUNT = 5

func process(ctx context.Context, dao *dao.Dao, m kafka.Message) error {
	video_danmu_consumer.Logger.INFO("OnConsuming:", zap.Any("KMsg", m))

	var dataStruct kafkaMsg.DanmuKMsg
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(m.Value, &dataStruct)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, union_var.TRACE_ID_KEY, utils.GetHeaderValue(m, union_var.TRACE_ID_KEY))

	err = dao.InsertDanmuIntoDBs(ctx, &dataStruct.Data)
	if err != nil {
		return err
	}

	return nil
}

func consumerLoop(ctx context.Context, dao *dao.Dao, r *kafka.Reader) {
	video_danmu_consumer.Logger.INFO("consumerLoop starting")
	defer video_danmu_consumer.Logger.INFO("consumerLoop stopped")

	for {
		select {
		case <-ctx.Done():
			video_danmu_consumer.Logger.INFO("context canceled, exiting consumerLoop")
			return
		default:
			video_danmu_consumer.Logger.INFO("waiting for message...")
			m, err := r.FetchMessage(ctx)
			if err != nil {
				video_danmu_consumer.Logger.INFO("FetchMessage failed", zap.Error(err))
				continue
			}
			video_danmu_consumer.Logger.INFO("message received",
				zap.Int64("offset", m.Offset),
				zap.Int("partition", m.Partition),
				zap.ByteString("key", m.Key),
				zap.ByteString("value", m.Value),
			)
			var processErr error
			for i := 0; i <= RETRY_COUNT; i++ {
				processErr = process(ctx, dao, m)
				if processErr == nil {
					break
				}
				video_danmu_consumer.Logger.WARN("process failed, retrying", zap.Int("attempt", i+1), zap.Error(processErr))
			}
			if processErr != nil {
				video_danmu_consumer.Logger.INFO("process failed after all retries", zap.Error(processErr))
			}
			var commitErr error
			for i := 0; i <= RETRY_COUNT; i++ {
				commitErr = r.CommitMessages(ctx, m)
				if commitErr == nil {
					break
				}
			}
			if commitErr != nil {
				video_danmu_consumer.Logger.INFO("CommitMessages failed", zap.Error(commitErr))
			}
		}
	}
}

func (r *ConsumerGroup) StartConsume() {
	r.ctx, r.cancel = context.WithCancel(context.Background())

	go consumerLoop(r.ctx, r.dao, r.kClient)
}

func (r *ConsumerGroup) StopConsume() error {
	r.cancel()
	return r.kClient.Close()
}
