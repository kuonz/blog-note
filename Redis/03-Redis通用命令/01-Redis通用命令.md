---
author: "kuonz"
draft: false
title: "Redis通用命令"
date: 2020-04-05
categories: ["Redis通用命令"]
---
  
## Key通用命令

| 命令                                       | 说明                        |
| ------------------------------------------ | --------------------------- |
| **del** *key*                              | 删除指定的key               |
| **exists** *key*                           | 判断key是否存在             |
| **type** *key*                             | 获取key的类型               |
| **expire** *key seconds*                   | 为key设置过期时间           |
| **pexpire** *key milliseconds*             | 为key设置过期时间           |
| **expireat** *key timestamp*               | 为key设置过期时间           |
| **pexpireat** *key milliseconds-timestamp* | 为key设置过期时间           |
| **ttl** *key*                              | 获取key的有效时间           |
| **pttl** *key*                             | 获取key的有效时间           |
| **persist** *key*                          | 设置key从时效性转换为永久性 |
| **rename** *key newkey*                    | 为key改名                   |
| **renamenx** *key newkey*                  | 为key改名                   |
| **sort** *key*                             | 对所有key排序               |
| **keys** *pattern*                         | 根据模式查询key             |

### pattern模式规则

| 规则 | 说明                   |
| ---- | ---------------------- |
| *    | 匹配任意数量的任意符号 |
| ?    | 配合一个任意符号       |
| []   | 匹配一个指定符号       |

| 示例               | 说明                                                |
| ------------------ | --------------------------------------------------- |
| **keys** *         | 查询所有                                            |
| **keys** re*       | 查询所有以re开头                                    |
| **keys** *dis      | 查询所有以dis结尾                                   |
| **keys** ??dis     | 查询所有前面两个字符任意，后面以dis结尾             |
| **keys** user:?    | 查询所有以user:开头，最后一个字符任意               |
| **keys** u[st]er:1 | 查询所有以u开头，以er:1结尾，中间包含一个字母，s或t |



## 数据库通用命令

| 命令                   | 说明                      |
| ---------------------- | ------------------------- |
| **select** *index*     | 切换数据库                |
| **quit**               | 退出服务器/客户端         |
| **ping**               | 检查客户端是否连接成功    |
| **echo** message       | 打印信息                  |
| **move** *key dbindex* | 数据移动                  |
| **dbsize**             | 数据库中键的数量          |
| **flushdb**            | 清空当前数据库的全部数据  |
| **flushall**           | 清空当前服务i去的全部数据 |

