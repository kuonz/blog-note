---
author: "kuonz"
draft: false
title: "MySQL变量"
date: 2020-04-05
categories: ["MySQL"]
---
  
## 变量分类

### 系统变量

全局变量：对于服务器所有的连接有效

会话变量：只在当前连接有效

### 自定义变量

用户变量：只在当前连接有效

局部变量：仅在 BEGIN-END 中有效



## 系统变量

查看所有的系统变量

```mysql
SHOW GLOBAL|SESSION VARIABLES;
```

查看某些的系统变量

```mysql
SHOW GLOBAL|SESSION VARIABLES LIKE '%char%';
```

查看指定的系统变量

```mysql
SELECT @@GOLBAL.系统变量名
SELECT @@SESSION.系统变量名
```

修改系统变量的值

```mysql
SET @@GLOBAL|SESSION.系统变量名 = 值
```



## 自定义变量

### 用户变量

声明并赋值

```mysql
SET @用户变量名 = 值
SET @用户变量名 := 值
SELECT @用户变量名 := 值
```

赋新值

```mysql
SET @用户变量名 = 值
SET @用户变量名 := 值
SELECT @用户变量名 := 值

SELECT 字段 INTO @变量名
FROM 表;
```

使用

```mysql
@变量名
SELECT @变量名
```

### 局部变量

声明并赋默认值

```mysql
BEGIN
  DECLARE 局部变量名 变量类型 DEFAULT 默认值;
END
```
