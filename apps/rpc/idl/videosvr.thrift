namespace go videosvr

struct VideoMataInfo {
  1: required string face_key
  2: required string minio_key
  3: required string title
  4: required string description
  5: required i64    view_num
}

struct VideoUserInfo {
  1: required i64 author_id
  2: required string author_name
  3: required string avatar_url
}

struct VideoDetail { //
  1: required i64 rvid
  2: required VideoMataInfo mata_info
  3: required VideoUserInfo user_info

  99: required bool in_judge
}

struct VideoListData { //
  1: required i64 rvid
  2: required string title
  3: required string face_key
  4: required VideoUserInfo user_info

  99: required bool in_judge
}

struct AddVideoData {
  1: required i64 rvid
  2: required i64 uid
  3: required string title
  4: required string description
}

struct GetVideoListData {
  1: required i64 total
  2: required list<VideoListData> videos
}

// 添加视频
struct AddVideoReq {
  1: required AddVideoData add_video_data
}

struct AddVideoResp {
  1: required i64 status
  2: required string info
}

// 删除视频
struct DelVideoReq {
  1: required i64 rvid
  2: required i64 uid
  3: required string role
}

struct DelVideoResp {
  1: required i64 status
  2: required string info
}

// 审核通过
struct JudgeAccessReq {
  1: required i64 rvid
}

struct JudgeAccessResp {
  1: required i64 status
  2: required string info
}

// 获取视频列表
struct GetVideoListReq {
  1: required i32 page
  2: required i32 page_size
}

struct GetVideoListResp {
  1: required i64 status
  2: required string info
  3: optional GetVideoListData data
}

// 获取预签名链接
struct GetPreSignedUrlReq {
  1: required i64 rvid
  2: required i64 uid
  3: required string role
}

struct GetPreSignedUrlResp {
  1: required i64 status
  2: required string info
  3: optional string data
}

// 获取审核列表
struct GetJudgeListReq {
  1: required i32 page
  2: required i32 page_size
}

struct GetJudgeListResp {
  1: required i64 status
  2: required string info
  3: optional GetVideoListData data
}

// 获取我的视频
struct GetMyVideoListReq {
  1: required i32 page
  2: required i32 page_size
  3: required i64 uid
}

struct GetMyVideoListResp {
  1: required i64 status
  2: required string info
  3: optional GetVideoListData data
}

// 视频播放量递增
struct InnocentViewNumReq {
  1: required i64 rvid
}

struct InnocentViewNumResp {
  1: required i64 status
  2: required string info
}

// 获取VideoDetail
struct GetVideoDetailReq {
  1: required i64 rvid
}

struct GetVideoDetailResp {
  1: required i64 status
  2: required string info
  3: optional VideoDetail data
}

service VideoSvr {
  AddVideoResp AddVideo(1: AddVideoReq req)
  DelVideoResp DelVideo(1: DelVideoReq req)
  JudgeAccessResp JudgeAccess(1: JudgeAccessReq req)
  GetJudgeListResp GetJudgeList(1: GetJudgeListReq req)
  GetVideoListResp GetVideoList(1: GetVideoListReq req)
  GetPreSignedUrlResp GetPreSignedUrl(1: GetPreSignedUrlReq req)
  GetMyVideoListResp GetMyVideoList(1: GetMyVideoListReq req)
  InnocentViewNumResp InnocentViewNum(1: InnocentViewNumReq req)
  GetVideoDetailResp GetVideoDetail(1: GetVideoDetailReq req)
}