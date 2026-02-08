namespace go danmusvr

struct DanmuMsg { // 弹幕结构体
  1: required i64  room_id
  2: required i64  user_id
  3: required string content
  4: required string color   = "#FFFFFF"
  5: required i64  ts        // 毫秒
}

// 发送弹幕
struct PubResp { // 发送弹幕响应
  1: required i64 status
  2: required string info
}

struct PubReq { // 发送弹幕请求
  1: required DanmuMsg danmuMsg
}

// 获取弹幕
struct GetResp { // 获取弹幕的响应
  1: required i64 status
  2: required string info
  3: required list<DanmuMsg> data
}

struct GetReq { // 获取弹幕的请求
  1: required i64 BV
}

// 获取Top1000条
struct GetTopResp { // 获取Top1000弹幕的响应
  1: required i64 status
  2: required string info
  3: required list<DanmuMsg> data
}

struct GetTopReq { // 获取Top1000弹幕的请求
  1: required i64 BV
}

// 发送Live弹幕
struct PubLiveResp { // 发送Live弹幕响应
  1: required i64 status
  2: required string info
}

struct PubLiveReq { // 发送Live弹幕请求
  1: required DanmuMsg danmuMsg
}

// 删除Live弹幕
struct DelLiveResp { // 删除Live弹幕响应
  1: required i64 status
  2: required string info
}

struct DelLiveReq { // 删除Live弹幕请求
  1: required DanmuMsg danmuMsg
}

// 删除Video弹幕
struct DelResp { // 删除Video弹幕响应
  1: required i64 status
  2: required string info
}

struct DelReq { // 删除Video弹幕请求
  1: required DanmuMsg danmuMsg
}

service DanmuSvr { // 服务方法
  PubResp PubDanmu(1: PubReq req)
  PubLiveResp PubLiveDanmu(1: PubLiveReq req)
  GetResp GetDanmu(1: GetReq req)
  GetTopResp GetTop(1: GetTopReq req)
  DelLiveResp DelLiveDanmu(1: DelLiveReq req)
  DelResp DelDanmu(1: DelReq req)
}