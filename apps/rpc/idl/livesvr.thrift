namespace go livesvr

struct LiveInfo {
  1: required i64 rvid
  2: required i64 ower_id
  3: required string title
  4: required string stream_name
  5: required string upstream_password
}

// 获取直播信息
struct GetLiveInfoReq {
  1: required i64 rvid
}

struct GetLiveInfoResp {
  1: required i64 status
  2: required string info
  3: optional LiveInfo data
}

// 获取直播列表
struct GetLiveListResp {
  1: required i64 status
  2: required string info
  3: optional list<LiveInfo> data
}

// 开启直播
struct StartLiveReq {
  1: required i64 ower_id
  2: required string stream_name
}

struct StartLiveResp {
  1: required i64 status
  2: required string info
  3: optional LiveInfo data
}

// 停止直播
struct StopLiveReq {
  1: required i64 rvid
}

struct StopLiveResp {
  1: required i64 status
  2: required string info
}

service LiveSvr {
  GetLiveInfoResp GetLiveInfo(1: GetLiveInfoReq req)
  GetLiveListResp GetLiveList()
  StartLiveResp StartLive(1: StartLiveReq req)
  StopLiveResp StopLive(1: StopLiveReq req)
}