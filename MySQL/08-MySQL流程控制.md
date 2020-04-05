---
author: "kuonz"
draft: false
title: "MySQL流程控制"
date: 2020-04-05
categories: ["MySQL"]
---
  
## 分支控制

### 分支控制分类

IF函数，IF语句，CASE语句

### IF函数

* IF(表达式1,表达式2,表达式3)
* 如果表达式1成立，则返回表达式2执行结果，否则返回执行表达式3执行结果，效果与三目运算符类似

### IF语句

```mysql
IF 条件
  THEN 语句1;
  ELSEIF 语句2;
  ELSE 语句3;
END IF;
```

### CASE语句【唯一的奇葩：一个分号都不带】

类似于 switch 语句

```mysql
# 都不用加分号

CASE 变量|表达式|字段
  WHEN 要判断的值 THEN 返回的值1
  WHEN 要判断的值 THEN 返回的值2
  ...
  WHEN 要判断的值 THEN 返回的值n
  ELSE 要返回的值n+1
END CASE
```

类似多重if语句

```mysql
# 都不用加分号

CASE
  WHEN 要判断的条件1 THEN 返回的值1
  WHEN 要判断的条件2 THEN 返回的值2
  ...
  WHEN 要判断的条件n THEN 返回的值n
  ELSE 要返回的值n+1
END
```

注意：ELSE可以省略，如果省略了且匹配不到，则返回NULL



## 循环控制

### 使用前提

只能在函数或存储过程中使用

### 分类

WHILE，LOOP，REPEAT

循环控制语句
* 必须配合标签使用
* ITERATE：类似continue
* LEAVE：类似break

### WHILE

```mysql
【标签:】WHILE 循环条件 DO

  循环体【每句都要带分号】
    
END WHILE【标签】;
```

### LOOP

```mysql
【标签:】 LOOP

  循环体【每句都要带分号】
    
END LOOP 【标签】;
```

### REPEAT

```mysql
【标签:】REPEAT

  循环体【每句都要带分号】
    
UNTIL 结束循环条件 END REPEAT【标签】;
```

