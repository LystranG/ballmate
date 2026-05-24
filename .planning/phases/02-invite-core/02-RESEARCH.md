# Phase 2: 邀约核心 - Research

**Researched:** 2026-05-24
**Domain:** 邀约 CRUD + 地图选点 + 日期时间选择 + 参与人管理
**Confidence:** HIGH

## Summary

Phase 2 在 Phase 1 已建立的 handler → service → model 后端分层和 Vue 3 + Vant 4 前端架构上，新增邀约（Invitation）核心功能。主要技术挑战集中在三个方面：(1) 高德地图 JS API 2.0 集成实现地图选点；(2) Vant 4 PickerGroup 组合 DatePicker + TimePicker 实现日期时间选择；(3) 邀约状态机设计（等待中→已召集/已终止/已过期）。

数据模型需要新增 `invitations` 和 `participations` 两张表，遵循 GORM 的 gorm.Model 模式（含软删除）。后端新增 invitation handler/service/model 三层文件，前端新增 create（发起邀约）和 invitation detail（邀约详情+参与人）两个视图。

**Primary recommendation:** 严格复用 Phase 1 的分层模式，地图选点封装为独立组件，PickerGroup 组合日期时间选择，状态判断在查询时计算（无定时任务）。

<user_constraints>
## User Constraints (from CONTEXT.md)

### Locked Decisions
- **D-01:** 地点选择使用高德地图 JS API 组件选点，用户在地图上点击选择坐标并搜索地址
- **D-02:** 时间选择使用 Vant DatePicker type='datetime' 一体选择器，精确到分钟
- **D-03:** 人数设置使用预设档位（2人/4人/6人/8人/10人），快速选择
- **D-04:** 球类选择复用 Phase 1 已有的球类标签列表，单选
- **D-05:** 四状态设计：等待中→已召集/已终止/已过期，终态不可逆
- **D-06:** "已过期"通过查询时判断实现，无需定时任务
- **D-07:** 终止和删除为链路关系，需确认弹窗
- **D-08:** 邀约卡片显示"N/M人"摘要
- **D-09:** 手机号脱敏在后端处理
- **D-10:** 只有发起人可查看参与人脱敏手机号
- **D-11:** 独立关联表 participations
- **D-12:** 地理坐标两个浮点字段

### Claude's Discretion
- 邀约表单的具体校验规则和错误提示
- 邀约详情页的具体布局
- API 路由设计和接口命名
- 数据库表的具体字段命名和索引策略
- 球类选择的 UI 交互细节（Tag 组件 vs Radio 组件）

### Deferred Ideas (OUT OF SCOPE)
None
</user_constraints>

<phase_requirements>
## Phase Requirements

| ID | Description | Research Support |
|----|-------------|------------------|
| INVITE-01 | 用户可以发起球类邀约（选择球类、时间、地点、人数要求） | 表单组件：Vant Tag/Radio + PickerGroup + AMap 选点 + 预设档位 |
| INVITE-02 | 发起邀约时支持地图选点确定地点坐标 | @amap/amap-jsapi-loader + AMap.PlaceSearch + AMap.Geocoder |
| INVITE-03 | 提交后邀约处于"等待中"状态 | 后端创建时 status 默认值 = "waiting" |
| INVITE-04 | 人员凑齐时自动变为"已召集"状态 | 加入参与时 service 层检查 count >= max_people |
| INVITE-05 | 发起人可以终止邀约 | PUT /invitations/:id/terminate，权限校验 creator_id |
| INVITE-06 | 发起人可以删除邀约 | DELETE /invitations/:id（GORM 软删除），前置条件：已终止 |
| PART-01 | 发起人可以查看当前邀约的参与人列表 | GET /invitations/:id/participants，权限校验 |
| PART-02 | 参与人信息包括姓名和手机号 | JOIN users 表查询 nickname + phone |
| PART-03 | 手机号中间4位用*号代替 | 后端 service 层脱敏处理后返回 |
</phase_requirements>

## Architectural Responsibility Map

| Capability | Primary Tier | Secondary Tier | Rationale |
|------------|-------------|----------------|-----------|
| 邀约表单 UI（球类/时间/人数） | Browser/Client | — | 纯前端交互，Vant 组件渲染 |
| 地图选点 | Browser/Client | — | 高德 JS API 在浏览器端运行，返回坐标给前端 |
| 邀约 CRUD 业务逻辑 | API/Backend | — | 数据校验、状态流转、权限控制在后端 |
| 状态机判断（已过期） | API/Backend | — | 查询时对比 activity_time 与当前时间 |
| 参与人管理 | API/Backend | — | JOIN 查询 + 手机号脱敏在后端完成 |
| 数据持久化 | Database/Storage | — | SQLite + GORM，invitations + participations 表 |
| 邀约详情展示 | Browser/Client | API/Backend | 前端渲染，后端提供聚合数据（含参与人数） |

## Standard Stack

### Core (已有，Phase 1 建立)
| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| Vue 3 | ^3.5.34 | 前端框架 | 项目约束 |
| Vant 4 | ^4.9.24 | 移动端 UI 组件库 | 项目约束，提供 DatePicker/TimePicker/PickerGroup/Tag/Dialog |
| Gin | v1.12.0 | Go HTTP 框架 | 项目约束 |
| GORM | v1.31.1 | Go ORM | 项目约束，提供软删除和自动迁移 |
| SQLite | — | 数据库 | 项目约束 |

### New Dependencies (Phase 2 新增)
| Library | Version | Purpose | When to Use |
|---------|---------|---------|-------------|
| @amap/amap-jsapi-loader | 1.0.1 | 高德地图 JS API 加载器 | 地图选点页面加载 AMap SDK |

### Alternatives Considered
| Instead of | Could Use | Tradeoff |
|------------|-----------|----------|
| @amap/amap-jsapi-loader | script 标签直接引入 | loader 更安全，避免异步加载错误，支持按需加载插件 |
| PickerGroup 组合 | 自定义 Picker columns | PickerGroup 是 Vant 4 官方方案，维护成本低 |
| Vant Tag 单选球类 | Vant Radio | Tag 视觉更紧凑适合多选项横排，Radio 更传统 |

**Installation (frontend):**
```bash
npm install @amap/amap-jsapi-loader --save
```

**Backend:** 无新依赖，复用现有 go.mod。

## Package Legitimacy Audit

*slopcheck was unavailable at research time. All packages tagged `[ASSUMED]` and planner must gate install behind checkpoint:human-verify.*

| Package | Registry | Age | Downloads | Source Repo | slopcheck | Disposition |
|---------|----------|-----|-----------|-------------|-----------|-------------|
| @amap/amap-jsapi-loader | npm | 4+ yrs | — | gitlab.alibaba-inc.com/amap-web/amap-jsapi-loader | N/A | Approved [ASSUMED] — official AMap/Alibaba package, maintainers are @alibaba-inc.com emails |

**Packages removed due to slopcheck [SLOP] verdict:** none
**Packages flagged as suspicious [SUS]:** none

## Architecture Patterns

### System Architecture Diagram

```
[用户浏览器]
    |
    |-- (1) 发起邀约表单 ---> [Vue 前端 /create]
    |       |                       |
    |       |-- 地图选点 ---------> [高德 JS API 2.0] --> 返回 lng/lat + address
    |       |-- 日期时间 ---------> [Vant PickerGroup]
    |       |-- 球类/人数 --------> [Vant Tag / Stepper]
    |       |
    |       v
    |   POST /api/v1/invitations --> [Gin Handler] --> [Service] --> [GORM Model]
    |                                                      |
    |                                                      v
    |                                               [SQLite: invitations]
    |
    |-- (2) 邀约详情 ------------> GET /api/v1/invitations/:id
    |                                   |
    |                                   v
    |                             [Service: 查询时判断过期状态]
    |                                   |
    |                                   v
    |                             返回 invitation + participant_count
    |
    |-- (3) 参与人列表 ----------> GET /api/v1/invitations/:id/participants
    |                                   |
    |                                   v
    |                             [Service: 权限校验 creator_id == userID]
    |                                   |
    |                                   v
    |                             [JOIN participations + users, 脱敏 phone]
    |
    |-- (4) 终止邀约 ------------> PUT /api/v1/invitations/:id/terminate
    |-- (5) 删除邀约 ------------> DELETE /api/v1/invitations/:id
```

### Recommended Project Structure

```
backend/
├── model/
│   ├── user.go              # 已有
│   ├── invitation.go        # 新增：Invitation 模型
│   └── participation.go     # 新增：Participation 关联模型
├── handler/
│   ├── auth.go              # 已有
│   ├── user.go              # 已有
│   ├── response.go          # 已有
│   └── invitation.go        # 新增：邀约相关 handler
├── service/
│   ├── auth.go              # 已有
│   ├── user.go              # 已有
│   └── invitation.go        # 新增：邀约业务逻辑
├── router/
│   └── router.go            # 修改：新增邀约路由组
└── main.go                  # 修改：AutoMigrate 新模型

frontend/src/
├── api/
│   ├── request.js           # 已有
│   ├── user.js              # 已有
│   └── invitation.js        # 新增：邀约 API 封装
├── views/
│   ├── create/
│   │   └── index.vue        # 新增：发起邀约页面
│   └── invitation/
│       ├── detail.vue       # 新增：邀约详情页
│       └── participants.vue # 新增：参与人列表页
├── components/
│   ├── TabBar.vue           # 已有
│   └── MapPicker.vue        # 新增：地图选点组件
└── router/
    └── index.js             # 修改：新增路由
```

### Pattern 1: 邀约状态机（查询时计算）
**What:** 邀约有四种状态，其中"已过期"不存储在数据库，而是查询时根据 activity_time < now() 动态判断
**When to use:** 每次返回邀约数据时
**Example:**
```go
// Source: 项目 D-06 决策 + GORM 查询模式 [ASSUMED]
func (s *InvitationService) GetByID(id uint) (*InvitationResponse, error) {
    var inv model.Invitation
    if err := DB.First(&inv, id).Error; err != nil {
        return nil, errors.New("邀约不存在")
    }
    
    // 查询时判断过期状态
    status := inv.Status
    if status == "waiting" && inv.ActivityTime.Before(time.Now()) {
        status = "expired"
    }
    
    // 查询参与人数
    var count int64
    DB.Model(&model.Participation{}).Where("invitation_id = ?", id).Count(&count)
    
    return &InvitationResponse{
        Invitation:       inv,
        Status:           status,
        ParticipantCount: int(count),
    }, nil
}
```

### Pattern 2: 高德地图选点组件封装
**What:** 将 AMap 初始化、地图点击事件、逆地理编码封装为独立 Vue 组件
**When to use:** 发起邀约表单中选择地点
**Example:**
```vue
<!-- Source: AMap JS API 2.0 官方文档 [CITED: lbs.amap.com/api/javascript-api-v2/guide/abc/load] -->
<template>
  <div>
    <div id="map-container" style="height: 300px"></div>
    <van-field v-model="address" readonly placeholder="点击地图选择地点" />
  </div>
</template>

<script setup>
import { shallowRef, onMounted, onUnmounted, defineEmits } from 'vue'
import AMapLoader from '@amap/amap-jsapi-loader'

const emit = defineEmits(['select'])
const map = shallowRef(null)
const address = ref('')

onMounted(() => {
  window._AMapSecurityConfig = { securityJsCode: 'YOUR_SECURITY_KEY' }
  
  AMapLoader.load({
    key: 'YOUR_KEY',
    version: '2.0',
    plugins: ['AMap.Geocoder', 'AMap.PlaceSearch']
  }).then((AMap) => {
    map.value = new AMap.Map('map-container', { zoom: 14 })
    const geocoder = new AMap.Geocoder()
    
    map.value.on('click', (e) => {
      const { lng, lat } = e.lnglat
      geocoder.getAddress([lng, lat], (status, result) => {
        if (status === 'complete') {
          address.value = result.regeocode.formattedAddress
          emit('select', { longitude: lng, latitude: lat, address: address.value })
        }
      })
    })
  })
})

onUnmounted(() => {
  map.value?.destroy()
})
</script>
```

### Pattern 3: Vant 4 PickerGroup 日期时间选择
**What:** 使用 PickerGroup 组合 DatePicker + TimePicker 实现一体化日期时间选择
**When to use:** 发起邀约表单中选择活动时间
**Example:**
```vue
<!-- Source: Vant 4 PickerGroup 官方 README [CITED: github.com/youzan/vant/blob/main/packages/vant/src/picker-group/README.md] -->
<van-popup v-model:show="showPicker" position="bottom">
  <van-picker-group
    title="选择活动时间"
    :tabs="['日期', '时间']"
    @confirm="onConfirm"
    @cancel="showPicker = false"
  >
    <van-date-picker
      v-model="selectedDate"
      :min-date="new Date()"
      :columns-type="['year', 'month', 'day']"
    />
    <van-time-picker
      v-model="selectedTime"
      :columns-type="['hour', 'minute']"
    />
  </van-picker-group>
</van-popup>
```

### Pattern 4: 手机号脱敏（后端处理）
**What:** 在 service 层返回参与人数据前，将手机号中间 4 位替换为 ****
**When to use:** 查询参与人列表时
**Example:**
```go
// Source: D-09 决策 [ASSUMED]
func maskPhone(phone string) string {
    if len(phone) != 11 {
        return phone
    }
    return phone[:3] + "****" + phone[7:]
}
```

### Anti-Patterns to Avoid
- **在前端做手机号脱敏:** 前端拿到完整手机号再遮盖是安全漏洞，必须后端处理
- **用定时任务判断过期:** 增加系统复杂度，查询时判断更简单可靠
- **地图实例用 ref() 包裹:** Vue 3 的 Proxy 会破坏 AMap 原生对象，必须用 shallowRef
- **在 handler 层写业务逻辑:** 保持 handler 只做参数校验和响应，逻辑放 service

## Don't Hand-Roll

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| 地图选点 | 自己画 canvas 地图 | 高德 JS API 2.0 | 地图渲染、POI 搜索、逆地理编码极其复杂 |
| 日期时间选择 | 自定义滚动列 | Vant PickerGroup + DatePicker + TimePicker | 处理闰年、月份天数、时区等边界情况 |
| 确认弹窗 | 自己写 modal | Vant Dialog.confirm | 已有标准交互模式 |
| 软删除 | 手动 is_deleted 字段 | GORM gorm.Model (DeletedAt) | 自动处理查询过滤和恢复 |

**Key insight:** 本 phase 的复杂度在于组件集成和状态流转，不在于底层实现。所有 UI 交互都有 Vant 现成组件，地图有高德 SDK，数据层有 GORM。

## Common Pitfalls

### Pitfall 1: Vue 3 Proxy 破坏 AMap 对象
**What goes wrong:** 用 `ref()` 存储 AMap.Map 实例，Vue 的 Proxy 拦截导致地图 API 调用异常
**Why it happens:** Vue 3 响应式系统用 Proxy 包裹对象，AMap 内部依赖 this 指向原始对象
**How to avoid:** 使用 `shallowRef()` 存储地图实例，避免深度代理
**Warning signs:** 地图渲染正常但事件回调报错，或 marker 添加失败

### Pitfall 2: 高德 API Key 安全配置遗漏
**What goes wrong:** 2021年12月后创建的 Key 必须配合安全密钥使用，否则 API 调用失败
**Why it happens:** 高德升级了安全策略，旧教程没有 securityJsCode 配置
**How to avoid:** 在 AMapLoader.load 之前设置 `window._AMapSecurityConfig = { securityJsCode: '...' }`
**Warning signs:** 控制台报 "INVALID_USER_KEY" 或 "USERKEY_PLAT_NOMATCH"

### Pitfall 3: 邀约状态判断遗漏
**What goes wrong:** 列表查询时忘记对每条记录做过期判断，导致已过期邀约仍显示"等待中"
**Why it happens:** 过期状态不存储在数据库，容易在新增查询接口时遗忘
**How to avoid:** 在 service 层封装统一的 `computeStatus()` 方法，所有返回邀约数据的接口都调用
**Warning signs:** 活动时间已过但状态仍为 waiting

### Pitfall 4: 参与人数竞态条件
**What goes wrong:** 两人同时加入导致超过人数上限
**Why it happens:** 先查询 count 再插入，中间有时间窗口
**How to avoid:** SQLite 单连接（MaxOpenConns=1）天然串行化写入，已在 main.go 配置。加入时在事务内检查
**Warning signs:** participant_count > max_people

### Pitfall 5: GORM 软删除影响关联查询
**What goes wrong:** 删除邀约后，participations 表的关联记录仍存在，查询时可能出现孤儿数据
**Why it happens:** GORM 软删除只标记 deleted_at，不级联删除关联表
**How to avoid:** 删除邀约时同时软删除对应的 participations 记录
**Warning signs:** 已删除邀约的参与人仍出现在某些统计中

## Code Examples

### 数据模型定义
```go
// Source: GORM 官方文档 + 项目 model/user.go 模式 [ASSUMED]
package model

import (
    "time"
    "gorm.io/gorm"
)

// Invitation 邀约模型
type Invitation struct {
    gorm.Model
    CreatorID    uint      `gorm:"not null;index" json:"creator_id"`
    SportType    string    `gorm:"size:20;not null" json:"sport_type"`
    ActivityTime time.Time `gorm:"not null" json:"activity_time"`
    Address      string    `gorm:"size:255;not null" json:"address"`
    Latitude     float64   `gorm:"not null" json:"latitude"`
    Longitude    float64   `gorm:"not null" json:"longitude"`
    MaxPeople    int       `gorm:"not null" json:"max_people"`
    Status       string    `gorm:"size:20;not null;default:waiting" json:"status"`
    // Status: waiting / gathered / terminated
    // expired 通过查询时计算，不存储
}

// Participation 参与记录
type Participation struct {
    gorm.Model
    InvitationID uint `gorm:"not null;index" json:"invitation_id"`
    UserID       uint `gorm:"not null;index" json:"user_id"`
}
```

### API 路由注册
```go
// Source: 项目 router/router.go 模式 [ASSUMED]
// 邀约路由组（需认证）
invitation := api.Group("/invitations")
invitation.Use(middleware.AuthMiddleware())
{
    invitation.POST("", handler.CreateInvitation)
    invitation.GET("/:id", handler.GetInvitation)
    invitation.PUT("/:id/terminate", handler.TerminateInvitation)
    invitation.DELETE("/:id", handler.DeleteInvitation)
    invitation.GET("/:id/participants", handler.GetParticipants)
}
```

### 前端 API 封装
```javascript
// Source: 项目 api/user.js 模式 [ASSUMED]
import request from './request'

export function createInvitation(data) {
  return request.post('/invitations', data)
}

export function getInvitation(id) {
  return request.get(`/invitations/${id}`)
}

export function terminateInvitation(id) {
  return request.put(`/invitations/${id}/terminate`)
}

export function deleteInvitation(id) {
  return request.delete(`/invitations/${id}`)
}

export function getParticipants(id) {
  return request.get(`/invitations/${id}/participants`)
}
```

### 自动状态流转（加入时检查）
```go
// Source: D-04/D-05 决策 [ASSUMED]
func JoinInvitation(invitationID, userID uint) error {
    var inv model.Invitation
    if err := DB.First(&inv, invitationID).Error; err != nil {
        return errors.New("邀约不存在")
    }
    
    // 检查状态
    if inv.Status != "waiting" || inv.ActivityTime.Before(time.Now()) {
        return errors.New("该邀约已结束或已过期")
    }
    
    // 检查是否已参与
    var existing model.Participation
    if err := DB.Where("invitation_id = ? AND user_id = ?", invitationID, userID).First(&existing).Error; err == nil {
        return errors.New("您已参与该邀约")
    }
    
    // 创建参与记录
    participation := &model.Participation{
        InvitationID: invitationID,
        UserID:       userID,
    }
    if err := DB.Create(participation).Error; err != nil {
        return errors.New("加入失败")
    }
    
    // 检查是否凑齐人数，自动更新状态
    var count int64
    DB.Model(&model.Participation{}).Where("invitation_id = ?", invitationID).Count(&count)
    if int(count) >= inv.MaxPeople {
        DB.Model(&inv).Update("status", "gathered")
    }
    
    return nil
}
```

## State of the Art

| Old Approach | Current Approach | When Changed | Impact |
|--------------|------------------|--------------|--------|
| Vant 3 DatetimePicker (单组件) | Vant 4 PickerGroup + DatePicker + TimePicker | Vant 4.0 (2022-11) | 更灵活的组合方式，DatetimePicker 已废弃 |
| 高德 JS API 1.x (script 标签) | @amap/amap-jsapi-loader + JS API 2.0 | 2021-12 | 必须配合 securityJsCode，loader 方式更安全 |
| GORM v1 | GORM v2 (gorm.io/gorm) | 2020 | 新 import 路径，API 有变化 |

**Deprecated/outdated:**
- `van-datetime-picker`: Vant 4 中已移除，用 PickerGroup 组合替代
- 高德 JS API 1.x 的 `AMap.plugin()` 动态加载方式：2.0 推荐在 loader 的 plugins 数组中声明

## Assumptions Log

| # | Claim | Section | Risk if Wrong |
|---|-------|---------|---------------|
| A1 | Vant 4 PickerGroup 支持 DatePicker + TimePicker 组合（基于 GitHub README） | Architecture Patterns | 如果 API 不同需调整组件写法 |
| A2 | @amap/amap-jsapi-loader 1.0.1 是当前最新版本 | Standard Stack | 低风险，npm view 已确认 |
| A3 | AMap.Geocoder 的 getAddress 回调格式为 (status, result) | Code Examples | 如果 API 变化需调整回调处理 |
| A4 | GORM gorm.Model 的 DeletedAt 字段自动过滤软删除记录 | Architecture Patterns | 极低风险，GORM 核心特性 |
| A5 | SQLite MaxOpenConns=1 可以避免写入竞态 | Common Pitfalls | 低风险，已在 main.go 中配置 |

## Open Questions

1. **高德地图 API Key**
   - What we know: 需要在高德开放平台申请 Web 端 Key + 安全密钥
   - What's unclear: 用户是否已有高德开发者账号和 Key
   - Recommendation: 计划中加入"配置 AMap Key"步骤，Key 值放环境变量或配置文件

2. **邀约创建者是否自动成为参与人**
   - What we know: CONTEXT.md 未明确说明
   - What's unclear: 创建者算不算 N/M 中的一个人
   - Recommendation: 创建者自动加入 participations 表（计入人数），这是最自然的用户预期

## Project Constraints (from CLAUDE.md)

- **Tech Stack**: Vue 3 + Vant 4 (frontend), Go (backend), SQLite (database)
- **Structure**: frontend/ 和 backend/ 分目录存放
- **Comments**: 复杂逻辑和函数需简短中文注释

## Sources

### Primary (HIGH confidence)
- AMap JS API 2.0 官方加载文档 — [lbs.amap.com/api/javascript-api-v2/guide/abc/load](https://lbs.amap.com/api/javascript-api-v2/guide/abc/load)
- Vant 4 PickerGroup README — [github.com/youzan/vant PickerGroup](https://github.com/youzan/vant/blob/main/packages/vant/src/picker-group/README.md)
- Vant 4 DatePicker README — [github.com/youzan/vant DatePicker](https://github.com/youzan/vant/blob/main/packages/vant/src/date-picker/README.md)
- Vant 4 TimePicker README — [github.com/youzan/vant TimePicker](https://github.com/youzan/vant/blob/main/packages/vant/src/time-picker/README.md)
- npm registry: @amap/amap-jsapi-loader v1.0.1 — [npmjs.com](https://www.npmjs.com/package/@amap/amap-jsapi-loader)

### Secondary (MEDIUM confidence)
- Vue 3 + AMap 集成实践 — [juejin.cn/post/7117556295703461924](https://juejin.cn/post/7117556295703461924)
- Vant 4 Release Notes (PickerGroup 介绍) — [develop365.gitlab.io/vant/en-US/release-note-v4/](https://develop365.gitlab.io/vant/en-US/release-note-v4/)

### Tertiary (LOW confidence)
- None

## Metadata

**Confidence breakdown:**
- Standard stack: HIGH - 项目已有技术栈，仅新增 1 个 npm 包（已验证存在）
- Architecture: HIGH - 严格复用 Phase 1 已建立的分层模式
- Pitfalls: HIGH - 基于官方文档和已知 Vue 3 + AMap 集成问题

**Research date:** 2026-05-24
**Valid until:** 2026-06-24 (stable stack, no fast-moving dependencies)
