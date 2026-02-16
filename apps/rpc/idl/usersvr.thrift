namespace go usersvr

struct RvUserInfo {
  1: required i64 uid
  2: required string user_name
  3: required string avatar_url
  4: optional string bio
  5: optional string role
}

struct LoginData {
  1: required string accessToken
  2: required string refreshToken
}

struct GetUserListData {
  1: required i64 total
  2: required list<RvUserInfo> users
}

// 登录
struct LoginReq {
  1: required RvUserInfo user_info
}

struct LoginResp {
  1: required i64 status
  2: required string info
  3: optional LoginData data
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

// 设置用户为ADMIN权限
struct SetAdminRoleReq {
  1: required i64 uid
}

struct SetAdminRoleResp {
  1: required i64 status
  2: required string info
}

// 获取ADMIN用户
struct GetAdminerReq {
  1: required i32 page
  2: required i32 page_size
}

struct GetAdminerResp {
  1: required i64 status
  2: required string info
  3: optional GetUserListData data
}

// 获取普通用户
struct GetUsersReq {
  1: required i32 page
  2: required i32 page_size
}

struct GetUsersResp {
  1: required i64 status
  2: required string info
  3: optional GetUserListData data
}

// 登出账号
struct LogoutReq {
  1: required i64 uid
}

struct LogoutResp {
  1: required i64 status
  2: required string info
}

// 微服务方法
service UserSvr {
  LoginResp UserLogin(1: LoginReq req)
  RefreshResp RefreshToken(1: RefreshReq req)
  GetUserInfoResp GetUserInfo(1: GetUserInfoReq req)
  SetAdminRoleResp SetAdminRole(1: SetAdminRoleReq req)
  GetAdminerResp GetAdminer(1: GetAdminerReq req)
  GetUsersResp GetUsers(1: GetUsersReq req)
  LogoutResp Logout(1: LogoutReq req)
}