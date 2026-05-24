# Phase 2: 邀约核心 - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-05-24
**Phase:** 2-邀约核心
**Areas discussed:** 发起邀约表单, 邀约状态流转, 参与人展示, 邀约数据模型

---

## 发起邀约表单

### 地图选点

| Option | Description | Selected |
|--------|-------------|----------|
| 地图组件选点 | 用户点击地图选择坐标，支持搜索地址。需要引入地图 SDK（如高德/腾讯） | ✓ |
| 文字输入+地理编码 | 用户输入地址文字，后端调用地理编码 API 转坐标。前端无需地图 SDK | |
| 文字+浏览器定位 | 用户手动填写地址文字，同时用浏览器定位获取当前坐标作为地点坐标。最简单但不精确 | |

**User's choice:** 地图组件选点

### 地图 SDK 选择

| Option | Description | Selected |
|--------|-------------|----------|
| 高德地图 | 免费额度充足，文档完善，支持 JS API 和 Web 服务 API | ✓ |
| 腾讯地图 | 免费额度充足，微信生态集成好，支持 JS API | |
| You decide | 哪个集成简单用哪个 | |

**User's choice:** 高德地图

### 时间选择

| Option | Description | Selected |
|--------|-------------|----------|
| 日期+时间分开选 | 用户分别选择日期和时间，精确到分钟。Vant DatePicker + TimePicker | |
| 日期时间一体选择器 | 一个组件同时选日期和时间。Vant DatePicker 支持 type='datetime' | ✓ |
| 快捷选项+自定义 | 提供"今天下午""明天上午"等快捷选项 + 自定义时间 | |

**User's choice:** 日期时间一体选择器

### 人数设置

| Option | Description | Selected |
|--------|-------------|----------|
| 精确人数 | 用户输入具体数字（如 6 人），凑齐后自动召集 | |
| 预设档位 | 提供预设选项（2人/4人/6人/8人/10人），快速选择 | ✓ |
| 人数范围 | 设置最少和最多人数范围，灵活但复杂 | |

**User's choice:** 预设档位

---

## 邀约状态流转

### 状态机设计

| Option | Description | Selected |
|--------|-------------|----------|
| 三状态线性 | 等待中→已召集（人数达标自动触发）、等待中→已终止（发起人手动）。已召集/已终止为终态。 | |
| 四状态（含过期） | 增加"已过期"状态，活动时间过后自动标记。需要定时任务或查询时判断。 | ✓ |
| 动态计算 | 不存状态字段，根据当前人数和时间实时计算状态。简单但查询时计算量大。 | |

**User's choice:** 四状态（含过期）

### 过期判断

| Option | Description | Selected |
|--------|-------------|----------|
| 查询时判断 | 查询邀约列表时，对比当前时间和活动时间，过期的标记为已过期。无需定时任务，简单可靠。 | ✓ |
| 定时任务扫描 | 后端定时任务扫描并更新状态字段。状态一致但需要额外机制。 | |

**User's choice:** 查询时判断

### 终止与删除

| Option | Description | Selected |
|--------|-------------|----------|
| 终止可见+删除隐藏 | 终止=标记为已终止（仍可见），删除=软删除（不再显示）。两个操作都需确认弹窗。 | |
| 终止→可删除链路 | 终止和删除都是软删除，只是状态不同。终止后可以再删除。 | ✓ |
| You decide | 你来决定合理的方式 | |

**User's choice:** 终止→可删除链路

---

## 参与人展示

### 参与人入口

| Option | Description | Selected |
|--------|-------------|----------|
| 详情页内嵌 | 邀约详情页底部直接展示参与人列表，无需额外跳转 | |
| 卡片摘要+详情查看 | 邀约卡片显示"N/M人"，点击进入详情页才看到具体参与人 | ✓ |
| You decide | 你来决定合理的展示方式 | |

**User's choice:** 卡片摘要+详情查看

### 脱敏规则

| Option | Description | Selected |
|--------|-------------|----------|
| 后端脱敏 | 后端返回时就已经脱敏（138****5678），前端直接展示。安全性更高。 | ✓ |
| 前端脱敏 | 后端返回完整手机号，前端显示时脱敏。简单但有泄露风险。 | |
| 发起人专属可见 | 只有发起人能看到脱敏手机号，其他参与人只能看到昵称 | |

**User's choice:** 后端脱敏

---

## 邀约数据模型

### 参与关系存储

| Option | Description | Selected |
|--------|-------------|----------|
| 关联表 | 独立的 participations 表（invitation_id, user_id, joined_at），标准多对多关系。查询灵活。 | ✓ |
| 嵌入字段 | 邀约表中存参与人 ID 列表（JSON 或逗号分隔）。简单但查询不便。 | |

**User's choice:** 关联表

### 坐标存储

| Option | Description | Selected |
|--------|-------------|----------|
| 两个浮点字段 | 邀约表中分别存 latitude 和 longitude 两个 float 字段。简单直接。 | ✓ |
| 单字符串字段 | 存为一个字符串字段（如 "39.9,116.4"）。简化存储但查询时需解析。 | |
| You decide | 你来决定 | |

**User's choice:** 两个浮点字段

---

## Claude's Discretion

- 邀约表单的具体校验规则和错误提示
- 邀约详情页的具体布局
- API 路由设计和接口命名
- 数据库表的具体字段命名和索引策略
- 球类选择的 UI 交互细节（Tag 组件 vs Radio 组件）

## Deferred Ideas

None — discussion stayed within phase scope
