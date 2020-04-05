---
author: "kuonz"
draft: false
title: "MySQL函数"
date: 2020-04-05
categories: ["MySQL"]
---
  
## 函数概述

### 概念

函数是一组预先编译好的SQL语句的集合

### 类比

存储过程类似于 Go 中的函数

### 优点

提高代码重用性，简化操作，减少编译次数



## 函数使用

### 创建函数

创建语法

```mysql
DELIMITER $

CREATE FUNCTION 函数名(参数列表) RETURNS 返回类型
BEGIN
  函数语句块
END;

$

DELIMITER ;
```

参数形式

```mysql
参数名 参数类型
```

注意事项

- 如果存储过程只有一句SQL，则可以省略BEGIN-END
- 存储过程中的SQL都必须以分号结尾，所以存储过程结尾可以用 DELIMITER 重新设置
- 必须有且只能有一个返回值，且必须有return语句
- 函数如果使用 SELECT，则必须配合INTO关键字，使用SELECT ... INTO ... 的结构，因为函数中不允许出现结果集
- 返回类型如果是varchar必须带长度

### 删除函数

```mysql
DROP FUNCTION myfunction;
```

### 查看有哪些函数

```mysql
SHOW FUNCTION STATUS;
```

### 查看函数创建

```mysql
SHOW CREATE FUNCTION myfunction;
```

### 调用函数

```mysql
SELECT 函数名(参数列表);
```



