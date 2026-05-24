# Phase 3: 附近约球与我的邀约 - Context

**Gathered:** 2026-05-24
**Status:** Ready for planning

<domain>
## Phase Boundary

用户可以基于浏览器定位发现附近邀约（按距离排序）、按时间排序和球类筛选、加入和退出邀约、查看个人邀约历史（发起的+参与的）。

</domain>

<decisions>
## Implementation Decisions

### 附近邀约列表展示
- **D-01:** 卡片式列表展示，每张卡片显示球类、地点、时间、人数进度、距离
- **D-02:** 下拉加载更多（Vant List 组件），不使用分页
- **D-03:** 顶部球类标签横向滚动栏筛选（全部/篮球/足球/...），点击切换
- **D-04:** 显示所有邀约（含已满员/已过期），不可加入的置灰标记

### 定位与距离计算
- **D-05:** 距离计算在后端完成（Haversine 公式），前端传经纬度参数，后端返回已排序列表
- **D-06:** 浏览器定位失败时，让用户手动选择位置（复用高德地图选点组件）
- **D-07:** 距离显示精确值（500m / 1.2km / 3.5km），不使用模糊范围
- **D-08:** 不限制距离范围，所有邀约都显示，距离远的自然排在后面

### 加入/退出交互
- **D-09:** 加入操作仅在详情页进行（非列表页），确保用户看过完整信息
- **D-10:** 加入需确认弹窗，确认后加入并显示 Toast 提示
- **D-11:** 退出需确认弹窗，确认后退出并显示 Toast 提示
- **D-12:** 详情页底部固定按钮，根据状态动态切换：未加入显示"加入"，已加入显示"退出"，已满/已过期置灰

### 我的邀约页面结构
- **D-13:** "我发起的"和"我参与的"使用 Tab 切换，同一页面内
- **D-14:** 从个人中心菜单进入"我的邀约"页面
- **D-15:** 复用附近邀约的卡片样式，但不显示距离，改为显示状态标签（等待中/已召集/已终止/已过期）

### Claude's Discretion
- 排序切换的具体 UI 形式（距离/时间切换按钮位置）
- 卡片内具体布局和样式细节
- 后端列表接口的分页参数设计（page/pageSize）
- 定位权限请求的时机和提示文案
- 空状态页面的具体展示

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### 项目规划
- `.planning/PROJECT.md` — 项目定义、约束、技术栈
- `.planning/REQUIREMENTS.md` — 完整需求列表（NEARBY-01~07, MY-01~02 属于本 phase）
- `.planning/ROADMAP.md` — Phase 定义和成功标准

### 前序 Phase 决策
- `.planning/phases/01-user-system/01-CONTEXT.md` — Phase 1 架构决策（API 风格、JWT、Axios 封装、项目结构）
- `.planning/phases/02-invite-core/02-CONTEXT.md` — Phase 2 邀约模型、状态流转、参与人设计、高德地图集成

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- `frontend/src/api/request.js` — Axios 封装，token 注入和错误处理
- `frontend/src/api/invitation.js` — 已有邀约 CRUD 接口，Phase 3 需新增列表/加入/退出接口
- `frontend/src/components/TabBar.vue` — 底部导航，首页 `/` 路由已预留
- `frontend/src/views/invitation/detail.vue` — 邀约详情页，需扩展加入/退出按钮
- `backend/handler/invitation.go` — 邀约 handler，需新增列表/加入/退出 handler
- `backend/service/invitation.go` — 邀约 service，需新增距离计算和列表查询逻辑
- `backend/model/invitation.go` — Invitation 模型，已有 latitude/longitude 字段
- `backend/model/participation.go` — Participation 模型，联合唯一索引已建立
- `frontend/src/views/create/index.vue` — 发起邀约页有 MapPicker 组件，定位失败时可复用

### Established Patterns
- 后端分层：handler → service → model（Phase 1 已建立）
- 前端结构：views/[feature]/index.vue + api/[feature].js
- GORM 模型使用 gorm.Model（含 ID, CreatedAt, UpdatedAt, DeletedAt 软删除）
- 前端路由守卫：meta.auth 控制需登录页面
- 统一响应格式 {code, message, data}
- 邀约状态：waiting/gathered/terminated 存储，expired 查询时计算

### Integration Points
- 首页 `/` 当前为空占位页，Phase 3 将其改造为附近邀约列表
- 详情页 `/invitation/:id` 已存在，需扩展底部加入/退出按钮
- 个人中心 `/profile` 菜单需新增"我的邀约"入口
- 后端路由 `/api/v1/invitations` 组需新增 GET 列表、POST join、DELETE leave 接口

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

*Phase: 3-附近约球与我的邀约*
*Context gathered: 2026-05-24*
