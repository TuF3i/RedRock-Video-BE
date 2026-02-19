package models

type WebsocketMsg struct {
	Msg      string                 `json:"msg"`
	Data     interface{}            `json:"data"`
	MataData map[string]interface{} `json:"mata_data"`
}

type VideoDanmuReq struct {
	RVID    int64  `json:"rvid"` // Rv123456
	UID     int64  `json:"uid"`
	Content string `json:"content"`
	Color   string `json:"color"`
	Ts      int64  `json:"ts"`
}

type LiveDanmuReq struct {
	RVID    int64  `json:"rvid"`
	UID     int64  `json:"uid"`
	Content string `json:"content"`
	Color   string `json:"color"`
	Ts      int64  `json:"ts"`
}
