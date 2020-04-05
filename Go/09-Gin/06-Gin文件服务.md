---
author: "kuonz"
draft: false
title: "Gin文件服务"
date: 2020-04-05
categories: ["Gin"]
---
  
## 单文件上传

### 概述

通过 `gin.Context` 的 `FormFile` 方法获取上传文件的内容，再通过 `SaveUploadedFile` 来保存文件内容

`HTTP` 的 `Content-Type` 为 `multipart/form-data`

### 方法签名

```go
FormFile(name string) (*multipart.FileHeader, error)
SaveUploadedFile(file *multipart.FileHeader, dst string)
```

### 示例代码

```go
engine.POST("/upload", func(ctx *gin.Context) {
  // 获取上传文件的内容
  file, err := ctx.FormFile("file")
  // 保存文件内容到本地
  dst := fmt.Sprintf("./tmp/%s", file.Filename)
  ctx.SaveUploadedFile(file, dst)
})
```



## 多文件上传

### 概述

通过 `gin.Context` 的 `MultipartForm` 方法获取 `Form` 对象，再通过 `Form` 对象获取文件列表，再遍历文件列表保存文件

### 方法签名

```go
MultipartForm() (*multipart.Form, error)
SaveUploadedFile(file *multipart.FileHeader, dst string)
```

### 示例代码

```go
engine.POST("/upload", func(ctx *gin.Context) {
  // 获取Form对象
  form, err := ctx.MultipartForm()
  // 通过Form对象获取文件列表
  files := form.File["file"]
  // 遍历文件列表
  for index, file := range files {
    dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
    // 保存文件内容到本地
    ctx.SaveUploadedFile(file, dst)
  }
})
```



## 文件下载

### 概述

文件下载服务等价于静态文件服务

使用 `gin.Context` 的 `Staic` 方法配置静态文件服务

### 方法签名

```go
Static(relativePath, root string)
```

relativePath：路由前缀

root：存放静态文件的文件夹路径

### 示例代码

```go
ctx.Static("/static", "./static")
```