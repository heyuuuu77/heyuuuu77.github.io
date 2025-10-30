---
title: "ClickHouse 性能之巅"
date: 2025-03-24T09:34:08+08:00
draft: false
description: "ClickHouse 读书笔记与实战经验"
tags: ["ClickHouse", "数据库", "书单"]
categories: ["后端开发"]
author: "Heyuuuu"
---



### Day1

#### OLTP & OLAP
- OLTP, OnLine Transaction Process 联机事务处理 (MySQL, PostgreSQL, Oracle)
    OLTP使用的是高反式的表。好处是数据结构不容易，事务处理速度快

- OLAP, OnLine Analyse Process 联机分析处理(Hive, HBase, ClickHouse)
    OLAP数据库被称为数据仓库(数仓), 它是反范式的。数据冗余度很高，正因为如此。它被用做数据分析，因为单个宽表就包含了很多字段，不需要进行连表(JOIN)查询。


#### 第零范式
不遵循第一范式，都被称之为第零范式，即字段中可以存储json或者文档

#### 第一范式
指数据表中的每个属性都不可再拆分。


### Day2
