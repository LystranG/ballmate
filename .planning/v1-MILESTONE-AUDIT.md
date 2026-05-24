---
milestone: v1
audited: 2026-05-24
status: gaps_found
scores:
  requirements: 17/26
  phases: 4/4
  integration: 24/26
  flows: 5/6
gaps:
  requirements:
    - id: "INVITE-01"
      status: "unsatisfied"
      phase: "Phase 2"
      claimed_by_plans: ["02-01-PLAN.md", "02-02-PLAN.md"]
      completed_by_plans: ["02-02-SUMMARY.md"]
      verification_status: "missing"
      evidence: "Frontend sends `capacity` field but backend expects `max_people` — request returns 400, invitation creation completely broken"
    - id: "INVITE-05"
      status: "partial"
      phase: "Phase 2"
      claimed_by_plans: ["02-03-PLAN.md"]
      completed_by_plans: ["02-03-SUMMARY.md"]
      verification_status: "missing"
      evidence: "Code exists but userInfo not auto-fetched after page refresh — creator cannot terminate if they refresh before visiting profile page"
    - id: "INVITE-06"
      status: "partial"
      phase: "Phase 2"
      claimed_by_plans: ["02-03-PLAN.md"]
      completed_by_plans: ["02-03-SUMMARY.md"]
      verification_status: "missing"
      evidence: "Same userInfo refresh issue — delete button not shown after page refresh"
    - id: "PART-01"
      status: "partial"
      phase: "Phase 2"
      claimed_by_plans: ["02-03-PLAN.md"]
      completed_by_plans: ["02-03-SUMMARY.md"]
      verification_status: "missing"
      evidence: "Participant list only loads for creator, but isCreator check fails after page refresh when userInfo is null"
  integration:
    - from: "Phase 2 frontend (create/index.vue)"
      to: "Phase 2 backend (handler/invitation.go)"
      issue: "Field name mismatch: frontend sends `capacity`, backend binds `max_people`"
      affected_requirements: ["INVITE-01"]
    - from: "Phase 1 (stores/user.js)"
      to: "Phase 2 (views/invitation/detail.vue)"
      issue: "userInfo not persisted to localStorage or auto-fetched on app init — null after refresh"
      affected_requirements: ["INVITE-05", "INVITE-06", "PART-01", "PART-02", "PART-03"]
  flows:
    - name: "Create Invitation"
      breaks_at: "Step 2: POST /invitations returns 400 due to field name mismatch"
      affected_requirements: ["INVITE-01", "INVITE-03", "INVITE-04"]
tech_debt:
  - phase: 01-user-system
    items:
      - "No VERIFICATION.md — phase was executed but never formally verified"
      - "REQUIREMENTS.md checkboxes not updated (AUTH-01~04, PROF-01~04 still unchecked)"
  - phase: 02-invite-core
    items:
      - "No VERIFICATION.md — phase was executed but never formally verified"
      - "SUMMARY files missing requirements-completed frontmatter field"
      - "REQUIREMENTS.md checkboxes not updated (INVITE-01~06, PART-01~03 still unchecked)"
  - phase: 03-nearby-and-my-invitations
    items:
      - "VERIFICATION.md is plan-verification (pre-execution), not post-execution verification"
      - "SUMMARY files missing requirements-completed frontmatter field"
  - phase: 04-report
    items:
      - "Report has placeholder content requiring manual user input ([用户填写：...] and [截图：...])"
---

# Milestone v1 Audit Report: 掌上约球

## Executive Summary

| Dimension | Score | Status |
|-----------|-------|--------|
| Requirements | 17/26 | Gaps found |
| Phases executed | 4/4 | All complete |
| Integration wiring | 24/26 | 1 blocker, 1 warning |
| E2E flows | 5/6 | 1 broken flow |

**Overall: GAPS FOUND** — 1 critical blocker prevents invitation creation, 1 warning degrades creator experience after page refresh.

---

## Requirements Cross-Reference (3-Source)

### Phase 1: 用户系统

| REQ-ID | Description | SUMMARY | VERIFICATION | REQUIREMENTS.md | Final Status |
|--------|-------------|---------|--------------|-----------------|--------------|
| AUTH-01 | 账号密码注册 | listed (01-01) | missing | `[ ]` Pending | partial |
| AUTH-02 | 账号密码登录 | listed (01-01) | missing | `[ ]` Pending | partial |
| AUTH-03 | 退出登录 | listed (01-02) | missing | `[ ]` Pending | partial |
| AUTH-04 | token持久化 | listed (01-01) | missing | `[ ]` Pending | partial |
| PROF-01 | 查看个人信息 | listed (01-02) | missing | `[ ]` Pending | partial |
| PROF-02 | 修改头像 | listed (01-02) | missing | `[ ]` Pending | partial |
| PROF-03 | 修改个人资料 | listed (01-02) | missing | `[ ]` Pending | partial |
| PROF-04 | 修改密码 | listed (01-02) | missing | `[ ]` Pending | partial |

**Note:** Integration checker confirms all 8 requirements are WIRED end-to-end. Gap is documentation only.

### Phase 2: 邀约核心

| REQ-ID | Description | SUMMARY | VERIFICATION | REQUIREMENTS.md | Final Status |
|--------|-------------|---------|--------------|-----------------|--------------|
| INVITE-01 | 发起球类邀约 | missing | missing | `[ ]` Pending | **unsatisfied** |
| INVITE-02 | 地图选点 | missing | missing | `[ ]` Pending | partial |
| INVITE-03 | 等待中状态 | missing | missing | `[ ]` Pending | partial |
| INVITE-04 | 已召集状态 | missing | missing | `[ ]` Pending | partial |
| INVITE-05 | 终止邀约 | missing | missing | `[ ]` Pending | **partial** (bug) |
| INVITE-06 | 删除邀约 | missing | missing | `[ ]` Pending | **partial** (bug) |
| PART-01 | 查看参与人列表 | missing | missing | `[ ]` Pending | **partial** (bug) |
| PART-02 | 参与人姓名手机号 | missing | missing | `[ ]` Pending | partial |
| PART-03 | 手机号脱敏 | missing | missing | `[ ]` Pending | partial |

### Phase 3: 附近约球与我的邀约

| REQ-ID | Description | SUMMARY | VERIFICATION | REQUIREMENTS.md | Final Status |
|--------|-------------|---------|--------------|-----------------|--------------|
| NEARBY-01 | 浏览器定位 | missing | passed | `[x]` Complete | satisfied |
| NEARBY-02 | 列出附近邀约 | missing | passed | `[x]` Complete | satisfied |
| NEARBY-03 | 按距离排序 | missing | passed | `[x]` Complete | satisfied |
| NEARBY-04 | 按时间排序 | missing | passed | `[x]` Complete | satisfied |
| NEARBY-05 | 按球类筛选 | missing | passed | `[x]` Complete | satisfied |
| NEARBY-06 | 加入邀约 | missing | passed | `[x]` Complete | satisfied |
| NEARBY-07 | 退出邀约 | missing | passed | `[x]` Complete | satisfied |
| MY-01 | 查看发起的邀约 | missing | passed | `[x]` Complete | satisfied |
| MY-02 | 查看参与的邀约 | missing | passed | `[x]` Complete | satisfied |

---

## Critical Gaps (Blockers)

### 1. INVITE-01: Field Name Mismatch Breaks Invitation Creation

**Severity:** BLOCKER
**Files:**
- `frontend/src/views/create/index.vue` (line ~196) — sends `capacity`
- `backend/handler/invitation.go` (line ~20) — expects `max_people`

**Impact:** No invitation can be created. The backend returns 400 on every create attempt. This cascades: without invitations in the database, the nearby listing (Phase 3) has no data to display.

**Fix:** Change frontend payload key from `capacity` to `max_people`.

### 2. userInfo Not Persisted After Page Refresh

**Severity:** WARNING (degraded UX, not total breakage)
**Files:**
- `frontend/src/stores/user.js` — userInfo initialized as null, not synced to localStorage
- `frontend/src/views/invitation/detail.vue` — isCreator computed uses userInfo?.id

**Impact:** After page refresh, creators cannot see participant list, terminate, or delete their invitations until they visit the profile page (which fetches userInfo).

**Fix:** Either persist userInfo to localStorage (like token) or auto-fetch profile on app init when token exists.

---

## E2E Flow Results

| Flow | Status | Break Point |
|------|--------|-------------|
| User Registration | COMPLETE | — |
| User Login | COMPLETE | — |
| Create Invitation | **BROKEN** | POST returns 400 (capacity vs max_people) |
| Browse Nearby | COMPLETE | — |
| Join/Leave Invitation | COMPLETE | — |
| My Invitations | COMPLETE | — |

---

## Tech Debt Summary

| Phase | Items | Severity |
|-------|-------|----------|
| 01-user-system | 2 | Low (documentation gaps) |
| 02-invite-core | 3 | Low (documentation gaps) |
| 03-nearby-and-my-invitations | 2 | Low (documentation gaps) |
| 04-report | 1 | Info (user action required) |

**Total: 8 items across 4 phases** — all documentation/process gaps, no code quality debt.

---

## Nyquist Compliance

**Status:** SKIPPED (workflow.nyquist_validation = false in config)

---

*Audited: 2026-05-24*
*Milestone: v1 掌上约球*
