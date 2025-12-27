# GoNow Blog

![GoNow Blog](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go)
![Vue](https://img.shields.io/badge/Vue.js-3.0-4FC08D?logo=vue.js)
![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?logo=mysql)
![TailwindCSS](https://img.shields.io/badge/Tailwind-3.0-38B2AC?logo=tailwindcss)

## 📖 项目介绍

**GoNow Blog** 是一个为开发者设计的全栈博客系统。它不仅提供了极简的阅读体验，还集成了强大的管理功能。

### 核心价值
- **极速性能**：后端 Go 语言高并发支持，前端 Vite 秒级启动。
- **创作自由**：支持沉浸式 Markdown 创作，自动同步文件系统。
- **深度定制**：管理员可在页面直接修改站点属性，实时生效。
- **互动友好**：无缝的评论流，支持匿名与实名评论。

---

## 🛠️ 本地开发指南

### 前置要求
- **Go**: 1.22+
- **Node.js**: 18+
- **MySQL**: 8.0+

### 1. 数据库准备
```sql
CREATE DATABASE go_now_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
在 `backend/DataBase/db.go` 中修改 DSN 连库配置。

### 2. 启动服务
- **后端**: `cd backend && go run main.go`
- **前端**: `cd frontend && npm install && npm run dev`

---

## 🚀 Linux 服务器部署流程

本指南假设您使用的是 Ubuntu/Debian/CentOS 系统。

### 1. 环境准备
安装 Nginx, MySQL, Go 和 Node.js 环境。

### 2. 后端部署 (Systemd)
1.  **编译程序**:
    ```bash
    cd backend
    go build -o gonow-server main.go
    ```
2.  **创建服务文件**: `sudo nano /etc/systemd/system/gonow-blog.service`
    ```ini
    [Unit]
    Description=GoNow Blog Backend
    After=network.target mysql.service

    [Service]
    Type=simple
    User=root
    WorkingDirectory=/var/www/go-now-blog/backend
    ExecStart=/var/www/go-now-blog/backend/gonow-server
    Restart=on-failure

    [Install]
    WantedBy=multi-user.target
    ```
3.  **启动后端**:
    ```bash
    sudo systemctl daemon-reload
    sudo systemctl enable gonow-blog
    sudo systemctl start gonow-blog
    ```

### 3. 前端部署 (Nginx)
1.  **打包前端**:
    ```bash
    cd frontend
    npm install
    npm run build
    ```
2.  **配置 Nginx**: `sudo nano /etc/nginx/sites-available/gonow-blog`
    ```nginx
    server {
        listen 80;
        server_name your_domain.com;

        # 前端静态文件
        location / {
            root /var/www/go-now-blog/frontend/dist;
            index index.html;
            try_files $uri $uri/ /index.html;
        }

        # 后端接口代理
        location /api {
            proxy_pass http://localhost:8080;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
    }
    ```
3.  **生效配置**:
    ```bash
    sudo ln -s /etc/nginx/sites-available/gonow-blog /etc/nginx/sites-enabled/
    sudo nginx -t
    sudo systemctl restart nginx
    ```

---

## 📝 运维须知

-   **日志查看**: `journalctl -u gonow-blog -f`
-   **初次登录**: 后端程序首次启动时会在控制台输出 `admin` 账户的随机初始密码，请及时修改。
-   **文件同步**: 文章会自动同步在 `backend/posts` 目录下，您可以直接通过 Git 提交 Markdown 文件来实现“GitOps”式的文章发布。

## 📄 开源协议
MIT License
