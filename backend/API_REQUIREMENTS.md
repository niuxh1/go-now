# GoNow Blog 后端 API 接口规范说明书

**版本**: v2.0 (Final Release)
**状态**: 正式生效
**维护者**: GoNow Team

## 1. 概述 (Overview)

本文档作为 `go-now-blog` 项目的**唯一权威后端接口规范**。任何后端重构或重写工作，必须严格遵守本文档中定义的路由、HTTP 方法、请求参数、响应结构及状态码标准。

本规范旨在确保后端服务与现有的 Vue 3 前端应用 (`go-now-blog-frontend`) 完美兼容，同时遵循 RESTful API 设计的最佳实践。

---

## 2. 通用技术标准 (General Standards)

### 2.1 协议与端点
- **协议**: HTTP/1.1
- **基础路径 (Base URL)**: `http://localhost:8080/api`
- **内容类型 (Content-Type)**: 请求与响应体均为 `application/json; charset=utf-8`

### 2.2 跨域资源共享 (CORS)
后端必须启用 CORS 中间件以支持前后端分离开发调试。
- **Access-Control-Allow-Origin**: `http://localhost:5173` (生产环境可配置为具体域名)
- **Access-Control-Allow-Methods**: `GET, POST, PUT, DELETE, OPTIONS`
- **Access-Control-Allow-Headers**: `Origin, Content-Type, Accept, Authorization, X-Requested-With`
- **Access-Control-Allow-Credentials**: `true` (必须开启，否则无法处理认证状态)
- **预检请求 (Preflight)**: 针对 `OPTIONS` 请求，必须直接返回 `200 OK` 或 `204 No Content`，不执行业务逻辑。

### 2.3 身份认证 (Authentication)
- **标准**: JWT (JSON Web Token)
- **签名算法**: 推荐使用 `HS256`。
- **传递方式**: HTTP 请求头 `Authorization`。
  - 格式: `Bearer <token>`
- **Token 载荷 (Payload)**:
  前端依赖 JWT 解析来获取当前用户的 ID 和角色。Token **必须**包含以下字段：
  ```json
  {
    "user_id": 1,       // (Number) 用户唯一标识
    "username": "admin",// (String) 用户名
    "role": "admin",    // (String) 角色权限标识，必须为 "admin" 或 "user"
    "exp": 1735689600   // (Number) 过期时间戳
  }
  ```

### 2.4 状态码规范 (Status Codes)
- **200 OK**: 请求成功。
- **201 Created**: 资源创建成功 (如注册、发帖)。
- **400 Bad Request**: 请求参数无效 (如缺少必填项、格式错误)。
- **401 Unauthorized**: 未提供 Token 或 Token 无效/过期。
- **403 Forbidden**: 权限不足 (如普通用户尝试发帖，或尝试重置管理员密码)。
- **404 Not Found**: 资源不存在。
- **500 Internal Server Error**: 服务器内部错误。

### 2.5 统一错误响应结构
当 HTTP 状态码 >= 400 时，响应体必须符合以下 JSON 结构，前端会直接显示 `error` 字段的内容给用户。
```json
{
  "error": "具体的错误描述信息 (例如：用户名已存在 / 密码长度不足)"
}
```

---

## 3. 接口详细定义 (API Endpoints)

### 3.1 认证模块 (Auth)

#### 3.1.1 注册用户
- **端点**: `POST /auth/register`
- **描述**: 创建一个新的普通用户账户。
- **请求体 (Request Body)**:
  | 字段 | 类型 | 必填 | 描述 | 校验规则 |
  | :--- | :--- | :--- | :--- | :--- |
  | `username` | String | 是 | 用户名 | 唯一，非空 |
  | `email` | String | 是 | 邮箱 | 唯一，合法邮箱格式 |
  | `password` | String | 是 | 密码 | 长度 >= 6 |
- **成功响应 (201 Created)**:
  ```json
  {
    "message": "User registered successfully",
    "user": {
      "id": 1,
      "username": "user1",
      "email": "user1@example.com",
      "role": "user",
      "created_at": "2025-01-01T12:00:00Z"
    }
  }
  ```

#### 3.1.2 用户登录
- **端点**: `POST /auth/login`
- **描述**: 验证凭据并颁发 JWT Token。
- **请求体**:
  | 字段 | 类型 | 必填 | 描述 |
  | :--- | :--- | :--- | :--- |
  | `username` | String | 是 | 用户名 |
  | `password` | String | 是 | 密码 |
- **成功响应 (200 OK)**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1Ni..." // JWT String
  }
  ```
- **错误响应**:
  - `401 Unauthorized`: `{"error": "invalid credentials"}`

#### 3.1.3 重置密码
- **端点**: `POST /auth/reset-password`
- **描述**: 普通用户忘记密码时，通过验证邮箱重置密码。
- **安全约束**: **严禁**通过此接口重置 `role="admin"` 的用户密码。
- **请求体**:
  | 字段 | 类型 | 必填 | 描述 |
  | :--- | :--- | :--- | :--- |
  | `username` | String | 是 | 用户名 |
  | `email` | String | 是 | 注册邮箱 |
  | `new_password` | String | 是 | 新密码 |
- **成功响应 (200 OK)**:
  ```json
  {
    "message": "Password updated successfully"
  }
  ```
- **错误响应**:
  - `403 Forbidden`: 尝试重置管理员密码。
  - `400 Bad Request`: 邮箱与用户名不匹配。

---

### 3.2 文章模块 (Articles)

#### 3.2.1 获取文章列表
- **端点**: `GET /articles`
- **描述**: 获取所有公开文章，用于首页流和归档页面。
- **排序**: 建议按 `date` 倒序排列 (最新的在前)。
- **响应 (200 OK)**:
  返回文章对象数组。列表页为节省流量，`content` 字段可为空字符串或截断。
  ```json
  [
    {
      "id": 10,
      "title": "Go Web 开发实战",
      "summary": "本文介绍了 Gin 框架的使用...", // 必须字段，前端卡片显示
      "content": "", // 可为空
      "date": "2025-12-08", // 格式: YYYY-MM-DD
      "category": "Tech",
      "views": 120
    },
    { ... }
  ]
  ```

#### 3.2.2 获取文章详情
- **端点**: `GET /articles/:id`
- **描述**: 获取指定 ID 的文章完整内容。
- **路径参数**: `id` (Integer)
- **响应 (200 OK)**:
  ```json
  {
    "id": 10,
    "title": "Go Web 开发实战",
    "content": "# 第一章 环境搭建\n\n...", // 必须包含完整的 Markdown 文本
    "date": "2025-12-08",
    "category": "Tech",
    "summary": "本文介绍了 Gin 框架的使用...",
    "views": 121,
    "created_at": "2025-12-08T09:00:00Z"
  }
  ```
- **错误响应**:
  - `404 Not Found`: `{"error": "Article not found"}`

#### 3.2.3 发布文章
- **端点**: `POST /articles`
- **权限**: **仅限管理员** (Role = "admin")
- **鉴权**: 需验证 `Authorization` 头。
- **请求体**:
  | 字段 | 类型 | 必填 | 描述 |
  | :--- | :--- | :--- | :--- |
  | `title` | String | 是 | 文章标题 |
  | `content` | String | 是 | Markdown 内容 |
  | `date` | String | 否 | 发布日期 (YYYY-MM-DD)，若空则取当前日期 |
- **成功响应 (201 Created)**:
  ```json
  {
    "id": 11,
    "title": "新文章",
    "date": "2025-12-09",
    ... // 其他字段同详情页
  }
  ```

---

### 3.3 评论模块 (Comments)

#### 3.3.1 获取文章评论
- **端点**: `GET /articles/:id/comments`
- **描述**: 获取指定文章的所有评论。
- **响应 (200 OK)**:
  ```json
  [
    {
      "id": 55,
      "content": "非常有用的教程！",
      "created_at": "2025-12-08T10:30:00Z",
      "user_id": 2,
      "user": { // 关联的用户信息对象
        "id": 2,
        "username": "developer_jack"
        // 邮箱等敏感信息不建议返回，除非有显示头像需求
      }
    }
  ]
  ```

#### 3.3.2 发布评论
- **端点**: `POST /articles/:id/comments`
- **权限**: **登录用户** (任意 Role)
- **鉴权**: 需验证 `Authorization` 头。
- **请求体**:
  ```json
  {
    "content": "期待更新下一篇！"
  }
  ```
- **逻辑**: 后端必须从 Token 中提取 `user_id` 并关联到该评论，不可信任前端传递的用户ID。
- **成功响应 (201 Created)**:
  ```json
  {
    "id": 56,
    "content": "期待更新下一篇！",
    "article_id": 10,
    "user_id": 3,
    "created_at": "2025-12-08T11:00:00Z"
  }
  ```

---

### 3.4 辅助模块 (Misc)

#### 3.4.1 获取标签
- **端点**: `GET /tags`
- **描述**: 获取所有标签列表（前端 API Client 定义了此接口）。
- **响应 (200 OK)**:
  ```json
  [
    { "id": 1, "name": "Golang" },
    { "id": 2, "name": "Frontend" }
  ]
  ```

---

## 4. 数据库初始化要求 (Initialization)

为了确保系统“开箱即用”，后端启动时应执行以下初始化逻辑：

1.  **自动迁移 (Auto Migration)**:
    - 确保 `users`, `articles`, `comments`, `tags` 表结构在数据库中已创建。
2.  **默认管理员 (Admin Seeder)**:
    - 检查数据库中是否存在 `username="admin"` 的用户。
    - **如果存在**: 更新其密码为一个新的**随机强密码**，并强制重置其 Role 为 `admin`。
    - **如果不存在**: 创建该用户。
    - **关键**: 必须在**控制台日志 (Stdout)** 中显式打印出生成的账号和密码，以便开发者/用户登录。
3.  **示例数据 (Article Seeder)**:
    - (可选) 扫描 `posts/` 目录下的 Markdown 文件并导入数据库，方便演示。

---

## 5. 数据库模型 (Database Schema)

以下为符合接口要求的数据库表结构定义。

### 5.1 用户表 (Users)

| 字段名 | 类型 | 必填 | 默认值 | 约束 | 描述 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `id` | INT | 是 | - | PRIMARY KEY, AUTO_INCREMENT | 用户唯一标识 |
| `username` | VARCHAR(255) | 是 | - | UNIQUE | 用户名 |
| `email` | VARCHAR(255) | 是 | - | UNIQUE | 电子邮箱 |
| `password` | VARCHAR(255) | 是 | - | - | 密码哈希 |
| `role` | VARCHAR(50) | 否 | `'user'` | - | 角色 ('admin' / 'user') |
| `created_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | - | 创建时间 |
| `updated_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | ON UPDATE CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | 否 | `NULL` | - | 软删除时间戳 |

### 5.2 文章表 (Articles)

| 字段名 | 类型 | 必填 | 默认值 | 约束 | 描述 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `id` | INT | 是 | - | PRIMARY KEY, AUTO_INCREMENT | 文章唯一标识 |
| `title` | VARCHAR(255) | 是 | - | - | 文章标题 |
| `summary` | VARCHAR(1000) | 否 | - | - | 文章摘要 (列表显示) |
| `content` | LONGTEXT | 否 | - | - | 完整 Markdown 内容 |
| `date` | DATE | 否 | - | - | 发布日期 (YYYY-MM-DD) |
| `category` | VARCHAR(100) | 否 | - | - | 文章分类 |
| `views` | INT | 否 | `0` | - | 浏览量 |
| `created_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | - | 创建时间 |
| `updated_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | ON UPDATE CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | 否 | `NULL` | - | 软删除时间戳 |

### 5.3 评论表 (Comments)

| 字段名 | 类型 | 必填 | 默认值 | 约束 | 描述 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `id` | INT | 是 | - | PRIMARY KEY, AUTO_INCREMENT | 评论唯一标识 |
| `content` | TEXT | 是 | - | - | 评论内容 |
| `article_id` | INT | 是 | - | FOREIGN KEY -> articles.id | 所属文章 ID |
| `user_id` | INT | 是 | - | FOREIGN KEY -> users.id | 评论者 ID |
| `created_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | - | 创建时间 |
| `updated_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | ON UPDATE CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | 否 | `NULL` | - | 软删除时间戳 |

### 5.4 标签表 (Tags)

| 字段名 | 类型 | 必填 | 默认值 | 约束 | 描述 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `id` | INT | 是 | - | PRIMARY KEY, AUTO_INCREMENT | 标签唯一标识 |
| `name` | VARCHAR(100) | 是 | - | UNIQUE | 标签名称 |
| `created_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | - | 创建时间 |
| `updated_at` | TIMESTAMP | 否 | `CURRENT_TIMESTAMP` | ON UPDATE CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | 否 | `NULL` | - | 软删除时间戳 |

### 5.5 文章-标签关联表 (Article Tags)

| 字段名 | 类型 | 必填 | 默认值 | 约束 | 描述 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `article_id` | INT | 是 | - | PK, FK -> articles.id | 文章 ID |
| `tag_id` | INT | 是 | - | PK, FK -> tags.id | 标签 ID |

**文档结束**。