---
author: "kuonz"
draft: false
title: "Redis启动"
date: 2020-04-05
categories: ["Redis启动与配置"]
---
  
## Redis服务端

| 操作         | 说明                                      |
| ------------ | ----------------------------------------- |
| 最简启动     | ${redis}/src/redis-server                 |
| 配置文件启动 | ${redis}/src/redis-server [配置文件]      |
| 动态参数启动 | ${redis}/src/redis-server --port [端口]   |
| 检查是否启动 | ps -ef \| grep redis                      |
| 关闭服务端   | ${redis}/src/redis-cli -p [端口] shutdown |



## Redis客户端

| 操作     | 说明                                           |
| -------- | ---------------------------------------------- |
| 连接     | ${redis}/src/redis-cli -h [主机地址] -p [端口] |
| 断开连接 | quit 或 exit                                   |

