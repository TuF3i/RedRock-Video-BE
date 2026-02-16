namespace go livesvr

include "./usersvr.thrift"

struct LiveInfo {
  1: required i64 rvid
  2: required i64 ower_id
  3: required string title
  4: required string stream_name
  5: optional string upstream_password
  6: optional usersvr.RvUserInfo user_info
}

struct GetLiveListData {
  1: required i64 total
  2: required list<LiveInfo> lives
}

// 获取直播信息
struct GetLiveInfoReq {
  1: required i64 rvid
  2: required i64 uid
}

struct GetLiveInfoResp {
  1: required i64 status
  2: required string info
  3: optional LiveInfo data
}

// 获取直播列表
struct GetLiveListReq {
  1: required i32 page
  2: required i32 pageSize
}

struct GetLiveListResp {
  1: required i64 status
  2: required string info
  3: optional GetLiveListData data
}

// 开启直播
struct StartLiveReq {
  1: required i64 ower_id
  2: required string title
}

struct StartLiveResp {
  1: required i64 status
  2: required string info
  3: optional LiveInfo data
}

// 停止直播
struct StopLiveReq {
  1: required i64 rvid
  2: required i64 uid
}

struct StopLiveResp {
  1: required i64 status
  2: required string info
}

service LiveSvr {
  GetLiveInfoResp GetLiveInfo(1: GetLiveInfoReq req)
  GetLiveListResp GetLiveList(1: GetLiveListReq req)
  StartLiveResp StartLive(1: StartLiveReq req)
  StopLiveResp StopLive(1: StopLiveReq req)
}