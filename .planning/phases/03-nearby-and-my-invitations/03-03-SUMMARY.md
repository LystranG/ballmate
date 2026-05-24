---
phase: 03-nearby-and-my-invitations
plan: "03"
subsystem: my-invitations
tags: [backend, frontend, api, vue, gin, sqlite]
dependency_graph:
  requires: [03-01]
  provides: [my-invitations-api, my-invitations-page]
  affects: [profile-page, router]
tech_stack:
  added: []
  patterns: [vant-tabs-independent-state, gin-query-binding, gorm-join-query]
key_files:
  created:
    - frontend/src/views/my-invitations/index.vue
  modified:
    - backend/service/invitation.go
    - backend/handler/invitation.go
    - backend/router/router.go
    - frontend/src/api/invitation.js
    - frontend/src/views/profile/index.vue
    - frontend/src/router/index.js
decisions:
  - "/my route registered before /:id to prevent Gin treating 'my' as id param"
  - "joined query excludes creator_id = userID to avoid showing own invitations in joined tab"
  - "each Tab maintains independent list/loading/finished/page state — no reset on tab switch"
metrics:
  duration: 8min
  completed: "2026-05-24"
  tasks: 2
  files: 6
---

# Phase 03 Plan 03: My Invitations Summary

**One-liner:** GET /invitations/my API with created/joined query + Vue Tabs page with independent per-tab infinite scroll, accessible from profile menu.

## Tasks Completed

| Task | Name | Commit | Files |
|------|------|--------|-------|
| 1 | 后端 MyInvitations API | 08d344a | service/invitation.go, handler/invitation.go, router/router.go |
| 2 | 前端我的邀约页面 + 个人中心入口 + 路由 | 8e0b575 | api/invitation.js, views/my-invitations/index.vue, views/profile/index.vue, router/index.js |

## What Was Built

**Backend:**
- `service.MyInvitations(userID, listType, page, pageSize)`: created mode uses `WHERE creator_id = userID`, joined mode uses JOIN on participations with `creator_id != userID` to exclude own invitations. Both return `[]InvitationResponse` with computed status and participant count.
- `handler.MyInvitations`: binds query params via `ShouldBindQuery`, caps PageSize at 20 (T-03-08 data isolation).
- Router: `GET /my` registered after `/nearby` and before `/:id` to prevent Gin param collision.

**Frontend:**
- `getMyInvitations(params)` added to `api/invitation.js`.
- `views/my-invitations/index.vue`: NavBar with back arrow, Vant Tabs with two tabs (我发起的 / 我参与的), each tab has its own `reactive` state object (list, loading, finished, page) so switching tabs does not reset the other tab's data. Uses `InvitationCard` with `mode="status"` to show status tags instead of distance. Empty state via `van-empty` when list is empty and finished.
- `views/profile/index.vue`: "我的邀约" cell added before "编辑资料".
- `router/index.js`: `/my-invitations` route with `meta: { auth: true }` added after `/profile`.

## Deviations from Plan

None - plan executed exactly as written.

## Known Stubs

None - all data is wired to real API calls.

## Threat Flags

No new security surface beyond what the plan's threat model covers. The `WHERE creator_id = userID` / `participations.user_id = userID` conditions enforce data isolation at the DB layer (T-03-08 mitigated).

## Self-Check: PASSED
