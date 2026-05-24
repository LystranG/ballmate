---
quick_id: 260524-pci
description: "修复 INVITE-01 字段名不匹配 + userInfo 刷新持久化"
mode: quick
tasks: 2
---

# Quick Task: 修复 INVITE-01 字段名不匹配 + userInfo 刷新持久化

来源：v1 Milestone Audit 发现的 1 个 BLOCKER + 1 个 WARNING

## Task 1: 修复 capacity → max_people 字段名不匹配

**Files:** frontend/src/views/create/index.vue
**Action:** 将 payload 中的 `capacity` 键改为 `max_people`，匹配后端 handler 的 JSON binding
**Done:** POST /api/v1/invitations 不再返回 400

## Task 2: userInfo 持久化 + 自动拉取

**Files:** frontend/src/stores/user.js, frontend/src/main.js
**Action:**
1. userInfo 同步到 localStorage（setUserInfo 时写入，logout 时清除）
2. 初始化时从 localStorage 恢复 userInfo
3. 添加 fetchUserIfNeeded()：有 token 无 userInfo 时自动调 getProfile
4. main.js 启动时调用 fetchUserIfNeeded()
**Done:** 页面刷新后 isCreator 判断正常，参与人列表可见
