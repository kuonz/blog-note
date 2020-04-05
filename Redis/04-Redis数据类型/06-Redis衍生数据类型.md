---
author: "kuonz"
draft: false
title: "Redis衍生数据类型"
date: 2020-04-05
categories: ["Redis"]
---

## bitmap

### bitmap是什么

bitmap译名为位图，其工作原理是用一个bit来标识一个数据的状态，状态只能存在两种，因为一个bit只有0或1两种值

### bitmap的作用

用于海量数据状态记录，比如

* 排序
* 去重
* 查询是否存在
* 记录状态

### bitmap本质/底层数据结构

bitmap的本质/底层数据结构是字符串

### bitmap的使用

#### 设置

设置某一位上的值，值只能是0或1，如果没设置默认补0

```shell
setbit key offset value
```

示例

```shell
setbit unique 0 1
setbit unique 5 1
setbit unique 11 1
setbit unique 15 1
setbit unique 20 1
```

#### 获取

获取某一位上的值

```shell
getbit key offset
```

示例

```shell
getbit unique 0
getbit unique 5
getbit unique 11
getbit unique 15
getbit unique 20
```

#### 长度

获取指定范围start，end字节中bit为1的个数，如果不指定start，end则为全部

```shell
bitcount key [start end]
bitcount key # 如果不指定start，end则为全部
```

示例

```shell
bitcount unique 1 2
```

#### 与或非异或

用于多个bitmap之间的与或非异或操作

```shell
bitop and|or|not|xor bitmap1 bitmap2
```

示例

```shell
bitop or unique1 unique2
bitop not unique1
```



## hyperloglog

### hyperloglog是什么

hyperloglog 是基于HyperLogLog算法实现的数据类型

### hyperloglog作用

hyperloglog用极小空间完成基数统计

基数统计：

| 数据集                | 基数集        | 基数 |
| --------------------- | ------------- | ---- |
| {1, 3, 5, 7, 5, 7, 8} | {1, 3, 5,7 8} | 5    |
| {1, 1, 1, 1, 2}       | {1, 2}        | 2    |

### hyperloglog使用

#### 添加

```shell
pfadd key element [element ...]
```

#### 统计数据

```shell
pfcount key [key ...]
```

#### 合并数据

```shell
pfmerge destkey sourcekey [sourcekey...]
```

### hyperloglog注意事项

1. 用于进行基数统计，不是集合，不保存数据，只记录数量而不是具体数据
2. 核心是基数估算算法，最终数值存在一定误差，误差范围：基数估计的结果是一个带有 0.81% 标准错误的近似值
3. 耗空间极小，每个hyperloglog key占用了12K的内存用于标记基数
4. pfadd命令不是一次性分配12K内存使用，会随着基数的增加内存逐渐增大
5. Pfmerge命令合并后占用的存储空间为12K，无论合并之前数据量多少



## geo

### geo是什么

geo是一种记录地理信息的经度和纬度，并提供地理相关操作API的数据类型

### geo作用

计算两地距离，范围等等

### geo使用

#### 添加坐标点

```shell
geoadd key longitude latitude member [longitude latitude member ...]
```

#### 获取坐标点

```shell
geopos key member [member ...]
```

#### 计算坐标距离

```shell
geodist key member1 member2 [unit]
```

### geo本质

geo类型本质使用zset实现

删除 API：zrem key 地方名