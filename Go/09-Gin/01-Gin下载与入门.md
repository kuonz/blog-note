---
author: "kuonz"
draft: false
title: "Gin下载与入门"
date: 2020-04-05
categories: ["Gin"]
---
  
## 介绍

`Gin` 是一个 `Go` 的 `Web` 框架

### 特点

`API` 友好：`API` 设计非常简洁和精妙，容易入门和使用

性能良好：路由底层使用了 `httprouter`，性能很好

支持中间件：有非常良好的自定义性和扩展性



## 下载

### 使用`GOPATH`方式

运行 `go get` 命令

```shell
$ go get -u github.com/gin-gonic/gin
```

### 使用`GOMODULE`方式

在源代码中写入导入语句

```go
import "github.com/gin-gonic/gin"
```

运行 `go mod tidy` 命令

```shell
$ go mod tidy
```



## 入门

`Gin` 框架版本的`Hello World`

```go
package main

import "github.com/gin-gonic/gin" // 导入Gin框架

func main() {
  engine := gin.Default() // 创建Gin使用默认配置的引擎

  engine.GET("/helloworld", func(ctx *gin.Context) { // 设置路由以及路由处理函数
    ctx.String(200, "Hello World")
  })

  engine.Run(":8080") // 启动引擎（启动服务器）
}
```

访问`127.0.0.1:8080/helloworld`即可看到浏览器显示`Hello World`