---
phase: 03-nearby-and-my-invitations
plan: "02"
subsystem: invitation-join-leave
tags: [backend, frontend, join, leave, participation]
dependency_graph:
  requires: [03-01]
  provides: [join-leave-api, detail-bottom-action]
  affects: [frontend/src/views/invitation/detail.vue, backend/service/invitation.go]
tech_stack:
  added: []
  patterns: [service-layer-validation, computed-status, soft-delete-participation]
key_files:
  created: []
  modified:
    - backend/service/invitation.go
    - backend/handler/invitation.go
    - backend/router/router.go
    - frontend/src/api/invitation.js
    - frontend/src/views/invitation/detail.vue
decisions:
  - gathered 状态下已加入的人不能退出（满员意味着活动确认，退出会破坏已召集状态）
  - LeaveInvitation 校验 computeStatus 而非 inv.Status，确保 expired 邀约也无法退出
  - 加入后满员自动流转 gathered，退出后若之前是 gathered 则恢复 waiting
  - 底部操作栏仅非创建者可见（创建者有独立的 action-section）
metrics:
  duration: 8min
  completed: "2026-05-24"
  tasks: 2
  files: 5
---

# Phase 3 Plan 02: 加入/退出邀约 Summary

加入/退出邀约完整垂直切片：后端 JoinInvitation/LeaveInvitation service + handler + 路由，前端详情页底部固定操作栏，含确认弹窗和状态动态切换。

## Tasks Completed

| Task | Name | Commit | Files |
|------|------|--------|-------|
| 1 | 后端 Join/Leave API + has_joined 字段 | ea1adda | backend/service/invitation.go, backend/handler/invitation.go, backend/router/router.go |
| 2 | 前端详情页底部加入/退出按钮 | ba39018 | frontend/src/api/invitation.js, frontend/src/views/invitation/detail.vue |

## What Was Built

**后端（Task 1）：**
- `service.JoinInvitation(invitationID, userID uint) error`：校验邀约存在、状态为 waiting、未满员，创建 Participation 记录（联合唯一索引防重复），加入后满员自动流转 gathered
- `service.LeaveInvitation(invitationID, userID uint) error`：校验邀约存在、非创建者、状态为 waiting，软删除 Participation，退出后若之前是 gathered 则恢复 waiting
- `service.GetInvitation` 签名改为 `(id, userID uint)`，新增 `has_joined bool` 字段
- `handler.JoinInvitation` / `handler.LeaveInvitation`：解析参数，区分 400（业务错误）/ 403（权限错误）
- 路由：`POST /:id/join`、`DELETE /:id/join`

**前端（Task 2）：**
- `api/invitation.js`：追加 `joinInvitation`、`leaveInvitation`
- `detail.vue`：底部固定 `.bottom-action` 栏，非创建者可见
  - 未加入：`van-button type="primary"`，`canJoin` 控制 disabled（waiting 且未满员）
  - 已加入：`van-button type="warning"`，`canLeave` 控制 disabled（仅 waiting 可退出）
  - 操作前 `showDialog` 确认弹窗，操作后 `showToast` 提示并 `loadDetail()` 刷新

## Deviations from Plan

None - plan executed exactly as written.

## Threat Model Coverage

| Threat ID | Mitigation | Status |
|-----------|-----------|--------|
| T-03-04 | service 层 COUNT 校验 + 加入后再次 COUNT 确认 | Implemented |
| T-03-06 | service.LeaveInvitation 校验 CreatorID != userID | Implemented |
| T-03-07 | 联合唯一索引 idx_inv_user，Create 失败返回友好错误 | Implemented |

## Known Stubs

None.

## Self-Check: PASSED

- backend/service/invitation.go: FOUND (JoinInvitation, LeaveInvitation, updated GetInvitation)
- backend/handler/invitation.go: FOUND (JoinInvitation, LeaveInvitation handlers)
- backend/router/router.go: FOUND (POST /:id/join, DELETE /:id/join routes)
- frontend/src/api/invitation.js: FOUND (joinInvitation, leaveInvitation)
- frontend/src/views/invitation/detail.vue: FOUND (bottom-action bar)
- Commit ea1adda: FOUND
- Commit ba39018: FOUND
- go build ./...: BUILD_OK
- npm run build: built in 186ms
