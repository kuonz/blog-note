---
author: "kuonz"
draft: false
title: "AOF"
date: 2020-04-05
categories: ["Redis持久化"]
---
  
## AOF是什么

AOF：Append Only File

AOF是一种将内存中数据保存到本地硬盘中的持久化方式，AOF记录的不是数据本身，而是记录产生数据的命令

AOF以独立日志的方式记录每次写命令，重启时再重新执行AOF文件中命令达到恢复数据的目的。与RDB相比可以简单描述为改记录数据为记录数据产生的过程

AOF的主要作用是解决了数据持久化的实时性，目前已经是Redis持久化的主流方式



## AOF三种触发方式

| 触发方式 | 特点                           | 优点             | 缺点              |
| -------- | ------------------------------ | ---------------- | ----------------- |
| always   | 每条记录都会将缓冲区写入硬盘   | 基本不会丢失数据 | IO开销大          |
| everysec | 每秒都会将缓冲区写入硬盘       | 每秒1次          | 可能会丢失1秒数据 |
| no       | OS自己决定什么将缓冲区写入硬盘 | 不用管理         | 不能管理          |

过程图

```
           1.发送指令                           
client  ---------------> redis  
                           |
                           | 
                           | 2.将命令放入AOF缓存区
                           |
                           v      3.将缓存区命令同步到AOF文件 
                        AOF缓存区 -------------------------> .aof
```





## AOF配置

| 配置项                      | 可选值                | 说明                        |
| --------------------------- | --------------------- | --------------------------- |
| appendonly                  | [yes\|no]             | 是否开启AOF                 |
| appendfilename              | [filename.aof]        | AOF文件名                   |
| appendfsync                 | [always\|everysc\|no] | AOF触发方式【默认everysec】 |
| auto-aof-rewrite-min-size   | [大小]MB              | AOF文件重写需要的尺寸       |
| auto-aof-rewrite-percentage | 百分比                | AOF文件重写增长率           |



## AOF重写

### 功能

随着命令不断写入AOF，文件会越来越大，为了解决这个问题，Redis引入了AOF重写机制压缩文件体积

AOF文件重写是将Redis进程内的数据转化为写命令同步到新AOF文件的过程，简单说就是将对同一个数据的若干个条命令执行结果转化成最终结果数据对应的指令进行记录

### 作用

* 降低磁盘占用量，提高磁盘利用率
* 提高持久化效率，降低持久化写时间，提高IO性能
* 降低数据恢复用时，提高数据恢复效率

### 重写规则

* 进程内已超时的数据不再写入文件

* 忽略无效指令，重写时使用进程内数据直接生成，这样新的AOF文件只保留最终数据的写入命令

  如`del key1、hdel key2、srem key3、set key4 111、set key4 222`等

* 对同一数据的多条写命令合并为一条命令

  如`lpush list1 a、lpush list1 b、 lpush list1 c` 可以转化为：`lpush list1 a b c`

* 为防止数据量过大造成客户端缓冲区溢出，对list、set、hash、zset等类型，每条指令最多写入64个元素

### 使用

#### 过程图

```
        bgrewriteaof
client <-------------> redis
              ok         |
                         |
                         v       AOF重写
                    redis子进程 -----------> AOF重写文件
```

#### 手动重写

命令：`redis-cli> bgrewriteaof`

#### 自动重写

通过设置配置

```
auto-aof-rewrite-min-size size
auto-aof-rewrite-percentage percentage
```

### 自动重写规则

自动重写触发条件设置

```
auto-aof-rewrite-min-size size
auto-aof-rewrite-percentage percent
```

自动重写触发比对参数

```shell
aof_current_size
aof_base_size
```

自动重写触发条件

![](/03-AOF-images/image-20200324153803995.png)



### 重写流程

![](/03-AOF-images/image-20200324153238990.png)



## AOF和RDB对比

| 对比项     | RDB    | AOF      |
| ---------- | ------ | -------- |
| 启动优先级 | 低     | 高       |
| 体积       | 小     | 大       |
| 恢复速度   | 快     | 慢       |
| 数据安全性 | 丢数据 | 根据策略 |
| 轻重       | 重     | 轻       |

