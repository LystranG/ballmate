# Phase 4: 报告撰写 - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-05-24
**Phase:** 4-报告撰写
**Areas discussed:** 报告输出格式, 报告内容深度, 报告结构编排, 自动化程度

---

## 报告输出格式

| Option | Description | Selected |
|--------|-------------|----------|
| Markdown + pandoc 转 PDF | 用 Markdown 写报告主体，截图用 ![placeholder] 占位，mermaid 图直接写在 md 中，最终用 pandoc 转 PDF | |
| 纯 txt 文件 | 直接生成 txt 文件，截图和图表用文字描述占位，手动转 PDF | ✓ |
| You decide | 由 Claude 决定最合适的格式 | |

**User's choice:** 纯 txt 文件
**Notes:** 用户表示"用纯txt文件即可，后面我手动来整"

### 截图和图表表示方式

| Option | Description | Selected |
|--------|-------------|----------|
| [截图：xxx] 中文标签 | 用 [截图：登录页面] 这样的中文标签标记，方便后续替换 | |
| mermaid 代码块 | 用 mermaid 代码块写图（用例图、E-R图、流程图），后面自己渲染截图替换 | |
| 混合方式 | 两者结合：能用 mermaid 画的用代码块，UI截图用文字占位 | ✓ |

**User's choice:** 混合方式

---

## 报告内容深度

### 核心代码展示范围

| Option | Description | Selected |
|--------|-------------|----------|
| 精选核心片段 | 每个模块选 1 个最关键的函数/组件，前后端各 3-4 段，每段不超过 30 行 | ✓ |
| 全模块覆盖 | 每个功能模块都展示前后端代码，每段 20-50 行 | |
| You decide | 由 Claude 判断哪些代码值得展示 | |

**User's choice:** 精选核心片段

### 展示模块选择

| Option | Description | Selected |
|--------|-------------|----------|
| 用户认证流程 | JWT 登录验证 + Axios 拦截器封装（前后端配合） | ✓ |
| 邀约核心功能 | 发起表单 + 高德地图选点 + 状态流转逻辑 | ✓ |
| 附近约球算法 | Haversine 距离计算 + 列表查询接口 | ✓ |
| 头像上传 | 文件上传 + 头像存储逻辑 | |

**User's choice:** 用户认证流程、邀约核心功能、附近约球算法

### Mermaid 图选择

| Option | Description | Selected |
|--------|-------------|----------|
| E-R 图 | 用 mermaid erDiagram 画 user/invitation/participation 三表关系 | ✓ |
| UML 用例图 | 用 mermaid 用例图展示主要参与者和功能 | ✓ |
| 状态流转图 | 用 mermaid flowchart 画邀约状态流转 | ✓ |
| 系统架构图 | 用 mermaid graph 画前后端模块分层架构 | ✓ |

**User's choice:** 全选

---

## 报告结构编排

### 章节结构

| Option | Description | Selected |
|--------|-------------|----------|
| 严格按模板 | 完全按模板的 5 章结构 | ✓ |
| 模板框架 + 灵活子节 | 保持模板大结构，子章节内根据实际内容调整 | |
| You decide | 由 Claude 判断怎么组织最合理 | |

**User's choice:** 严格按模板

### AI 辅助描述

| Option | Description | Selected |
|--------|-------------|----------|
| 坦诚描述 AI 辅助 | 如实描述使用 AI 辅助代码生成和调试 | |
| 模糊提及 | 只提及"参考了相关技术文档和工具"，不具体展开 | ✓ |
| 不提及 | 不在报告中提及 AI 辅助 | |

**User's choice:** 模糊提及

### 开源章节

| Option | Description | Selected |
|--------|-------------|----------|
| 包含（加分项） | 包含开源章节，简要描述在 Gitee/GitHub 上创建项目的过程 | ✓ |
| 不包含 | 跳过开源章节 | |

**User's choice:** 包含

---

## 自动化程度

| Option | Description | Selected |
|--------|-------------|----------|
| 尽可能自动生成 | 表结构、接口列表、mermaid 图、核心代码片段都从代码提取填入报告 | |
| 部分自动 | 只自动生成数据库表结构和 mermaid 图，文字描述和代码说明用户手写 | ✓ |
| 纯框架 | 只给章节框架和提示，用户自己填内容 | |

**User's choice:** 部分自动
**Notes:** 用户初始不理解"自动化"含义，解释后选择部分自动——表结构和 mermaid 图自动生成，文字描述手写

---

## Claude's Discretion

- 报告中文字描述的具体措辞和行文风格
- mermaid 图的具体布局和节点命名
- 代码片段的精确截取范围（在 30 行限制内）
- 各章节的篇幅分配

## Deferred Ideas

None
