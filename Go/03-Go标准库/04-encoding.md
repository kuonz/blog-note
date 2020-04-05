---
author: "kuonz"
draft: false
title: "encoding"
date: 2020-04-05
categories: ["Go标准库"]
---
  
## JSON/序列化/反序列化

### JSON概念以及作用

`JSON`全称 `JavaScript Object Notation`，是一种轻量级的数据交换格式，易于阅读和书写，也方便机器进行生成和解析

在数据传输前，传输的数据会经过处理变为`JSON`形式的字符串后再通过网络传输，接收方接收到 `JSON` 字符串后会进行处理，将 `JSON`字符串转为原始的数据

* 序列化：原始数据 -> `JSON`字符串

* 反序列化：`JSON`字符串 -> 原始数据

### 序列化所需的包以及函数

```go
import "encoding/json"

func Marshal(v interface{}) ([]byte, error)
```

### 序列化例子

```go
import "encoding/json"

func main() {
  var m map[string]interface{}
  m["name"] = "golang"
  m["age"] = 10
  m["address"] = "google"
    
  data,err := json.Marshal(m)
  if err != nil {
    fmt.Println("序列化失败", err)
  } else {
    fmt.Println("序列化后：", string(data))
  }
}
```

### 结构体序列化

结构体序列化时，非导出（以小写字母开头）的字段不会被序列化，因为非导出字段在`encoding/json`包中无法访问，所以要进行序列化的字段必须是导出的（以大写字母开头）

可以使用`tag`自定义序列化后`key`值的名称来实现序列化后键名是小写字母开头的

```go
// 反引号是必须的
type Monster struct {
  Name string `json:"name"`
  Age int     `json:"age"`
  BirthDay string `json:"birthday"`
  Sal float64 `json:"sal"`
  Skill string `json:"skill"`
}
```

### 反序列化所需的包以及函数

```go
import "encoding/json"

func Unmarshal(s []byte, v interface{}) (err error) 
```

### 反序列化例子

```go
package main

import "encoding/json"
import "fmt"

func main() {
  str := `{"address":"address","age":10,"name":"golang"}`

  var m map[string]interface{}

  err := json.Unmarshal([]byte(str), &m)

  if err != nil {
    fmt.Println("序列化失败", err)
  } else {
    fmt.Println("序列化后：", m)
  }
}
```



