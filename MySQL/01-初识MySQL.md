---
author: "kuonz"
draft: false
title: "初识MySQL"
date: 2020-04-05
categories: ["MySQL"]
---
  
## 关系型数据库

### 关系型数据库的作用

* 可以将数据持久化到本地
* 可以结构化查询



### 概念对比

* DB：数据库，存放数据的容器
* DBMS：数据库管理系统，用于管理DB
* SQL：结构化查询语言，是关系型数据库的通用语言



### 常用的DBMS

* MySQL
* Oracle
* SQL Server



### MySQL介绍

* MySQL是一个关系型数据库
* MySQL开源免费
* MySQL由瑞典 MySQL AB公司开发，现在被Oracle收购



## MySQL下载与安装

直接使用Docker进行MySQL的下载与安装

```shell
docker search mysql

docker pull mysql:lastest

docker run \
-id \
-name mysql \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=123456 \
--restart always mysql:lastest
```



## MySQL配置

### 配置文件位置

```
sudo vim /etc/my.cnf
```

### 配置介绍

```
[mysql]
# 设置mysql客户端默认字符集
default-character-set=utf8 
socket=/var/lib/mysql/mysql.sock

[mysqld]
skip-name-resolve
#设置3306端口
port = 3306 
socket=/var/lib/mysql/mysql.sock
# 设置mysql的安装目录
basedir=/usr/local/mysql
# 设置mysql数据库的数据的存放目录
datadir=/usr/local/mysql/data
# 允许最大连接数
max_connections=200
# 服务端使用的字符集默认为8比特编码的latin1字符集
character-set-server=utf8
# 创建新表时将使用的默认存储引擎
default-storage-engine=INNODB
lower_case_table_name=1
max_allowed_packet=16M
```



## MySQL存储引擎

### 查看存储引擎

```mysql
SHOW ENGINES;
```

### 查看默认存储引擎与当前存储引擎

```mysql
SHOW VARIABLES LIKE '%storage_engine%';
```

### InnoDB和MyISAM对比

| 对比项 | MyISAM                                                 | InnoDB                                     |
| ------ | ------------------------------------------------------ | ------------------------------------------ |
| 主外键 | 不支持                                                 | 支持                                       |
| 事务   | 不支持                                                 | 支持                                       |
| 行表锁 | 表锁；即使操作一条记录也会锁住整张表，不适合高并发操作 | 行锁；操作时只会锁住某一行，适合高并发操作 |
| 缓存   | 只缓存索引，不缓存真实数据，内存要求小                 | 索引和真实数据都可以缓存，内存要求高       |
| 表空间 | 小                                                     | 大                                         |
| 关注点 | 性能                                                   | 事务                                       |



## MySQL常用命令

| 功能                   | 命令                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 进入MySQL命令行        | mysql -u 用户名 -p                                           |
| 查看MySQL的版本        | mysql --version<br />mysql -V<br />SELECT VERSION();         |
| 查看所有的数据库       | SHOW DATABASES;                                              |
| 打开指定的数据库       | USE tables;                                                  |
| 查看当前数据库所有的表 | SHOW TABLES;                                                 |
| 查看其他库所有的表     | SHOW TABLES FROM 库名;                                       |
| 查看表结构             | DESC 表名;                                                   |
| 执行 sql 脚本          | SOURCE ./path/xxx.sql; -- 路径是相对打开MySql时所处的路径    |
| 清空屏幕               | Linux：clear<br />Windows：无法，可以exit -> cls -> 重新进入 |
| 退出MySql终端          | exit                                                         |