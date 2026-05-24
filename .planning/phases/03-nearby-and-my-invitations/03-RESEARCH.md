# Phase 3: 附近约球与我的邀约 - Research

**Researched:** 2026-05-24
**Domain:** Geolocation + Distance Queries + Mobile List UI
**Confidence:** HIGH

## Summary

Phase 3 的核心技术挑战有三个：(1) 浏览器 Geolocation API 获取用户位置并处理各种失败场景；(2) 后端使用 Haversine 公式在 SQLite 中计算距离并排序；(3) 前端使用 Vant 4 的 List/Tabs/Tag 组件构建无限滚动列表和 Tab 切换页面。

项目已有完整的邀约 CRUD、MapPicker 组件、Axios 封装和后端分层架构。Phase 3 在此基础上新增列表查询、加入/退出接口和两个新页面（首页改造 + 我的邀约页）。

**Primary recommendation:** 后端使用 GORM Raw SQL + Haversine 公式实现距离排序，前端用 `useGeolocation` composable 封装定位逻辑，Vant List 组件实现无限滚动加载。

<user_constraints>
## User Constraints (from CONTEXT.md)

### Locked Decisions
- D-01: 卡片式列表展示，每张卡片显示球类、地点、时间、人数进度、距离
- D-02: 下拉加载更多（Vant List 组件），不使用分页
- D-03: 顶部球类标签横向滚动栏筛选（全部/篮球/足球/...），点击切换
- D-04: 显示所有邀约（含已满员/已过期），不可加入的置灰标记
- D-05: 距离计算在后端完成（Haversine 公式），前端传经纬度参数，后端返回已排序列表
- D-06: 浏览器定位失败时，让用户手动选择位置（复用高德地图选点组件）
- D-07: 距离显示精确值（500m / 1.2km / 3.5km），不使用模糊范围
- D-08: 不限制距离范围，所有邀约都显示，距离远的自然排在后面
- D-09: 加入操作仅在详情页进行（非列表页），确保用户看过完整信息
- D-10: 加入需确认弹窗，确认后加入并显示 Toast 提示
- D-11: 退出需确认弹窗，确认后退出并显示 Toast 提示
- D-12: 详情页底部固定按钮，根据状态动态切换：未加入显示"加入"，已加入显示"退出"，已满/已过期置灰
- D-13: "我发起的"和"我参与的"使用 Tab 切换，同一页面内
- D-14: 从个人中心菜单进入"我的邀约"页面
- D-15: 复用附近邀约的卡片样式，但不显示距离，改为显示状态标签

### Claude's Discretion
- 排序切换的具体 UI 形式（距离/时间切换按钮位置）
- 卡片内具体布局和样式细节
- 后端列表接口的分页参数设计（page/pageSize）
- 定位权限请求的时机和提示文案
- 空状态页面的具体展示

### Deferred Ideas (OUT OF SCOPE)
None
</user_constraints>

<phase_requirements>
## Phase Requirements

| ID | Description | Research Support |
|----|-------------|------------------|
| NEARBY-01 | 基于浏览器定位获取用户当前位置 | useGeolocation composable + Geolocation API 错误处理 |
| NEARBY-02 | 列出附近其他人发出的球类邀约 | 后端 ListNearby API + Haversine 距离计算 |
| NEARBY-03 | 列表默认按地点从近到远排列 | SQL ORDER BY distance ASC |
| NEARBY-04 | 可按时间排序 | sort_by 查询参数切换 ORDER BY |
| NEARBY-05 | 可按球类筛选 | sport_type 查询参数 + WHERE 条件 |
| NEARBY-06 | 用户可以加入一个邀约 | POST /invitations/:id/join + 业务校验 |
| NEARBY-07 | 用户可以退出已加入的邀约 | DELETE /invitations/:id/join + 业务校验 |
| MY-01 | 用户可以查看自己发起的邀约列表 | GET /invitations/my?type=created |
| MY-02 | 用户可以查看自己参与的邀约列表 | GET /invitations/my?type=joined |
</phase_requirements>

## Architectural Responsibility Map

| Capability | Primary Tier | Secondary Tier | Rationale |
|------------|-------------|----------------|-----------|
| 浏览器定位 | Browser/Client | — | Geolocation API 是纯浏览器端 API |
| 定位失败回退（地图选点） | Browser/Client | — | 复用已有 MapPicker 组件 |
| 距离计算与排序 | API/Backend | — | D-05 锁定后端计算，避免前端大量数据传输 |
| 邀约列表查询（筛选/排序/分页） | API/Backend | Database | 业务逻辑在 service 层，SQL 在 GORM |
| 加入/退出邀约 | API/Backend | Database | 需要事务保证和并发控制 |
| 列表 UI 展示 | Browser/Client | — | Vant List + 卡片组件 |
| Tab 切换（我的邀约） | Browser/Client | — | Vant Tabs 纯前端切换 |
| 状态计算（expired） | API/Backend | — | 已有 computeStatus 函数 |

## Standard Stack

### Core (已有，无需新增依赖)

| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| Vue 3 | ^3.5.34 | 前端框架 | 项目已用 |
| Vant 4 | ^4.9.24 | 移动端 UI 组件库 | 项目已用，提供 List/Tabs/Tag/Dialog |
| Gin | v1.12.0 | Go Web 框架 | 项目已用 |
| GORM | v1.31.1 | Go ORM | 项目已用，支持 Raw SQL |
| SQLite | — | 数据库 | 项目已用 |
| @amap/amap-jsapi-loader | ^1.0.1 | 高德地图加载器 | 项目已用 |

### Supporting (无需新增)

本 Phase 不需要引入任何新的外部依赖。所有功能均可通过已有依赖实现：
- 距离计算：Go 标准库 `math` 包（sin/cos/asin/sqrt）
- 浏览器定位：原生 `navigator.geolocation` API
- UI 组件：Vant 4 已有 List、Tabs、Tag、Dialog、Toast、Empty

**Installation:** 无需安装新包。

## Architecture Patterns

### System Architecture Diagram

```
[用户浏览器]
    │
    ├── navigator.geolocation.getCurrentPosition()
    │       │
    │       ▼ (lat, lng)
    │
    ├── GET /api/v1/invitations/nearby?lat=X&lng=Y&page=1&page_size=10&sort=distance&sport_type=篮球
    │       │
    │       ▼
    │   [Gin Handler: ListNearby]
    │       │
    │       ▼
    │   [Service: ListNearbyInvitations]
    │       │ 1. 构建 bounding box WHERE 条件（可选优化）
    │       │ 2. Haversine SQL 计算 distance
    │       │ 3. 筛选 sport_type
    │       │ 4. ORDER BY distance/time
    │       │ 5. LIMIT/OFFSET 分页
    │       ▼
    │   [SQLite: invitations + participations JOIN]
    │       │
    │       ▼ (列表数据 + distance 字段)
    │
    ├── POST /api/v1/invitations/:id/join
    │       │
    │       ▼
    │   [Service: JoinInvitation]
    │       │ 1. 校验邀约状态（waiting + 未过期 + 未满员）
    │       │ 2. 校验用户未重复加入
    │       │ 3. 创建 Participation 记录
    │       │ 4. 检查是否满员 → 更新状态为 gathered
    │       ▼
    │
    └── DELETE /api/v1/invitations/:id/join
            │
            ▼
        [Service: LeaveInvitation]
            │ 1. 校验邀约状态（waiting）
            │ 2. 校验用户已加入且非创建者
            │ 3. 删除 Participation 记录
            ▼
```


### Recommended Project Structure

```
frontend/src/
├── views/
│   ├── home/index.vue          # 改造为附近邀约列表页
│   └── my-invitations/index.vue # 新增：我的邀约页（Tab 切换）
├── composables/
│   └── useGeolocation.js       # 新增：浏览器定位 composable
├── components/
│   ├── InvitationCard.vue      # 新增：邀约卡片组件（列表复用）
│   └── MapPicker.vue           # 已有：定位失败时复用
├── api/
│   └── invitation.js           # 扩展：新增 nearby/join/leave/my 接口
└── stores/
    └── user.js                 # 已有：无需修改

backend/
├── handler/
│   └── invitation.go           # 扩展：新增 ListNearby/Join/Leave/MyInvitations handler
├── service/
│   └── invitation.go           # 扩展：新增距离计算、列表查询、加入退出逻辑
├── model/
│   ├── invitation.go           # 已有：无需修改
│   └── participation.go        # 已有：无需修改
└── router/
    └── router.go               # 扩展：新增路由
```

### Pattern 1: Haversine Distance Calculation in Go + GORM Raw SQL

**What:** 使用 Go 标准库 math 包实现 Haversine 公式，通过 GORM Raw SQL 在查询中计算距离并排序。

**When to use:** 需要按距离排序返回邀约列表时。

**Example:**
```go
// Source: Haversine formula + GORM Raw pattern [ASSUMED based on math formula + GORM docs]
import "math"

const earthRadiusKm = 6371.0

// haversineDistance 计算两点间距离（公里）
func haversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
    dLat := (lat2 - lat1) * math.Pi / 180
    dLng := (lng2 - lng1) * math.Pi / 180
    lat1Rad := lat1 * math.Pi / 180
    lat2Rad := lat2 * math.Pi / 180

    a := math.Sin(dLat/2)*math.Sin(dLat/2) +
        math.Cos(lat1Rad)*math.Cos(lat2Rad)*
        math.Sin(dLng/2)*math.Sin(dLng/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
    return earthRadiusKm * c
}

// ListNearbyInvitations 查询附近邀约（带距离排序）
func ListNearbyInvitations(lat, lng float64, sportType, sortBy string, page, pageSize int) ([]NearbyInvitationResponse, int64, error) {
    // Haversine SQL 表达式（SQLite 支持基本数学函数）
    // 注意：SQLite 没有 RADIANS 函数，需要手动 * pi/180
    distanceExpr := `(6371 * 2 * asin(sqrt(
        pow(sin((latitude - ?) * 3.141592653589793 / 180 / 2), 2) +
        cos(? * 3.141592653589793 / 180) * cos(latitude * 3.141592653589793 / 180) *
        pow(sin((longitude - ?) * 3.141592653589793 / 180 / 2), 2)
    )))`

    query := DB.Table("invitations").
        Select("invitations.*, "+distanceExpr+" as distance", lat, lat, lng).
        Where("invitations.deleted_at IS NULL")

    // 球类筛选
    if sportType != "" {
        query = query.Where("sport_type = ?", sportType)
    }

    // 总数统计
    var total int64
    countQuery := DB.Table("invitations").Where("deleted_at IS NULL")
    if sportType != "" {
        countQuery = countQuery.Where("sport_type = ?", sportType)
    }
    countQuery.Count(&total)

    // 排序
    if sortBy == "time" {
        query = query.Order("activity_time ASC")
    } else {
        query = query.Order("distance ASC")
    }

    // 分页
    offset := (page - 1) * pageSize
    query = query.Offset(offset).Limit(pageSize)

    var results []struct {
        model.Invitation
        Distance float64
    }
    query.Scan(&results)

    // 构建响应...
    return responses, total, nil
}
```

### Pattern 2: useGeolocation Composable

**What:** Vue 3 composable 封装浏览器 Geolocation API，提供响应式位置数据和错误状态。

**When to use:** 首页加载时获取用户位置。

**Example:**
```javascript
// Source: Browser Geolocation API [CITED: developer.mozilla.org/en-US/docs/Web/API/Geolocation_API]
import { ref } from 'vue'

export function useGeolocation() {
  const latitude = ref(null)
  const longitude = ref(null)
  const error = ref(null)
  const loading = ref(false)

  function getPosition() {
    return new Promise((resolve, reject) => {
      if (!navigator.geolocation) {
        error.value = 'NOT_SUPPORTED'
        reject(error.value)
        return
      }

      loading.value = true
      navigator.geolocation.getCurrentPosition(
        (position) => {
          latitude.value = position.coords.latitude
          longitude.value = position.coords.longitude
          loading.value = false
          resolve({ lat: latitude.value, lng: longitude.value })
        },
        (err) => {
          loading.value = false
          // err.code: 1=PERMISSION_DENIED, 2=POSITION_UNAVAILABLE, 3=TIMEOUT
          switch (err.code) {
            case 1:
              error.value = 'PERMISSION_DENIED'
              break
            case 2:
              error.value = 'POSITION_UNAVAILABLE'
              break
            case 3:
              error.value = 'TIMEOUT'
              break
            default:
              error.value = 'UNKNOWN'
          }
          reject(error.value)
        },
        {
          enableHighAccuracy: true,
          timeout: 10000,
          maximumAge: 300000 // 5分钟缓存
        }
      )
    })
  }

  return { latitude, longitude, error, loading, getPosition }
}
```


### Pattern 3: Vant List 无限滚动加载

**What:** 使用 Vant List 组件实现上拉加载更多。

**When to use:** 附近邀约列表和我的邀约列表。

**Example:**
```vue
<!-- Source: Vant 4 List docs [CITED: develop365.gitlab.io/vant/en-US/list/] -->
<template>
  <van-list
    v-model:loading="loading"
    :finished="finished"
    finished-text="没有更多了"
    @load="onLoad"
  >
    <InvitationCard
      v-for="item in list"
      :key="item.id"
      :invitation="item"
      @click="goDetail(item.id)"
    />
  </van-list>
</template>

<script setup>
import { ref } from 'vue'

const list = ref([])
const loading = ref(false)
const finished = ref(false)
const page = ref(1)
const pageSize = 10

async function onLoad() {
  try {
    const res = await fetchNearbyList({ page: page.value, page_size: pageSize, ... })
    list.value.push(...res.data.list)
    page.value++
    // 判断是否加载完毕
    if (list.value.length >= res.data.total) {
      finished.value = true
    }
  } catch (err) {
    // 错误处理
  } finally {
    loading.value = false
  }
}
</script>
```

### Pattern 4: Vant Tabs 切换独立数据

**What:** 使用 Vant Tabs 组件实现"我发起的"和"我参与的"两个 Tab 独立加载数据。

**When to use:** 我的邀约页面。

**Example:**
```vue
<!-- Source: Vant 4 Tabs docs [CITED: develop365.gitlab.io/vant/en-US/tab/] -->
<template>
  <van-tabs v-model:active="activeTab" @change="onTabChange">
    <van-tab title="我发起的">
      <van-list v-model:loading="createdLoading" :finished="createdFinished" @load="loadCreated">
        <InvitationCard v-for="item in createdList" :key="item.id" :invitation="item" mode="status" />
        <van-empty v-if="createdFinished && createdList.length === 0" description="暂无发起的邀约" />
      </van-list>
    </van-tab>
    <van-tab title="我参与的">
      <van-list v-model:loading="joinedLoading" :finished="joinedFinished" @load="loadJoined">
        <InvitationCard v-for="item in joinedList" :key="item.id" :invitation="item" mode="status" />
        <van-empty v-if="joinedFinished && joinedList.length === 0" description="暂无参与的邀约" />
      </van-list>
    </van-tab>
  </van-tabs>
</template>
```

### Pattern 5: 详情页底部固定按钮（加入/退出）

**What:** 详情页底部固定操作按钮，根据用户与邀约的关系动态切换。

**When to use:** 邀约详情页扩展。

**Example:**
```vue
<template>
  <!-- 底部固定操作栏（非创建者可见） -->
  <div class="bottom-action" v-if="!isCreator">
    <van-button
      v-if="!hasJoined"
      type="primary"
      block
      round
      :disabled="!canJoin"
      @click="handleJoin"
    >
      {{ canJoin ? '加入邀约' : (isFull ? '已满员' : '已过期') }}
    </van-button>
    <van-button
      v-else
      type="warning"
      block
      round
      @click="handleLeave"
    >
      退出邀约
    </van-button>
  </div>
</template>

<script setup>
import { showConfirmDialog, showToast } from 'vant'

// hasJoined: 当前用户是否已加入
// canJoin: 邀约状态为 waiting 且未满员
const handleJoin = () => {
  showConfirmDialog({
    title: '加入邀约',
    message: '确定要加入这个邀约吗？'
  }).then(async () => {
    await joinInvitation(id)
    showToast('已加入')
    loadDetail() // 刷新详情
  }).catch(() => {})
}

const handleLeave = () => {
  showConfirmDialog({
    title: '退出邀约',
    message: '确定要退出这个邀约吗？'
  }).then(async () => {
    await leaveInvitation(id)
    showToast('已退出')
    loadDetail()
  }).catch(() => {})
}
</script>

<style scoped>
.bottom-action {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 12px 16px;
  background: #fff;
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.06);
}
</style>
```

### Anti-Patterns to Avoid

- **前端计算距离:** 不要在前端用 JavaScript 计算距离再排序。数据量大时性能差，且无法做服务端分页。D-05 已锁定后端计算。
- **SQLite 自定义函数:** 不要尝试注册 SQLite 自定义函数（如 `sqlite3_create_function`），Go 的 mattn/go-sqlite3 虽然支持但增加复杂度。直接在 SQL 中用数学表达式即可。
- **Vant List 不设 loading=false:** 每次 onLoad 回调结束后必须设置 `loading.value = false`，否则 List 不会再次触发加载。
- **Tab 切换重复请求:** 使用 Vant Tabs 的 `lazy-render` 特性（默认开启），避免未激活 Tab 的内容被渲染和请求。

## Don't Hand-Roll

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| 距离计算公式 | 自己推导球面距离公式 | Haversine 标准公式 | 数学精度问题，边界情况（极点、日期变更线） |
| 无限滚动 | 自己监听 scroll 事件 | Vant List 组件 | 处理了节流、边界检测、loading 状态管理 |
| 确认弹窗 | 自己写 modal | Vant showConfirmDialog | Promise 化 API，统一样式 |
| 定位权限处理 | 自己管理权限状态 | useGeolocation composable | 封装错误码映射和状态管理 |
| 距离格式化 | 到处写 if/else | formatDistance 工具函数 | 统一 500m / 1.2km 格式 |

**Key insight:** 本 Phase 的所有 UI 交互都有 Vant 4 现成组件支持，核心技术难点仅在后端距离查询。

## Common Pitfalls

### Pitfall 1: SQLite 缺少数学函数

**What goes wrong:** SQLite 默认不支持 `RADIANS()`、`ASIN()` 等函数，直接写 MySQL 风格的 Haversine SQL 会报错。

**Why it happens:** SQLite 是轻量级数据库，内置数学函数有限。但它支持基本的 `sin()`、`cos()`、`asin()`、`sqrt()`、`pow()` 通过编译选项 `-DSQLITE_ENABLE_MATH_FUNCTIONS`（Go 的 mattn/go-sqlite3 默认启用）。

**How to avoid:** 使用 `sin(x * 3.141592653589793 / 180)` 替代 `sin(radians(x))`。或者确认 mattn/go-sqlite3 编译时启用了数学函数（默认启用），则可直接用 `asin()`、`sqrt()` 等。

**Warning signs:** SQL 执行报 "no such function: radians" 错误。

### Pitfall 2: Geolocation API 在 HTTP 下不可用

**What goes wrong:** 现代浏览器要求 HTTPS 才能使用 Geolocation API（localhost 除外）。

**Why it happens:** 安全策略限制，防止中间人攻击获取用户位置。

**How to avoid:** 开发环境使用 localhost（自动允许），生产环境确保 HTTPS。D-06 的回退方案（手动选点）也能覆盖此场景。

**Warning signs:** `getCurrentPosition` 直接触发 PERMISSION_DENIED 错误，但用户并未拒绝。

### Pitfall 3: Vant List 首次加载触发问题

**What goes wrong:** Vant List 组件挂载后会自动触发一次 `load` 事件（`immediate-check` 默认为 true）。如果此时定位还未完成，会发出无经纬度的请求。

**Why it happens:** List 组件检测到内容未填满视口，自动触发加载。

**How to avoid:** 先完成定位，再渲染 List 组件（用 `v-if="hasLocation"` 控制）。或设置 `immediate-check="false"` 手动控制首次加载时机。

**Warning signs:** 页面加载时发出 lat=undefined 的 API 请求。

### Pitfall 4: 加入邀约的并发竞态

**What goes wrong:** 多个用户同时加入同一邀约，可能超过 max_people 限制。

**Why it happens:** 检查人数和创建记录不是原子操作。

**How to avoid:** 利用 Participation 表的联合唯一索引防止重复加入。在加入后重新 COUNT 检查是否超员，超员则回滚。SQLite 单连接写入（项目已设置 `SetMaxOpenConns(1)`）天然串行化。

**Warning signs:** participant_count 超过 max_people。

### Pitfall 5: 距离排序时的分页一致性

**What goes wrong:** 用户滚动加载下一页时，如果有新邀约插入，可能导致重复或遗漏。

**Why it happens:** OFFSET 分页在数据变动时不稳定。

**How to avoid:** 对于校园应用数据量小，OFFSET 分页完全可接受。如果需要更严格一致性，可用 cursor-based 分页（按 distance + id 组合游标），但本项目不需要。

**Warning signs:** 列表中出现重复卡片。


## Code Examples

### 后端：附近邀约列表 Handler

```go
// Source: 基于项目已有 handler 模式 [VERIFIED: codebase]
type listNearbyRequest struct {
    Latitude  float64 `form:"lat" binding:"required"`
    Longitude float64 `form:"lng" binding:"required"`
    Page      int     `form:"page,default=1"`
    PageSize  int     `form:"page_size,default=10"`
    SortBy    string  `form:"sort_by,default=distance"` // distance | time
    SportType string  `form:"sport_type"`
}

func ListNearby(c *gin.Context) {
    userID := c.GetUint("userID")
    var req listNearbyRequest
    if err := c.ShouldBindQuery(&req); err != nil {
        Error(c, http.StatusBadRequest, "请求参数错误")
        return
    }
    if req.PageSize > 20 {
        req.PageSize = 20
    }

    list, total, err := service.ListNearbyInvitations(
        userID, req.Latitude, req.Longitude,
        req.SportType, req.SortBy, req.Page, req.PageSize,
    )
    if err != nil {
        Error(c, http.StatusInternalServerError, err.Error())
        return
    }
    Success(c, gin.H{"list": list, "total": total})
}
```

### 后端：加入邀约 Service

```go
// Source: 基于项目已有 service 模式 [VERIFIED: codebase]
func JoinInvitation(invitationID, userID uint) error {
    var inv model.Invitation
    if err := DB.First(&inv, invitationID).Error; err != nil {
        return errors.New("邀约不存在")
    }

    // 校验状态
    if computeStatus(inv) != "waiting" {
        return errors.New("该邀约已无法加入")
    }

    // 校验是否已满员
    var count int64
    DB.Model(&model.Participation{}).Where("invitation_id = ?", invitationID).Count(&count)
    if int(count) >= inv.MaxPeople {
        return errors.New("邀约人数已满")
    }

    // 创建参与记录（联合唯一索引防重复）
    participation := model.Participation{InvitationID: invitationID, UserID: userID}
    if err := DB.Create(&participation).Error; err != nil {
        return errors.New("您已加入该邀约")
    }

    // 加入后检查是否满员
    DB.Model(&model.Participation{}).Where("invitation_id = ?", invitationID).Count(&count)
    if int(count) >= inv.MaxPeople {
        DB.Model(&inv).Update("status", "gathered")
    }

    return nil
}
```

### 后端：退出邀约 Service

```go
// Source: 基于项目已有 service 模式 [VERIFIED: codebase]
func LeaveInvitation(invitationID, userID uint) error {
    var inv model.Invitation
    if err := DB.First(&inv, invitationID).Error; err != nil {
        return errors.New("邀约不存在")
    }

    // 创建者不能退出自己的邀约
    if inv.CreatorID == userID {
        return errors.New("创建者不能退出自己的邀约")
    }

    // 只有 waiting 状态可以退出
    if computeStatus(inv) != "waiting" {
        return errors.New("该邀约状态不允许退出")
    }

    // 删除参与记录
    result := DB.Where("invitation_id = ? AND user_id = ?", invitationID, userID).
        Delete(&model.Participation{})
    if result.RowsAffected == 0 {
        return errors.New("您未加入该邀约")
    }

    // 如果之前是 gathered 状态（有人退出后不再满员），恢复为 waiting
    if inv.Status == "gathered" {
        DB.Model(&inv).Update("status", "waiting")
    }

    return nil
}
```

### 后端：我的邀约列表 Service

```go
// Source: 基于项目已有 service 模式 [VERIFIED: codebase]
func MyInvitations(userID uint, listType string, page, pageSize int) ([]InvitationResponse, int64, error) {
    var invitations []model.Invitation
    var total int64
    offset := (page - 1) * pageSize

    if listType == "created" {
        DB.Model(&model.Invitation{}).Where("creator_id = ?", userID).Count(&total)
        DB.Where("creator_id = ?", userID).
            Order("created_at DESC").
            Offset(offset).Limit(pageSize).
            Find(&invitations)
    } else {
        // joined: 通过 participation 表关联查询
        DB.Model(&model.Invitation{}).
            Joins("JOIN participations ON participations.invitation_id = invitations.id").
            Where("participations.user_id = ? AND participations.deleted_at IS NULL", userID).
            Where("invitations.creator_id != ?", userID). // 排除自己创建的
            Count(&total)
        DB.Joins("JOIN participations ON participations.invitation_id = invitations.id").
            Where("participations.user_id = ? AND participations.deleted_at IS NULL", userID).
            Where("invitations.creator_id != ?", userID).
            Order("invitations.created_at DESC").
            Offset(offset).Limit(pageSize).
            Find(&invitations)
    }

    responses := make([]InvitationResponse, len(invitations))
    for i, inv := range invitations {
        responses[i] = *buildResponse(inv)
    }
    return responses, total, nil
}
```

### 前端：距离格式化工具函数

```javascript
// Source: D-07 要求精确值显示 [VERIFIED: CONTEXT.md]
/**
 * 格式化距离显示
 * @param {number} distanceKm - 距离（公里）
 * @returns {string} 格式化后的距离字符串
 */
export function formatDistance(distanceKm) {
  if (distanceKm < 1) {
    return `${Math.round(distanceKm * 1000)}m`
  }
  return `${distanceKm.toFixed(1)}km`
}
```

### 前端：API 接口扩展

```javascript
// Source: 基于项目已有 api 模式 [VERIFIED: codebase]
import request from './request'

// 获取附近邀约列表
export function getNearbyInvitations(params) {
  return request.get('/invitations/nearby', { params })
}

// 加入邀约
export function joinInvitation(id) {
  return request.post(`/invitations/${id}/join`)
}

// 退出邀约
export function leaveInvitation(id) {
  return request.delete(`/invitations/${id}/join`)
}

// 获取我的邀约列表
export function getMyInvitations(params) {
  return request.get('/invitations/my', { params })
}
```

### 后端：路由注册扩展

```go
// Source: 基于项目已有 router 模式 [VERIFIED: codebase]
// 在 invitation 路由组中新增：
invitation.GET("/nearby", handler.ListNearby)
invitation.GET("/my", handler.MyInvitations)
invitation.POST("/:id/join", handler.JoinInvitation)
invitation.DELETE("/:id/join", handler.LeaveInvitation)
```

## State of the Art

| Old Approach | Current Approach | When Changed | Impact |
|--------------|------------------|--------------|--------|
| Vant 3 Dialog.confirm() | Vant 4 showConfirmDialog() | Vant 4.0 | 函数式调用，import 方式变化 |
| Geolocation watchPosition | getCurrentPosition + 缓存 | — | 单次定位足够，避免持续消耗电量 |
| 前端计算距离 | 后端 SQL 计算 | — | 支持服务端分页，减少数据传输 |

## Project Constraints (from CLAUDE.md)

- **Tech Stack**: Vue 3 + Vant 4 (frontend), Go (backend), SQLite (database)
- **Structure**: frontend/ 和 backend/ 分目录存放
- **Comments**: 复杂逻辑和函数需简短中文注释

## Assumptions Log

| # | Claim | Section | Risk if Wrong |
|---|-------|---------|---------------|
| A1 | mattn/go-sqlite3 默认启用 SQLITE_ENABLE_MATH_FUNCTIONS 编译选项 | Pitfall 1 | sin/cos/asin 函数不可用，需改用 Go 层计算距离 |
| A2 | SQLite 支持 pow() 和 sqrt() 函数 | Pattern 1 | SQL 表达式报错，需改用替代写法或 Go 层计算 |
| A3 | Vant 4 List 组件 v-model:loading 在 Vant ^4.9 中可用 | Pattern 3 | 可能需要用 :loading 替代 |

**Mitigation for A1/A2:** 如果 SQLite 数学函数不可用，回退方案是在 Go service 层用 `haversineDistance()` 函数计算距离，先查出所有邀约再在内存中排序分页。对于校园应用数据量（百级），性能完全可接受。

## Open Questions

1. **SQLite 数学函数可用性**
   - What we know: mattn/go-sqlite3 通常编译时启用数学函数扩展
   - What's unclear: 当前项目的 go-sqlite3 版本是否确实启用
   - Recommendation: 实现时先测试 `SELECT sin(1)` 是否可执行，不可用则回退到 Go 层计算

2. **详情页是否需要返回当前用户的参与状态**
   - What we know: 详情页需要根据用户是否已加入显示不同按钮
   - What's unclear: 是在详情接口中返回 `has_joined` 字段，还是前端单独请求
   - Recommendation: 在详情接口中增加 `has_joined` 字段（需要传 userID），一次请求解决

## Sources

### Primary (HIGH confidence)
- Vant 4 List docs: https://develop365.gitlab.io/vant/en-US/list/
- Vant 4 Tabs docs: https://develop365.gitlab.io/vant/en-US/tab/
- Vant 4 Dialog docs: https://develop365.gitlab.io/vant/en-US/dialog/
- MDN Geolocation API: https://developer.mozilla.org/en-US/docs/Web/API/GeolocationPositionError/code
- Project codebase: backend/service/invitation.go, frontend/src/views/invitation/detail.vue

### Secondary (MEDIUM confidence)
- GORM distance query pattern: https://forum.golangbridge.org/t/distance-searching-with-raw-gorm-query-solved/5924
- Haversine formula: https://thelinuxcode.com/program-for-distance-between-two-points-on-earth-a-practical-production-ready-guide/

### Tertiary (LOW confidence)
- SQLite math functions availability in mattn/go-sqlite3 (training knowledge)

## Metadata

**Confidence breakdown:**
- Standard stack: HIGH - 所有依赖已在项目中使用，无需新增
- Architecture: HIGH - 完全遵循已有分层模式，代码模式已验证
- Pitfalls: MEDIUM - SQLite 数学函数可用性需实际验证

**Research date:** 2026-05-24
**Valid until:** 2026-06-24 (stable stack, no fast-moving dependencies)
