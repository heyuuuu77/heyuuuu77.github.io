---
title: "Hugo 静态博客搭建"
date: 2025-02-08T17:06:08+08:00
draft: false
description: "Hugo 建站教程与使用技巧"
tags: ["Hugo", "博客"]
categories: ["工具"]
author: "Heyuuuu"
---


使用Hugo搭建 github page越到的坑。 首先 hugo 命令会生成静态文件。 
本地查看的时候执行的命令是 `hugo server`。 这个时候，草稿是不会暂时出来的。如果需要暂时草稿， 需要加上参数 `-D`。 

创建新的文章的命令是: `hugo new posts/your-post-name.md`. 注意这个时候自动生成的 markdown文件中的 draft 默认为true，表示当前是草稿模式。
如果要可以见，需要将其设置成false。

接下来如果本地看了没问题，重点是需要关闭本地服务, 即`hugo server` 打开的服务，并重新执行`hugo`。 
不这样做，hugo会将public下的静态文件全部编译成本地。 这时候提交到github，静态文件的跳转链接会乱掉。全部指向本地链接。

提交至github page时，不用把hugo项目全部提交上去，而是进入到hugo项目的public 路径下，执行git init， 然后添加 remote。 在提交就OK了

