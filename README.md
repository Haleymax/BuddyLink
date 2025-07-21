# BuddyLink

一个基于 Go + Vue 3 的现代化全栈 Web 应用，提供用户认证、文件存储等核心功能。

## 技术栈

**后端**
- Go 1.23 + Gin 框架
- GORM + MySQL 数据库
- Redis 缓存
- MinIO 对象存储
- JWT 身份验证
- SMTP 邮件服务

**前端**
- Vue 3 + TypeScript
- Vite 构建工具
- Naive UI 组件库
- Vue Router 路由管理
- Axios HTTP 客户端

## 功能特性

- 🔐 用户注册/登录系统
- 🛡️ JWT 身份验证中间件
- 📁 文件上传与存储
- 📧 邮件服务集成
- 🗄️ Redis 缓存支持
- 🎨 响应式前端界面

## 快速开始

### 环境要求

- Go 1.23+
- Node.js 16+
- MySQL 5.7+
- Redis 6.0+
- MinIO

### 后端启动

```bash
cd backend
cp config/config_template.yaml config/config.yaml
# 编辑配置文件
go mod tidy
go run main.go
```

### 前端启动

```bash
cd front
npm install
npm run dev
```

## 配置说明

复制 `backend/config/config_template.yaml` 为 `config.yaml`，并配置以下服务：

- MySQL 数据库连接
- Redis 缓存服务  
- MinIO 对象存储
- SMTP 邮件服务