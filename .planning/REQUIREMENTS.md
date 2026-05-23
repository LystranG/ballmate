# Requirements: 掌上约球

**Defined:** 2026-05-23
**Core Value:** 让用户能快速找到附近想打球的人，凑齐人数开始运动

## v1 Requirements

Requirements for initial release. Each maps to roadmap phases.

### 认证与用户管理

- [ ] **AUTH-01**: 用户可以使用账号密码注册，填写手机号、常住地、球类兴趣
- [ ] **AUTH-02**: 用户可以使用账号密码登录系统
- [ ] **AUTH-03**: 用户可以退出登录
- [ ] **AUTH-04**: 用户登录状态在刷新后保持（token持久化）

### 个人信息

- [ ] **PROF-01**: 用户可以查看个人信息页面
- [ ] **PROF-02**: 用户可以修改头像
- [ ] **PROF-03**: 用户可以修改个人资料（昵称、常住地、球类兴趣）
- [ ] **PROF-04**: 用户可以修改密码

### 发起邀约

- [ ] **INVITE-01**: 用户可以发起球类邀约（选择球类、时间、地点、人数要求）
- [ ] **INVITE-02**: 发起邀约时支持地图选点确定地点坐标
- [ ] **INVITE-03**: 提交后邀约处于"等待中"状态
- [ ] **INVITE-04**: 人员凑齐时自动变为"已召集"状态
- [ ] **INVITE-05**: 发起人可以终止邀约
- [ ] **INVITE-06**: 发起人可以删除邀约

### 附近约球

- [ ] **NEARBY-01**: 基于浏览器定位获取用户当前位置
- [ ] **NEARBY-02**: 列出附近其他人发出的球类邀约
- [ ] **NEARBY-03**: 列表默认按地点从近到远排列
- [ ] **NEARBY-04**: 可按时间排序
- [ ] **NEARBY-05**: 可按球类筛选
- [ ] **NEARBY-06**: 用户可以加入一个邀约
- [ ] **NEARBY-07**: 用户可以退出已加入的邀约

### 参与人管理

- [ ] **PART-01**: 发起人可以查看当前邀约的参与人列表
- [ ] **PART-02**: 参与人信息包括姓名和手机号
- [ ] **PART-03**: 手机号中间4位用*号代替，不显示完整手机号

### 我的邀约

- [ ] **MY-01**: 用户可以查看自己发起的邀约列表
- [ ] **MY-02**: 用户可以查看自己参与的邀约列表

## v2 Requirements

### 增强功能

- **ENH-01**: 邀约评价与评分
- **ENH-02**: 球友推荐（基于兴趣匹配）
- **ENH-03**: 消息通知（邀约状态变更提醒）
- **ENH-04**: 历史记录与统计

## Out of Scope

| Feature | Reason |
|---------|--------|
| 微信小程序 | 使用Vue+Vant H5方案 |
| 实时聊天 | 复杂度高，非核心功能 |
| 支付功能 | 约球场景不涉及线上支付 |
| 推送通知 | v1不做 |
| OAuth第三方登录 | 账号密码足够 |

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| AUTH-01 | Phase 1 | Pending |
| AUTH-02 | Phase 1 | Pending |
| AUTH-03 | Phase 1 | Pending |
| AUTH-04 | Phase 1 | Pending |
| PROF-01 | Phase 1 | Pending |
| PROF-02 | Phase 1 | Pending |
| PROF-03 | Phase 1 | Pending |
| PROF-04 | Phase 1 | Pending |
| INVITE-01 | Phase 2 | Pending |
| INVITE-02 | Phase 2 | Pending |
| INVITE-03 | Phase 2 | Pending |
| INVITE-04 | Phase 2 | Pending |
| INVITE-05 | Phase 2 | Pending |
| INVITE-06 | Phase 2 | Pending |
| NEARBY-01 | Phase 3 | Pending |
| NEARBY-02 | Phase 3 | Pending |
| NEARBY-03 | Phase 3 | Pending |
| NEARBY-04 | Phase 3 | Pending |
| NEARBY-05 | Phase 3 | Pending |
| NEARBY-06 | Phase 3 | Pending |
| NEARBY-07 | Phase 3 | Pending |
| PART-01 | Phase 2 | Pending |
| PART-02 | Phase 2 | Pending |
| PART-03 | Phase 2 | Pending |
| MY-01 | Phase 3 | Pending |
| MY-02 | Phase 3 | Pending |

**Coverage:**
- v1 requirements: 26 total
- Mapped to phases: 26
- Unmapped: 0

---
*Requirements defined: 2026-05-23*
*Last updated: 2026-05-23 after initial definition*
