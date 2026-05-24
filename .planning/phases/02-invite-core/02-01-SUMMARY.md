---
phase: 02-invite-core
plan: 01
status: complete
started: 2026-05-24T13:15:00Z
completed: 2026-05-24T13:22:00Z
---

# Summary: 02-01 后端基础 + AMap 包安装

## What Was Built

邀约核心后端 CRUD API（5 个端点）和前端高德地图加载器安装。

## Key Files

### Created
- `backend/model/invitation.go` — Invitation GORM 模型
- `backend/model/participation.go` — Participation GORM 模型（联合唯一索引）
- `backend/service/invitation.go` — 邀约业务逻辑（含 computeStatus、maskPhone）
- `backend/handler/invitation.go` — 5 个 HTTP handler

### Modified
- `backend/main.go` — AutoMigrate 追加 Invitation + Participation
- `backend/router/router.go` — 注册 /api/v1/invitations 路由组
- `frontend/package.json` — 添加 @amap/amap-jsapi-loader@1.0.1

## Decisions Made

- 创建者自动加入 participations（计入人数），maxPeople=1 时创建即 gathered
- computeStatus 在查询时计算 expired，数据库不存储该状态
- 软删除邀约时同步软删除对应 participations

## Self-Check: PASSED

- [x] go build ./... 通过
- [x] 5 个 API 端点已注册（gin debug 输出确认）
- [x] 未认证请求返回 "未登录"（auth middleware 生效）
- [x] @amap/amap-jsapi-loader 已安装
