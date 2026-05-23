# Phase 1: 用户系统 - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-05-23
**Phase:** 1-用户系统
**Areas discussed:** 注册信息设计, 头像与个人资料, 页面流转与布局, 项目基础架构

---

## 注册信息设计

### 手机号验证

| Option | Description | Selected |
|--------|-------------|----------|
| 仅格式校验 | 只验证格式是否正确（11位数字），不发短信 | ✓ |
| 短信验证码 | 接入短信服务发送验证码 | |
| 模拟验证码 | 有验证码UI但后端不真正发短信 | |

**User's choice:** 仅格式校验

### 球类兴趣

| Option | Description | Selected |
|--------|-------------|----------|
| 多选标签 | 预设常见球类，用户可多选 | ✓ |
| 单选 | 用户只能选择一个主要兴趣球类 | |
| 自由输入 | 用户自己输入感兴趣的球类 | |

**User's choice:** 多选标签

### 常住地输入

| Option | Description | Selected |
|--------|-------------|----------|
| 自由输入框 | 用户直接输入地址文字 | |
| 省市区选择器 | 省-市-区三级联动选择器 | ✓ |
| 地址搜索联想 | 输入时提供地址候选列表 | |

**User's choice:** 省市区选择器

### 必填字段

| Option | Description | Selected |
|--------|-------------|----------|
| 核心必填+其余选填 | 账号、密码、手机号必填，常住地和球类兴趣选填 | ✓ |
| 全部必填 | 所有字段都在注册时必填 | |
| 最简注册 | 只需账号密码即可注册 | |

**User's choice:** 核心必填+其余选填

---

## 头像与个人资料

### 头像方案

| Option | Description | Selected |
|--------|-------------|----------|
| 相册上传 | 用户从手机相册选图上传 | ✓ |
| 预设头像选择 | 提供一组预设头像供选择 | |
| 上传+裁剪 | 支持相册选图+圆形裁剪 | |

**User's choice:** 相册上传

### 头像存储

| Option | Description | Selected |
|--------|-------------|----------|
| 文件存储 | 图片存为文件，数据库只存路径 | ✓ |
| Base64存数据库 | 图片转base64存入SQLite | |

**User's choice:** 文件存储

### 资料编辑方式

| Option | Description | Selected |
|--------|-------------|----------|
| 统一表单页 | 一个编辑页面，所有字段一起修改后保存 | ✓ |
| 逐项编辑 | 每个字段单独编辑 | |

**User's choice:** 统一表单页

---

## 页面流转与布局

### 登录注册页

| Option | Description | Selected |
|--------|-------------|----------|
| 分页 | 登录和注册各自独立页面 | ✓ |
| Tab切换 | 同一页面上方tab切换 | |

**User's choice:** 分页

### 注册后行为

| Option | Description | Selected |
|--------|-------------|----------|
| 自动登录 | 注册成功后直接登录跳转首页 | ✓ |
| 跳转登录页 | 注册成功后提示并跳转登录页 | |

**User's choice:** 自动登录

### 个人中心布局

| Option | Description | Selected |
|--------|-------------|----------|
| 列表式菜单 | 顶部头像+昵称，下方列表菜单 | ✓ |
| 卡片式 | 所有信息在一个卡片内 | |

**User's choice:** 列表式菜单

### 底部导航

| Option | Description | Selected |
|--------|-------------|----------|
| 有底部导航 | Phase 1搭建底部导航栏框架 | ✓ |
| 无底部导航 | 不设底部导航，通过链接跳转 | |

**User's choice:** 有底部导航

---

## 项目基础架构

### API风格

| Option | Description | Selected |
|--------|-------------|----------|
| RESTful + 统一响应 | 统一响应格式{code,message,data}，RESTful路径 | ✓ |
| 简单路径 | 不严格遵循REST规范 | |

**User's choice:** RESTful + 统一响应

### Go框架

| Option | Description | Selected |
|--------|-------------|----------|
| Gin | 轻量级HTTP框架，生态丰富 | ✓ |
| 标准库 | Go标准库net/http | |
| Fiber | 类似Express的轻量框架 | |

**User's choice:** Gin

### Token策略

| Option | Description | Selected |
|--------|-------------|----------|
| JWT + localStorage | JWT存localStorage，请求时放Authorization header | ✓ |
| JWT + Cookie | JWT存httpOnly cookie | |
| Session | 服务端维护session | |

**User's choice:** JWT + localStorage

### 前端请求

| Option | Description | Selected |
|--------|-------------|----------|
| Axios封装 | 封装axios实例，统一处理token注入、响应拦截、错误提示 | ✓ |
| 原生fetch | 使用浏览器原生fetch API | |

**User's choice:** Axios封装

---

## Claude's Discretion

- JWT 过期时间、刷新策略
- 前端目录结构和路由组织
- 数据库表结构设计
- Go 项目分层结构
- Vant 组件具体使用方式

## Deferred Ideas

None — discussion stayed within phase scope
