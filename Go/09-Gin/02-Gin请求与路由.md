---
author: "kuonz"
draft: false
title: "Gin请求与路由"
date: 2020-04-05
categories: ["Gin"]
---
  
## 请求方式

`Gin` 支持使用下列九种常见的 `HTTP` 请求方式

| 请求方式 | 说明                                                      |
| -------- | --------------------------------------------------------- |
| GET      | 请求指定页面的信息，并返回实体主体                        |
| POST     | 向指定资源提交数据进行处理请求，数据被包含在请求体中      |
| PUT      | 从客户端向服务器传送的数据取代指定的文档的内容            |
| PATCH    | 是对 PUT 方法的补充，用来对已知资源进行局部更新           |
| DELETE   | 请求服务器删除指定内容                                    |
| HEAD     | 类似GET请求，只不过返回的响应中没有具体内容，用于获取报头 |
| OPTIONS  | 允许客户端查看服务器的性能                                |
| CONNECT  | HTTP/1.1 协议中预留给能够将连接改为管道方式的代理服务器   |
| TRACE    | 回显服务器收到的请求，主要用于测试或诊断                  |



## 路由

`Gin` 定义路由时需要指定该路由的请求方式以及请求处理函数

### 定义形式

```go
engine.[HTTP请求方法]("路由地址", 请求处理函数)
```

### 请求处理函数

请求处理函数必须是 `gin.HandleFunc` 类型的代码，而 `gin.HandleFunc` 类型的定义如下

```go
type HandleFunc func(*Context)
```

### 示例代码

```go
engine.GET("/path", func(ctx *gin.Context){ /* 请求处理代码 */ })     // 路由：/path，请求方式：GET
engine.POST("/path", func(ctx *gin.Context){ /* 请求处理代码 */ })    // 路由：/path，请求方式：POST
engine.PUT("/path", func(ctx *gin.Context){ /* 请求处理代码 */ })     // 路由：/path，请求方式：PUT
engine.PATCH("/path", func(ctx *gin.Context){ /* 请求处理代码 */ })   // 路由：/path，请求方式：PATCH
engine.DELETE("/path", func(ctx *gin.Context){ /* 请求处理代码 */ })  // 路由：/path，请求方式：DELETE
engine.HEAD("/path", func(ctx *gin.Context){ /* 请求处理代码 */ })    // 路由：/path，请求方式：HEAD
engine.TRACE("/path", func(ctx *gin.Context){ /* 请求处理代码 */ })   // 路由：/path，请求方式：TRACE
engine.OPTIONS("/path", func(ctx *gin.Context){ /* 请求处理代码 */ }) // 路由：/path，请求方式：OPTIONS
engine.CONNECT("/path", func(ctx *gin.Context){ /* 请求处理代码 */ }) // 路由：/path，请求方式：CONNECT
```

### 匹配所有方法

一个路由如果想支持所有请求方式，则可以使用 `Any` 方法

```go
// 路由：/path
// 请求方式：GET POST PUT PATCH DELETE HEAD TRACE OPTIONS CONNECT
engine.Any("/path", func(ctx *gin.Context){ /* 请求处理代码 */ }) 
```

`Any` 方法可以配合 `swtich` 语句使用

```go
engine.Any("/path", func(ctx *gin.Context) {
  switch ctx.Request.Method {
    case "GET": /* 处理GET请求 */
    case "POST": /* 处理POST请求 */
    case "PUT": /* 处理PUT请求 */
    case "PATCH": /* 处理PATCH请求 */
    case "DELETE": /* 处理DELETE请求 */
    case "HEAD": /* 处理HEAD请求 */
    case "CONNECT": /* 处理CONNECT请求 */
    case "OPTIONS": /* 处理OPTIONS请求 */
    case "TRACE": /* 处理TRACE请求 */
  }
})
```



### 处理未定义路由

`Gin` 中可以给未定义的路由绑定一个统一的请求处理函数

使用 `Gin` 引擎的 `NoRoute` 方法进行绑定

```go
engine.NoRoute(func(ctx *gin.Context){
  /* 未定义路由处理代码 */
})
```

当用户访问没有在程序中定义的路由时，会执行**未定义路由处理代码**



## 路由组

### 使用需求

对于多个路由，可能它们的路由地址的前缀是相同的，比如下列路由定义，它们路由地址前缀都是 `/article`

```go
engine.GET("/article/add", func(ctx *gin.Context){
  // 文章添加处理代码
})

engine.GET("/article/delete", func(ctx *gin.Context){
  // 文章删除处理代码
})

engine.GET("/article/update", func(ctx *gin.Context){
  // 文章更新处理代码
})
```

为了便于编写和维护，可以使用路由组

### 路由组的定义

`Gin` 中路由组定义使用到 `Group` 方法，方法签名如下

```go
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
```

示例代码

```go
articleGroup := engine.Group("/article")
```

### 路由组的使用

每个路由组都可绑定不同的路由，绑定的路由地址会自动加上该路由组的前缀

```go
articleGroup := engine.Group("/article")

articleGroup.GET("/add", func(ctx *gin.Context) { // 实际路由地址为 /article/add
  // 文章添加处理代码
})

articleGroup.GET("/delete", func(ctx *gin.Context) { // 实际路由地址为 /article/delete
  // 文章删除处理代码
})

articleGroup.GET("/update", func(ctx *gin.Context) { // 实际路由地址为 /article/update
  // 文章更新处理代码
})
```

### 路由组的优雅写法

在使用路由组时，可以使用`{}`和缩进让代码更加清晰

```go
articleGroup := engine.Group("/article")
{
  articleGroup.GET("/add", func(ctx *gin.Context) { // 实际路由地址为 /article/add
    // 文章添加处理代码
  })

  articleGroup.GET("/delete", func(ctx *gin.Context) { // 实际路由地址为 /article/delete
    // 文章删除处理代码
  })

  articleGroup.GET("/update", func(ctx *gin.Context) { // 实际路由地址为 /article/update
    // 文章更新处理代码
  })
}
```

这种写法只是为了代码更加清晰，功能本质无改变，不强制使用该写法

### 路由组的嵌套

路由组是支持嵌套的

```go
blogGroup := engine.Group("/blog")
{
  v1Group := blogGroup.Group("/v1")
  v2Group := blogGroup.Group("/v2")
  
  v1Group.GET("/index") // 实际路由地址为 /blog/v1/index
  v1Group.GET("/article") // 实际路由地址为 /blog/v1/article
  
  v2Group.GET("/index") // 实际路由地址为 /blog/v2/index
  v2Group.GET("/article") // 实际路由地址为 /blog/v2/article
}
```



## 请求重定向和转发

### HTTP重定向

`HTTP`重定向分为：`301`永久重定向和`302`暂时重定向，区别在于搜索引擎会不会替换地址

`HTTP`重定向请求的地址并会发生改变

使用 `gin.Context` 的 `Redirect` 方法进行`HTTP` 重定向

```go
ctx.Redirect(http.StatusMovePermanently, "重定向地址") //  301永久重定向
ctx.Redirect(301, "重定向地址") //  301永久重定向

ctx.Redirect(http.StatusFound, "重定向地址") // 302暂时重定向
ctx.Redirect(302, "重定向地址") // 302暂时重定向
```

### 路由转发

路由重定向是在程序内部进行请求的转发，请求的地址并不会发生改变

先设置请求的 `URL` 再使用 `gin.Context` 的 `HandleContext` 进行路由重定向

```go
engine.GET("/old", func(ctx *gin.Context) {
  ctx.Request.URL.Path = "/new" // 指定转发的URL
  r.HandleContext(ctx) // 进行转发
})

engine.GET("/new", func(ctx *gin.Context) {
  ctx.String(http.StatusOK, "New Content")
})
```



## 路由原理

`Gin` 的路由系统使用的是 `httprouter` 库，[点击查看github仓库](https://github.com/julienschmidt/httprouter)

`httprouter` 底层数据结构使用的是前缀树