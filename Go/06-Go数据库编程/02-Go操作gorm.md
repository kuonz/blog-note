---
author: "kuonz"
draft: false
title: "Go操作gorm"
date: 2020-04-05
categories: ["Go数据库编程"]
---
  
## ORM是什么

### ORM全称及含义

**O**bject **R**elational **M**apping：对象关系映射

| O      | R          | M       |
| ------ | ---------- | ------- |
| Object | Relational | Mapping |
| 对象   | 关系       | 映射    |

### ORM作用

**ORM**本质上程序中的实例与数据库表之间的映射关系

数据表==结构体

数据表的行==结构体实例

数据表的列==结构体字段

### ORM优点缺点

优点

* 提高开发效率
* 能够轻易切换对应数据库而不用修改代码
* 更安全，不会发生SQL注入的情况

缺点

* 执行性能较差
* 灵活性较差
* 弱化开发者的SQL能力



## GORM

