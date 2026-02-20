# RedRock Video Backend API 文档

## 概述

RedRock Video Backend 是一个基于 Hertz + Kitex 的微服务架构项目，提供视频、直播、弹幕、用户等功能。

## 架构

- **Gateway 层**：HTTP API 网关
  - user_gateway：用户服务网关
  - video_gateway：视频服务网关
  - danmu_gateway：弹幕服务网关
  - live_gateway：直播服务网关

- **RPC 层**：微服务
  - usersvr：用户服务
  - videosvr：视频服务
  - danmusvr：弹幕服务
  - livesvr：直播服务

- **Consumer 层**：Kafka 消费者
  - videoDanmu：视频弹幕消费者
  - liveDanmu：直播弹幕消费者

---

## 接口文档

### 1. 用户服务 (user_gateway)

#### 1.1 GitHub OAuth2 登录

**接口路径：** `GET /user/oauth2/github`

**描述：** GitHub OAuth2 登录接口，重定向到 GitHub 授权页面

**请求参数：** 无

**响应：** 302 重定向到 GitHub

---

#### 1.2 GitHub OAuth2 回调

**接口路径：** `GET /user/oauth2/github/callback`

**描述：** GitHub OAuth2 回调接口，处理授权码并获取用户信息

**请求参数：**
- `code`：GitHub 授权码（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "accessToken": "xxx",
    "refreshToken": "yyy"
  }
}
```

---

#### 1.3 刷新 Access Token

**接口路径：** `POST /user/refresh`

**描述：** 使用 refresh_token 刷新 access_token

**请求头：**
- `Authorization: Bearer <refresh_token>`

**请求参数：** 无

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "accessToken": "xxx",
    "refreshToken": "yyy"
  }
}
```

---

#### 1.4 获取用户信息

**接口路径：** `GET /user/info`

**描述：** 获取当前登录用户的信息

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：** 无

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "uid": 123456,
    "userName": "username",
    "avatarUrl": "https://xxx",
    "bio": "user bio",
    "role": "user"
  }
}
```

---

#### 1.5 设置管理员权限

**接口路径：** `POST /user/admin/set`

**描述：** 设置用户为管理员（需要管理员权限）

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `uid`：用户ID（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 1.6 获取管理员列表

**接口路径：** `GET /user/admin/list`

**描述：** 获取管理员列表（需要管理员权限）

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `page`：页码（Query 参数）
- `pageSize`：每页数量（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "total": 100,
    "list": [
      {
        "uid": 123456,
        "userName": "admin",
        "avatarUrl": "https://xxx",
        "bio": "admin bio",
        "role": "admin"
      }
    ]
  }
}
```

---

#### 1.7 获取用户列表

**接口路径：** `GET /user/list`

**描述：** 获取用户列表（需要管理员权限）

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `page`：页码（Query 参数）
- `pageSize`：每页数量（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "total": 1000,
    "list": [
      {
        "uid": 123456,
        "userName": "user",
        "avatarUrl": "https://xxx",
        "bio": "user bio",
        "role": "user"
      }
    ]
  }
}
```

---

#### 1.8 用户登出

**接口路径：** `POST /user/logout`

**描述：** 用户登出，清除 access_token 和 refresh_token

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：** 无

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

### 2. 视频服务 (video_gateway)

#### 2.1 上传视频文件

**接口路径：** `POST /video/upload/:rvid`

**描述：** 上传视频文件到 MinIO

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `rvid`：视频ID（RVID编码）

**请求参数：**
- `file`：视频文件（Form Data）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 2.2 上传视频封面

**接口路径：** `POST /video/face/:rvid`

**描述：** 上传视频封面到 MinIO

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `rvid`：视频ID（RVID编码）

**请求参数：**
- `file`：封面文件（Form Data）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 2.3 创建视频记录

**接口路径：** `POST /video/add`

**描述：** 创建视频记录

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
```json
{
  "rvid": 123456,
  "title": "视频标题",
  "description": "视频描述"
}
```

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 2.4 删除视频

**接口路径：** `DELETE /video/delete/:rvid`

**描述：** 删除视频

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `rvid`：视频ID（RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 2.5 获取视频信息

**接口路径：** `GET /video/info/:rvid`

**描述：** 获取视频详细信息

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `rvid`：视频ID（RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "rvid": 123456,
    "uid": 789012,
    "title": "视频标题",
    "description": "视频描述",
    "faceUrl": "https://xxx",
    "viewNum": 1000,
    "inJudge": false
  }
}
```

---

#### 2.6 获取视频列表

**接口路径：** `GET /video/list`

**描述：** 获取视频列表

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `page`：页码（Query 参数）
- `pageSize`：每页数量（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "total": 100,
    "list": [
      {
        "rvid": 123456,
        "uid": 789012,
        "title": "视频标题",
        "description": "视频描述",
        "faceUrl": "https://xxx",
        "viewNum": 1000,
        "inJudge": false
      }
    ]
  }
}
```

---

#### 2.7 获取用户视频列表

**接口路径：** `GET /video/user/list`

**描述：** 获取当前用户的视频列表

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `page`：页码（Query 参数）
- `pageSize`：每页数量（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "total": 10,
    "list": [
      {
        "rvid": 123456,
        "uid": 789012,
        "title": "视频标题",
        "description": "视频描述",
        "faceUrl": "https://xxx",
        "viewNum": 1000,
        "inJudge": false
      }
    ]
  }
}
```

---

#### 2.8 获取待审核视频列表

**接口路径：** `GET /video/judge/list`

**描述：** 获取待审核视频列表（需要管理员权限）

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `page`：页码（Query 参数）
- `pageSize`：每页数量（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "total": 10,
    "list": [
      {
        "rvid": 123456,
        "uid": 789012,
        "title": "视频标题",
        "description": "视频描述",
        "faceUrl": "https://xxx",
        "viewNum": 0,
        "inJudge": true
      }
    ]
  }
}
```

---

#### 2.9 审核通过视频

**接口路径：** `POST /video/judge/access/:rvid`

**描述：** 审核通过视频（需要管理员权限）

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `rvid`：视频ID（RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 2.10 生成视频预签名URL

**接口路径：** `GET /video/url`

**描述：** 生成视频文件的预签名URL

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `rvid`：视频ID（Query 参数，RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "url": "https://minio.example.com/xxx?signature=yyy"
  }
}
```

---

#### 2.11 生成新的RVID

**接口路径：** `GET /video/new/rvid`

**描述：** 生成新的视频ID

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：** 无

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "rvid": 123456789,
    "raw": "abc123"
  }
}
```

---

#### 2.12 增加视频播放量

**接口路径：** `POST /video/view/incr`

**描述：** 增加视频播放量

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
```json
{
  "rvid": 123456
}
```

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

### 3. 弹幕服务 (danmu_gateway)

#### 3.1 获取热门弹幕

**接口路径：** `GET /danmu/hot/:rvid`

**描述：** 获取视频的热门弹幕

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `rvid`：视频ID（RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "danmuList": [
      {
        "danID": 123456,
        "userID": 789012,
        "content": "弹幕内容",
        "video": true,
        "videoTime": 123.45,
        "liveTime": 0,
        "user": {
          "uid": 789012,
          "userName": "username",
          "avatarUrl": "https://xxx",
          "bio": "user bio",
          "role": "user"
        }
      }
    ]
  }
}
```

---

#### 3.2 获取完整弹幕

**接口路径：** `GET /danmu/full/:rvid`

**描述：** 获取视频的完整弹幕列表

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `rvid`：视频ID（RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "danmuList": [
      {
        "danID": 123456,
        "userID": 789012,
        "content": "弹幕内容",
        "video": true,
        "videoTime": 123.45,
        "liveTime": 0,
        "user": {
          "uid": 789012,
          "userName": "username",
          "avatarUrl": "https://xxx",
          "bio": "user bio",
          "role": "user"
        }
      }
    ]
  }
}
```

---

#### 3.3 发送视频弹幕

**接口路径：** `POST /danmu/video`

**描述：** 发送视频弹幕

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
```json
{
  "rvid": 123456,
  "content": "弹幕内容",
  "videoTime": 123.45
}
```

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 3.4 删除弹幕

**接口路径：** `DELETE /danmu/delete/:danid`

**描述：** 删除弹幕

**请求头：**
- `Authorization: Bearer <access_token>`

**路径参数：**
- `danid`：弹幕ID

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 3.5 WebSocket 连接

**接口路径：** `GET /danmu/ws`

**描述：** WebSocket 连接，用于实时接收弹幕

**请求头：**
- `Authorization: Bearer <access_token>`

**协议：** WebSocket

---

### 4. 直播服务 (live_gateway)

#### 4.1 开始直播

**接口路径：** `POST /live/start`

**描述：** 开始直播

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
```json
{
  "title": "直播标题"
}
```

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "rvid": 123456,
    "raw": "abc123",
    "pushKey": "stream_key_xxx"
  }
}
```

---

#### 4.2 停止直播

**接口路径：** `POST /live/stop`

**描述：** 停止直播

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `rvid`：直播ID（Query 参数，RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": null
}
```

---

#### 4.3 获取直播信息

**接口路径：** `GET /live/info`

**描述：** 获取直播详细信息

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `rvid`：直播ID（Query 参数，RVID编码）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "rvid": 123456,
    "uid": 789012,
    "title": "直播标题",
    "streamer": {
      "uid": 789012,
      "userName": "username",
      "avatarUrl": "https://xxx",
      "bio": "user bio",
      "role": "user"
    },
    "status": "living"
  }
}
```

---

#### 4.4 获取直播列表

**接口路径：** `GET /live/list`

**描述：** 获取直播列表

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：**
- `page`：页码（Query 参数）
- `pageSize`：每页数量（Query 参数）

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "total": 10,
    "list": [
      {
        "rvid": 123456,
        "uid": 789012,
        "title": "直播标题",
        "streamer": {
          "uid": 789012,
          "userName": "username",
          "avatarUrl": "https://xxx",
          "bio": "user bio",
          "role": "user"
        },
        "status": "living"
      }
    ]
  }
}
```

---

#### 4.5 获取我的直播列表

**接口路径：** `GET /live/my/list`

**描述：** 获取当前用户的直播列表

**请求头：**
- `Authorization: Bearer <access_token>`

**请求参数：** 无

**响应格式：**
```json
{
  "status": 0,
  "info": "success",
  "data": {
    "list": [
      {
        "rvid": 123456,
        "uid": 789012,
        "title": "直播标题",
        "streamer": {
          "uid": 789012,
          "userName": "username",
          "avatarUrl": "https://xxx",
          "bio": "user bio",
          "role": "user"
        },
        "status": "living"
      }
    ]
  }
}
```

---

#### 4.6 SRS 直播认证

**接口路径：** `GET /live/srs/auth`

**描述：** SRS 直播流认证接口

**请求参数：**
- `stream`：流名称（Query 参数）
- `key`：推流密钥（Query 参数）

**响应格式：**
```json
{
  "code": 0
}
```

---

## 错误码说明

| 状态码 | 说明 |
|--------|------|
| 0 | 操作成功 |
| -1 | 无效的UID |
| -2 | 无效的用户名 |
| -3 | 无效的头像URL |
| -4 | 无效的简介 |
| -5 | 服务器内部错误 |
| -6 | 无效的RVID |
| -7 | 无效的标题 |
| -8 | 无效的描述 |
| -9 | 空的RVID |
| -10 | 你没有权限 |
| -11 | 视频不存在 |
| -12 | 直播不存在 |
| -13 | 无效的弹幕内容 |
| -14 | 弹幕不存在 |
| -15 | 请求验证失败 |

---

## 备注

- 所有需要认证的接口都需要在请求头中携带 `Authorization: Bearer <access_token>`
- RVID 是经过编码的视频ID，使用 `utils.RVIDDecoder()` 解码
- WebSocket 连接用于实时接收弹幕
- 直播使用 SRS 进行推流和拉流
