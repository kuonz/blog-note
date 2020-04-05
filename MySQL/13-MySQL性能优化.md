---
author: "kuonz"
draft: false
title: "MySQL性能优化"
date: 2020-04-05
categories: ["MySQL"]
---
  
## 常用性能优化方式

| 方式            | 说明                                                         |
| --------------- | ------------------------------------------------------------ |
| 服务器硬件优化  | 加机器，加内存                                               |
| MySQL服务器优化 | 更改参数，增加缓冲等等                                       |
| SQL本身优化     | 减少子查询，减少连接查询的使用                               |
| 反范式设计优化  | 为了减少连接查询使用，可以允许适量数据冗余，使用空间换时间   |
| 物理设计优化    | 选择更好的数据类型：数值 > 时间日期 > 字符类型；同级别数据类型，优先选择占用空间少的数据类型<br />选择合适的存储引擎：MyiSAM和Memory的性能都比InnoDB要好 |
| 添加索引优化    |                                                              |



## SQL执行加载顺序

### 人写

```mysql
SELECT
FROM
JOIN ON
WHERE
GROUP BY
HAVING
ORDER BY
LIMIT
```

### 机读

```mysql
FROM
JOIN ON
WHERE
GROUP BY
HAVING
SELECT
ORDER BY
LIMIT
```

