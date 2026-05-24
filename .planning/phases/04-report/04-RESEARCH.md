# Phase 4: 报告撰写 - Research

**Researched:** 2026-05-24
**Domain:** 期末考查报告生成（txt格式，含mermaid图和代码片段）
**Confidence:** HIGH

## Summary

本阶段为纯文档生成任务，不涉及代码开发或外部依赖安装。核心工作是按照模板严格的5章结构，从已完成的代码库中提取数据模型、核心代码片段和架构信息，组装成一份完整的期末考查报告（txt格式）。

报告模板已通过 pandoc 解析确认，结构为：第1章考查题目与要求、第2章项目背景与需求（含UML用例图）、第3章设计与开发（含E-R图、数据库表结构、前后端核心代码）、第4章项目开源（可选加分）、第5章总结与展望。

**Primary recommendation:** 按模板5章结构生成txt文件，mermaid图直接嵌入代码块，UI截图用中文占位符标记，核心代码精选6-8段（前后端各3-4段，每段不超30行）。

<user_constraints>
## User Constraints (from CONTEXT.md)

### Locked Decisions
- D-01: 源文件使用纯 txt 格式，用户后续手动排版转 PDF
- D-02: 截图使用 [截图：xxx] 中文标签占位，用户后续替换为实际截图
- D-03: 图表采用混合方式：能用 mermaid 画的（E-R图、用例图、状态图、架构图）写 mermaid 代码块，UI 截图用文字占位
- D-04: 核心代码精选展示，前后端各 3-4 段，每段不超过 30 行
- D-05: 展示模块选择：用户认证流程（JWT + Axios 拦截器）、邀约核心功能（发起表单 + 状态流转）、附近约球算法（Haversine 距离计算 + 列表查询）
- D-06: 需要的 mermaid 图：E-R 图（三表关系）、UML 用例图（参与者和功能）、系统架构图（前后端分层）、邀约状态流转图
- D-07: 严格按模板 5 章结构
- D-08: 总结部分对 AI 辅助采用模糊提及方式
- D-09: 包含"项目开源"章节（可选加分项）
- D-10: 部分自动生成：数据库表结构从 model/*.go 提取，mermaid 图从代码关系生成
- D-11: 文字描述和代码说明部分由用户手写，工具只提供框架和提示

### Claude's Discretion
- 报告中文字描述的具体措辞和行文风格
- mermaid 图的具体布局和节点命名
- 代码片段的精确截取范围（在 30 行限制内）
- 各章节的篇幅分配

### Deferred Ideas (OUT OF SCOPE)
None
</user_constraints>

## Template Requirements (from 期末考查报告-模板.docx)

### Exact Chapter Structure

| 章节 | 标题 | 内容要求 |
|------|------|----------|
| 1 | 考查题目与要求 | 题目名称 + 主要功能列表（直接引用题目要求） |
| 2 | 项目背景与需求 | 2.1 背景与意义 + 2.2 功能需求（需UML用例图） |
| 3 | 设计与开发 | 3.1 数据库设计（E-R图 + 表结构表格）+ 3.2 前端开发 + 3.3 后端开发 |
| 4 | 项目开源（可选） | 描述在Gitee上创建并维护开源项目的流程 |
| 5 | 总结与展望 | 收获与能力提升 + AI辅助总结 |

### Formatting Rules (from template)
- 封面信息：学院、专业、姓名、学号、班级、日期
- 数据库表结构格式：字段 / 类型 / 注释 / 外键关系 / 备注（5列表格）
- 代码展示规则：按功能模块描述，核心代码+功能截图+关键代码含义说明
- 连续代码不能超过1页（约30-40行）
- 注意事项（需在正式报告中删除）

### Grading Criteria (from 考查题目及说明)
1. **功能完善性** — 完整应用功能 + 良好UI/UX + 可选特色功能
2. **代码质量** — 编程规范 + 结构清晰 + 注释充分 + 可读性/可维护性
3. **开源（可选加分）** — GitHub/Gitee开源 + 合适许可证

### Submission Requirements
- PDF格式：学号-姓名-班级-期末考查报告.pdf
- 源代码ZIP：学号-姓名-班级-源代码.zip（不含node_modules和依赖包）
- 5分钟讲解视频（本人出镜，功能演示+核心代码讲解）

## Code Snippets Catalog

### Backend Snippets (3-4 segments)

| # | File | Lines | What It Shows | Recommended Excerpt |
|---|------|-------|---------------|---------------------|
| B1 | `backend/utils/jwt.go` | 12-37 | JWT token签发与解析（完整文件26行） | 全文件：GenerateToken + ParseToken |
| B2 | `backend/middleware/auth.go` | 12-40 | JWT认证中间件（Bearer token提取+验证） | 全文件：AuthMiddleware函数 |
| B3 | `backend/service/invitation.go` | 17-28 | Haversine距离计算公式 | haversineDistance函数（12行） |
| B4 | `backend/service/invitation.go` | 319-344 | 加入邀约+事务保证+自动状态流转 | JoinInvitation函数核心逻辑（约25行） |

### Frontend Snippets (3-4 segments)

| # | File | Lines | What It Shows | Recommended Excerpt |
|---|------|-------|---------------|---------------------|
| F1 | `frontend/src/api/request.js` | 1-34 | Axios封装+JWT注入+401自动跳转 | 全文件（34行，略超但逻辑完整） |
| F2 | `frontend/src/views/create/index.vue` | 109-206 | 发起邀约表单script（reactive表单+地图选点+提交） | script setup部分精选（约30行） |
| F3 | `frontend/src/router/index.js` | 51-68 | 路由守卫（token检查+自动跳转） | beforeEach导航守卫（约15行） |
| F4 | `frontend/src/views/home/index.vue` | 90-170 | 附近列表页核心逻辑（定位+分页加载+筛选排序） | script setup精选（约28行） |

### Snippet Selection Rationale
- **D-05 模块1（JWT认证）**: B1 + B2 + F1 完整展示前后端认证链路
- **D-05 模块2（邀约核心）**: B4 + F2 展示创建+加入的完整流程
- **D-05 模块3（附近算法）**: B3 + F4 展示距离计算+列表查询

## Data Model Summary

### User 表 (users)

| 字段 | 类型 | 注释 | 外键关系 | 备注 |
|------|------|------|----------|------|
| id | uint (PK) | 主键 | — | GORM自动生成 |
| username | varchar(50) | 用户名 | — | 唯一索引，不可重复 |
| password | varchar(100) | 密码 | — | bcrypt加密存储 |
| phone | varchar(11) | 手机号 | — | 11位数字 |
| nickname | varchar(50) | 昵称 | — | 默认为用户名 |
| avatar | varchar(255) | 头像路径 | — | 上传文件路径 |
| location | varchar(100) | 常住地 | — | — |
| sports | varchar(255) | 球类兴趣 | — | 逗号分隔 |
| created_at | datetime | 创建时间 | — | GORM自动 |
| updated_at | datetime | 更新时间 | — | GORM自动 |
| deleted_at | datetime | 软删除时间 | — | GORM软删除 |

### Invitation 表 (invitations)

| 字段 | 类型 | 注释 | 外键关系 | 备注 |
|------|------|------|----------|------|
| id | uint (PK) | 主键 | — | GORM自动生成 |
| creator_id | uint | 创建者ID | → users.id | 非空，有索引 |
| sport_type | varchar(20) | 球类名称 | — | 篮球/足球/羽毛球等 |
| activity_time | datetime | 活动时间 | — | 非空 |
| address | varchar(255) | 活动地点 | — | 文字地址 |
| latitude | float64 | 纬度 | — | 用于距离计算 |
| longitude | float64 | 经度 | — | 用于距离计算 |
| max_people | int | 最大人数 | — | 非空 |
| status | varchar(20) | 状态 | — | waiting/gathered/terminated |
| created_at | datetime | 创建时间 | — | GORM自动 |
| updated_at | datetime | 更新时间 | — | GORM自动 |
| deleted_at | datetime | 软删除时间 | — | GORM软删除 |

### Participation 表 (participations)

| 字段 | 类型 | 注释 | 外键关系 | 备注 |
|------|------|------|----------|------|
| id | uint (PK) | 主键 | — | GORM自动生成 |
| invitation_id | uint | 邀约ID | → invitations.id | 联合唯一索引 |
| user_id | uint | 用户ID | → users.id | 联合唯一索引 |
| created_at | datetime | 加入时间 | — | GORM自动 |
| updated_at | datetime | 更新时间 | — | GORM自动 |
| deleted_at | datetime | 软删除时间 | — | GORM软删除 |

**Key relationships:**
- User 1:N Invitation (creator_id)
- User N:M Invitation (through Participation)
- 联合唯一索引 `idx_inv_user` 防止重复参与

## Architecture Map

### System Layers

```
┌─────────────────────────────────────────────────┐
│                   用户浏览器                       │
│  Vue 3 + Vant 4 + Vue Router + Axios + 高德地图   │
└──────────────────────┬──────────────────────────┘
                       │ HTTP (RESTful API)
                       │ Authorization: Bearer <JWT>
┌──────────────────────▼──────────────────────────┐
│                   Go 后端                         │
│  Gin Framework + CORS中间件 + JWT认证中间件        │
│  ┌─────────┐  ┌──────────┐  ┌────────────┐     │
│  │ Handler │→ │ Service  │→ │   Model    │     │
│  │ (路由)  │  │ (业务)   │  │  (GORM)    │     │
│  └─────────┘  └──────────┘  └─────┬──────┘     │
└────────────────────────────────────┼────────────┘
                                     │ GORM ORM
                       ┌─────────────▼─────────────┐
                       │     SQLite 数据库           │
                       │  users / invitations /     │
                       │  participations            │
                       └───────────────────────────┘
```

### Frontend Architecture
- **框架**: Vue 3 (Composition API, `<script setup>`)
- **UI组件库**: Vant 4 (按需导入 via unplugin-vue-components)
- **路由**: Vue Router 5 (History模式 + 导航守卫)
- **HTTP**: Axios (请求/响应拦截器)
- **地图**: 高德地图 JS API 2.0 (@amap/amap-jsapi-loader)
- **构建**: Vite 8

### Backend Architecture
- **语言**: Go 1.26
- **Web框架**: Gin 1.12
- **ORM**: GORM 1.31 + SQLite driver
- **认证**: golang-jwt/jwt v5 (HS256)
- **密码**: bcrypt (golang.org/x/crypto)
- **跨域**: gin-contrib/cors

### API Routes Summary

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | /api/v1/auth/register | No | 用户注册 |
| POST | /api/v1/auth/login | No | 用户登录 |
| GET | /api/v1/user/profile | Yes | 获取个人信息 |
| PUT | /api/v1/user/profile | Yes | 修改个人资料 |
| PUT | /api/v1/user/password | Yes | 修改密码 |
| POST | /api/v1/user/avatar | Yes | 上传头像 |
| POST | /api/v1/invitations | Yes | 创建邀约 |
| GET | /api/v1/invitations/nearby | Yes | 附近邀约列表 |
| GET | /api/v1/invitations/my | Yes | 我的邀约 |
| GET | /api/v1/invitations/:id | Yes | 邀约详情 |
| PUT | /api/v1/invitations/:id/terminate | Yes | 终止邀约 |
| DELETE | /api/v1/invitations/:id | Yes | 删除邀约 |
| GET | /api/v1/invitations/:id/participants | Yes | 参与人列表 |
| POST | /api/v1/invitations/:id/join | Yes | 加入邀约 |
| DELETE | /api/v1/invitations/:id/join | Yes | 退出邀约 |

### Frontend Pages

| Route | Component | Description |
|-------|-----------|-------------|
| /login | views/login/index.vue | 登录页 |
| /register | views/register/index.vue | 注册页 |
| / | views/home/index.vue | 首页（附近约球列表） |
| /create | views/create/index.vue | 发起邀约 |
| /invitation/:id | views/invitation/detail.vue | 邀约详情 |
| /my-invitations | views/my-invitations/index.vue | 我的邀约 |
| /profile | views/profile/index.vue | 个人中心 |
| /profile/edit | views/profile/edit.vue | 编辑资料 |
| /profile/password | views/profile/password.vue | 修改密码 |

## State Flow Analysis

### Invitation Status Transitions

```
States: waiting / gathered / terminated / expired(动态计算)

                    ┌──────────────────────────────────┐
                    │                                  │
    创建邀约         ▼         人员凑齐                  │ 有人退出
  ──────────→  [waiting] ──────────────→ [gathered] ───┘
                    │
                    │ 发起人终止
                    ▼
              [terminated]
                    │
                    │ 发起人删除
                    ▼
               (soft delete)

    [waiting] + activity_time < now → [expired] (动态计算，不存储)
```

### Transition Rules (from service/invitation.go)

| From | To | Trigger | Code Location |
|------|----|---------|---------------|
| — | waiting | 创建邀约 | CreateInvitation: `Status: "waiting"` |
| waiting | gathered | 参与人数 >= max_people | JoinInvitation: 事务内count检查 |
| gathered | waiting | 有人退出导致人数不足 | LeaveInvitation: `Update("status", "waiting")` |
| waiting | terminated | 发起人主动终止 | TerminateInvitation: 仅创建者+仅waiting |
| terminated | (deleted) | 发起人删除 | DeleteInvitation: 仅创建者+仅terminated |
| waiting | expired | activity_time已过（动态） | computeStatus: 不写入DB |

### Key Business Rules
- 创建者自动加入参与人（CreateInvitation中自动创建Participation）
- 创建者不能退出自己的邀约
- 加入邀约使用事务保证并发安全（防止超员）
- expired状态不存储，查询时根据activity_time动态计算
- 手机号脱敏：`138****5678`（maskPhone函数）

## Mermaid Diagrams Specification

### D-06-1: E-R Diagram (3 tables)
- User ←1:N→ Invitation (creator_id)
- User ←N:M→ Invitation (through Participation)
- 展示所有字段和关系

### D-06-2: UML Use Case Diagram
- 参与者：普通用户、发起人（同一用户的两种角色）
- 用例：注册/登录、浏览附近邀约、加入邀约、发起邀约、管理邀约、查看参与人、个人信息维护

### D-06-3: System Architecture Diagram
- 三层：前端(Vue3+Vant4) → 后端(Go+Gin) → 数据库(SQLite)
- 标注通信协议和认证方式

### D-06-4: Invitation State Flow Diagram
- 四状态 + 触发条件（如上State Flow Analysis所述）

## Risks/Gaps

| # | Risk | Impact | Mitigation |
|---|------|--------|------------|
| 1 | txt格式无法真正渲染mermaid图 | 用户需手动将mermaid代码粘贴到渲染工具生成图片 | 报告中标注"请将以下mermaid代码渲染为图片后插入" |
| 2 | 代码片段30行限制可能截断逻辑 | 部分函数需精心裁剪 | 选择自包含的短函数，或截取核心逻辑段 |
| 3 | request.js为34行略超30行限制 | 可能需要删减注释 | 保留核心逻辑，压缩空行 |
| 4 | 用户需手动填写封面信息 | 工具无法获取学号等个人信息 | 用占位符标记 [请填写：学号] |
| 5 | 开源章节需要实际Gitee截图 | 工具无法生成 | 用 [截图：Gitee仓库页面] 占位 |

## Assumptions Log

| # | Claim | Section | Risk if Wrong |
|---|-------|---------|---------------|
| A1 | mermaid erDiagram语法支持中文标签 | Mermaid Diagrams | 图表渲染失败，需改用英文标签 |
| A2 | 报告txt文件使用UTF-8编码 | Template Requirements | 中文乱码 |

## Sources

### Primary (HIGH confidence)
- `期末考查报告-模板.docx` — pandoc解析确认5章结构和表格格式
- `移动软件开发技术_考查题目及说明 (1).docx` — pandoc解析确认评分标准和提交要求
- 项目源代码直接读取 — 所有代码片段和数据模型均从实际文件提取

### Secondary (MEDIUM confidence)
- `.planning/CONTEXT.md` — Phase 4 决策记录
- `.planning/PROJECT.md` — 项目约束和技术栈定义

## Metadata

**Confidence breakdown:**
- Template structure: HIGH - 直接从docx解析获得
- Code snippets: HIGH - 直接从源文件读取
- Data models: HIGH - 直接从Go model文件提取
- Architecture: HIGH - 从main.go/router.go/package.json确认
- State flow: HIGH - 从service层代码逻辑分析

**Research date:** 2026-05-24
**Valid until:** 2026-06-24 (stable — report template and codebase are frozen)
