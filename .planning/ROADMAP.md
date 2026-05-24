# Roadmap: 掌上约球

**Created:** 2026-05-23
**Phases:** 3
**Mode:** MVP (vertical slices)

---

## Phase Overview

| # | Phase | Goal | Requirements | Success Criteria |
|---|-------|------|--------------|------------------|
| 1 | 用户系统 | 完整的注册登录和个人信息管理 | AUTH-01~04, PROF-01~04 | 4 |
| 2 | 邀约核心 | 发起邀约、管理邀约、查看参与人 | INVITE-01~06, PART-01~03 | 4 |
| 3 | 附近约球与我的邀约 | 1/3 | In Progress|  |

---

### Phase 1: 用户系统

**Goal:** 用户可以注册、登录、管理个人信息，前后端基础架构搭建完成
**Mode:** mvp
**Plans:** 2 plans
**Success Criteria:**

1. 用户可以注册账号并登录，刷新后保持登录状态
2. 用户可以查看和编辑个人资料（头像、昵称、常住地、球类兴趣）
3. 用户可以修改密码和退出登录
4. 前端Vue+Vant项目结构和后端Go+SQLite项目结构搭建完成

**Requirements:** AUTH-01, AUTH-02, AUTH-03, AUTH-04, PROF-01, PROF-02, PROF-03, PROF-04

**UI hint:** yes

Plans:

- [x] 01-01-PLAN.md — 基础架构 + 注册/登录垂直切片（前后端骨架、注册、登录、token 持久化） ✓ 2026-05-23
- [x] 01-02-PLAN.md — 个人中心垂直切片（查看/编辑资料、头像上传、修改密码、退出登录） ✓ 2026-05-23

---

### Phase 2: 邀约核心

**Goal:** 用户可以发起球类邀约，管理邀约状态，查看参与人信息
**Mode:** mvp
**Plans:** 3 plans
**Success Criteria:**

1. 用户可以创建邀约（选择球类、时间、地点、人数）并提交
2. 邀约状态正确流转（等待中→已召集/已终止）
3. 发起人可以查看参与人列表（手机号脱敏显示）
4. 发起人可以终止或删除自己的邀约

**Requirements:** INVITE-01, INVITE-02, INVITE-03, INVITE-04, INVITE-05, INVITE-06, PART-01, PART-02, PART-03

**UI hint:** yes

Plans:
**Wave 1**

- [x] 02-01-PLAN.md — 后端基础（数据模型 + service + handler + 路由）+ AMap 包安装确认 ✓ 2026-05-24

**Wave 2** *(blocked on Wave 1 completion)*

- [x] 02-02-PLAN.md — 发起邀约垂直切片（MapPicker 组件 + 发起邀约表单页） ✓ 2026-05-24
- [x] 02-03-PLAN.md — 邀约详情垂直切片（详情展示 + 状态操作 + 参与人列表） ✓ 2026-05-24

---

### Phase 3: 附近约球与我的邀约

**Goal:** 用户可以发现附近邀约、加入退出邀约、查看个人邀约历史
**Mode:** mvp
**Plans:** 1/3 plans executed
**Success Criteria:**

1. 浏览器定位获取用户位置，按距离排序展示附近邀约
2. 支持按时间排序和按球类筛选
3. 用户可以加入和退出邀约
4. 用户可以查看自己发起和参与的邀约列表

**Requirements:** NEARBY-01, NEARBY-02, NEARBY-03, NEARBY-04, NEARBY-05, NEARBY-06, NEARBY-07, MY-01, MY-02

**UI hint:** yes

Plans:
**Wave 1**

- [x] 03-01-PLAN.md — 附近邀约列表垂直切片（后端 Haversine 距离查询 + 前端定位/筛选/排序/无限滚动）
- [ ] 03-02-PLAN.md — 加入/退出邀约垂直切片（后端 Join/Leave API + 前端详情页底部操作按钮）

**Wave 2** *(depends on 03-01 for InvitationCard component)*

- [ ] 03-03-PLAN.md — 我的邀约垂直切片（后端 My API + 前端 Tab 切换页面 + 个人中心入口）

---

## Phase 4: 报告撰写

**Goal:** 完成期末考查报告
**Mode:** mvp
**Success Criteria:**

1. 报告包含项目背景、需求分析、系统设计、核心代码说明
2. 报告输出为txt格式，含截图占位符和mermaid图
3. 报告结构符合模板要求

**Requirements:** (report deliverable, not functional requirement)

**UI hint:** no

---

*Created: 2026-05-23*
*Updated: 2026-05-24 — Phase 3 planned (3 plans, 2 waves)*
