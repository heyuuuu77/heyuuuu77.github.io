# Heyuuuu 的个人博客

[![Hugo](https://img.shields.io/badge/Hugo-0.145-blue.svg)](https://gohugo.io)
[![Theme](https://img.shields.io/badge/Theme-PaperMod-green.svg)](https://github.com/adityatelange/hugo-PaperMod)
[![License](https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-red.svg)](https://creativecommons.org/licenses/by-nc-sa/4.0/)

## 📝 简介

这是一个基于 Hugo 构建的个人技术博客，记录学习笔记和生活点滴。

**在线访问**: https://heyuuuu77.github.io/

## 🎨 主题

使用 [PaperMod](https://github.com/adityatelange/hugo-PaperMod) 主题，特点：
- ✅ 简洁现代的设计
- ✅ 完美支持中文
- ✅ 响应式布局
- ✅ 深色/浅色主题切换
- ✅ 全文搜索功能
- ✅ 自动目录生成
- ✅ 代码高亮
- ✅ SEO 优化

## 🚀 快速开始

### 本地开发

```bash
# 启动本地服务器
hugo serve

# 启动本地服务器（包含草稿）
hugo serve -D

# 访问
open http://localhost:1313
```

### 创建新文章

```bash
# 创建新文章
hugo new posts/my-new-post.md

# 编辑文章
vim content/posts/my-new-post.md
```

### 构建静态文件

```bash
# 生成静态文件到 public/ 目录
hugo

# 清理并重新构建
rm -rf public resources && hugo
```

## 📁 项目结构

```
blog/
├── archetypes/          # 内容模板
├── assets/              # 资源文件
├── code/                # 代码示例
│   ├── golang/          # Go 示例
│   └── python/          # Python 示例
├── content/             # 博客内容 ⭐
│   ├── posts/           # 文章目录
│   ├── archives.md      # 归档页
│   └── search.md        # 搜索页
├── static/              # 静态资源
├── themes/              # 主题
│   └── PaperMod/        # 当前使用主题
├── hugo.toml            # 配置文件
└── README.md            # 本文件
```

## ✍️ 文章格式

### Front Matter 示例

```yaml
---
title: "文章标题"
date: 2025-10-30T14:00:00+08:00
draft: false
author: "Heyuuuu"
description: "文章简介"
tags: ["Go", "Python"]
categories: ["技术"]
toc: true
cover:
    image: "/images/cover.jpg"
    alt: "封面图片"
    caption: "图片说明"
---
```

### 常用 Front Matter 字段

- `title`: 文章标题 (必填)
- `date`: 发布日期 (必填)
- `draft`: 是否为草稿，true=草稿，false=发布
- `author`: 作者
- `description`: 文章描述（SEO）
- `tags`: 标签列表
- `categories`: 分类列表
- `toc`: 是否显示目录
- `cover`: 封面图片配置

### 内容编写技巧

#### 1. 使用提示框

```markdown
> 💡 **提示**: 这是一个提示信息

> ⚠️ **警告**: 这是一个警告信息

> ✅ **成功**: 这是一个成功信息

> ❌ **错误**: 这是一个错误信息
```

#### 2. 代码块

````markdown
```go
func main() {
    fmt.Println("Hello, World!")
}
```
````

#### 3. 表格

```markdown
| 功能 | 说明 |
|------|------|
| TOC  | 自动生成目录 |
| 搜索 | 全文搜索 |
```

## 🔧 配置说明

### 主要配置项 (hugo.toml)

```toml
# 基础配置
baseURL = "https://heyuuuu77.github.io/"
languageCode = "zh-cn"
theme = "PaperMod"
title = "Heyuuuu"

# PaperMod 主题配置
[params]
ShowReadingTime = true       # 显示阅读时间
ShowShareButtons = true      # 显示分享按钮
ShowPostNavLinks = true      # 显示上下篇导航
ShowBreadCrumbs = true       # 显示面包屑
ShowCodeCopyButtons = true   # 代码复制按钮
ShowToc = true               # 显示目录
```

## 📊 内容统计

- 📝 文章总数: 23篇
- 🏷️ 标签数: 15+
- 🗂️ 分类: 技术、生活、随笔

### 主要技术栈

- **后端**: Go, Python, Django, FastAPI
- **数据库**: Redis, Elasticsearch, Kafka
- **工具**: Docker, Hugo, Git
- **算法**: LeetCode, 设计模式

## 🔗 相关链接

- [Hugo 官方文档](https://gohugo.io/documentation/)
- [PaperMod 主题文档](https://github.com/adityatelange/hugo-PaperMod/wiki)
- [Markdown 语法](https://www.markdownguide.org/)

## 📝 待办事项

- [ ] 为文章添加分类
- [ ] 为文章添加封面图
- [ ] 启用评论系统
- [ ] 优化 SEO
- [ ] 添加更多代码示例

## 🤝 贡献

欢迎提出建议和反馈！

## 📄 许可证

内容采用 [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/) 许可协议。

---

**Made with ❤️ by Heyuuuu**

