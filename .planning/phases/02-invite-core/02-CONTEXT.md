# Phase 2: 邀约核心 - Context

**Gathered:** 2026-05-24
**Status:** Ready for planning

<domain>
## Phase Boundary

用户可以发起球类邀约（选择球类、时间、地点、人数），管理邀约状态（等待中→已召集/已终止/已过期），查看参与人信息（手机号脱敏显示），终止或删除自己的邀约。

</domain>

<decisions>
## Implementation Decisions

### 发起邀约表单
- **D-01:** 地点选择使用高德地图 JS API 组件选点，用户在地图上点击选择坐标并搜索地址
- **D-02:** 时间选择使用 Vant DatePicker type='datetime' 一体选择器，精确到分钟
- **D-03:** 人数设置使用预设档位（2人/4人/6人/8人/10人），快速选择
- **D-04:** 球类选择复用 Phase 1 已有的球类标签列表（篮球、足球、羽毛球、乒乓球、网球、排球等），单选

### 邀约状态流转
- **D-05:** 四状态设计：等待中→已召集（人数达标自动触发）、等待中→已终止（发起人手动）、等待中→已过期（活动时间已过）。已召集/已终止/已过期为终态
- **D-06:** "已过期"通过查询时判断实现（对比当前时间和活动时间），无需定时任务
- **D-07:** 终止和删除为链路关系：等待中的邀约可以终止，终止后可以删除（软删除，不再显示）。两个操作都需确认弹窗

### 参与人展示
- **D-08:** 邀约卡片显示"N/M人"摘要，点击进入详情页查看具体参与人列表
- **D-09:** 手机号脱敏在后端处理（返回时已是 138****5678 格式），前端直接展示
- **D-10:** 只有发起人可以查看参与人的脱敏手机号（PART-01~03 要求）

### 邀约数据模型
- **D-11:** 邀约与参与人使用独立关联表（participations），标准多对多关系
- **D-12:** 地理坐标使用两个浮点字段（latitude, longitude）存储，便于 Phase 3 距离计算

### Claude's Discretion
- 邀约表单的具体校验规则和错误提示
- 邀约详情页的具体布局
- API 路由设计和接口命名
- 数据库表的具体字段命名和索引策略
- 球类选择的 UI 交互细节（Tag 组件 vs Radio 组件）

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### 项目规划
- `.planning/PROJECT.md` — 项目定义、约束、技术栈
- `.planning/REQUIREMENTS.md` — 完整需求列表（INVITE-01~06, PART-01~03 属于本 phase）
- `.planning/ROADMAP.md` — Phase 定义和成功标准

### 前序 Phase 决策
- `.planning/phases/01-user-system/01-CONTEXT.md` — Phase 1 架构决策（API 风格、JWT、Axios 封装、项目结构）

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- `frontend/src/api/request.js` — Axios 封装，token 注入和错误处理，直接复用
- `frontend/src/components/TabBar.vue` — 底部导航已预留 `/create` 路由
- `backend/router/router.go` — 路由注册模式，新增邀约路由组
- `backend/handler/response.go` — 统一响应格式
- `backend/middleware/auth.go` — JWT 认证中间件，邀约接口需认证
- `backend/model/user.go` — User 模型，邀约关联 user_id

### Established Patterns
- 后端分层：handler → service → model（Phase 1 已建立）
- 前端结构：views/[feature]/index.vue + api/[feature].js
- GORM 模型使用 gorm.Model（含 ID, CreatedAt, UpdatedAt, DeletedAt 软删除）
- 前端路由守卫：meta.auth 控制需登录页面

### Integration Points
- 邀约创建后关联当前登录用户（从 JWT 中获取 user_id）
- 参与人查询需 JOIN users 表获取昵称和手机号
- TabBar `/create` 路由需要对应发起邀约页面
- Phase 3 将复用邀约列表查询接口（加距离排序）

</code_context>

<specifics>
## Specific Ideas

No specific requirements — open to standard approaches

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 2-邀约核心*
*Context gathered: 2026-05-24*
