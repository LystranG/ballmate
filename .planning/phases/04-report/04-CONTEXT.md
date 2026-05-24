# Phase 4: 报告撰写 - Context

**Gathered:** 2026-05-24
**Status:** Ready for planning

<domain>
## Phase Boundary

完成期末考查报告（txt 格式），包含项目背景、需求分析、系统设计、核心代码说明、开源说明和总结展望。报告严格按模板结构编排，最终由用户手动转为 PDF 提交。

</domain>

<decisions>
## Implementation Decisions

### 报告输出格式
- **D-01:** 源文件使用纯 txt 格式，用户后续手动排版转 PDF
- **D-02:** 截图使用 [截图：xxx] 中文标签占位，用户后续替换为实际截图
- **D-03:** 图表采用混合方式：能用 mermaid 画的（E-R图、用例图、状态图、架构图）写 mermaid 代码块，UI 截图用文字占位

### 报告内容深度
- **D-04:** 核心代码精选展示，前后端各 3-4 段，每段不超过 30 行
- **D-05:** 展示模块选择：用户认证流程（JWT + Axios 拦截器）、邀约核心功能（发起表单 + 状态流转）、附近约球算法（Haversine 距离计算 + 列表查询）
- **D-06:** 需要的 mermaid 图：E-R 图（三表关系）、UML 用例图（参与者和功能）、系统架构图（前后端分层）、邀约状态流转图

### 报告结构编排
- **D-07:** 严格按模板 5 章结构：1-考查题目与要求、2-项目背景与需求、3-设计与开发、4-项目开源、5-总结与展望
- **D-08:** 总结部分对 AI 辅助采用模糊提及方式（"参考了相关技术文档和工具"），不具体展开
- **D-09:** 包含"项目开源"章节（可选加分项），简要描述在 Gitee/GitHub 上创建项目的过程

### 自动化程度
- **D-10:** 部分自动生成：数据库表结构从 model/*.go 提取生成表格，mermaid 图从代码关系生成
- **D-11:** 文字描述和代码说明部分由用户手写，工具只提供框架和提示

### Claude's Discretion
- 报告中文字描述的具体措辞和行文风格
- mermaid 图的具体布局和节点命名
- 代码片段的精确截取范围（在 30 行限制内）
- 各章节的篇幅分配

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### 报告模板与要求
- `期末考查报告-模板.docx` — 报告结构模板（5 章结构、表格格式、注意事项）
- `移动软件开发技术_考查题目及说明 (1).docx` — 考查题目要求、评分标准、提交规范

### 项目规划
- `.planning/PROJECT.md` — 项目定义、约束、技术栈
- `.planning/REQUIREMENTS.md` — 完整需求列表（报告需求分析章节的来源）
- `.planning/ROADMAP.md` — Phase 定义和成功标准

### 前序 Phase 决策
- `.planning/phases/01-user-system/01-CONTEXT.md` — Phase 1 架构决策（项目结构、技术选型理由）
- `.planning/phases/02-invite-core/02-CONTEXT.md` — Phase 2 邀约模型、状态流转设计
- `.planning/phases/03-nearby-and-my-invitations/03-CONTEXT.md` — Phase 3 距离算法、列表展示决策

</canonical_refs>

<code_context>
## Existing Code Insights

### 数据库模型（表结构提取来源）
- `backend/model/user.go` — User 表结构定义
- `backend/model/invitation.go` — Invitation 表结构定义（含 latitude/longitude）
- `backend/model/participation.go` — Participation 表结构定义（联合唯一索引）

### 核心代码展示来源
- `backend/handler/` — 后端接口 handler（认证、邀约、参与）
- `backend/service/invitation.go` — Haversine 距离计算、邀约状态逻辑
- `frontend/src/api/request.js` — Axios 封装 + JWT token 注入
- `frontend/src/views/create/index.vue` — 发起邀约表单 + 高德地图选点
- `frontend/src/views/home/index.vue` — 附近约球列表页

### 路由与架构信息
- `frontend/src/router/index.js` — 前端路由定义（页面结构来源）
- `backend/router.go` 或 `backend/main.go` — 后端路由注册（API 列表来源）

</code_context>

<specifics>
## Specific Ideas

- 代码片段需配简短中文注释说明关键逻辑
- E-R 图展示三表关系及外键
- 用例图区分"普通用户"和"发起人"两个角色
- 状态流转图展示四状态（等待中→已召集/已终止/已过期）及触发条件

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 4-报告撰写*
*Context gathered: 2026-05-24*
