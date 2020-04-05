---
author: "kuonz"
draft: false
title: "NoSQL与Redis"
date: 2020-04-05
categories: ["Redis"]
---

## NoSQL

### NoSQL诞生背景

传统的关系型数据库由于是基于磁盘存储的，且数据与数据间可能存在复杂关系，所以在高并发量和海量用户的情况下性能较差，且关系型数据库的扩容和伸缩能力不佳，不便于大规模集群

### NoSQL简介

`NoSQL` 是 `Not-Only SQL` 的简称，泛指非关系型的数据库

### NoSQL特征

1. 易扩容，易伸缩
2. 大数据下保持高性能
3. 灵活的数据模型
4. 高可用

### 常见NoSQL数据库

1. Redis
2. MongoDB
3. Memcache
4. HBase



## Redis

### Redis简介

`Redis` 是 `REmote DIctionary Server` 的缩写

`Redis` 是用 C 语言开发的开源的高性能键值对数据库

### Redis特征

1. 数据与数据间没有必然的关系（区别于关系型数据库）
2. 单进程，单线程，原子操作
3. 高性能
4. 多数据类型支持
5. 支持持久化
6. 支持主从复制，哨兵模式，集群模式

### Redis应用

1. 加速热点数据的查询
2. 分布式锁
3. 分布式数据共享
4. 消息队列
5. 任务队列
6. 即时信息查询
7. 时效性信息控制