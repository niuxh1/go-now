# Neon Ink Blog (霓虹·墨)

这是一个基于 Vue 3 和 Go (Gin) 的前后端分离博客项目，融合了赛博朋克与中国传统美学风格。

## 项目结构

- `frontend/`: Vue 3 + Vite 前端项目
- `backend/`: Go + Gin 后端项目

## 快速开始

### 1. 启动后端 (Backend)

确保你已经安装了 Go 1.18+。

```bash
cd backend
go mod tidy
go run main.go
```

后端服务将在 `http://localhost:8080` 启动。

### 2. 启动前端 (Frontend)

确保你已经安装了 Node.js (推荐 v16+)。

```bash
cd frontend
npm install
npm run dev
```

前端服务将在 `http://localhost:5173` 启动。

## 设计理念

- **视觉风格**: 深色背景 (`#0a0a0c`) 搭配霓虹青 (`#00f3ff`) 与品红 (`#ff0055`)，点缀以玉石绿 (`#00ff9d`)。
- **排版**: 结合现代无衬线字体与传统衬线字体（如楷体）的韵味。
- **交互**: 极简的卡片式布局，带有微交互动画。

## 技术栈

- **Frontend**: Vue 3, Vite, Axios, Lucide Icons
- **Backend**: Go, Gin, CORS Middleware
