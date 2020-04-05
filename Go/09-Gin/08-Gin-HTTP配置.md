---
author: "kuonz"
draft: false
title: "Gin-HTTP配置"
date: 2020-04-05
categories: ["Gin"]
---
  
### 概述

因为 `Gin` 使用的是`Go` 内置的 `http` 库，所以可以通过 `Go` 的 `net/http` 包的来设置 `HTTP` 配置

### 示例代码

```go
func main() {
  engine := gin.Default()

  s := &http.Server{
    Addr:           ":8080",
    Handler:        router,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  
  s.ListenAndServe() // 不要使用 engine.Run
}
```

