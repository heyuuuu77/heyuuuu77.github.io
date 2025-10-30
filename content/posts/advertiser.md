---
title: "广告系统基础"
date: 2025-04-14T09:53:08+08:00
draft: false
description: "广告计费模式与投放策略"
tags: ["广告", "商业模式"]
categories: ["业务知识"]
author: "Heyuuuu"
---



#### 广告的计费模式有哪几种
1. CPM (Cost Per Mile)千次展示计费
2. CPC (Cost Per Click) 点击收费
3. CPA (Cost Per Action) 行动收费 广告主在用户完成某个特定行为才会支付费用
4. CTP (Cost Per Time) 按时长收费
5. CPI (Cost Per Install) 按安装收费等


#### 现代广告系统架构
- 广告投放平台（DSP, Demand-Side Platform） 广告主通过DSP进行投放，DSP负责竞价，定向，优化等功能
- 供应方平台（SSP, Supply-Side Platform） 媒体方通过SSP管理广告库存，广告库存就是广告位。如果自己开发一个app，想接广告。不需要自己找广告主，直接把广告位卖给SSP就行

- 广告交易市场(Ad Exchange), 一个开发的广告竞价平台，链接DSP和SSP。 
- 数据管理平台(DMP, Data Managerment Platform), 收集，分析用户数据，提高广告定向能力等

- 广告监测平台(Ad Monitor), 监测广告投放效果，比如广告投放的曝光量，点击量，转化等，归因分析。