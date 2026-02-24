package models

type HStartLiveRequest struct {
	Title string `json:"title"`
}

type SRSCallback struct {
	ServerId  string `json:"server_id"`
	Action    string `json:"action"`
	ClientId  string `json:"client_id"`
	Ip        string `json:"ip"`
	Vhost     string `json:"vhost"`
	App       string `json:"app"`
	TcUrl     string `json:"tcUrl"`
	Stream    string `json:"stream"`
	Param     string `json:"param"`
	StreamUrl string `json:"stream_url"`
	StreamId  string `json:"stream_id"`
}
