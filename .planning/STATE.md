---
gsd_state_version: 1.0
milestone: v1.0
milestone_name: milestone
status: executing
last_updated: "2026-05-24T09:00:00.000Z"
progress:
  total_phases: 4
  completed_phases: 2
  total_plans: 8
  completed_plans: 6
  percent: 62
---

# Project State: 掌上约球

## Current Phase

**Phase:** 3 — 附近约球与我的邀约
**Status:** Executing Phase 03
**Started:** 2026-05-24

## Project Reference

See: .planning/PROJECT.md (updated 2026-05-23)

**Core value:** 让用户能快速找到附近想打球的人，凑齐人数开始运动
**Current focus:** Phase 03 — nearby-and-my-invitations (Plan 01 complete)

## Progress

| Phase | Name | Status | Plans |
|-------|------|--------|-------|
| 1 | 用户系统 | ● Complete | 2/2 |
| 2 | 邀约核心 | ● Complete | 3/3 |
| 3 | 附近约球与我的邀约 | ◐ In Progress | 1/3 |
| 4 | 报告撰写 | ○ Not Started | 0/0 |

## Decisions Log

- JWT 过期时间 7 天，无刷新 token（课程项目简化）
- Sports 字段用逗号分隔字符串存储
- 头像文件名使用 time.Now().UnixNano() 生成唯一名称
- 文件上传前后端双重校验（前端 beforeRead + 后端扩展名白名单+大小限制）
- 创建者自动加入 participations（计入人数），maxPeople=1 时创建即 gathered
- computeStatus 查询时计算 expired，数据库不存储该状态
- 邀约详情路由使用 detail.vue（非 index.vue）
- Haversine SQL 表达式使用手动 pi/180 替代 RADIANS（SQLite 不支持 RADIANS）
- Go 层 haversineDistance 回退：当 SQLite 数学函数不可用时在内存中计算距离排序
- GET /invitations/nearby 在 /:id 之前注册，防止 Gin 将 "nearby" 当作 id 参数
- PageSize 上限 20（T-03-02 DoS 防护）
- v-if="hasLocation" 控制 Vant List 渲染时机，防止定位前发出无坐标请求

## Performance Metrics

| Phase | Plan | Duration | Tasks | Files |
|-------|------|----------|-------|-------|
| 01 | 01 | 6min | 2 | 32 |
| 01 | 02 | 3min | 2 | 7 |
| 02 | 01 | 7min | 3 | 7 |
| 02 | 02 | 2min | 4 | 6 |
| 02 | 03 | 3min | 1 | 3 |
| 03 | 01 | 12min | 2 | 8 |

---
*Last updated: 2026-05-24 after phase 3 plan 01 completion*
