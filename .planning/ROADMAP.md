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
| 3 | 附近约球与我的邀约 | 定位发现附近邀约、加入退出、个人邀约管理 | NEARBY-01~07, MY-01~02 | 4 |

---

### Phase 1: 用户系统
**Goal:** 用户可以注册、登录、管理个人信息，前后端基础架构搭建完成
**Mode:** mvp
**Success Criteria:**
1. 用户可以注册账号并登录，刷新后保持登录状态
2. 用户可以查看和编辑个人资料（头像、昵称、常住地、球类兴趣）
3. 用户可以修改密码和退出登录
4. 前端Vue+Vant项目结构和后端Go+SQLite项目结构搭建完成

**Requirements:** AUTH-01, AUTH-02, AUTH-03, AUTH-04, PROF-01, PROF-02, PROF-03, PROF-04

**UI hint:** yes

---

### Phase 2: 邀约核心
**Goal:** 用户可以发起球类邀约，管理邀约状态，查看参与人信息
**Mode:** mvp
**Success Criteria:**
1. 用户可以创建邀约（选择球类、时间、地点、人数）并提交
2. 邀约状态正确流转（等待中→已召集/已终止）
3. 发起人可以查看参与人列表（手机号脱敏显示）
4. 发起人可以终止或删除自己的邀约

**Requirements:** INVITE-01, INVITE-02, INVITE-03, INVITE-04, INVITE-05, INVITE-06, PART-01, PART-02, PART-03

**UI hint:** yes

---

### Phase 3: 附近约球与我的邀约
**Goal:** 用户可以发现附近邀约、加入退出邀约、查看个人邀约历史
**Mode:** mvp
**Success Criteria:**
1. 浏览器定位获取用户位置，按距离排序展示附近邀约
2. 支持按时间排序和按球类筛选
3. 用户可以加入和退出邀约
4. 用户可以查看自己发起和参与的邀约列表

**Requirements:** NEARBY-01, NEARBY-02, NEARBY-03, NEARBY-04, NEARBY-05, NEARBY-06, NEARBY-07, MY-01, MY-02

**UI hint:** yes

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
