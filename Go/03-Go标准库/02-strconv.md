---
author: "kuonz"
draft: true
title: "strconv"
date: 2020-04-05
categories: ["Go标准库"]
---

> strconv用于实现字符串与其他数据类型转换的包

> 看博客

### 基础数据类型转为string

```go
strconv.Itoa(i)
```

也可以使用`fmt.Sprintf`

```go
i := int32(97)
ret := fmt.Sprintf("%d", i)
```

### string转为基础数据类型

```go
b, err := strconv.ParseBool(str)
i, err := strconv.ParseInt(str, 10, 64) // 10进制，int64
f, err := strconv.ParseFloat(str, 64)   // float64

strconv.Itoa

如果转换失败，则返回为默认值，还有err
```

```go
i := int32(2500)
ret := string(i)
fmt.Println(ret) // 此时 ret 是 UTF-8编码中2500对应的符号，而不是字符粗 "2500"
```

