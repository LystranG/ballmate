# Phase 1: 用户系统 - Context

**Gathered:** 2026-05-23
**Status:** Ready for planning

<domain>
## Phase Boundary

完整的用户注册、登录、个人信息管理功能，同时搭建前后端基础项目架构（Vue 3 + Vant 4 前端、Go + Gin 后端、SQLite 数据库）。

</domain>

<decisions>
## Implementation Decisions

### 注册信息设计
- **D-01:** 手机号仅做格式校验（11位数字），不发送短信验证码
- **D-02:** 球类兴趣使用多选标签方式，预设选项：篮球、足球、羽毛球、乒乓球、网球、排球、棒球、橄榄球等
- **D-03:** 常住地使用省市区三级联动选择器（Vant Area 组件）
- **D-04:** 必填字段：账号、密码、手机号；选填字段：常住地、球类兴趣（注册后可在个人中心补充）

### 头像与个人资料
- **D-05:** 头像通过相册选图上传，后端以文件形式存储，数据库存路径
- **D-06:** 个人资料编辑使用统一表单页，所有字段一起修改后保存

### 页面流转与布局
- **D-07:** 登录和注册为独立页面，通过链接互相跳转
- **D-08:** 注册成功后自动登录并跳转首页
- **D-09:** 个人中心采用列表式菜单布局（顶部头像+昵称，下方菜单项）
- **D-10:** Phase 1 搭建底部导航栏框架（首页/发起约球/我的），后续 phase 填充内容

### 项目基础架构
- **D-11:** 后端 API 采用 RESTful 风格，统一响应格式 {code, message, data}，路径如 /api/v1/auth/login
- **D-12:** Go 后端使用 Gin 框架
- **D-13:** 登录状态使用 JWT，存储在 localStorage，请求时放入 Authorization header
- **D-14:** 前端使用 Axios 封装请求层，统一处理 token 注入、响应拦截、错误提示

### Claude's Discretion
- JWT 过期时间、刷新策略
- 前端目录结构和路由组织
- 数据库表结构设计
- Go 项目分层结构（handler/service/model）
- Vant 组件具体使用方式

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

No external specs — requirements fully captured in ROADMAP.md 和 REQUIREMENTS.md 中，实现决策在上方 decisions 中。

### 项目规划
- `.planning/PROJECT.md` — 项目定义、约束、技术栈
- `.planning/REQUIREMENTS.md` — 完整需求列表（AUTH-01~04, PROF-01~04 属于本 phase）
- `.planning/ROADMAP.md` — Phase 定义和成功标准

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- 无现有代码（greenfield 项目）

### Established Patterns
- 无（本 phase 将建立项目基础模式）

### Integration Points
- 本 phase 搭建的基础架构（路由、请求封装、token 管理、底部导航）将被后续 phase 复用

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

*Phase: 1-用户系统*
*Context gathered: 2026-05-23*
