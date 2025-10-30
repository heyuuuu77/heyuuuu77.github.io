---
title: "Docker 使用踩坑记录"
date: 2025-09-17T11:36:54+08:00
draft: false
description: "Docker 实战中遇到的问题与解决方案"
tags: ["Docker", "容器化"]
categories: ["运维"]
author: "Heyuuuu"
---



#### 上下文
线上一个小型服务，直接使用docker-compose 管理。 docker-compose up 的时候首先报错
```
[+] Building 16.4s (1/1) FINISHED                                                                                                                                                                                                                docker-container:linux-builder
 => ERROR [web internal] booting buildkit                                                                                                                                                                                                                                 16.4s
 => => pulling image moby/buildkit:buildx-stable-1                                                                                                                                                                                                                        16.4s
------
[+] Running 0/1l] booting buildkit:
 ⠸ Service web  Building                                                                                                                                                                                                                                                  16.4s
Error response from daemon: Get "https://registry-1.docker.io/v2/": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)

```

查看 /etc/docker/daemon.json配置，显示了
```
{
  "data-root": "/var/lib/docker",
  "registry-mirrors": [
    "https://docker.xuanyuan.me",
    "https://docker.aityp.com",
    "https://mirror.baidubce.com"
  ]
}
```

一开始的解决思路就是新增国内镜像。但是问题是出在 buildKit 上。 Docker BuildKit 是新一代的镜像构建引擎。它会使用 DockerHub 官方 registry（https://registry-1.docker.io/v2/）
拉取镜像。因此会设计到网络访问问题，已知国内的云服务器访问不了外网。
解决这个问题的方式有两种，但是目的是同一个，就是禁用 buildKit. 

##### 方案一 （临时禁用BuildKit）：
```
export DOCKER_BUILDKIT=0
export COMPOSE_DOCKER_CLI_BUILD=0
```
执行上述命令后再执行docker-compose up 
或者直接执行 `DOCKER_BUILDKIT=0 docker-compose up --build`

##### 方案二 （修改～/.bashrc 或 ~/.zshrc）

```
echo 'export DOCKER_BUILDKIT=0' >> ~/.zshrc
source ~/.zshrc

echo 'export DOCKER_BUILDKIT=0' >> ~/.bashrc
echo 'export COMPOSE_DOCKER_CLI_BUILD=0' >> ~/.bashrc
source ~/.bashrc
```

修改 daemon.json 之后。又报错
```
Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: invalid rootfs: not an absolute path, or a symlink: unknown
```

这个问题是容器运行时 (runc) 要求根文件系统 (rootfs) 的路径必须是：
1. 绝对路径 
2. 真实路径

如果是 相对路径 或者 符号链接 就会触发这个错误。

因为 "data-root": "/var/lib/docker" 中的 /var/lib/docker 实际是 /data/logs/docker-data 的软链接。 最后改成 /data/logs/docker-data 就OK了。


#### Tips

docker-compose up --build 中的 --build 会覆盖缓存，重新构建，确保是最新代码。当Dockerfile 或者构建目录中的源代码/依赖配置(package.json, requirement.txt) 有修改时，
或者怀疑镜像缓存导致代码未更新，可以强制重构，避免缓存干扰