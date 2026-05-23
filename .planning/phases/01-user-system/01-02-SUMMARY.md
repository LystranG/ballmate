---
phase: 01-user-system
plan: 02
subsystem: profile
tags: [gin, gorm, file-upload, vant4, vue3, pinia, area-picker, uploader]

requires:
  - "Go + Gin + GORM + SQLite 后端骨架（handler/service/model 三层）"
  - "POST /api/v1/auth/register 和 POST /api/v1/auth/login 接口"
  - "JWT 签发/解析工具 + Auth 中间件"
  - "统一响应格式 {code, message, data}"
  - "Vue 3 + Vant 4 + Pinia + Vue Router 前端骨架"
  - "Axios 封装含 token 注入和 401 拦截"
provides:
  - "GET /api/v1/user/profile 获取当前用户信息"
  - "PUT /api/v1/user/profile 更新昵称、常住地、球类兴趣"
  - "PUT /api/v1/user/password 验证旧密码后修改密码"
  - "POST /api/v1/user/avatar 头像上传含类型和大小校验"
  - "个人中心页面（列表式菜单布局 + 退出登录确认弹窗）"
  - "编辑资料页面（头像上传 + 省市区选择 + 球类兴趣多选）"
  - "修改密码页面（旧密码验证 + 新密码确认）"
affects: [02-invitation-core, 03-nearby-invitations]

tech-stack:
  added: []
  patterns: [file-upload-validation, area-picker-popup, checkbox-group-sports, confirm-dialog-logout]

key-files:
  created:
    - backend/handler/user.go
    - backend/service/user.go
  modified:
    - backend/router/router.go
    - backend/main.go
    - frontend/src/views/profile/index.vue
    - frontend/src/views/profile/edit.vue
    - frontend/src/views/profile/password.vue

key-decisions:
  - "头像文件名使用 time.Now().UnixNano() 生成，避免冲突和路径遍历"
  - "文件上传前后端双重校验：前端 beforeRead + 后端 filepath.Ext 和 file.Size"
  - "省市区选择结果存储为斜杠分隔文本（如 广东/深圳/南山区）"

patterns-established:
  - "文件上传模式：前端 van-uploader beforeRead 校验 → FormData → 后端 c.FormFile + 扩展名白名单 + 大小限制"
  - "showConfirmDialog 确认弹窗模式：退出登录等破坏性操作需二次确认"
  - "表单初始化模式：onMounted 从 Pinia store 读取已有数据填充表单"

requirements-completed: [AUTH-03, PROF-01, PROF-02, PROF-03, PROF-04]

duration: 3min
completed: 2026-05-23
---

# Phase 1 Plan 02: 个人中心垂直切片 Summary

**用户信息查看/编辑/头像上传/改密/退出登录全链路，含文件上传双重校验和省市区三级联动**

## Performance

- **Duration:** 3 min
- **Started:** 2026-05-23T13:32:50Z
- **Completed:** 2026-05-23T13:35:54Z
- **Tasks:** 2
- **Files modified:** 7

## Accomplishments
- 后端用户信息 API 全部实现：GET/PUT profile、PUT password、POST avatar
- 头像上传含文件类型白名单（jpg/jpeg/png/gif）和 5MB 大小限制双重校验
- 个人中心页面实现列表式菜单布局，退出登录含 showConfirmDialog 确认弹窗
- 编辑资料页面实现头像上传、昵称编辑、省市区三级联动选择、球类兴趣多选
- 修改密码页面实现旧密码验证 + 新密码一致性前端校验
- Phase 1 全部需求（AUTH-01~04, PROF-01~04）覆盖完成

## Task Commits

1. **Task 1: 后端用户信息 API（查看/编辑/改密/头像）** - `2a5fa6b` (feat)
2. **Task 2: 前端个人中心三页面（查看/编辑/改密）** - `387798b` (feat)

## Files Created/Modified
- `backend/handler/user.go` - GetProfile/UpdateProfile/UpdatePassword/UploadAvatar handler
- `backend/service/user.go` - 用户信息业务逻辑层
- `backend/router/router.go` - 替换占位 handler 为真实实现
- `backend/main.go` - 设置 MaxMultipartMemory=8MB
- `frontend/src/views/profile/index.vue` - 个人中心主页（头像+菜单+退出登录）
- `frontend/src/views/profile/edit.vue` - 编辑资料（头像上传+省市区+球类兴趣）
- `frontend/src/views/profile/password.vue` - 修改密码（旧密码验证+新密码确认）

## Decisions Made
- 头像文件名使用 time.Now().UnixNano() 生成唯一名称，防止冲突和路径遍历
- 文件上传前后端双重校验确保安全性
- 省市区选择结果存储为斜杠分隔文本字符串

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered

None.

## User Setup Required

None - no external service configuration required.

## Next Phase Readiness
- Phase 1 用户系统全部完成，AUTH-01~04 和 PROF-01~04 需求全覆盖
- 后端 handler/service 分层模式可直接复用到 Phase 2 邀约核心
- 前端页面模式（表单提交、文件上传、弹窗确认）可复用到后续功能

## Self-Check: PASSED

- All 7 key files exist on disk
- Commit 2a5fa6b (Task 1) found in git log
- Commit 387798b (Task 2) found in git log

---
*Phase: 01-user-system*
*Completed: 2026-05-23*
