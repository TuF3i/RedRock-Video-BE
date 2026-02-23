namespace go livesvr

struct LiveDetail {
  1: required i64 rvid
  2: required i64 ower_id
  3: required string title
  4: required string stream_name
  5: required string upstream_password
}

struct LiveListInfo {
  1: required i64 rvid
  2: required string title
  3: required string stream_name
  4: required UserInfo user_info
}

struct UserInfo {
  1: required i64 uid
  2: required string user_name
  3: required string avatar_url
}

struct GetLiveListData {
  1: required i64 total
  2: required list<LiveListInfo> lives
}

struct GetMyLiveListData {
  1: required i64 total
  2: required list<LiveDetail> lives
}

// 获取直播信息
struct GetLiveInfoReq {
  1: required i64 rvid
  2: required i64 uid
}

struct GetLiveInfoResp {
  1: required i64 status
  2: required string info
  3: optional LiveDetail data
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
  3: optional LiveDetail data
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

// SRS校验
struct SRSAuthReq {
  1: required i64 rvid
  2: required string password
}

struct SRSAuthResp {
  1: required i32 ok = 1
}

// 获取自己的直播
struct GetMyLiveListReq {
  1: required i64 uid
}

struct GetMyLiveListResp {
  1: required i64 status
  2: required string info
  3: optional GetMyLiveListData data
}

service LiveSvr {
  GetLiveInfoResp GetLiveInfo(1: GetLiveInfoReq req)
  GetLiveListResp GetLiveList(1: GetLiveListReq req)
  StartLiveResp StartLive(1: StartLiveReq req)
  StopLiveResp StopLive(1: StopLiveReq req)
  SRSAuthResp SRSAuth(1: SRSAuthReq req)
  GetMyLiveListResp GetMyLiveList(1: GetMyLiveListReq req)
}