package kafka

import "LiveDanmu/apps/public/models/dao"

const (
	LIVE_DANMU_PUB_TOPIC  = "danmu.live"
	VIDEO_DANMU_PUB_TOPIC = "danmu.video"
)

type DanmuKMsg struct {
	RVID int64
	Data dao.DanmuData
}
