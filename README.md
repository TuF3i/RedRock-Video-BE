# RedRock-Video-BE

> **红岩网校后端寒假考核项目**

## 1. 业务模块划分

- **用户模块**
- **视频模块**
- **直播模块**
- **弹幕模块**

## 2. 架构图

![arch](D:\Project\RedRock-Video-BE\docs\image\arch.png)

## 3. 接口设计

### 3.0 公共接口

#### 3.0.1 GitHub OAuth2 跳转

```
GET /user/auth
```

- 请求方法：`GET`

- 请求体为`null`

- 响应体data为`null`


#### 3.0.2 GitHub OAuth2 回调

```
GET /user/auth/callback
```

- 请求方法：`GET`

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

  - 字段含义：

    | 字段          | 数据类型 | 备注                                     |
    | ------------- | -------- | ---------------------------------------- |
    | accessToken   | String   | 认证Token,用于操作鉴权                   |
    | refreshToken  | String   | 刷新Token,用于access_token过期后进行刷新 |

#### 3.0.3 刷新AccessToken

```
GET /user/refresh
```

- 请求方法: `GET`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Refresh Token  |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  ```

  - 字段含义：

    | 字段          | 数据类型 | 备注                                     |
    | ------------- | -------- | ---------------------------------------- |
    | data          | String   | 新的认证Token,用于操作鉴权                |


### 3.1 用户管理模块 - 用户

#### 3.1.1 获取用户信息

```
GET /user/info/user
```

- 请求方法：`GET`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "uid": 122974688,
      "user_name": "TuF3i",
      "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4",
      "bio": "@DaWesen 😋",
      "role": "jwt_role_admin"
  }
  ```

  - 字段含义：

    | 字段       | 数据类型 | 备注         |
    | ---------- | -------- | ------------ |
    | uid        | Integer  | 用户ID       |
    | user_name  | String   | 用户名       |
    | avatar_url | String   | 用户头像URL  |
    | bio        | String   | 用户简介     |
    | role       | String   | 用户权限角色 |

#### 3.1.2 获取用户列表

```
GET /user/info/users
```

- 请求方法：`GET`

- 请求参数：

  | 名称      | 位置  | 类型   | 必选 | 备注     |
  | --------- | ----- | ------ | ---- | -------- |
  | page      | query | string | 否   | 页码     |
  | pageSize  | query | string | 否   | 每页条数 |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "total": 1,
      "users": [
          {
              "uid": 122974688,
              "user_name": "TuF3i",
              "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4",
              "bio": "@DaWesen 😋",
              "role": "jwt_role_admin"
          }
      ]
  }
  ```

  - 字段含义：

    | 字段       | 数据类型 | 备注         |
    | ---------- | -------- | ------------ |
    | total      | Integer  | 总条数       |
    | users      | List     | 用户列表     |
    | uid        | Integer  | 用户ID       |
    | user_name  | String   | 用户名       |
    | avatar_url | String   | 用户头像URL  |
    | bio        | String   | 用户简介     |
    | role       | String   | 用户权限角色 |

#### 3.1.3 获取管理员列表

```
GET /user/info/adminer
```

- 请求方法：`GET`

- 请求参数：

  | 名称      | 位置  | 类型   | 必选 | 备注     |
  | --------- | ----- | ------ | ---- | -------- |
  | page      | query | string | 否   | 页码     |
  | pageSize  | query | string | 否   | 每页条数 |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "total": 1,
      "users": [
          {
              "uid": 122974688,
              "user_name": "TuF3i",
              "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4",
              "bio": "@DaWesen 😋",
              "role": "jwt_role_admin"
          }
      ]
  }
  ```

  - 字段含义：

    | 字段       | 数据类型 | 备注         |
    | ---------- | -------- | ------------ |
    | total      | Integer  | 总条数       |
    | users      | List     | 用户列表     |
    | uid        | Integer  | 用户ID       |
    | user_name  | String   | 用户名       |
    | avatar_url | String   | 用户头像URL  |
    | bio        | String   | 用户简介     |
    | role       | String   | 用户权限角色 |

#### 3.1.4 设置管理员权限

```
GET /user/set/adminer
```

- 请求方法：`GET`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注   |
  | ---- | ----- | ------ | ---- | ------ |
  | uid  | query | string | 否   | 用户ID |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段为`null`

#### 3.1.5 用户登出

```
GET /user/auth/logout
```

- 请求方法：`GET`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段为`null`


### 3.2 视频管理模块 - 用户

#### 3.2.1 获取视频列表

```
GET /video/list
```

- 请求方法：`GET`

- 请求参数：

  | 名称      | 位置  | 类型   | 必选 | 备注     |
  | --------- | ----- | ------ | ---- | -------- |
  | page      | query | string | 否   | 页码     |
  | pageSize  | query | string | 否   | 每页条数 |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "total": 1,
      "videos": [
          {
              "rvid": 2553177934,
              "title": "GuGu GaGa !!!",
              "face_key": "Rv2553177934",
              "user_info": {
                  "author_id": 122974688,
                  "author_name": "TuF3i",
                  "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4"
              },
              "in_judge": false
          }
      ]
  }
  ```

  - 字段含义：

    | 字段          | 数据类型 | 备注           |
    | ------------- | -------- | -------------- |
    | total         | Integer  | 总条数         |
    | videos        | List     | 视频列表       |
    | rvid          | Integer  | 视频ID         |
    | title         | String   | 视频标题       |
    | face_key      | String   | 封面minioKey   |
    | user_info     | Object   | 用户信息       |
    | author_id     | Integer  | 作者ID         |
    | author_name   | String   | 作者名称       |
    | avatar_url    | String   | 头像URL        |
    | in_judge      | Boolean  | 是否在审核     |

#### 3.2.2 获取我的视频列表

```
GET /video/list/my
```

- 请求方法：`GET`

- 请求参数：

  | 名称      | 位置  | 类型   | 必选 | 备注     |
  | --------- | ----- | ------ | ---- | -------- |
  | page      | query | string | 是   | 页码     |
  | pageSize  | query | string | 是   | 每页条数 |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "total": 1,
      "videos": [
          {
              "rvid": 2553177934,
              "title": "GuGu GaGa !!!",
              "face_key": "Rv2553177934",
              "user_info": {
                  "author_id": 122974688,
                  "author_name": "TuF3i",
                  "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4"
              },
              "in_judge": true
          }
      ]
  }
  ```

  - 字段含义：

    | 字段          | 数据类型 | 备注             |
    | ------------- | -------- | ---------------- |
    | total         | Integer  | 总条数           |
    | videos        | List     | 视频列表         |
    | rvid          | Integer  | 视频ID           |
    | title         | String   | 视频标题         |
    | face_key      | String   | 封面minioKey     |
    | user_info     | Object   | 用户信息         |
    | author_id     | Integer  | 作者ID           |
    | author_name   | String   | 作者名称         |
    | avatar_url    | String   | 头像URL          |
    | in_judge      | Boolean  | 是否处于审核中   |

#### 3.2.3 获取新RVID

```
GET /video/new/rvid
```

- 请求方法：`GET`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  3925799209
  ```

  - 字段含义：

    | 字段 | 数据类型 | 备注     |
    | ---- | -------- | -------- |
    | data | Integer  | 新的RVID |

#### 3.2.4 查看视频详情

```
GET /video/{rvid}/detail
```

- 请求方法：`GET`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "rvid": 2553177934,
      "mata_info": {
          "face_key": "Rv2553177934",
          "minio_key": "Rv2553177934",
          "title": "GuGu GaGa !!!",
          "description": "Out",
          "view_num": 0
      },
      "user_info": {
          "author_id": 122974688,
          "author_name": "TuF3i",
          "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4"
      },
      "in_judge": true
  }
  ```

  - 字段含义：

    | 字段          | 数据类型 | 备注           |
    | ------------- | -------- | -------------- |
    | rvid          | Integer  | 视频ID         |
    | mata_info     | Object   | 元数据信息     |
    | face_key      | String   | 封面minioKey   |
    | minio_key     | String   | 视频minioKey   |
    | title         | String   | 视频标题       |
    | description   | String   | 视频描述       |
    | view_num      | Integer  | 观看次数       |
    | user_info     | Object   | 用户信息       |
    | author_id     | Integer  | 作者ID         |
    | author_name   | String   | 作者名称       |
    | avatar_url    | String   | 头像URL        |
    | in_judge      | Boolean  | 是否在审核     |

#### 3.2.5 生成预签名URL

```
GET /video/{rvid}/play-url
```

- 请求方法：`GET`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  "http://101.36.123.131:9000/video/Rv2553177934?X-Amz-Algorithm=AWS4-HMAC-SHA256..."
  ```

  - 字段含义：

    | 字段 | 数据类型 | 备注           |
    | ---- | -------- | -------------- |
    | data | String   | 视频播放URL    |

#### 3.2.6 发布新视频

```
POST /video/new
```

- 请求方法：`POST`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体JSON:

  ```json
  {
      "rvid": 2553177934,
      "title": "GuGu GaGa !!!",
      "description": "Out"
  }
  ```

  - 字段含义：

    | 字段         | 数据类型 | 备注                     |
    | ------------ | -------- | ------------------------ |
    | rvid         | Integer  | 视频ID                   |
    | title        | String   | 视频标题                 |
    | description  | String   | 视频简介                 |

- 响应体data字段为`null`

#### 3.2.7 上传视频

```
POST /video/{rvid}/upload/video
```

- 请求方法：`POST`

- 请求参数：

  | 名称 | 位置  | 类型          | 必选 | 备注     |
  | ---- | ----- | ------------- | ---- | -------- |
  | rvid | path  | string        | 是   | 视频ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`multipart/form-data`:

  | 名称 | 类型          | 必选 | 备注           |
  | ---- | ------------- | ---- | -------------- |
  | file | file(binary) | 否   | 视频文件，小于1GB |

- 响应体data字段为`null`

#### 3.2.8 上传封面

```
POST /video/{rvid}/upload/cover
```

- 请求方法：`POST`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`multipart/form-data`:

  | 名称 | 类型          | 必选 | 备注     |
  | ---- | ------------- | ---- | -------- |
  | file | file(binary) | 否   | 封面文件 |

- 响应体data字段为`null`

#### 3.2.9 删除视频

```
DELETE /video/{rvid}
```

- 请求方法：`DELETE`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段为`null`

#### 3.2.10 播放自增

```
PATCH /video/{rvid}/innocent
```

- 请求方法：`PATCH`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求体为`null`

- 响应体data字段为`null`


### 3.3 视频管理模块 - Admin

#### 3.3.1 查看审核列表

```
GET /video/judge/list
```

- 请求方法：`GET`

- 请求参数：

  | 名称      | 位置  | 类型   | 必选 | 备注     |
  | --------- | ----- | ------ | ---- | -------- |
  | page      | query | string | 否   | 页码     |
  | pageSize  | query | string | 否   | 每页条数 |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "total": 1,
      "videos": [
          {
              "rvid": 2553177934,
              "title": "GuGu GaGa !!!",
              "face_key": "Rv2553177934",
              "user_info": {
                  "author_id": 0,
                  "author_name": "",
                  "avatar_url": ""
              },
              "in_judge": true
          }
      ]
  }
  ```

  - 字段含义：

    | 字段          | 数据类型 | 备注           |
    | ------------- | -------- | -------------- |
    | total         | Integer  | 总条数         |
    | videos        | List     | 视频列表       |
    | rvid          | Integer  | 视频ID         |
    | title         | String   | 视频标题       |
    | face_key      | String   | 封面minioKey  |
    | user_info     | Object   | 用户信息       |
    | author_id     | Integer  | 作者ID         |
    | author_name   | String   | 作者名称       |
    | avatar_url    | String   | 作者头像Url    |
    | in_judge      | Boolean  | 是否处于审核状态 |

#### 3.3.2 通过审核

```
PATCH /video/judge/{rvid}
```

- 请求方法：`PATCH`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段为`null`


### 3.4 直播管理模块 - 用户

#### 3.4.1 获取自己的直播列表

```
GET /live/list/my
```

- 请求方法：`GET`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "total": 2,
      "lives": [
          {
              "rvid": 2559677982,
              "ower_id": 122974688,
              "title": "GuGu GaGa!!!",
              "stream_name": "Rv2559677982",
              "upstream_password": "e044ce44"
          },
          {
              "rvid": 2837762687,
              "ower_id": 122974688,
              "title": "GuGu GaGa!!!",
              "stream_name": "Rv2837762687",
              "upstream_password": "5e998a11"
          }
      ]
  }
  ```

  - 字段含义：

    | 字段              | 数据类型 | 备注             |
    | ----------------- | -------- | ---------------- |
    | total             | Integer  | 总条数           |
    | lives             | List     | 直播列表         |
    | rvid              | Integer  | 直播ID           |
    | ower_id           | Integer  | 主播ID           |
    | title             | String   | 直播标题         |
    | stream_name       | String   | 直播名           |
    | upstream_password | String   | 推流密码         |

#### 3.4.2 获取自己的直播信息

```
GET /live/info
```

- 请求方法：`GET`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | query | string | 否   | 视频ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "rvid": 2837762687,
      "ower_id": 122974688,
      "title": "GuGu GaGa!!!",
      "stream_name": "Rv2837762687",
      "upstream_password": "5e998a11"
  }
  ```

  - 字段含义：

    | 字段              | 数据类型 | 备注             |
    | ----------------- | -------- | ---------------- |
    | rvid              | Integer  | 直播ID           |
    | ower_id           | Integer  | 主播ID           |
    | title             | String   | 直播标题         |
    | stream_name       | String   | 直播名称         |
    | upstream_password | String   | 推流密码         |

#### 3.4.3 获取所有直播

```
GET /live/list
```

- 请求方法：`GET`

- 请求参数：

  | 名称      | 位置  | 类型   | 必选 | 备注       |
  | --------- | ----- | ------ | ---- | ---------- |
  | page      | query | string | 否   | 页码       |
  | pageSize  | query | string | 否   | 每页的条数 |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "total": 2,
      "lives": [
          {
              "rvid": 2559677982,
              "title": "GuGu GaGa!!!",
              "stream_name": "Rv2559677982",
              "user_info": {
                  "uid": 122974688,
                  "user_name": "TuF3i",
                  "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4"
              }
          },
          {
              "rvid": 2837762687,
              "title": "GuGu GaGa!!!",
              "stream_name": "Rv2837762687",
              "user_info": {
                  "uid": 122974688,
                  "user_name": "TuF3i",
                  "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4"
              }
          }
      ]
  }
  ```

  - 字段含义：

    | 字段          | 数据类型 | 备注             |
    | ------------- | -------- | ---------------- |
    | total         | Integer  | 总条数           |
    | lives         | List     | 直播列表         |
    | rvid          | Integer  | 直播ID           |
    | title         | String   | 直播标题         |
    | stream_name   | String   | 直播名称         |
    | user_info     | Object   | 用户信息         |
    | uid           | Integer  | 用户ID           |
    | user_name     | String   | 用户名称         |
    | avatar_url    | String   | 用户头像URL      |

#### 3.4.4 开启直播

```
POST /live/start
```

- 请求方法：`POST`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体JSON:

  ```json
  {
      "title": "GuGu GaGa!!!"
  }
  ```

  - 字段含义：

    | 字段  | 数据类型 | 备注       |
    | ----- | -------- | ---------- |
    | title | String   | 直播标题   |

- 响应体data字段结构：

  ```json
  {
      "rvid": 3959610205,
      "ower_id": 122974688,
      "title": "GuGu GaGa!!!",
      "stream_name": "Rv3959610205",
      "upstream_password": "f9df954b"
  }
  ```

  - 字段含义：

    | 字段              | 数据类型 | 备注             |
    | ----------------- | -------- | ---------------- |
    | rvid              | Integer  | 直播ID           |
    | ower_id           | Integer  | 主播ID           |
    | title             | String   | 直播标题         |
    | stream_name       | String   | 直播名称         |
    | upstream_password | String   | 推流密码         |

#### 3.4.5 关闭直播

```
GET /live/stop
```

- 请求方法：`GET`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | query | string | 否   | 直播ID   |

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体为`null`

- 响应体data字段为`null`

#### 3.4.6 SRS媒体服务器鉴权接口

```
POST /live/srs/auth
```

- 请求方法：`POST`

- 请求参数：

  | 名称    | 位置  | 类型   | 必选 | 备注           |
  | ------- | ----- | ------ | ---- | -------------- |
  | key     | query | string | 否   | 推流密码       |
  | action  | query | string | 否   | 操作类型       |
  | tcUrl   | query | string | 否   | 推流地址       |
  | stream  | query | string | 否   | 直播名称       |
  | param   | query | string | 否   | 额外参数       |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  {
      "code": 1
  }
  ```

  - 字段含义：

    | 字段  | 数据类型 | 备注                   |
    | ----- | -------- | ---------------------- |
    | code  | Integer  | 验证状态(1:不可推流;0:可以推流) |


### 3.5 弹幕管理模块 - 用户 (Auth-required)

#### 3.5.1 获取视频所有弹幕

```
GET /danmu/full/{rvid}
```

- 请求方法：`GET`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  [
      {
          "dan_id": 32726716,
          "rvid": 2553177934,
          "content": "GOGO GAGA!!!",
          "color": "#FFFFFF",
          "time_stamp": 0,
          "user_info": {
              "uid": 122974688,
              "user_name": "TuF3i",
              "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4"
          }
      }
  ]
  ```

  - 字段含义：

    | 字段        | 数据类型 | 备注         |
    | ----------- | -------- | ------------ |
    | dan_id      | Integer  | 弹幕ID       |
    | rvid        | Integer  | 视频ID       |
    | content     | String   | 内容         |
    | color       | String   | 颜色         |
    | time_stamp  | Integer  | 时间戳(毫秒) |
    | user_info   | Object   | 用户信息     |
    | uid         | Integer  | 用户ID       |
    | user_name   | String   | 用户名称     |
    | avatar_url  | String   | 用户头像URL  |

#### 3.5.2 获取热门弹幕(前1000条)

```
GET /danmu/hot/{rvid}
```

- 请求方法：`GET`

- 请求参数：

  | 名称 | 位置  | 类型   | 必选 | 备注     |
  | ---- | ----- | ------ | ---- | -------- |
  | rvid | path  | string | 是   | 视频ID   |

- 请求体为`null`

- 响应体data字段结构：

  ```json
  [
      {
          "dan_id": 32726716,
          "rvid": 2553177934,
          "content": "GOGO GAGA!!!",
          "color": "#FFFFFF",
          "time_stamp": 0,
          "user_info": {
              "uid": 122974688,
              "user_name": "TuF3i",
              "avatar_url": "https://avatars.githubusercontent.com/u/122974688?v=4"
          }
      }
  ]
  ```

  - 字段含义：

    | 字段        | 数据类型 | 备注         |
    | ----------- | -------- | ------------ |
    | dan_id      | Integer  | 弹幕ID       |
    | rvid        | Integer  | 视频ID       |
    | content     | String   | 内容         |
    | color       | String   | 颜色         |
    | time_stamp  | Integer  | 时间戳(毫秒) |
    | user_info   | Object   | 用户信息     |
    | uid         | Integer  | 用户ID       |
    | user_name   | String   | 用户名称     |
    | avatar_url  | String   | 用户头像URL  |

#### 3.5.3 发布视频弹幕

```
POST /danmu/video
```

- 请求方法：`POST`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体JSON:

  ```json
  {
      "rvid": 2553177934,
      "content": "1234",
      "color": "#FFFFFF",
      "ts": 0
  }
  ```

  - 字段含义：

    | 字段    | 数据类型 | 备注       |
    | ------- | -------- | ---------- |
    | rvid    | Integer  | 视频ID     |
    | content | String   | 内容       |
    | color   | String   | 颜色       |
    | ts      | Integer  | 时间戳     |

- 响应体data字段为`null`

#### 3.5.4 发布直播弹幕

```
POST /danmu/live
```

- 请求方法：`POST`

- 请求Header：

  | 字段          | 数据类型 | 备注           |
  | ------------- | -------- | -------------- |
  | Authorization | String   | Access Token   |

- 请求体JSON:

  ```json
  {
      "rvid": 123456,
      "content": "1234",
      "color": "#FFFFFF",
      "ts": 0
  }
  ```

  - 字段含义：

    | 字段    | 数据类型 | 备注       |
    | ------- | -------- | ---------- |
    | rvid    | Integer  | 直播ID     |
    | content | String   | 内容       |
    | color   | String   | 颜色       |
    | ts      | Integer  | 时间戳     |

- 响应体data字段为`null`

### 4. 依赖

| 依赖                             | 备注         |
| -------------------------------- | ------------ |
| **Hertz**                        | API引擎      |
| **Kitex**                        | RPC引擎      |
| **ZooKeeper**                    | 注册中心     |
| **PostgreSQL**                   | 关系型数据库 |
| **Redis Cluster**                | 缓存         |
| **Kafka**                        | 消息队列     |
| **MinIO**                        | 对象存储     |
| **SRS (Simple Realtime Server)** | 直播服务器   |
| **Zap**                          | 日志         |
| **JWT**                          | 鉴权         |
| **Grafana**                      | 可观测性     |
| **Grafana-Loki**                 | 日志服务器   |

### 5. 状态码

#### 5.1 **Gateway (网关层)**

| 状态码 | 英文信息 | 中文说明 |
|--------|---------|---------|
| 10200 | Operation Success | 执行成功 |
| 10001 | Empty JWT String | 空 JWT 字符串 |
| 10002 | JWT Not Registered In Redis | JWT 未注册或已过期 |
| 10003 | Validate Request Fail | 校验请求失败 |
| 10004 | Empty RVID | 空 RVID |
| 10005 | You Do Not Have Access | 你没有权限 |
| 10006 | Nil Github User Info | 空 Github 用户信息 |
| 500 | (Internal Error) | 服务器内部错误 |

---

#### 5.2 **DanmuSvr (弹幕服务)**

| 状态码 | 英文信息 | 中文说明 |
|--------|---------|---------|
| 0 | Operation Success | 执行成功 |
| 40001 | Invalid RoomID | 无效的房间 ID |
| 40002 | Invalid UserID | 无效的用户 ID |
| 40003 | Invalid Color | 无效的颜色 |
| 40004 | Invalid Content | 无效的内容 |
| 40005 | Invalid DanID | 无效的弹幕 ID |
| 40006 | Danmu Not Exist | 弹幕不存在 |
| 40007 | No Permission | 没有权限 |
| 49999 | (Server Internal Error) | 服务器内部错误 |

---

#### 5.3 **LiveSvr (直播服务)**

| 状态码 | 英文信息 | 中文说明 |
|--------|---------|---------|
| 0 | Operation Success | 执行成功 |
| 90001 | Invalid RVID | 无效的 RVID |
| 90002 | Invalid UID | 无效的用户 ID |
| 90003 | No Permission | 没有权限 |
| 90004 | Live Not Exist | 直播不存在 |
| 99999 | (Server Internal Error) | 服务器内部错误 |

#### 5.4 **UserSvr (用户服务)**

| 状态码 | 英文信息 | 中文说明 |
|--------|---------|---------|
| 0 | Operation Success | 执行成功 |
| 80001 | No User Exist | 用户不存在 |
| 80002 | Invalid RefreshToken | 无效的刷新令牌 |
| 80003 | Invalid UID | 无效的用户 ID |
| 80004 | Invalid User Name | 无效的用户名 |
| 80005 | Invalid AvatarURL | 无效的头像 URL |
| 80006 | Invalid Bio | 无效的简介 |
| 89999 | (Server Internal Error) | 服务器内部错误 |

---

#### 5.5 **VideoSvr (视频服务)**

| 状态码 | 英文信息 | 中文说明 |
|--------|---------|---------|
| 0 | Operation Success | 执行成功 |
| 50001 | Invalid RVID | 无效的 RVID |
| 50002 | Invalid FaceUrl | 无效的封面 URL |
| 50003 | Invalid MinioKey | 无效的 MinIO 键 |
| 50004 | Invalid Description | 无效的描述 |
| 50005 | Invalid Uid | 无效的用户 ID |
| 50006 | Invalid Title | 无效的标题 |
| 50007 | Invalid AuthorName | 无效的作者名 |
| 50008 | No Permission | 没有权限 |
| 59999 | (Server Internal Error) | 服务器内部错误 |

----

#### 5.6 **状态码范围说明**

| 服务 | 范围 | 说明 |
|------|------|------|
| Gateway | 10000-10999 | 网关层通用错误 |
| DanmuSvr | 40000-40999 | 弹幕服务错误 |
| LiveSvr | 90000-90999 | 直播服务错误 |
| UserSvr | 80000-80999 | 用户服务错误 |
| VideoSvr | 50000-50999 | 视频服务错误 |
| 通用成功 | 0 | 操作成功 |
| 内部错误 | 500, 59999, 89999, 49999, 99999 | 服务器内部错误 |

### 6. 命令参数

#### 6.1 所有可用命令

| Unit     | Name       | 说明           | 完整命令示例                 |
| -------- | ---------- | -------------- | ---------------------------- |
| gateway  | danmu      | 弹幕网关       | `rv run gateway danmu`       |
| gateway  | live       | 直播网关       | `rv run gateway live`        |
| gateway  | user       | 用户网关       | `rv run gateway user`        |
| gateway  | video      | 视频网关       | `rv run gateway video`       |
| rpc      | danmusvr   | 弹幕 RPC 服务  | `rv run rpc danmusvr`        |
| rpc      | livesvr    | 直播 RPC 服务  | `rv run rpc livesvr`         |
| rpc      | usersvr    | 用户 RPC 服务  | `rv run rpc usersvr`         |
| rpc      | videosvr   | 视频 RPC 服务  | `rv run rpc videosvr`        |
| consumer | liveDanmu  | 直播弹幕消费者 | `rv run consumer liveDanmu`  |
| consumer | videoDanmu | 视频弹幕消费者 | `rv run consumer videoDanmu` |
| (init)   | db-init    | 数据库初始化   | `rv run db-init`             |

---

#### 6.2 杂项

```bash
# 查看可用命令
./rv help

# 查看 run 子命令帮助
./rv run help
```

### 7. 环境变量表格

#### 7.1 Danmu Gateway (弹幕网关)

| 环境变量名                         | 默认值        | 中文注释             |
| ---------------------------------- | ------------- | -------------------- |
| DANMU_GATEWAY_HERTZ_LISTENADDR     | 0.0.0.0       | Hertz 服务器监听地址 |
| DANMU_GATEWAY_HERTZ_LISTENPORT     | 8080          | Hertz 服务器监听端口 |
| DANMU_GATEWAY_HERTZ_MONITORINGPORT | 8081          | Hertz 监控端口       |
| DANMU_GATEWAY_ETCD_SERVICENAME     | zookeeper     | Etcd 服务名          |
| DANMU_GATEWAY_ETCD_NAMESPACE       | *             | Etcd 命名空间        |
| DANMU_GATEWAY_LOKI_SERVICE         | DANMU_GATEWAY | Loki 服务标签        |
| DANMU_GATEWAY_LOKI_ENV             | proc          | Loki 环境标识        |
| DANMU_GATEWAY_LOKI_LEVEL           | INFO          | Loki 日志级别        |
| DANMU_GATEWAY_REDIS_SERVICENAME    | redis         | Redis 服务名         |
| DANMU_GATEWAY_REDIS_NAMESPACE      | *             | Redis 命名空间       |
| DANMU_GATEWAY_REDIS_PASSWORD       | *             | Redis 密码           |
| DANMU_GATEWAY_KAFKA_SERVICENAME    | kafka         | Kafka 服务名         |
| DANMU_GATEWAY_KAFKA_NAMESPACE      | *             | Kafka 命名空间       |
| DANMU_GATEWAY_POD_UID              | (UUID)        | Pod 唯一标识         |

----

#### 7.2 Danmu RPC (弹幕 RPC)

| 环境变量名                  | 默认值    | 中文注释            |
| --------------------------- | --------- | ------------------- |
| DANMU_RPC_ETCD_SERVICENAME  | zookeeper | Etcd 服务名         |
| DANMU_RPC_ETCD_NAMESPACE    | *         | Etcd 命名空间       |
| DANMU_RPC_KAFKA_SERVICENAME | kafka     | Kafka 服务名        |
| DANMU_RPC_KAFKA_NAMESPACE   | *         | Kafka 命名空间      |
| DANMU_RPC_PGSQL_SERVICENAME | pgpool    | PostgreSQL 服务名   |
| DANMU_RPC_PGSQL_NAMESPACE   | *         | PostgreSQL 命名空间 |
| DANMU_RPC_PGSQL_USER        | root      | PostgreSQL 用户名   |
| DANMU_RPC_PGSQL_PASSWORD    | *         | PostgreSQL 密码     |
| DANMU_RPC_PGSQL_DBNAME      | rvideo    | PostgreSQL 数据库名 |
| DANMU_RPC_REDIS_SERVICENAME | redis     | Redis 服务名        |
| DANMU_RPC_REDIS_NAMESPACE   | *         | Redis 命名空间      |
| DANMU_RPC_REDIS_PASSWORD    | *         | Redis 密码          |
| DANMU_RPC_LOKI_SERVICE      | DANMU_RPC | Loki 服务标签       |
| DANMU_RPC_LOKI_ENV          | proc      | Loki 环境标识       |
| DANMU_RPC_LOKI_LEVEL        | INFO      | Loki 日志级别       |
| DANMU_RPC_POD_UID           | (UUID)    | Pod 唯一标识        |

----

#### 7.3 User Gateway (用户网关)

| 环境变量名                        | 默认值                                   | 中文注释             |
| --------------------------------- | ---------------------------------------- | -------------------- |
| USER_GATEWAY_HERTZ_LISTENADDR     | 0.0.0.0                                  | Hertz 服务器监听地址 |
| USER_GATEWAY_HERTZ_LISTENPORT     | 8080                                     | Hertz 服务器监听端口 |
| USER_GATEWAY_HERTZ_MONITORINGPORT | 8081                                     | Hertz 监控端口       |
| USER_GATEWAY_ETCD_SERVICENAME     | zookeeper                                | Etcd 服务名          |
| USER_GATEWAY_ETCD_NAMESPACE       | *                                        | Etcd 命名空间        |
| USER_GATEWAY_LOKI_SERVICE         | USER_GATEWAY                             | Loki 服务标签        |
| USER_GATEWAY_LOKI_ENV             | proc                                     | Loki 环境标识        |
| USER_GATEWAY_LOKI_LEVEL           | INFO                                     | Loki 日志级别        |
| USER_GATEWAY_REDIS_SERVICENAME    | redis                                    | Redis 服务名         |
| USER_GATEWAY_REDIS_NAMESPACE      | *                                        | Redis 命名空间       |
| USER_GATEWAY_REDIS_PASSWORD       | *                                        | Redis 密码           |
| USER_GATEWAY_POD_UID              | (UUID)                                   | Pod 唯一标识         |
| USER_GATEWAY_CLIENT_ID            | *                                        | OAuth2 客户端ID      |
| USER_GATEWAY_CLIENT_SECRET        | *                                        | OAuth2 客户端密钥    |
| USER_GATEWAY_REDIRECT_URL         | http://127.0.0.1:8080/user/auth/callback | OAuth2 回调URL       |

----

#### 7.4 USER_RPC (用户RPC服务)

| 环境变量名                 | 默认值    | 中文注释                 |
| -------------------------- | --------- | ------------------------ |
| USER_RPC_ETCD_SERVICENAME  | zookeeper | Etcd 服务名              |
| USER_RPC_ETCD_NAMESPACE    | *         | Etcd 命名空间            |
| USER_RPC_PGSQL_SERVICENAME | pgpool    | PostgreSQL 服务名        |
| USER_RPC_PGSQL_NAMESPACE   | *         | PostgreSQL 命名空间      |
| USER_RPC_PGSQL_USER        | root      | PostgreSQL 用户名        |
| USER_RPC_PGSQL_PASSWORD    | *         | PostgreSQL 密码          |
| USER_RPC_PGSQL_DBNAME      | rvideo    | PostgreSQL 数据库名      |
| USER_RPC_REDIS_SERVICENAME | redis     | Redis 服务名             |
| USER_RPC_REDIS_NAMESPACE   | *         | Redis 命名空间           |
| USER_RPC_REDIS_PASSWORD    | *         | Redis 密码               |
| USER_RPC_LOKI_SERVICE      | USER_RPC  | Loki 服务标签            |
| USER_RPC_LOKI_ENV          | proc      | Loki 环境标识            |
| USER_RPC_LOKI_LEVEL        | INFO      | Loki 日志级别            |
| USER_RPC_ADMIN             | *         | 管理员用户 **Github ID** |
| USER_RPC_POD_UID           | (UUID)    | Pod 唯一标识             |

---

#### 7.5 VIDEO_GATEWAY (视频网关)

| 环境变量名                         | 默认值        | 中文注释             |
| ---------------------------------- | ------------- | -------------------- |
| VIDEO_GATEWAY_HERTZ_LISTENADDR     | 0.0.0.0       | Hertz 服务器监听地址 |
| VIDEO_GATEWAY_HERTZ_LISTENPORT     | 8080          | Hertz 服务器监听端口 |
| VIDEO_GATEWAY_HERTZ_MONITORINGPORT | 8081          | Hertz 监控端口       |
| VIDEO_GATEWAY_ETCD_SERVICENAME     | zookeeper     | Etcd 服务名          |
| VIDEO_GATEWAY_ETCD_NAMESPACE       | *             | Etcd 命名空间        |
| VIDEO_GATEWAY_LOKI_SERVICE         | VIDEO_GATEWAY | Loki 服务标签        |
| VIDEO_GATEWAY_LOKI_ENV             | proc          | Loki 环境标识        |
| VIDEO_GATEWAY_LOKI_LEVEL           | INFO          | Loki 日志级别        |
| VIDEO_GATEWAY_REDIS_SERVICENAME    | redis         | Redis 服务名         |
| VIDEO_GATEWAY_REDIS_NAMESPACE      | *             | Redis 命名空间       |
| VIDEO_GATEWAY_REDIS_PASSWORD       | *             | Redis 密码           |
| VIDEO_GATEWAY_MINIO_SERVICENAME    | minio         | MinIO 服务名         |
| VIDEO_GATEWAY_MINIO_NAMESPACE      | *             | MinIO 命名空间       |
| VIDEO_GATEWAY_MINIO_USESSL         | false         | MinIO 是否使用SSL    |
| VIDEO_GATEWAY_MINIO_ACCESSKEY      | *             | MinIO 访问密钥       |
| VIDEO_GATEWAY_MINIO_SECRETKEY      | *             | MinIO 秘密密钥       |
| VIDEO_GATEWAY_MINIO_BLANKETNAME    | video         | MinIO 视频存储桶名   |
| VIDEO_GATEWAY_MINIO_PICBLANKETNAME | videoface     | MinIO 封面存储桶名   |
| VIDEO_GATEWAY_POD_UID              | (UUID)        | Pod 唯一标识         |

---

#### 7.6 VIDEO_RPC (视频RPC服务)

| 环境变量名                     | 默认值    | 中文注释            |
| ------------------------------ | --------- | ------------------- |
| VIDEO_RPC_ETCD_SERVICENAME     | zookeeper | Etcd 服务名         |
| VIDEO_RPC_ETCD_NAMESPACE       | *         | Etcd 命名空间       |
| VIDEO_RPC_LOKI_SERVICE         | VIDEO_RPC | Loki 服务标签       |
| VIDEO_RPC_LOKI_ENV             | proc      | Loki 环境标识       |
| VIDEO_RPC_LOKI_LEVEL           | INFO      | Loki 日志级别       |
| VIDEO_RPC_REDIS_SERVICENAME    | redis     | Redis 服务名        |
| VIDEO_RPC_REDIS_NAMESPACE      | *         | Redis 命名空间      |
| VIDEO_RPC_REDIS_PASSWORD       | *         | Redis 密码          |
| VIDEO_RPC_MINIO_SERVICENAME    | minio     | MinIO 服务名        |
| VIDEO_RPC_MINIO_NAMESPACE      | *         | MinIO 命名空间      |
| VIDEO_RPC_MINIO_USESSL         | false     | MinIO 是否使用SSL   |
| VIDEO_RPC_MINIO_ACCESSKEY      | *         | MinIO 访问密钥      |
| VIDEO_RPC_MINIO_SECRETKEY      | *         | MinIO 秘密密钥      |
| VIDEO_RPC_MINIO_BLANKETNAME    | video     | MinIO 视频存储桶名  |
| VIDEO_RPC_MINIO_PICBLANKETNAME | videoface | MinIO 封面存储桶名  |
| VIDEO_RPC_PGSQL_SERVICENAME    | pgpool    | PostgreSQL 服务名   |
| VIDEO_RPC_PGSQL_NAMESPACE      | *         | PostgreSQL 命名空间 |
| VIDEO_RPC_PGSQL_USER           | root      | PostgreSQL 用户名   |
| VIDEO_RPC_PGSQL_PASSWORD       | *         | PostgreSQL 密码     |
| VIDEO_RPC_PGSQL_DBNAME         | rvideo    | PostgreSQL 数据库名 |
| VIDEO_RPC_POD_UID              | (UUID)    | Pod 唯一标识        |

---

#### 7.7 LIVE_GATEWAY (直播网关)

| 环境变量名                        | 默认值       | 中文注释             |
| --------------------------------- | ------------ | -------------------- |
| LIVE_GATEWAY_HERTZ_LISTENADDR     | 0.0.0.0      | Hertz 服务器监听地址 |
| LIVE_GATEWAY_HERTZ_LISTENPORT     | 8080         | Hertz 服务器监听端口 |
| LIVE_GATEWAY_HERTZ_MONITORINGPORT | 8099         | Hertz 监控端口       |
| LIVE_GATEWAY_ETCD_SERVICENAME     | zookeeper    | Etcd 服务名          |
| LIVE_GATEWAY_ETCD_NAMESPACE       | *            | Etcd 命名空间        |
| LIVE_GATEWAY_LOKI_SERVICE         | LIVE_GATEWAY | Loki 服务标签        |
| LIVE_GATEWAY_LOKI_ENV             | proc         | Loki 环境标识        |
| LIVE_GATEWAY_LOKI_LEVEL           | INFO         | Loki 日志级别        |
| LIVE_GATEWAY_REDIS_SERVIVENAME    | redis        | Redis 服务名         |
| LIVE_GATEWAY_REDIS_NAMESPACE      | *            | Redis 命名空间       |
| LIVE_GATEWAY_REDIS_PASSWORD       | *            | Redis 密码           |
| LIVE_GATEWAY_POD_UID              | (UUID)       | Pod 唯一标识         |

----

#### 7.8 LIVE_RPC (直播RPC服务)

| 环境变量名                 | 默认值    | 中文注释            |
| -------------------------- | --------- | ------------------- |
| LIVE_RPC_ETCD_SERVICENAME  | zookeeper | Etcd 服务名         |
| LIVE_RPC_ETCD_NAMESPACE    | *         | Etcd 命名空间       |
| LIVE_RPC_PGSQL_SERVICENAME | pgpool    | PostgreSQL 服务名   |
| LIVE_RPC_PGSQL_NAMESPACE   | *         | PostgreSQL 命名空间 |
| LIVE_RPC_PGSQL_USER        | root      | PostgreSQL 用户名   |
| LIVE_RPC_PGSQL_PASSWORD    | *         | PostgreSQL 密码     |
| LIVE_RPC_PGSQL_DBNAME      | rvideo    | PostgreSQL 数据库名 |
| LIVE_RPC_REDIS_SERVICENAME | redis     | Redis 服务名        |
| LIVE_RPC_REDIS_NAMESPACE   | *         | Redis 命名空间      |
| LIVE_RPC_REDIS_PASSWORD    | *         | Redis 密码          |
| LIVE_RPC_LOKI_SERVICE      | LIVE_RPC  | Loki 服务标签       |
| LIVE_RPC_LOKI_ENV          | proc      | Loki 环境标识       |
| LIVE_RPC_LOKI_LEVEL        | INFO      | Loki 日志级别       |
| LIVE_RPC_KAFKA_NAMESPACE   | *         | Kafka 命名空间      |
| LIVE_RPC_KAFKA_SERVICENAME | kafka     | Kafka 服务名        |
| LIVE_RPC_POD_UID           | (UUID)    | Pod 唯一标识        |

#### 7.9 DB_INIT (数据库初始化工具)

| 环境变量名                | 默认值 | 中文注释            |
| ------------------------- | ------ | ------------------- |
| DB_INIT_PGSQL_SERVICENAME | pgsql  | PostgreSQL 服务名   |
| DB_INIT_PGSQL_NAMESPACE   | *      | PostgreSQL 命名空间 |
| DB_INIT_PGSQL_USER        | root   | PostgreSQL 用户名   |
| DB_INIT_PGSQL_PASSWORD    | *      | PostgreSQL 密码     |
| DB_INIT_PGSQL_DBNAME      | rvideo | PostgreSQL 数据库名 |

---

#### 7.10 RV_DEBUG（调试标志位）

| 环境变量名 | 默认值 | 中文注释   |
| ---------- | ------ | ---------- |
| RV_DEBUG   | *      | 调试标志位 |

### 8. 数据流

#### 弹幕发送流程

```
客户端 → Danmu Gateway → DanmuSvr (RPC) → Kafka → Consumer → Redis缓存
```

#### 视频上传流程

```
客户端 → Video Gateway (获取RVID) → VideoSvr (记录DB, 携带RVID) → Video Gateway → MinIO (上传文件, 携带RVID)
```

#### 直播流程

```
开始直播: 客户端 → Live Gateway → LiveSvr → PostgreSQL
弹幕发送: 客户端 → Danmu Gateway → DanmuSvr (RPC) → Kafka → Consumer → Redis缓存 → Kafka通知 → WebSocket推送
结束直播: 客户端 → Live Gateway → LiveSvr → PostgreSQL删除 → Kafka通知
```

---

### 9. 部署 （重点讲一下调试模式部署）

> [!important]
>
> **调试模式下，在配置环境变量时，请在将其他环境变量配置好后将`RV_DEBUG`置为`1`**

#### 9.1 调试模式下各个组件的地址配置：

**调试模式下各个组件的地址配置在文件：`RedRock-Video-BE\apps\shared\config\debug-conf.go`中：**

```go
package config

var (
	ETCD_ADDRS          = []string{""} // 127.0.0.1:2181
	REDIS_CLUSTER_ADDRS = []string{""} // 127.0.0.1:6379
	KAFKA_CLUSTER_ADDRS = []string{""} // 127.0.0.1:9092
	PGSQL_ADDRS         = []string{""} // 127.0.0.1
	MINIO_ADDRS         = []string{""} // 127.0.0.1:9000
)
```

#### 9.2 调试模式下启用主机地址服务发现：

```go
	// 调试模式下改为你实际的IP地址
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8888")
	if err != nil {
		l.Error("Resolve TCPAddr Error: %v", err.Error())
		os.Exit(1)
	}

	svr = danmusvr.NewServer(
		new(handle.DanmuSvrImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: union_var.DANMU_SVR,
		}),
		server.WithRegistry(registry),
		server.WithServiceAddr(addr),
		// 容器环境下默认使用容器dns发现，无需在Registry指定IP, 调试模式下去掉注释
		// server.WithRegistryInfo(&rinfo.Info{ServiceName: union_var.DANMU_SVR, Addr: addr}),
		server.WithMiddleware(middleware.PreInit),
		server.WithMiddleware(middleware.DanmuPoolReleaseMiddleware),
	)

	// 启动服务
	err = svr.Run()
	if err != nil {
		l.Error("Run DanmuSvr Error: %v", err.Error())
		os.Exit(1)
	}
```


### 10. 附件

| 名称           | 路径                                        |
| -------------- | ------------------------------------------- |
| ApiFox接口文件 | [ApiFox](docs/RedRockVideo.apifox.json)     |
| 接口文档       | [RedRockVideo.html](docs\RedRockVideo.html) |


> [!note]
>
> **许可证：MIT LICENSE**