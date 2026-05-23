# 掌上约球

## What This Is

一个面向校园和社区的球友召集平台移动Web应用。用户可以发起球类运动邀约、浏览附近的约球活动并加入，方便快速凑齐球友一起运动。

## Core Value

让用户能快速找到附近想打球的人，凑齐人数开始运动。

## Requirements

### Validated

(None yet — ship to validate)

### Active

- [ ] 用户注册登录（账号密码、手机号、常住地、球类兴趣）
- [ ] 附近约球列表（基于浏览器定位，按距离/时间/球类筛选排序）
- [ ] 发起邀约（球类、时间、地点、人数，支持地图选点）
- [ ] 邀约状态管理（等待中/已召集/已终止）
- [ ] 查看参与人信息（手机号中间4位脱敏）
- [ ] 个人信息维护（头像、资料、密码、退出登录）

### Out of Scope

- 微信小程序 — 本项目使用Vue+Vant的H5方案
- 实时聊天 — 复杂度高，非核心功能
- 支付功能 — 约球场景不涉及线上支付
- 推送通知 — v1不做，后续可加

## Context

- 课程期末考查项目，需要完整实现功能并撰写报告
- 前端使用Vue 3 + Vant 4组件库，移动端H5界面
- 后端使用Go语言，数据库SQLite
- 定位方案：浏览器Geolocation API获取经纬度
- 展示方案：列表为主，发起邀约时支持地图选点定位
- 复杂逻辑需要简短中文注释
- 需要同时产出期末考查报告（txt格式）

## Constraints

- **Tech Stack**: Vue 3 + Vant 4 (frontend), Go (backend), SQLite (database) — 课程要求
- **Structure**: frontend/ 和 backend/ 分目录存放
- **Comments**: 复杂逻辑和函数需简短中文注释
- **Report**: 需产出txt格式报告，含截图占位符和mermaid图

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Vue 3 + Vant 4 | 成熟的移动端组件库，适合H5开发 | — Pending |
| Go + SQLite | 轻量后端，部署简单，适合课程项目 | — Pending |
| 浏览器定位 | 简单直接，不依赖第三方地图SDK | — Pending |
| 列表为主展示 | 降低复杂度，核心是信息展示而非地图交互 | — Pending |

## Evolution

This document evolves at phase transitions and milestone boundaries.

**After each phase transition** (via `/gsd-transition`):
1. Requirements invalidated? → Move to Out of Scope with reason
2. Requirements validated? → Move to Validated with phase reference
3. New requirements emerged? → Add to Active
4. Decisions to log? → Add to Key Decisions
5. "What This Is" still accurate? → Update if drifted

**After each milestone** (via `/gsd:complete-milestone`):
1. Full review of all sections
2. Core Value check — still the right priority?
3. Audit Out of Scope — reasons still valid?
4. Update Context with current state

---
*Last updated: 2026-05-23 after initialization*
