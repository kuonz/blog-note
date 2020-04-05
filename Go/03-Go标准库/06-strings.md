---
author: "kuonz"
draft: true
title: "strings"
date: 2020-04-05
categories: ["Go标准库"]
---

### Split

用于字符串切割，返回的`[]string`类型

```go
var s string = "1-2-3-4-5-6-7-8-9"
ret := strings.Split(s, "-")
fmt.Println(ret) // ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
```

### Join

用于拼接字符串

```go
var s []string = []string{"a", "b", "c"}
ret := strings.Join(s, "-")
fmt.Println(ret) // a-b-c
```

### Contains

是否包含子串

```go
var s = "Hello World"
fmt.Println(strings.Contains(s, "Hello")) // true
fmt.Println(strings.Contains(s, "Go"))    // false
```

### HasPrefix

是否前缀

```go
var s = "Hello World"
fmt.Println(strings.HasPrefix(s, "Hel")) // true
fmt.Println(strings.HasPrefix(s, "abc")) // false
```

### HasSuffix

是否后缀

```go
var s = "Hello World"
fmt.Println(strings.HasSuffix(s, "rld")) // true
fmt.Println(strings.HasSuffix(s, "abc")) // false
```

### Index

### LastIndex

### TrimSpace