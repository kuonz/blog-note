---
author: "kuonz"
draft: false
title: "Gin请求响应"
date: 2020-04-05
categories: ["Gin"]
---
  
## 响应模板数据

### 概述

`Gin` 框架的模板引擎使用的是 `Go` 标准库中的 `html/template`，`html/template` 的具体使用方法参考文章

### 流程

1. 定义模板文件
2. 解析模板文件
3. 渲染模板内容

### 解析模板文件

通过 `gin.Context` 的 `LoadHTMLFiles` 和 `LoadHTMLGlob` 来解析模板文件

方法签名

```go
LoadHTMLFiles(files ...string)
LoadHTMLGlob(pattern string)
```

示例代码

```go
engine.LoadHTMLFiles("./template/index.tmpl", "./template/login.tmpl")
engine.LoadHTMLGlob("./template/*.tmpl")
```

### 渲染模板内容

通过 `gin.Context` 的 `HTML` 方法来渲染模板内容

方法签名

```go
HTML(code int, name string, obj interface{})
```

示例代码

```go
engine.HTML(http.StatusOK, "index.tmpl", data)
```

### 自定义模板函数

通过 `gin.Context` 的 `SetFuncMap` 方法来自定义模板函数

方法签名

```go
SetFuncMap(funcMap template.FuncMap)
```

示例代码

```go
engine.SetFuncMap(template.FuncMap{
  "safe": func(str string) template.HTML {
    return template.HTML(str)
  }
})
```



## 响应字符串数据

通过 `gin.Context` 的 `String` 方法来返回字符串数据

### 方法签名

```go
String(code int, format string, values ...interface{})
```

### 示例代码

```go
engine.GET("/string", func (ctx *gin.Context) {
  ctx.String(200, "This is a string content")
})
```



## 响应JSON数据

通过 `gin.Context` 的 `JSON` 方法来返回 `JSON` 格式的数据

### 方法签名

```go
JSON(code int, obj interface{})
```

### 使用方式1：传递map

```go
engine.GET("/json", func (ctx *gin.Context) {
  ctx.JSON(http.StatusOK, map[string]interface{} {
    "status_code": 200,
    "message": "This is a JSON content"
  })
})
```

注意：为了简化写法，`gin.H` 基于 `map[string]interface{}` 自定义了类型，即

```go
type H map[string]interface{}
```

所以上述例子也可以写成

```go
engine.GET("/json", func (ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H {
    "status_code": 200,
    "message": "This is a JSON content"
  })
})
```

### 使用方式2：传递结构体

`Gin`序列化使用的库是`Go`官方的`encoding/json`，所以可以通过反射设置键名

```go
engine.GET("/json", func (ctx *gin.Context) {
  var msg = struct {
    Message string `json:"message"`
    StatusCode int `json:"status_code"`
  } {
    Message: "This is a JSON content",
    StatusCode: 200,
  }
  
  ctx.JSON(http.StatusOK, msg)
})
```

