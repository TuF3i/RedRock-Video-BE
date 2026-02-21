package kafka

import (
	"LiveDanmu/apps/consumer/live_danmu_consumer"
	"LiveDanmu/apps/consumer/live_danmu_consumer/dao"
	"LiveDanmu/apps/consumer/live_danmu_consumer/kafka_boardcast"
	"LiveDanmu/apps/public/union_var"
	"LiveDanmu/apps/public/utils"
	"context"

	kafkaMsg "LiveDanmu/apps/public/models/kafka"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const RETRY_COUNT = 5

func process(ctx context.Context, dao *dao.Dao, boardCast *kafka_boardcast.BoardCast, m kafka.Message) error {
	live_danmu_consumer.Logger.INFO("OnConsuming:", zap.Any("KMsg", m))

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

	err = boardCast.SendDanmuMsg(ctx, &dataStruct)
	if err != nil {
		return err
	}

	return nil
}

func consumerLoop(ctx context.Context, dao *dao.Dao, boardCast *kafka_boardcast.BoardCast, r *kafka.Reader) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			//装一个新的context
			ctx := context.Background()
			// 自动提交偏移量
			m, err := r.FetchMessage(ctx)
			if err != nil {
				continue
			}
			// 处理消息
			for i := 0; i < RETRY_COUNT; i++ {
				err := process(ctx, dao, boardCast, m)
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
}

func (r *ConsumerGroup) StartConsume() {
	r.ctx, r.cancel = context.WithCancel(context.Background())

	go consumerLoop(r.ctx, r.dao, r.kBoardCast, r.kClient)
}

func (r *ConsumerGroup) StopConsume() error {
	r.cancel()
	return r.kClient.Close()
}
