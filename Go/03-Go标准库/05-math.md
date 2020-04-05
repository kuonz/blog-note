---
author: "kuonz"
draft: true
title: "math"
date: 2020-04-05
categories: ["Go标准库"]
---

### 最大值与最小值

最大整数

```go
math.MaxInt8
math.MaxInt16
math.MaxInt32
math.MaxInt64
==============
math.MaxUint8
math.MaxUint16
math.MaxUint32
math.MaxUint64
```

最小整数

```go
math.MinInt8
math.MinInt16
math.MinInt32
math.MinInt64
```

最大浮点数

```go
math.MaxFloat32
math.MaxFloat64
```

最小浮点数

```go
math.MinFloat32
math.MinFloat64
```



### 随机rand

rand.seed(time.Now().UnixNano()) // 设置随机数种子

rand.Intn(100) // 0-99的随机整数 【0 <= x < 100】