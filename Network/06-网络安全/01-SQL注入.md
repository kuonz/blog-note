---
author: "kuonz"
draft: false
title: "SQL注入"
date: 2020-04-05
categories: ["网络安全"]
---
  
### SQL注入原理

`SQL`注入针对的是数据库和`SQL`拼接语句，如果建立 `SQL` 语句的方法不够安全，那么应用程序就很容易遭到 `SQL`注入攻击

### 常用SQL注入语句

```
' or 1='1
```

使用示例

```SQL
"select * from userinfo where password='"+ password +"'"
-- 如果传入的 password 是 ' or 1='1，则拼接后SQL语句为
select * from userinfo where password='' or 1='1'
-- 1='1' 永远为真，所以就验证通过了
```

### SQL注入危害

* 探知数据库
* 泄露数据，尤其是机密信息，账户信息
* 取的更高的权限

### SQL注入防御方法

`SQL` 注入的根本性的防御策略是使用参数化查询，也叫做预处理语句

参数化查询在建立一个包含用户输入的SQL语句时分为两步：

* 指定查询结构，用户输入预留占位符
* 指定占位符内容

使用参数化参数后，就算用户输入类似 `'or 1=1'`，也只会认为其是用户输入的内容，而不参数SQL语句的构建





