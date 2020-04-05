---
author: "kuonz"
draft: false
title: "MySQL触发器"
date: 2020-04-05
categories: ["MySQL"]
---
  
## 触发器概述

触发器是一类特殊的事务，可以监视某种DML操作，并触发相关DML操作



## 触发器使用

### 创建触发器

```mysql
DELIMITER $

CREATE TRIGGER 触发器名字
AFTER|BEFORE INSERT|UPDATE|DELETE ON 表名
FOR EACH ROW  # 固定写法
BEGIN
  一句或多句DML操作SQL语句
END;

$
```

### 使用旧值与新值

| 操作        | 说明                                   | 例子                     |
| ----------- | -------------------------------------- | ------------------------ |
| 对于 INSERT | 用new表示插入的新行                    | new.name，new.id         |
| 对于 UPDATE | 用new表示修改后的行，old表示修改前的行 | new.g_count，old.g_count |
| 对于 DELETE | 用old表示已经删除的行                  | old.name，old.id         |

### 删除触发器

```mysql
DROP TRIGGER 触发器名
```

### 查看触发器

查看所有的触发器

```mysql
SHOW TRIGGER STATUS;
```

查看某个触发器的创建语句

```mysql
SHOW CREATE TRIGGER 触发器名;
```

### 调用触发器

触发器是自动触发的，无法手动触发