# Phase 3: 附近约球与我的邀约 - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-05-24
**Phase:** 3-附近约球与我的邀约
**Areas discussed:** 附近邀约列表展示, 定位与距离计算, 加入/退出交互, 我的邀约页面结构

---

## 附近邀约列表展示

| Option | Description | Selected |
|--------|-------------|----------|
| 卡片式列表 | 每张卡片显示球类图标+地点+时间+人数进度+距离，视觉丰富 | ✓ |
| 简洁列表 | 紧凑单行展示，信息密度高 | |

**User's choice:** 卡片式列表

| Option | Description | Selected |
|--------|-------------|----------|
| 下拉加载更多 | 用户下拉到底部自动加载更多，流畅体验 | ✓ |
| 分页 | 底部分页按钮 | |
| 下拉刷新+上拉加载 | Vant List 组件原生支持 | |

**User's choice:** 下拉加载更多

| Option | Description | Selected |
|--------|-------------|----------|
| 标签栏筛选 | 页面顶部球类标签横向滚动栏 | ✓ |
| 下拉筛选 | 下拉菜单选择球类 | |

**User's choice:** 标签栏筛选

| Option | Description | Selected |
|--------|-------------|----------|
| 只显示可加入的 | 只显示等待中且未过期的邀约 | |
| 显示所有，置灰不可加入的 | 显示所有邀约，已满/已过期的置灰标记 | ✓ |

**User's choice:** 显示所有，置灰不可加入的

---

## 定位与距离计算

| Option | Description | Selected |
|--------|-------------|----------|
| 后端计算 | 后端接收经纬度，Haversine 公式计算距离并排序 | ✓ |
| 前端计算 | 后端返回所有邀约坐标，前端计算距离 | |

**User's choice:** 后端计算

| Option | Description | Selected |
|--------|-------------|----------|
| 降级为时间排序 | 不开启定位则不显示距离，按时间排序 | |
| 手动选择位置 | 让用户手动选择位置（复用高德地图选点） | ✓ |

**User's choice:** 手动选择位置

| Option | Description | Selected |
|--------|-------------|----------|
| 精确距离 | 显示 500m / 1.2km / 3.5km | ✓ |
| 模糊范围 | 显示 500m内 / 1km内 / 3km内 | |

**User's choice:** 精确距离

| Option | Description | Selected |
|--------|-------------|----------|
| 不限制 | 显示所有邀约，距离远的排后面 | ✓ |
| 有范围限制 | 默认显示 5km/10km 内 | |

**User's choice:** 不限制

---

## 加入/退出交互

| Option | Description | Selected |
|--------|-------------|----------|
| 列表页直接加入 | 卡片上直接有加入按钮 | |
| 仅详情页加入 | 必须进详情页才能加入 | ✓ |

**User's choice:** 仅详情页加入

| Option | Description | Selected |
|--------|-------------|----------|
| 直接加入 + Toast | 点击直接加入，显示 Toast | |
| 确认弹窗后加入 | 弹出确认弹窗，确认后加入 | ✓ |

**User's choice:** 确认弹窗后加入

| Option | Description | Selected |
|--------|-------------|----------|
| 确认弹窗退出 | 退出需确认弹窗 | ✓ |
| 直接退出 | 无确认直接退出 | |

**User's choice:** 确认弹窗退出

| Option | Description | Selected |
|--------|-------------|----------|
| 底部固定按钮 | 详情页底部固定按钮，根据状态动态切换 | ✓ |
| 内容区域内按钮 | 按钮放在页面内容区域内 | |

**User's choice:** 底部固定按钮

---

## 我的邀约页面结构

| Option | Description | Selected |
|--------|-------------|----------|
| Tab 切换 | 一个页面顶部 Tab 切换"我发起的"/"我参与的" | ✓ |
| 分开两个页面 | 两个独立页面，从个人中心分别进入 | |

**User's choice:** Tab 切换

| Option | Description | Selected |
|--------|-------------|----------|
| TabBar "我的"直接进入 | 底部导航"我的"就是我的邀约页 | |
| 个人中心菜单入口 | 个人中心菜单里加"我的邀约"入口 | ✓ |

**User's choice:** 个人中心菜单入口

| Option | Description | Selected |
|--------|-------------|----------|
| 复用卡片+状态标签 | 复用附近邀约卡片样式，不显示距离，改显示状态标签 | ✓ |
| 简化列表 | 只显示球类+时间+状态 | |

**User's choice:** 复用卡片+状态标签

---

## Claude's Discretion

- 排序切换的具体 UI 形式
- 卡片内具体布局和样式细节
- 后端列表接口的分页参数设计
- 定位权限请求的时机和提示文案
- 空状态页面的具体展示

## Deferred Ideas

None — discussion stayed within phase scope
