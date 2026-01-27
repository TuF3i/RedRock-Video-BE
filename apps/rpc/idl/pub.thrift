namespace go pub

struct DanmuMsg {
  1: required i64  room_id
  2: required i64  user_id
  3: required string content
  4: required string color   = "#FFFFFF"
  5: required i64  ts        // 毫秒
}

struct PubResp {
  1: required i64 status
  2: required string info
}

struct PubReq {
  1: required DanmuMsg danmuMsg
}

service Pub {
    PubResp pub(1: PubReq req)
}