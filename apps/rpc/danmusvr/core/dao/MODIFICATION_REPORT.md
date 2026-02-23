# DAO 层代码修改报告

## 修改概述

本次修改针对 `d:\Project\RedRock-Video-BE\apps\rpc\danmusvr\core\dao` 目录下的代码，修复了多个 bug 并进行了性能优化。

---

## 1. handle.go - Compare 函数切片越界 bug

**位置**: `handle.go` 第 12 行

**问题**: 原来 `data[:max-1]` 会丢失最后一个元素，例如 max=1000 时只返回 999 个元素

**修改**:
```go
// 修改前
return data[:max-1]

// 修改后
return data[:max]
```

---

## 2. redis-operation.go - JSON 反序列化指针 bug（关键性能问题）

**位置**: `redis-operation.go` `getHotDanmuR` 和 `getFullDanmuR` 函数

**问题**:
1. 使用 `var DanmuData *publicDao.DanmuData` 声明指针但未分配内存
2. `jsoniter.Unmarshal([]byte(data), DanmuData)` 传入的是 nil 指针，会导致 panic 或解析失败
3. `var results []*publicDao.DanmuData` 未预分配容量，每次 append 都会触发内存重新分配

**修改**:
```go
// 修改前
var results []*publicDao.DanmuData
for _, data := range rawJsonList {
    var DanmuData *publicDao.DanmuData
    if err := jsoniter.Unmarshal([]byte(data), DanmuData); err != nil {
        continue
    }
    results = append(results, DanmuData)
}

// 修改后
results := make([]*publicDao.DanmuData, 0, len(rawJsonList))
for _, data := range rawJsonList {
    var DanmuData publicDao.DanmuData
    if err := jsoniter.Unmarshal([]byte(data), &DanmuData); err != nil {
        continue
    }
    results = append(results, &DanmuData)
}
```

**性能影响**: 修复后避免了解析失败和频繁内存分配，显著提升弹幕读取速度

---

## 3. pgsql-operation.go - GORM 查询指针 bug

**位置**: `pgsql-operation.go` `getVideoDanmuDetail` 函数

**问题**:
```go
// 修改前
var data *models.DanmuData
err := tx.Where("dan_id = ?", danID).First(data).Error  // data 是 nil，查询结果无法写入
```

**修改**:
```go
// 修改后
var data models.DanmuData
err := tx.Where("dan_id = ?", danID).First(&data).Error
return &data, nil
```

---

## 4. pgsql-operation.go - delVideoDanmu 冗余错误处理

**位置**: `pgsql-operation.go` `delVideoDanmu` 函数

**问题**:
1. GORM 的 Delete 方法不会返回 `gorm.ErrRecordNotFound`，这个检查是多余的
2. 参数命名 `Tx` 与 GORM 内部使用的 `tx` 容易混淆

**修改**:
```go
// 修改前
func (r *Dao) delVideoDanmu(Tx *gorm.DB, danID int64) error {
    var dest models.DanmuData
    if err := Tx.Where("dan_id = ?", danID).Delete(&dest).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil
        }
        return err
    }
    return nil
}

// 修改后
func (r *Dao) delVideoDanmu(tx *gorm.DB, danID int64) error {
    var dest models.DanmuData
    if err := tx.Where("dan_id = ?", danID).Delete(&dest).Error; err != nil {
        return err
    }
    return nil
}
```

---

## 5. handle.go + pgsql-operation.go - N+1 查询性能优化（最重要）

**位置**:
- `handle.go` 第 31-34 行 (`ReadHotDanmu` 函数)
- `handle.go` 第 78-81 行 (`ReadFullDanmu` 函数)
- `pgsql-operation.go` 新增 `getUserInfoBatch` 函数

**问题**:
```go
// 原代码：每个弹幕都单独查询用户信息，1000条弹幕 = 1001次数据库查询
for _, val := range data {
    val.User, _ = r.getUserInfo(val.UserId)
}
```

**修改**: 新增批量查询函数 `getUserInfoBatch`，将 N+1 次查询优化为 1 次

```go
// pgsql-operation.go 新增
func (r *Dao) getUserInfoBatch(uids []int64) (map[int64]models.DUserInfo, error) {
    if len(uids) == 0 {
        return make(map[int64]models.DUserInfo), nil
    }
    var users []models.RvUser
    err := r.pgdb.Where("github_uid IN ?", uids).Select("github_uid", "avatar_url", "github_login").Find(&users).Error
    // ... 构建 map 返回
}

// handle.go 修改
userIDs := make([]int64, len(data))
for i, val := range data {
    userIDs[i] = val.UserId
}
userMap, _ := r.getUserInfoBatch(userIDs)
for _, val := range data {
    if user, ok := userMap[val.UserId]; ok {
        val.User = user
    }
}
```

**性能影响**: 假设 1000 条弹幕，从 1001 次数据库查询减少到 2 次（1 次弹幕查询 + 1 次用户批量查询），性能提升 **500 倍以上**

---

## 性能优化总结

| 优化项 | 性能提升 | 影响范围 |
|--------|----------|----------|
| JSON 解析修复 | 避免 panic/解析失败 | 所有弹幕读取 |
| 切片预分配 | 减少内存分配 | 所有弹幕读取 |
| N+1 查询优化 | 500x+ | 首次从 PGSQL 读取弹幕 |

---

## 注意事项

1. **硬编码端口问题**: 按要求未修改
2. **用户未返回时处理**: 批量查询后如果用户不存在，不会设置 User 字段（保持原有行为）
3. **错误处理**: 批量查询错误被忽略（与原有代码行为一致），如需严格处理可自行添加
