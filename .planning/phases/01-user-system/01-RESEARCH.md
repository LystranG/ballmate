# Phase 1: 用户系统 - Research

**Researched:** 2026-05-23
**Domain:** 用户认证、个人信息管理、前后端基础架构（Vue 3 + Vant 4 / Go + Gin + SQLite）
**Confidence:** HIGH

## Summary

本 phase 是 greenfield 项目的基础搭建，涵盖前端 Vue 3 + Vant 4 移动端应用和后端 Go + Gin + SQLite RESTful API。核心功能包括用户注册/登录（JWT 认证）、个人资料管理（含头像上传）、密码修改。

技术栈成熟稳定，所有库均为各自生态的主流选择，版本活跃维护中。前端使用 Vite 构建、Pinia 状态管理、Vue Router 路由；后端使用 Gin 框架、GORM ORM、golang-jwt 做 token 签发验证。

**Primary recommendation:** 采用标准的前后端分离架构，前端 Vite + Vue 3 + Vant 4 + Pinia + Axios，后端 Gin + GORM + SQLite + JWT，按 handler/service/model 三层组织代码。

<user_constraints>
## User Constraints (from CONTEXT.md)

### Locked Decisions
- **D-01:** 手机号仅做格式校验（11位数字），不发送短信验证码
- **D-02:** 球类兴趣使用多选标签方式，预设选项：篮球、足球、羽毛球、乒乓球、网球、排球、棒球、橄榄球等
- **D-03:** 常住地使用省市区三级联动选择器（Vant Area 组件）
- **D-04:** 必填字段：账号、密码、手机号；选填字段：常住地、球类兴趣（注册后可在个人中心补充）
- **D-05:** 头像通过相册选图上传，后端以文件形式存储，数据库存路径
- **D-06:** 个人资料编辑使用统一表单页，所有字段一起修改后保存
- **D-07:** 登录和注册为独立页面，通过链接互相跳转
- **D-08:** 注册成功后自动登录并跳转首页
- **D-09:** 个人中心采用列表式菜单布局（顶部头像+昵称，下方菜单项）
- **D-10:** Phase 1 搭建底部导航栏框架（首页/发起约球/我的），后续 phase 填充内容
- **D-11:** 后端 API 采用 RESTful 风格，统一响应格式 {code, message, data}，路径如 /api/v1/auth/login
- **D-12:** Go 后端使用 Gin 框架
- **D-13:** 登录状态使用 JWT，存储在 localStorage，请求时放入 Authorization header
- **D-14:** 前端使用 Axios 封装请求层，统一处理 token 注入、响应拦截、错误提示

### Claude's Discretion
- JWT 过期时间、刷新策略
- 前端目录结构和路由组织
- 数据库表结构设计
- Go 项目分层结构（handler/service/model）
- Vant 组件具体使用方式

### Deferred Ideas (OUT OF SCOPE)
None
</user_constraints>

<phase_requirements>
## Phase Requirements

| ID | Description | Research Support |
|----|-------------|------------------|
| AUTH-01 | 用户可以使用账号密码注册，填写手机号、常住地、球类兴趣 | Vant Form + Field 组件做表单，Area 组件做地区选择，后端 bcrypt 加密密码存储 |
| AUTH-02 | 用户可以使用账号密码登录系统 | Gin handler 验证密码 + 签发 JWT token |
| AUTH-03 | 用户可以退出登录 | 前端清除 localStorage token + Pinia 状态重置 |
| AUTH-04 | 用户登录状态在刷新后保持（token持久化） | localStorage 存储 JWT，Axios 拦截器自动注入 Authorization header |
| PROF-01 | 用户可以查看个人信息页面 | GET /api/v1/user/profile 接口 + Vant Cell 组件展示 |
| PROF-02 | 用户可以修改头像 | Vant Uploader 组件 + Gin multipart file upload + 文件存储 |
| PROF-03 | 用户可以修改个人资料（昵称、常住地、球类兴趣） | PUT /api/v1/user/profile 接口 + Vant Form 统一编辑页 |
| PROF-04 | 用户可以修改密码 | PUT /api/v1/user/password 接口，验证旧密码后更新 |
</phase_requirements>

## Architectural Responsibility Map

| Capability | Primary Tier | Secondary Tier | Rationale |
|------------|-------------|----------------|-----------|
| 用户注册/登录表单 | Browser/Client | — | 表单渲染、输入校验、交互反馈在客户端完成 |
| JWT 签发与验证 | API/Backend | — | 安全敏感操作必须在服务端执行 |
| 密码加密存储 | API/Backend | — | bcrypt 哈希必须在服务端，客户端不处理 |
| Token 持久化 | Browser/Client | — | localStorage 是浏览器存储 |
| 头像文件上传 | API/Backend | Browser/Client | 客户端选图预览，服务端接收存储 |
| 头像文件服务 | API/Backend | — | Gin 静态文件服务提供图片访问 |
| 数据持久化 | Database/Storage | — | SQLite 存储用户数据 |
| 路由守卫 | Browser/Client | — | 前端 vue-router 导航守卫控制页面访问 |
| 请求拦截/错误处理 | Browser/Client | — | Axios 拦截器统一处理 |

## Standard Stack

### Core - Frontend

| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| vue | 3.5.34 | 前端框架 | 课程要求，生态成熟 |
| vant | 4.9.24 | 移动端 UI 组件库 | 课程要求，专为移动 Web 设计 |
| vue-router | 5.0.7 | 前端路由 | Vue 官方路由方案 |
| pinia | 3.0.4 | 状态管理 | Vue 官方推荐，替代 Vuex |
| axios | 1.16.1 | HTTP 请求 | 用户决策 D-14 指定 |
| @vant/area-data | 2.1.0 | 省市区数据 | Vant Area 组件配套数据包 |
| vite | 8.0.14 | 构建工具 | Vue 生态标准构建工具 |
| @vitejs/plugin-vue | 6.0.7 | Vite Vue 插件 | Vite + Vue 必需 |

### Core - Backend

| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| github.com/gin-gonic/gin | v1.12.0 | Web 框架 | 用户决策 D-12 指定 |
| gorm.io/gorm | v1.31.1 | ORM | Go 生态最流行的 ORM |
| gorm.io/driver/sqlite | v1.6.0 | SQLite 驱动 | GORM 官方 SQLite 驱动 |
| github.com/golang-jwt/jwt/v5 | v5.3.1 | JWT 处理 | Go 生态标准 JWT 库 |
| golang.org/x/crypto | v0.52.0 | bcrypt 密码加密 | Go 官方扩展库 |
| github.com/gin-contrib/cors | v1.7.7 | CORS 中间件 | 前后端分离必需 |

### Supporting

| Library | Version | Purpose | When to Use |
|---------|---------|---------|-------------|
| unplugin-vue-components | 32.1.0 | Vant 组件自动导入 | 避免手动 import 每个 Vant 组件 |
| @vant/auto-import-resolver | 1.3.0 | Vant 自动导入解析器 | 配合 unplugin-vue-components |

### Alternatives Considered

| Instead of | Could Use | Tradeoff |
|------------|-----------|----------|
| GORM | database/sql + 手写 SQL | 更轻量但开发效率低，对课程项目不划算 |
| Pinia | Vuex 4 | Vuex 已不推荐用于新项目 |
| Axios | fetch API | fetch 原生但缺少拦截器等便利功能 |

**Installation:**

Frontend:
```bash
npm create vite@latest frontend -- --template vue
cd frontend
npm install vue-router pinia axios vant @vant/area-data
npm install -D unplugin-vue-components @vant/auto-import-resolver
```

Backend:
```bash
mkdir backend && cd backend
go mod init ballmate
go get github.com/gin-gonic/gin@v1.12.0
go get gorm.io/gorm@v1.31.1
go get gorm.io/driver/sqlite@v1.6.0
go get github.com/golang-jwt/jwt/v5@v5.3.1
go get golang.org/x/crypto@v0.52.0
go get github.com/gin-contrib/cors@v1.7.7
```

## Package Legitimacy Audit

> slopcheck 不可用，所有包标记为 [ASSUMED]。Planner 应在安装前进行人工确认。

| Package | Registry | Age | Downloads | Source Repo | slopcheck | Disposition |
|---------|----------|-----|-----------|-------------|-----------|-------------|
| vue | npm | 10+ yrs | 5M+/wk | github.com/vuejs/core | N/A | [ASSUMED] Approved |
| vant | npm | 7+ yrs | 100K+/wk | github.com/youzan/vant | N/A | [ASSUMED] Approved |
| vue-router | npm | 10+ yrs | 3M+/wk | github.com/vuejs/router | N/A | [ASSUMED] Approved |
| pinia | npm | 4+ yrs | 2M+/wk | github.com/vuejs/pinia | N/A | [ASSUMED] Approved |
| axios | npm | 10+ yrs | 40M+/wk | github.com/axios/axios | N/A | [ASSUMED] Approved |
| @vant/area-data | npm | 3+ yrs | 30K+/wk | github.com/youzan/vant | N/A | [ASSUMED] Approved |
| gin-gonic/gin | Go module | 9+ yrs | — | github.com/gin-gonic/gin | N/A | [ASSUMED] Approved |
| gorm.io/gorm | Go module | 7+ yrs | — | github.com/go-gorm/gorm | N/A | [ASSUMED] Approved |
| golang-jwt/jwt | Go module | 5+ yrs | — | github.com/golang-jwt/jwt | N/A | [ASSUMED] Approved |
| gin-contrib/cors | Go module | 7+ yrs | — | github.com/gin-contrib/cors | N/A | [ASSUMED] Approved |

**Packages removed due to slopcheck [SLOP] verdict:** none
**Packages flagged as suspicious [SUS]:** none

*slopcheck 不可用，所有包标记为 [ASSUMED]。这些均为各自生态的主流知名库，风险极低。*

## Architecture Patterns

### System Architecture Diagram

```
Browser (Mobile Web)
  |
  |-- Login Page / Register Page / Profile Pages
  |         |
  |         v
  |   Axios Request Layer (token inject / error intercept)
  |         |
  |   Pinia Store (user state, token <-> localStorage)
  |
  | HTTP (JSON / multipart)
  v
Go Backend (Gin)
  |
  |-- CORS Middleware --> JWT Middleware (protected routes)
  |         |
  |         v
  |   Router (/api/v1/auth/*, /api/v1/user/*)
  |         |
  |         v
  |   Handlers (parse request, call service, return response)
  |         |
  |         v
  |   Services (business logic: password hash, token gen)
  |         |
  |         v
  |   Models (GORM entities, DB operations)
  |         |
  |         v
  |   SQLite Database (ballmate.db)
  |
  |-- Static File Server (/uploads/ -> avatar files)
```

### Recommended Project Structure

```
frontend/
├── index.html
├── vite.config.js
├── package.json
├── src/
│   ├── main.js              # 应用入口，注册插件
│   ├── App.vue              # 根组件
│   ├── router/
│   │   └── index.js         # 路由配置 + 导航守卫
│   ├── stores/
│   │   └── user.js          # 用户状态（token、用户信息）
│   ├── api/
│   │   ├── request.js       # Axios 实例封装
│   │   └── user.js          # 用户相关 API 调用
│   ├── views/
│   │   ├── login/
│   │   │   └── index.vue    # 登录页
│   │   ├── register/
│   │   │   └── index.vue    # 注册页
│   │   ├── profile/
│   │   │   ├── index.vue    # 个人中心主页
│   │   │   ├── edit.vue     # 编辑资料页
│   │   │   └── password.vue # 修改密码页
│   │   └── home/
│   │       └── index.vue    # 首页（占位）
│   └── components/
│       └── TabBar.vue       # 底部导航栏

backend/
├── go.mod
├── go.sum
├── main.go                  # 入口，启动服务
├── config/
│   └── config.go            # 配置（JWT密钥、数据库路径等）
├── middleware/
│   ├── cors.go              # CORS 中间件
│   └── auth.go              # JWT 认证中间件
├── router/
│   └── router.go            # 路由注册
├── handler/
│   ├── auth.go              # 注册/登录 handler
│   └── user.go              # 用户信息 handler
├── service/
│   ├── auth.go              # 认证业务逻辑
│   └── user.go              # 用户业务逻辑
├── model/
│   └── user.go              # User GORM 模型
├── utils/
│   └── jwt.go               # JWT 工具函数
├── uploads/                 # 头像文件存储目录
└── data/
    └── ballmate.db          # SQLite 数据库文件
```


### Pattern 1: Axios 请求封装 + Token 注入
**What:** 创建 Axios 实例，通过请求拦截器自动注入 JWT，响应拦截器统一处理错误
**When to use:** 所有 API 请求
**Example:**
```javascript
// Source: Axios official docs + 社区标准模式 [ASSUMED]
import axios from 'axios'
import { showToast } from 'vant'
import router from '@/router'

const request = axios.create({
  baseURL: '/api/v1',
  timeout: 10000
})

// 请求拦截器：注入 token
request.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// 响应拦截器：统一错误处理
request.interceptors.response.use(
  response => response.data,
  error => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      router.push('/login')
      showToast('登录已过期，请重新登录')
    } else {
      showToast(error.response?.data?.message || '网络错误')
    }
    return Promise.reject(error)
  }
)

export default request
```

### Pattern 2: Pinia User Store + Token 持久化
**What:** 用 Pinia 管理用户状态，token 同步到 localStorage
**When to use:** 用户登录状态管理
**Example:**
```javascript
// Source: Pinia official docs pattern [ASSUMED]
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref(null)

  const isLoggedIn = computed(() => !!token.value)

  function setToken(newToken) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function setUserInfo(info) {
    userInfo.value = info
  }

  function logout() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
  }

  return { token, userInfo, isLoggedIn, setToken, setUserInfo, logout }
})
```

### Pattern 3: Go Gin 统一响应格式
**What:** 封装统一的 JSON 响应结构 {code, message, data}
**When to use:** 所有 API 响应
**Example:**
```go
// Source: Gin 社区标准模式 [ASSUMED]
package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Code:    0,
        Message: "success",
        Data:    data,
    })
}

func Error(c *gin.Context, httpCode int, message string) {
    c.JSON(httpCode, Response{
        Code:    -1,
        Message: message,
        Data:    nil,
    })
}
```

### Pattern 4: JWT 中间件
**What:** Gin 中间件验证 Authorization header 中的 Bearer token
**When to use:** 需要认证的路由组
**Example:**
```go
// Source: golang-jwt/jwt 官方用法 + Gin middleware 模式 [ASSUMED]
package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "ballmate/config"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "未登录"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(config.JWTSecret), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "token无效"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "token解析失败"})
            c.Abort()
            return
        }

        c.Set("userID", uint(claims["user_id"].(float64)))
        c.Next()
    }
}
```

### Pattern 5: GORM User 模型 + bcrypt 密码
**What:** 定义 User 模型，使用 bcrypt 加密密码
**When to use:** 用户数据持久化
**Example:**
```go
// Source: GORM official docs + golang.org/x/crypto/bcrypt [ASSUMED]
package model

import (
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex;size:50;not null" json:"username"`
    Password string `gorm:"size:100;not null" json:"-"`
    Phone    string `gorm:"size:11;not null" json:"phone"`
    Nickname string `gorm:"size:50" json:"nickname"`
    Avatar   string `gorm:"size:255" json:"avatar"`
    Location string `gorm:"size:100" json:"location"`
    Sports   string `gorm:"size:255" json:"sports"`
}

func (u *User) SetPassword(password string) error {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hash)
    return nil
}

func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}
```

### Pattern 6: Gin 文件上传（头像）
**What:** 接收 multipart 文件，生成唯一文件名，保存到 uploads/ 目录
**When to use:** 头像上传接口
**Example:**
```go
// Source: gin-gonic.com/en/docs/routing/upload-file/single-file/ [CITED]
func UploadAvatar(c *gin.Context) {
    file, err := c.FormFile("avatar")
    if err != nil {
        Error(c, http.StatusBadRequest, "请选择文件")
        return
    }

    ext := filepath.Ext(file.Filename)
    filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
    dst := filepath.Join("uploads", filename)

    if err := c.SaveUploadedFile(file, dst); err != nil {
        Error(c, http.StatusInternalServerError, "文件保存失败")
        return
    }

    userID := c.GetUint("userID")
    // 更新数据库中的头像路径...
    Success(c, gin.H{"avatar": "/uploads/" + filename})
}
```

### Pattern 7: Vue Router 导航守卫
**What:** 未登录用户访问受保护页面时重定向到登录页
**When to use:** 路由配置
**Example:**
```javascript
// Source: Vue Router official docs pattern [ASSUMED]
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/login', component: () => import('@/views/login/index.vue'), meta: { guest: true } },
    { path: '/register', component: () => import('@/views/register/index.vue'), meta: { guest: true } },
    { path: '/', component: () => import('@/views/home/index.vue'), meta: { auth: true } },
    { path: '/profile', component: () => import('@/views/profile/index.vue'), meta: { auth: true } },
    { path: '/profile/edit', component: () => import('@/views/profile/edit.vue'), meta: { auth: true } },
    { path: '/profile/password', component: () => import('@/views/profile/password.vue'), meta: { auth: true } },
  ]
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.auth && !token) {
    next('/login')
  } else if (to.meta.guest && token) {
    next('/')
  } else {
    next()
  }
})

export default router
```

### Anti-Patterns to Avoid
- **前端存储明文密码:** 密码只在提交时传输，不存储在前端任何位置
- **JWT 存储敏感信息:** token payload 只放 user_id 和过期时间，不放密码等敏感数据
- **不校验文件类型:** 头像上传必须校验文件扩展名和 MIME 类型，只允许图片格式
- **CORS 全开:** 生产环境应限制 AllowOrigins，开发时可用通配符
- **同步密码比较:** 必须使用 bcrypt.CompareHashAndPassword，不能用 == 比较
- **硬编码 JWT 密钥:** 密钥应放在配置文件中，不要写死在代码里

## Don't Hand-Roll

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| 密码加密 | 自定义哈希算法 | golang.org/x/crypto/bcrypt | bcrypt 自带盐值和时间成本，防彩虹表 |
| JWT 签发/验证 | 手动拼接 token | golang-jwt/jwt/v5 | 标准实现，处理签名验证、过期检查 |
| 表单校验 | 手写正则逻辑 | Vant Form rules | 内置校验规则，支持异步校验 |
| 省市区数据 | 自己维护地区编码 | @vant/area-data | 官方维护，数据准确且定期更新 |
| CORS 处理 | 手写 header 设置 | gin-contrib/cors | 处理 preflight、credentials 等边界情况 |
| HTTP 请求封装 | 原生 fetch 封装 | Axios | 拦截器、取消请求、自动 JSON 转换 |
| 文件上传处理 | 手动解析 multipart | Gin c.FormFile | 框架内置，处理内存限制和临时文件 |

**Key insight:** 这些都是看似简单但边界情况多的问题。用成熟库可以避免安全漏洞和兼容性问题。

## Common Pitfalls

### Pitfall 1: SQLite 并发写入锁
**What goes wrong:** 多个请求同时写入 SQLite 导致 "database is locked" 错误
**Why it happens:** SQLite 使用文件级锁，同一时刻只允许一个写入
**How to avoid:** GORM 连接配置中设置 `db.SetMaxOpenConns(1)` 或使用 WAL 模式
**Warning signs:** 并发测试时出现间歇性 500 错误

### Pitfall 2: Vite 开发代理配置遗漏
**What goes wrong:** 前端请求后端 API 时出现 CORS 错误或 404
**Why it happens:** 开发环境前后端端口不同，未配置代理
**How to avoid:** vite.config.js 中配置 server.proxy 将 /api 请求代理到后端端口
**Warning signs:** 浏览器控制台出现 CORS 或网络错误

### Pitfall 3: JWT token 过期后前端无限循环
**What goes wrong:** token 过期 -> 401 -> 跳转登录 -> 如果登录页也发请求 -> 循环
**Why it happens:** 响应拦截器对所有 401 都跳转登录，包括登录接口本身
**How to avoid:** 登录/注册接口不需要 token，拦截器中排除这些路径
**Warning signs:** 登录页面不断刷新或白屏

### Pitfall 4: 头像上传文件名冲突
**What goes wrong:** 不同用户上传同名文件导致覆盖
**Why it happens:** 直接使用原始文件名保存
**How to avoid:** 使用时间戳或 UUID 生成唯一文件名
**Warning signs:** 用户头像显示为其他人的图片

### Pitfall 5: GORM AutoMigrate 不删除列
**What goes wrong:** 修改模型删除字段后，数据库中旧列仍然存在
**Why it happens:** AutoMigrate 只添加不删除，这是设计决策
**How to avoid:** 对于课程项目，开发阶段可以删除 .db 文件重建；生产环境需手动迁移
**Warning signs:** 查询返回意外的旧数据

### Pitfall 6: go-sqlite3 需要 CGO
**What goes wrong:** `go build` 报错 cgo 相关错误
**Why it happens:** mattn/go-sqlite3 是 CGO 绑定，需要 C 编译器
**How to avoid:** macOS 自带 clang 所以通常没问题；确保 CGO_ENABLED=1（默认值）
**Warning signs:** 编译时出现 "cgo: C compiler not found" 错误

## Code Examples

### API 路由注册示例
```go
// Source: Gin 路由分组模式 [ASSUMED]
func SetupRouter(r *gin.Engine) {
    api := r.Group("/api/v1")

    // 公开路由（无需认证）
    auth := api.Group("/auth")
    {
        auth.POST("/register", handler.Register)
        auth.POST("/login", handler.Login)
    }

    // 需要认证的路由
    user := api.Group("/user")
    user.Use(middleware.AuthMiddleware())
    {
        user.GET("/profile", handler.GetProfile)
        user.PUT("/profile", handler.UpdateProfile)
        user.PUT("/password", handler.UpdatePassword)
        user.POST("/avatar", handler.UploadAvatar)
    }

    // 静态文件服务（头像访问）
    r.Static("/uploads", "./uploads")
}
```

### Vant Area 省市区选择器使用
```vue
<!-- Source: Vant Area 组件文档 [ASSUMED] -->
<template>
  <van-field
    v-model="fieldValue"
    is-link
    readonly
    label="常住地"
    placeholder="请选择省市区"
    @click="showArea = true"
  />
  <van-popup v-model:show="showArea" position="bottom">
    <van-area
      :area-list="areaList"
      @confirm="onConfirm"
      @cancel="showArea = false"
    />
  </van-popup>
</template>

<script setup>
import { ref } from 'vue'
import { areaList } from '@vant/area-data'

const showArea = ref(false)
const fieldValue = ref('')
const areaCode = ref('')

const onConfirm = ({ selectedOptions }) => {
  showArea.value = false
  fieldValue.value = selectedOptions.map(item => item.text).join('/')
  areaCode.value = selectedOptions[selectedOptions.length - 1].value
}
</script>
```

### Vant Uploader 头像上传
```vue
<!-- Source: Vant Uploader 组件模式 [ASSUMED] -->
<template>
  <van-uploader
    v-model="fileList"
    :max-count="1"
    :after-read="onAfterRead"
    :before-read="beforeRead"
  />
</template>

<script setup>
import { ref } from 'vue'
import { showToast } from 'vant'
import request from '@/api/request'

const fileList = ref([])

const beforeRead = (file) => {
  if (!/\.(jpg|jpeg|png|gif)$/i.test(file.name)) {
    showToast('请上传图片文件')
    return false
  }
  if (file.size > 5 * 1024 * 1024) {
    showToast('图片大小不能超过5MB')
    return false
  }
  return true
}

const onAfterRead = async (file) => {
  const formData = new FormData()
  formData.append('avatar', file.file)
  try {
    const res = await request.post('/user/avatar', formData)
    showToast('头像上传成功')
  } catch (e) {
    showToast('上传失败')
  }
}
</script>
```

### Vite 代理配置
```javascript
// vite.config.js [ASSUMED]
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import { VantResolver } from '@vant/auto-import-resolver'

export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [VantResolver()]
    })
  ],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      },
      '/uploads': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  },
  resolve: {
    alias: {
      '@': '/src'
    }
  }
})
```

## State of the Art

| Old Approach | Current Approach | When Changed | Impact |
|--------------|------------------|--------------|--------|
| Vuex 4 | Pinia 3 | 2023+ | Pinia 是 Vue 官方推荐的状态管理方案 |
| Options API | Composition API + script setup | Vue 3.2+ | 更好的 TypeScript 支持和代码组织 |
| dgrijalva/jwt-go | golang-jwt/jwt/v5 | 2021 | 原库不再维护，社区 fork 为 golang-jwt |
| Vue CLI (webpack) | Vite | 2022+ | 开发速度快 10-100 倍 |
| 手动导入 Vant 组件 | unplugin-vue-components 自动导入 | Vant 4+ | 减少样板代码 |

**Deprecated/outdated:**
- **Vuex:** 官方已推荐 Pinia 替代
- **dgrijalva/jwt-go:** 已归档，使用 golang-jwt/jwt/v5
- **Vue CLI:** 新项目应使用 Vite

## Assumptions Log

| # | Claim | Section | Risk if Wrong |
|---|-------|---------|---------------|
| A1 | JWT 过期时间建议 7 天，无刷新 token 机制（课程项目简化） | Claude's Discretion | 低 - 课程项目不需要复杂的 token 刷新 |
| A2 | Sports 字段用逗号分隔字符串存储（非 JSON 数组） | Pattern 5 | 低 - 简单查询足够，不需要复杂查询 |
| A3 | Vant Area 组件 confirm 事件返回 selectedOptions 数组 | Code Examples | 中 - API 可能有变化，需实际验证 |
| A4 | unplugin-vue-components 配合 VantResolver 可自动导入 | Supporting Stack | 低 - 这是 Vant 4 官方推荐方式 |
| A5 | gorm.io/driver/sqlite 在 macOS arm64 上无需额外配置 | Pitfall 6 | 低 - macOS 自带 clang |

## Open Questions

1. **Vant Area 组件 API 细节**
   - What we know: 使用 @vant/area-data 提供省市区数据，通过 Popup + Area 组合使用
   - What's unclear: Vant 4 最新版本的 confirm 事件回调参数格式
   - Recommendation: 实现时参考 node_modules/vant 中的类型定义确认

2. **头像文件大小限制**
   - What we know: 需要限制上传文件大小和类型
   - What's unclear: 具体限制值（建议 5MB）
   - Recommendation: 前端 beforeRead 校验 + 后端 MaxMultipartMemory 双重限制

## Environment Availability

| Dependency | Required By | Available | Version | Fallback |
|------------|------------|-----------|---------|----------|
| Node.js | Frontend build | Yes | v24.16.0 | — |
| npm | Package management | Yes | 11.13.0 | — |
| Go | Backend | Yes | 1.26.3 | — |
| C compiler (CGO) | go-sqlite3 | Yes (macOS clang) | — | — |

**Missing dependencies with no fallback:** None
**Missing dependencies with fallback:** None

## Security Domain

### Applicable ASVS Categories

| ASVS Category | Applies | Standard Control |
|---------------|---------|-----------------|
| V2 Authentication | yes | bcrypt 密码哈希 + JWT token |
| V3 Session Management | yes | JWT 过期时间 + localStorage |
| V4 Access Control | yes | JWT middleware 保护路由 |
| V5 Input Validation | yes | Vant Form rules + 后端参数校验 |
| V6 Cryptography | yes | bcrypt (golang.org/x/crypto) |

### Known Threat Patterns

| Pattern | STRIDE | Standard Mitigation |
|---------|--------|---------------------|
| 密码暴力破解 | Spoofing | bcrypt 高成本 + 可选登录频率限制 |
| JWT 伪造 | Tampering | HMAC-SHA256 签名验证 |
| 路径遍历（文件上传） | Information Disclosure | filepath.Base() 清理文件名 |
| XSS via 用户输入 | Tampering | Vue 默认转义 + 后端不信任输入 |
| CORS 配置不当 | Information Disclosure | 限制 AllowOrigins |

## Sources

### Primary (HIGH confidence)
- npm registry — 版本验证 (vue 3.5.34, vant 4.9.24, vue-router 5.0.7, pinia 3.0.4, axios 1.16.1)
- Go module proxy — 版本验证 (gin v1.12.0, gorm v1.31.1, jwt v5.3.1)
- [Gin official docs - file upload](https://gin-gonic.com/en/docs/routing/upload-file/single-file/) — 文件上传模式

### Secondary (MEDIUM confidence)
- [Vant Area component docs](https://develop365.gitlab.io/vant/en-US/area/) — Area 组件数据格式
- [Gin JWT middleware patterns](https://leapcell.io/blog/secure-your-apis-with-jwt-authentication-in-gin-middleware) — JWT 中间件实现参考
- [GORM documentation](https://deepwiki.com/go-gorm/gorm) — ORM 使用模式

### Tertiary (LOW confidence)
- WebSearch results for Vue 3 project structure patterns — 目录组织参考

## Metadata

**Confidence breakdown:**
- Standard stack: HIGH - 所有版本通过 npm/go module 验证，均为生态主流库
- Architecture: HIGH - handler/service/model 是 Go Web 项目标准分层
- Pitfalls: MEDIUM - 基于社区经验和文档，部分为训练数据

**Research date:** 2026-05-23
**Valid until:** 2026-06-23 (stable stack, 30 days)
