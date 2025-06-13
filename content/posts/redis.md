+++
date = '2025-02-05T09:39:00+08:00'
draft = false
title = 'Redis'
tags = ["Redis"]
UseHugoToc = true
+++


#### 数据类型
- String: 字符串
- Hash: 哈希, 是一个键值对集合，类似于 Python 中的字典。适合存储对象，一个哈希可以包含多个字段和对应的值，每个哈希最多可存储2^32 -1个键值对。
    - 常见命令:
        1. 设置单个字段值：HSET key field value，例如 HSET user:1 name "Bob"。
        2. 获取单个字段值：HGET key field，例如 HGET user:1 name 会返回 "Bob"。
        3. 获取所有字段和值：HGETALL key。

- List: 列表
    - 常见命令:
        LPUSH RPUSH LPOP RPOP LINDEX LRANGE LTRIM 等
        是一个有序，可以重复的线性数据接口，基于双向链表，支持快速的插入和删除操作。
        1. LPUSH key value [value ...]：将一个或多个值插入到列表的头部。
        2. RPUSH key value [value ...]：将一个或多个值插入到列表的尾部。
        3. LPOP key：移除并返回列表的第一个元素。
        4. RPOP key：移除并返回列表的最后一个元素。
        5. LINDEX key index：返回列表中指定索引的元素。
        6. LRANGE key start stop：返回列表中指定范围内的元素。
        7. LTRIM key start stop：截取列表，只保留指定范围内的元素。

    - 消息队列: FIFO (First In First Out)
        1. LPUSH + RPOP 
        2. LPUSH + BRPOP(阻塞式右取出)

    - 栈: LIFO (Last In First Out)
        1. LPUSH + LPOP
        2. RPUSH + RPOP
        
        
- Set: 集合
    - 常见命令:
        SADD SMEMBERS SREM SISMEMBER SCARD 等
        是一个无序，不重复的集合数据接口，基于哈希表实现，支持快速的插入、删除和查找操作。
        1. SADD key member [member...]：向集合中添加一个或多个成员。
    - 常用场景:
        点赞/投票， 关注/取消关注， 共同关注， 推荐系统。
        1. 去重：集合可以用于存储唯一的元素，例如用户 ID 列表。
        2. 交集、并集、差集：可以对多个集合进行交集、并集、差集操作，例如找出两个用户共同关注的话题。
        3. ZRANGE key start stop：返回有序集合中指定范围内的元素。
        4. ZREM key member：移除有序集合中的一个或多个成员。
        5. ZSCORE key member：返回有序集合中指定成员的分数。
        6. ZINCRBY key increment member：增加有序集合中指定成员的分数。
        7. SISMEMBER key member：判断成员是否在集合中. 

- Sorted Set: 有序集合
    - 核心用途
        - 排行榜
        ```
            ZADD leaderboard 1000 "player1"  # 玩家1得分1000
            ZADD leaderboard 800 "player2"  # 玩家2得分800
            ZREVRANGE leaderboard 0 2  # 获取排行榜前3名（降序）
        ```
        - 消息队列
        - 延时任务
        - 计数器
