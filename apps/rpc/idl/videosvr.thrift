namespace go videosvr

struct VideoInfo {
  1: required i64 rvid
  2: required string face_url
  3: required string minio_key    // 注意：json tag 是 m3u8_url，字段名是 MinioKey
  4: required string title
  5: required string description
  6: required i64 view_num

  7: required bool use_face
  8: required bool in_judge
}

// 添加视频
struct AddVideoReq {
  1: required VideoInfo video_info
}

struct AddVideoResp {
  1: required i64 status
  2: required string info
}

// 删除视频
struct DelVideoReq {
  1: required i64 rvid
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
  3: required list<VideoInfo> data
}

// 获取预签名链接
struct GetPreSignedUrlReq {
  1: required string rvid
  2: required i64 uid
  3: required string role
}

struct GetPreSignedUrlResp {
  1: required i64 status
  2: required string info
  3: required string url
}

service VideoSvr {
  AddVideoResp AddVideo(1: AddVideoReq req)
  DelVideoResp DelVideo(1: DelVideoReq req)
  JudgeAccessResp JudgeAccess(1: JudgeAccessReq req)
  GetVideoListResp GetVideoList(1: GetPreSignedUrlReq req)
  GetPreSignedUrlResp GetPreSignedUrl(1: GetPreSignedUrlReq req)
}