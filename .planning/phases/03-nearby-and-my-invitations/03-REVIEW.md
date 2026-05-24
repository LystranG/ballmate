---
phase: 03-nearby-and-my-invitations
reviewed: 2026-05-24T00:00:00Z
depth: standard
files_reviewed: 12
files_reviewed_list:
  - backend/service/invitation.go
  - backend/handler/invitation.go
  - backend/router/router.go
  - frontend/src/api/invitation.js
  - frontend/src/composables/useGeolocation.js
  - frontend/src/utils/format.js
  - frontend/src/components/InvitationCard.vue
  - frontend/src/views/home/index.vue
  - frontend/src/views/invitation/detail.vue
  - frontend/src/views/my-invitations/index.vue
  - frontend/src/views/profile/index.vue
  - frontend/src/router/index.js
findings:
  critical: 2
  warning: 6
  info: 3
  total: 11
status: issues_found
---

# Phase 03: Code Review Report

**Reviewed:** 2026-05-24
**Depth:** standard
**Files Reviewed:** 12
**Status:** issues_found

## Summary

Reviewed the full phase-03 implementation covering nearby invitations, join/leave flows, my-invitations page, and supporting utilities. The routing and API layer are clean. The main concerns are: a dead-code logic contradiction in `LeaveInvitation` that makes the gathered→waiting status recovery unreachable; a TOCTOU race condition in `JoinInvitation` with no transaction; pervasive unchecked DB errors in the service layer; missing coordinate bounds validation; and a duplicated `formatTime` implementation in `detail.vue`.

---

## Critical Issues

### CR-01: TOCTOU Race Condition in JoinInvitation — Overbooking Possible

**File:** `backend/service/invitation.go:327-334`
**Issue:** The capacity check (count < maxPeople) and the participation insert are two separate DB operations with no transaction wrapping them. Under concurrent requests, two users can both pass the count check simultaneously and both insert, exceeding `max_people`. The unique index only prevents the same user joining twice; it does not prevent two different users from both slipping through the capacity gate at the same time.

**Fix:**
```go
func JoinInvitation(invitationID, userID uint) error {
    return DB.Transaction(func(tx *gorm.DB) error {
        var inv model.Invitation
        if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&inv, invitationID).Error; err != nil {
            return errors.New("邀约不存在")
        }
        if computeStatus(inv) != "waiting" {
            return errors.New("该邀约已无法加入")
        }
        var count int64
        tx.Model(&model.Participation{}).Where("invitation_id = ? AND deleted_at IS NULL", invitationID).Count(&count)
        if int(count) >= inv.MaxPeople {
            return errors.New("邀约人数已满")
        }
        if err := tx.Create(&model.Participation{InvitationID: invitationID, UserID: userID}).Error; err != nil {
            return errors.New("您已加入该邀约")
        }
        var newCount int64
        tx.Model(&model.Participation{}).Where("invitation_id = ? AND deleted_at IS NULL", invitationID).Count(&newCount)
        if int(newCount) >= inv.MaxPeople {
            tx.Model(&inv).Update("status", "gathered")
        }
        return nil
    })
}
```

---

### CR-02: Dead Code Makes gathered→waiting Recovery Unreachable in LeaveInvitation

**File:** `backend/service/invitation.go:355-365`
**Issue:** The function comment (line 345-346) states that leaving a gathered invitation should restore it to `waiting`. However, the guard at line 355 calls `computeStatus(inv)`, which returns `"gathered"` for a gathered invitation — not `"waiting"`. This causes the function to return an error before ever reaching the recovery code at line 364. The `if inv.Status == "gathered"` block is dead code and the stated behavior is never executed.

The frontend `canLeave` comment in `detail.vue:146` says "gathered 状态已满员，不允许退出", which contradicts the backend comment. One of the two is wrong; the dead code proves the backend implementation does not match its own stated intent.

**Fix — Option A (block leaving gathered, remove dead code):**
```go
// LeaveInvitation 退出邀约
// 校验：邀约存在、非创建者、状态为 waiting、已加入
// gathered 状态不允许退出
func LeaveInvitation(invitationID, userID uint) error {
    // ... existing checks ...
    if computeStatus(inv) != "waiting" {
        return errors.New("该邀约状态不允许退出")
    }
    result := DB.Where("invitation_id = ? AND user_id = ?", invitationID, userID).Delete(&model.Participation{})
    if result.RowsAffected == 0 {
        return errors.New("您未加入该邀约")
    }
    // Remove the dead gathered->waiting block entirely
    return nil
}
```

**Fix — Option B (allow leaving gathered, fix the guard):**
```go
if inv.Status != "waiting" && inv.Status != "gathered" {
    return errors.New("该邀约状态不允许退出")
}
// ... soft delete ...
if inv.Status == "gathered" {
    DB.Model(&inv).Update("status", "waiting")
}
```

Pick one intent and remove the contradiction.

---

## Warnings

### WR-01: Pervasive Unchecked DB Errors in Service Layer

**File:** `backend/service/invitation.go` — multiple lines
**Issue:** Numerous DB operations silently swallow errors. If any of these fail (disk full, connection lost, constraint violation), the function continues with stale or zero-value data and returns a success response to the caller.

Key unchecked calls:
- Line 63: `countQuery.Count(&total)` — silent failure returns `total=0`, causing frontend to think there are no more pages
- Line 285: `DB.Create(&model.Participation{...})` — creator auto-join failure is silently ignored
- Line 291: `DB.Model(&inv).Update("status", "gathered")` — status update failure is silently ignored
- Line 339: same pattern in `JoinInvitation`
- Line 365: `DB.Model(&inv).Update("status", "waiting")` in `LeaveInvitation`
- Line 382: `DB.Model(&inv).Update("status", "terminated")` in `TerminateInvitation`
- Lines 399-400: `DeleteInvitation` — both deletes unchecked
- Line 419: `GetParticipants` scan unchecked

**Fix:** Check `.Error` on every terminal DB call. For state-mutating operations, return the error to the caller:
```go
if err := DB.Model(&inv).Update("status", "gathered").Error; err != nil {
    return fmt.Errorf("状态更新失败: %w", err)
}
```

---

### WR-02: No Coordinate Bounds Validation — Haversine Produces NaN for Invalid Input

**File:** `backend/handler/invitation.go:49-50` and `backend/service/invitation.go:15-26`
**Issue:** Latitude and longitude are only validated as `required` (non-zero). No range check enforces lat ∈ [-90, 90] and lng ∈ [-180, 180]. Values outside these ranges cause `math.Asin` to receive an argument > 1 or < -1, returning `NaN`. A `NaN` distance propagates silently through sorting and response serialization.

**Fix:**
```go
// In handler, after ShouldBindQuery:
if req.Latitude < -90 || req.Latitude > 90 {
    Error(c, http.StatusBadRequest, "纬度范围无效（-90 到 90）")
    return
}
if req.Longitude < -180 || req.Longitude > 180 {
    Error(c, http.StatusBadRequest, "经度范围无效（-180 到 180）")
    return
}
```
Apply the same check in `CreateInvitation` handler.

---

### WR-03: No Upper Bound on MaxPeople — Arbitrary Large Values Accepted

**File:** `backend/handler/invitation.go:20`
**Issue:** `MaxPeople` is validated with `min=1` but has no `max=` constraint. A client can submit `max_people: 999999`, which is stored and displayed without error. The capacity check in `JoinInvitation` will never trigger, and the progress bar in `InvitationCard.vue` will always show near-zero fill.

**Fix:**
```go
MaxPeople int `json:"max_people" binding:"required,min=1,max=100"`
```

---

### WR-04: No Past-Time Validation for ActivityTime in CreateInvitation

**File:** `backend/handler/invitation.go:33-37`
**Issue:** After parsing `activityTime`, there is no check that it is in the future. A user can create an invitation with an activity time in the past. `computeStatus` will immediately return `"expired"` for such an invitation, making it permanently un-joinable the moment it is created.

**Fix:**
```go
activityTime, err := time.Parse(time.RFC3339, req.ActivityTime)
if err != nil {
    Error(c, http.StatusBadRequest, "时间格式错误，请使用 RFC3339 格式")
    return
}
if activityTime.Before(time.Now()) {
    Error(c, http.StatusBadRequest, "活动时间不能早于当前时间")
    return
}
```

---

### WR-05: formatDistance Does Not Guard Against null/undefined/NaN Input

**File:** `frontend/src/utils/format.js:7-12`
**Issue:** `formatDistance(distanceKm)` calls `distanceKm.toFixed(1)` without checking whether the input is a valid number. If `distanceKm` is `null`, `undefined`, or `NaN` (which can happen if the backend returns a null distance or the SQL haversine produces NaN), the call throws a TypeError or renders `"NaNkm"`.

**Fix:**
```js
export function formatDistance(distanceKm) {
  if (distanceKm == null || isNaN(distanceKm) || distanceKm < 0) return '--'
  if (distanceKm < 1) {
    return `${Math.round(distanceKm * 1000)}m`
  }
  return `${distanceKm.toFixed(1)}km`
}
```

---

### WR-06: my-invitations List Is Stale After Returning from Detail Page

**File:** `frontend/src/views/my-invitations/index.vue:67-79`
**Issue:** The `createdState` and `joinedState` lists are initialized once and never refreshed when the user navigates back from the detail page (e.g., after joining or leaving an invitation). The status shown in the list will be out of date until the user manually refreshes. There is no `onActivated` hook or route watch to trigger a reload.

**Fix:** Add an `onActivated` hook (requires `<keep-alive>` in the router view) or a `watch` on the route to reset and reload both lists when the page becomes active:
```js
import { onActivated } from 'vue'

onActivated(() => {
  // Reset and reload the active tab
  if (activeTab.value === 0) {
    Object.assign(createdState, { list: [], page: 1, finished: false })
    loadCreated()
  } else {
    Object.assign(joinedState, { list: [], page: 1, finished: false })
    loadJoined()
  }
})
```

---

## Info

### IN-01: formatTime Duplicated in detail.vue Instead of Importing from utils

**File:** `frontend/src/views/invitation/detail.vue:153-161`
**Issue:** `formatTime` is defined locally in `detail.vue` (lines 153-161) with identical logic to `frontend/src/utils/format.js:19-27`. The utility file exists precisely to avoid this duplication. `InvitationCard.vue` already imports from `@/utils/format`.

**Fix:**
```js
// Remove the local formatTime definition and add to imports:
import { formatTime } from '@/utils/format'
```

---

### IN-02: statusConfig Duplicated Between detail.vue and InvitationCard.vue

**File:** `frontend/src/views/invitation/detail.vue:127-132` and `frontend/src/components/InvitationCard.vue:71-76`
**Issue:** The `statusConfig` object (colors and labels for waiting/gathered/terminated/expired) is defined identically in both files. Any future status addition requires updating two places.

**Fix:** Extract to a shared constant, e.g. `frontend/src/utils/statusConfig.js`, and import in both files.

---

### IN-03: No Minimum Page Validation — page=0 Produces Negative Offset

**File:** `backend/handler/invitation.go:51` and `backend/handler/invitation.go:150`
**Issue:** `page` defaults to 1 but has no `min=1` binding constraint. A client sending `page=0` produces `offset = (0-1) * pageSize = -10`. GORM passes this negative offset to SQLite, which treats a negative OFFSET as 0 — so it silently returns the first page instead of an error. This is a latent correctness issue.

**Fix:**
```go
Page int `form:"page,default=1" binding:"min=1"`
```

---

_Reviewed: 2026-05-24_
_Reviewer: Claude (gsd-code-reviewer)_
_Depth: standard_
