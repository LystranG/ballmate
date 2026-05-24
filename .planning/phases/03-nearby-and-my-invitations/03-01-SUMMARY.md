---
phase: 03-nearby-and-my-invitations
plan: "01"
subsystem: nearby-invitations
tags: [geolocation, haversine, vant-list, infinite-scroll, filtering]
dependency_graph:
  requires: [phase-02-invitation-crud]
  provides: [nearby-invitations-api, invitation-card-component, geolocation-composable]
  affects: [home-page, invitation-detail]
tech_stack:
  added: []
  patterns:
    - Haversine SQL expression in GORM (SQLite-compatible, no RADIANS)
    - Go-layer fallback distance calculation when SQL math functions unavailable
    - useGeolocation composable pattern for browser Geolocation API
    - Vant List infinite scroll with v-if guard against premature load
key_files:
  created:
    - backend/service/invitation.go (NearbyInvitationResponse, ListNearbyInvitations, haversineDistance)
    - frontend/src/composables/useGeolocation.js
    - frontend/src/utils/format.js
    - frontend/src/components/InvitationCard.vue
  modified:
    - backend/handler/invitation.go (ListNearby handler, listNearbyRequest)
    - backend/router/router.go (GET /nearby registered before /:id)
    - frontend/src/api/invitation.js (getNearbyInvitations added)
    - frontend/src/views/home/index.vue (full rewrite)
decisions:
  - "Haversine SQL expression uses manual pi/180 conversion instead of RADIANS() — SQLite lacks RADIANS"
  - "Go-layer fallback implemented for SQLite math function unavailability (all-zero distance detection)"
  - "v-if='hasLocation' on Vant List prevents premature load before geolocation completes"
  - "PageSize capped at 20 in handler (T-03-02 DoS mitigation)"
  - "GET /nearby registered before /:id in router to prevent Gin treating 'nearby' as id param"
metrics:
  duration: 12min
  completed_date: "2026-05-24"
  tasks_completed: 2
  files_changed: 8
---

# Phase 3 Plan 01: 附近邀约列表完整垂直切片 Summary

**One-liner:** Haversine SQL distance query API + Vue 3 geolocation composable + Vant List infinite scroll home page with sport filtering and sort toggle.

## Tasks Completed

| Task | Name | Commit | Files |
|------|------|--------|-------|
| 1 | 后端 ListNearby API（Haversine 距离查询 + 排序 + 筛选 + 分页） | 158997e | backend/service/invitation.go, backend/handler/invitation.go, backend/router/router.go |
| 2 | 前端首页改造（定位 + 邀约卡片 + 筛选标签 + 排序 + 无限滚动） | e2095f0 | frontend/src/composables/useGeolocation.js, frontend/src/utils/format.js, frontend/src/api/invitation.js, frontend/src/components/InvitationCard.vue, frontend/src/views/home/index.vue |

## What Was Built

**Backend:**
- `NearbyInvitationResponse` struct embedding `InvitationResponse` + `Distance float64` field
- `ListNearbyInvitations(userID, lat, lng, sportType, sortBy, page, pageSize)` — excludes current user's own invitations, Haversine SQL expression (SQLite-compatible), sport_type filter, distance/time sort, OFFSET/LIMIT pagination
- `haversineDistance()` Go-layer fallback using `math` package for when SQLite math functions are unavailable
- `ListNearby` handler with `ShouldBindQuery`, PageSize cap at 20
- Route `GET /invitations/nearby` registered before `/:id` to prevent Gin param collision

**Frontend:**
- `useGeolocation` composable: wraps `navigator.geolocation.getCurrentPosition`, error code mapping (1→PERMISSION_DENIED, 2→POSITION_UNAVAILABLE, 3→TIMEOUT), 5-minute cache
- `formatDistance` util: < 1km shows meters (rounded), >= 1km shows km (1 decimal)
- `formatTime` util: formats ISO string to `MM月DD日 HH:mm`
- `InvitationCard` component: `mode` prop switches between distance display (nearby list) and status tag (my invitations list); inactive greying for gathered/terminated/expired
- Home page rewrite: NavBar + horizontal sport filter tabs + sort toggle buttons + Vant List with `v-if="hasLocation"` guard + MapPicker popup fallback for geolocation failure

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 1 - Bug] Duplicate function declarations in handler/invitation.go**
- **Found during:** Task 1 verification (go build)
- **Issue:** Edit tool replaced only the top portion of the file, leaving original GetInvitation/TerminateInvitation/DeleteInvitation/GetParticipants functions duplicated
- **Fix:** Removed the duplicate block at end of file
- **Files modified:** backend/handler/invitation.go
- **Commit:** 158997e (fixed before commit)

## Known Stubs

None — all data flows are wired. The home page fetches real data from the backend API.

## Threat Flags

No new security-relevant surface beyond what was planned. The `GET /invitations/nearby` endpoint is covered by T-03-01, T-03-02, T-03-03 in the plan's threat model.

## Self-Check: PASSED

- [x] backend/service/invitation.go — contains `func ListNearbyInvitations` and `func haversineDistance`
- [x] backend/handler/invitation.go — contains `func ListNearby`
- [x] backend/router/router.go — contains `invitation.GET("/nearby"`
- [x] frontend/src/composables/useGeolocation.js — contains `export function useGeolocation`
- [x] frontend/src/utils/format.js — contains `export function formatDistance`
- [x] frontend/src/components/InvitationCard.vue — contains `InvitationCard`
- [x] frontend/src/views/home/index.vue — contains `van-list`
- [x] Commits 158997e and e2095f0 exist in git log
- [x] Backend `go build ./...` passes
- [x] Frontend `npm run build` passes (built in 176ms)
