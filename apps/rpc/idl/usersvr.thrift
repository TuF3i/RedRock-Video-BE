namespace go usersvr

struct RvUserInfo {
  1: required i64 uid
  2: required string user_name
  3: required string avatar_url
  4: required string bio
}

// 登录
struct LoginReq {
  1: required RvUserInfo user_info
}

struct LoginResp {
  1: required i64 status
  2: required string info
  3: optional string data
}

// 刷新AccessToken
struct RefreshReq {
  1: required string refreshToken
}

struct RefreshResp {
  1: required i64 status
  2: required string info
  3: optional string data // accessToken
}

// 获取用户信息
struct GetUserInfoReq {
  1: required i64 uid
}

struct GetUserInfoResp {
  1: required i64 status
  2: required string info
  3: optional RvUserInfo data
}

// 微服务方法
service UserSvr {
  LoginResp UserLogin(1: LoginReq req)
  RefreshResp RefreshToken(1: RefreshReq req)
  GetUserInfoResp GetUserInfo(1: GetUserInfoReq req)
}