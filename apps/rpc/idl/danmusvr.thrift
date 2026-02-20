namespace go danmusvr

struct PubDanmuData { // 弹幕结构体
  1: required i64  dan_id
  2: required i64  rvid
  3: required i64  uid
  4: required string content
  5: required string color
  6: required i64 time_stamp
}

struct GetDanmuData {
  1: required i64    dan_id
  2: required i64    rvid
  3: required string content
  4: required string color
  5: required i64 time_stamp
  6: required UserInfo user_info
}

struct UserInfo {
  1: required i64 uid
  2: required string user_name
  3: required string avatar_url
}

// 发送弹幕
struct PubVideoResp { // 发送弹幕响应
  1: required i64 status
  2: required string info
}

struct PubVideoReq { // 发送弹幕请求
  1: required PubDanmuData danmuMsg
}

// 获取弹幕
struct GetFullResp { // 获取弹幕的响应
  1: required i64 status
  2: required string info
  3: optional list<GetDanmuData> data
}

struct GetFullReq { // 获取弹幕的请求
  1: required i64 rvid
}

// 获取Top1000条
struct GetTopResp { // 获取Top1000弹幕的响应
  1: required i64 status
  2: required string info
  3: optional list<GetDanmuData> data
}

struct GetTopReq { // 获取Top1000弹幕的请求
  1: required i64 rvid
}

// 发送Live弹幕
struct PubLiveResp { // 发送Live弹幕响应
  1: required i64 status
  2: required string info
}

struct PubLiveReq { // 发送Live弹幕请求
  1: required PubDanmuData danmuMsg
}

// 删除Live弹幕
struct DelLiveResp { // 删除Live弹幕响应
  1: required i64 status
  2: required string info
}

struct DelLiveReq { // 删除Live弹幕请求
  1: required i64 dan_id
}

// 删除Video弹幕
struct DelResp { // 删除Video弹幕响应
  1: required i64 status
  2: required string info
}

struct DelReq { // 删除Video弹幕请求
  1: required i64 dan_id
  2: required i64 uid
}

service DanmuSvr { // 服务方法
  PubVideoResp PubVideoDanmu(1: PubVideoReq req)
  PubLiveResp PubLiveDanmu(1: PubLiveReq req)
  GetFullResp GetDanmu(1: GetFullReq req)
  GetTopResp GetTop(1: GetTopReq req)
  DelLiveResp DelLiveDanmu(1: DelLiveReq req)
  DelResp DelDanmu(1: DelReq req)
}