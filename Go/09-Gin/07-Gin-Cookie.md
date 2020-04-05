---
author: "kuonz"
draft: false
title: "Gin-Cookie"
date: 2020-04-05
categories: ["Gin"]
---
  
## 设置Cookie

通过 `gin.Context` 的 `SetCookie` 方法设置 `Cookie`

### 方法签名

```go
SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
```

### 示例代码

```go
ctx.SetCookie("gin_cookie", "gin_cookie_value", 3600, "/", "localhost", false, true)
```



## 获取Cookie

通过 `gin.Context` 的 `Cookie` 方法获取 `Cookie`

### 方法签名

```go
Cookie(name string) (string, error)
```

### 示例代码

```go
cookie, err := ctx.Cookie("gin_cookie")
```

