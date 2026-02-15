package SRS

var (
	UPSTREAM_REJECT = 1
	UPSTREAM_ALLOW  = 0
)

// CallbackReq SRS 回调请求结构体
type CallbackReq struct {
	Action   string `form:"action"`
	ClientID string `form:"client_id"`
	IP       string `form:"ip"`
	VHost    string `form:"vhost"`
	App      string `form:"app"`
	Stream   string `form:"stream"`
	Param    string `form:"param"` // ?key=xxx
}

// CallbackResp SRS 回调响应结构体
type CallbackResp struct {
	Code int    `json:"code"` // 0=允许推流，非0=拒绝
	Msg  string `json:"msg"`  // 提示信息
}

type SRS struct {
}
