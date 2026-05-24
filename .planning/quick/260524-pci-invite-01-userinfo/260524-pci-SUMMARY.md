---
phase: quick
plan: 260524-pci
status: complete
started: 2026-05-24T10:15:00Z
completed: 2026-05-24T10:18:00Z
requirements-completed: [INVITE-01, INVITE-05, INVITE-06, PART-01]
---

# Quick Task 260524-pci: 修复 INVITE-01 字段名不匹配 + userInfo 刷新持久化

## Accomplishments

1. **修复 capacity → max_people 字段名** — 前端发起邀约 payload 中 `capacity` 改为 `max_people`，匹配后端 JSON binding，解除创建邀约 400 错误
2. **userInfo 持久化** — setUserInfo 时同步写入 localStorage，初始化时从 localStorage 恢复，logout 时清除
3. **自动拉取 userInfo** — 添加 fetchUserIfNeeded()，有 token 无 userInfo 时自动调 getProfile API；main.js 启动时调用

## Files Modified

- `frontend/src/views/create/index.vue` — payload key capacity → max_people
- `frontend/src/stores/user.js` — localStorage 持久化 + fetchUserIfNeeded
- `frontend/src/main.js` — 启动时调用 fetchUserIfNeeded

## Verification

- Frontend build: PASS (vite build 成功)
- Backend build: PASS (go build 成功)
