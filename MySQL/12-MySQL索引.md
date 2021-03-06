---
author: "kuonz"
draft: false
title: "MySQL索引"
date: 2020-04-05
categories: ["MySQL"]
---
  
## 索引概述

### 概念

索引是一种帮助MySQL高效查询数据和排序数据的数据结构

### 本质

在数据之外，数据库系统还维护一种数据结构，该数据结构称为索引，作用是能够帮助MySQL快速查找和排序数据

### 底层

MySQL索引的底层实现是 B+ 树，此外还有Hash索引，FULL-TEXT全文索引，R-Tree索引  

### 注意

索引本身也很大，不可能全部存储到内存中，因此索引往往以索引文件的形式存储到磁盘上

mysql默认存储引擎innodb只支持B+树形式的索引，不支持哈希索引

### 优缺点

| 优点                           | 缺点               |
| ------------------------------ | ------------------ |
| 提高检索效率，降低数据库IO成本 | 索引占的空间比较大 |
| 降低排序成本，降低CPU的消耗    | DML操作的速度变慢  |

### 索引的分类

| 分类       | 说明                                             |
| ---------- | ------------------------------------------------ |
| 普通索引   | 一个索引包含单个列，一个表可以有多个单列索引     |
| 唯一索引   | 基于普通索引，只是列中的值必须唯一，但允许有空值 |
| 复合索引   | 一个索引可以包含多个列，推荐使用复合索引         |
| 聚集索引   | 如果索引和数据放在同一个文件，称该索引为聚集索引 |
| 非聚集索引 | 如果索引和数据放在不同文件，称该索引为非聚集索引 |



## 索引的增删改查

### 创建

```mysql
CREATE [UNIQUE] INDEX 索引名 
ON 表名(字段名);
```

### 删除

```mysql
DROP INDEX [索引名] 
ON 表名;
```

### 查看

```mysql
SHOW INDEX FROM 表名;
```



## 索引的使用场景

### 适合建立索引

* 频繁查询的字段
* 用于筛选和排序的字段
* 主键自动建立唯一索引
* 外键字段

### 不适合建立索引

* 表记录太少
* 频繁DML操作的字段
* 重复太多值的字段
* WHERE中用不到的字段
* GROUP BY中用不到的字段

