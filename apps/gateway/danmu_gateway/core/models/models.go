package models

type WebsocketMsg struct {
	Msg      string                 `json:"msg"`
	Data     interface{}            `json:"data"`
	MataData map[string]interface{} `json:"mata_data"`
}
