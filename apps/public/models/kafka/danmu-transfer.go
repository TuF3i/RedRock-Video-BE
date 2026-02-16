package kafka

import "LiveDanmu/apps/public/models/dao"

const (
	LIVE_DANMU_PUB_TOPIC       = "danmusvr.live"
	LIVE_DANMU_BOARDCAST_TOPIC = "danmusvr.live.boardcast"
	VIDEO_DANMU_PUB_TOPIC      = "danmusvr.video"
)

// OPS
const (
	PUB_LIVE_DANMU = "live.danmu.pub"
	CLOSE_LIVE     = "live.off"
)

type DanmuKMsg struct {
	RVID int64
	OP   string
	Data dao.DanmuData
}
