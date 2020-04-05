---
author: "kuonz"
draft: false
title: "Gin请求参数获取"
date: 2020-04-05
categories: ["Gin"]
---
  
## HTTP请求常用参数形式

| 参数形式            | 说明                                             |
| ------------------- | ------------------------------------------------ |
| QueryString附带参数 | 在请求地址的QueryString中定义参数                |
| Path附带参数        | 在请求地址的Path中定义参数                       |
| Form参数            | 在网页表单中填写数据，通过提交表单的形式发送参数 |
| JSON形式参数        | 在HTTP请求体中定义JSON格式数据                   |



## QueryString参数

### QueryString是什么

`QueryString` 指的是请求地址中的 `?` 之后的字符串，参数是键值对形式，多个键值对之间以 `&` 进行分隔

比如：对于 `https://google.com/query?name=alice&age=20&gender=female` 这个请求地址的 `QueryString` 是`name=alice&age=20&gender=female`，其参数有`name=alice`，`age=20` 和 `gender=female`

### Gin获取QueryString参数

通过 `gin.Context` 的 `Query` 系列方法来获取 `QueryString` 中的参数

| Query系列方法                                        | 说明                                                         |
| ---------------------------------------------------- | ------------------------------------------------------------ |
| Query(key string) string                             | 获取QueryString参数的值                                      |
| DefaultQuery(key string, defaultValue string) string | 如果QueryString参数不存在，则返回设置的默认值                |
| GetQuery(key string) (string, bool)                  | 如果QueryString参数不存在，则返回值1为类型零值，返回值2为false |
| QueryArray(key string) []string                      | 获取QueryString中数组形式的参数                              |
| GetQueryArray(key string) ([]string, bool)           | 获取QueryString中数数组形式的参数，如果不存在则返回值1为类型零值，返回值2为false |
| QueryMap(key string) map[string]string               | 获取QueryString中数映射形式的参数                            |
| GetQueryMap(key string) (map[string]string, bool)    | 获取QueryString中数映射形式的参数，如果不存在则返回值1为类型零值，返回值2为false |
| BindQuery                                            | 将QueryString参数的值映射到结构体中，类似 ShouldBind         |
| ShouldBindQuery                                      | 将QueryString参数的值映射到结构体中，类似 ShouldBind         |

### 示例代码

```go
engine.GET("/login", func (ctx *gin.Context) {
  name := ctx.Query("name")
  age := ctx.Query("age")
  gender := ctx.Query("gender")
})

engine.GET("/login", func (ctx *gin.Context) {
  name := ctx.DefaultQuery("name", "Bob")
  age := ctx.DefaultQuery("age", 18)
  gender := ctx.DefaultQuery("gender", "male")
})

engine.GET("/login", func (ctx *gin.Context) {
  name, ok := ctx.GetQuery("name")
  age, ok := ctx.GetQuery("age")
  gender, ok := ctx.GetQuery("gender")
})
```



## Path参数

### 概述

`Path`参数是指 `URL` 中的参数

### Gin指定Path参数

在指定路由时，使用`:参数名`的方式来定义`Path`参数，比如

```go
engine.GET("/topics/:name/:num", func (ctx *gin.Context){}) 
```

### Gin获取Path参数

通过 `gin.Context` 的 `Param` 方法来获取 `Path` 中的参数

方法签名

```go
func (ctx *Context) Param(key string) string
```

### 示例代码

```go
engine.GET("/topic/:name/:num", func (ctx *gin.Context) {
  name := ctx.Param("name")
  age := ctx.Param("age")
  ctx.JSON(http.StatusOK, gin.H{
    "name": name,
    "age": age,
  })
})
```

### 注意事项

在定义`Path`参数时注意可能发生的冲突，比如

```go
engine.GET("/topics/:name/:num", func (ctx *gin.Context){}) 
// 与
engine.GET("/topics/:id/:star", func (ctx *gin.Context){}) 
// 是会发生冲突的
```

为了避免这种冲突，需要前缀不一样，比如

```go
engine.GET("/topics/:name/:num", func (ctx *gin.Context){}) 
// 与
engine.GET("/users/:id/:star", func (ctx *gin.Context){}) 
// 就不会发生冲突
```



## Form参数

### 概述

`Form` 参数是指由表单进行提交的数据

`HTTP`请求的 `Content-Type` 为 `application/x-www-form-urlencoded`

### Gin获取Form参数

通过 `gin.Context` 的 `PostForm` 系列方法来获取 `Form` 中的参数

| PostForm系列方法                                     | 说明                                                         |
| ---------------------------------------------------- | ------------------------------------------------------------ |
| PostForm(key string) string                          | 获取Form参数的值                                             |
| DefaultPostForm(key string, defaultValue string)     | 如果Form参数不存在，则返回设置的默认值                       |
| GetPostForm(key string) (string, bool)               | 如果Form参数不存在，则返回值1为类型零值，返回值2为false      |
| PostFormArraykey string) []string                    | 获取Form中数组形式的参数                                     |
| GetPostFormArray(key string) ([]string, bool)        | 获取Form中数组形式的参数，如果不存在则返回值1为类型零值，返回值2为false |
| PostFormMap(key string) map[string]string            | 获取Form中映射形式的参数                                     |
| GetPostFormMap(key string) (map[string]string, bool) | 获取Form中映射形式的参数，如果不存在则返回值1为类型零值，返回值2为false |

### 示例代码

```go
engine.POST("/login", func (ctx *gin.Context) {
  username := ctx.PostForm("username")
  password := ctx.PostForm("password")
})

engine.POST("/login", func (ctx *gin.Context) {
  username := ctx.DefaultPostForm("username", "emptyUsername")
  password := ctx.DefaultPostForm("password", "emptyPassword")
})

engine.POST("/login", func (ctx *gin.Context) {
  username, ok := ctx.GetPostForm("username")
  password, ok := ctx.GetPostForm("password")
})
```



## JSON参数

### 概述

`JSON` 参数是指 `HTTP` 请求体中的数据是 `JSON` 格式的

`HTTP` 请求的 `Content-Type` 为 `application/json`

### Gin获取JSON参数

通过 `gin.Context` 的 `BindJSON` 方法来获取 `JSON` 参数

方法签名

```go
func (c *Context) BindJSON(obj interface{}) error
```

使用该方法前需要定义用于接收 `JSON` 参数的结构体对象

### 示例代码

假设请求参数如下

```json
{
  "article_id": 1002,
  "article_title": "This is the title of the article",
  "article_content": "This is the content of the article"
}
```

定义接收 `JSON` 参数的结构体

```go
type ArticleParams struct {
  ArticleId int `json:"article_id"`
  ArticleTitle string `json:"article_title"`
  ArticleContent string `json:"article_content"`
}
```

使用 `BindJSON` 方法接收参数

```go
engine.POST("/article/update", func (ctx *gin.Context) {
  var articleParams ArticleParams // 创建接收参数的结构体对象
  ctx.BindJSON(&articleParams) // 必须传递指针，因为要改变该对象本身的值
})
```



## 参数自动绑定

### 概述

`Gin` 提供了一个的方法 `ShouldBind` 来自动获取参数，该方法会自动识别参数的类型，并将参数值赋值到指定结构体中

### 方法签名

```go
func (c *Context) ShouldBind(obj interface{}) error
```

### 结构体要求

在使用 `ShouldBind` 时，需要指定用于接收参数的结构体对象，该结构体的字段必须使用 `tag` 注明它所对应的键名，其中 `tag` 的名字为 `form`

```go
type ArticleParams struct {
  ArticleId int `form:"article_id"` // 绑定键名为 article_id 的参数
  ArticleTitle string `form:"article_title"` // 绑定键名为 article_title 的参数
  ArticleContent string `form:"article_content"` // 绑定键名为 article_content 的参数
}
```

### 示例代码

```go
engine.POST("/article/update", func (ctx *gin.Context) {
  var articleParams ArticleParams // 创建接收参数的结构体对象
  ctx.ShouldBind(&articleParams) // 必须传递指针，因为要改变该对象本身的值
})
```



## 自定义参数获取方法

### 概述

`HTTP` 参数传递的方式常用的是4种：QueryString，Path，Form，JSON，但也有可能出现其他形式的参数传递方式

`HTTP` 请求可以分为请求头和请求体，只要获取了请求头和请求体就可以自定义参数的获取方法了

### 请求头和请求体获取方法

通过 `gin.Context` 的 `Request` 字段可以获取请求头和请求体

| API                | 说明     |
| ------------------ | -------- |
| ctx.Request.Method | 请求方法 |
| ctx.Request.Header | 请求头   |
| ctx.Request.Body   | 请求体   |

通过读取分析请求头和请求体的内容就可以自定义任意参数类型的获取了