---
author: "kuonz"
draft: true
title: "context"
date: 2020-04-05
categories: ["Go标准库"]
---
  
## 引例

问题：在多个`goroutine`中，如何在`goroutine A`中让`goroutine B` 结束？

### 方式1：通过指针

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

var wg sync.WaitGroup

// goroutine A 中执行的任务
func fnA(exit *bool) {
  defer wg.Done()
  
  for {
    select {
    case <-time.After(time.Duration(5) * time.Second): // 5秒后让 flag 为 false
      fmt.Println("[goroutine A] : exit 设置为 true， [goroutine B] 将要结束")
      *exit = true
      return
    }
  }
}

// goroutine B 中执行的任务
func fnB(exit *bool) {
  defer wg.Done()
  
  // 创建一个间隔为500ms的定时器
  t := time.Tick(time.Duration(500) * time.Millisecond)
  
  for {
    select {
      case <-t:
        fmt.Println("This is goroutine B...") // 每500ms输出一次
    }
    if (*exit == true) { // 如果 exit 为 true，则结束goroutine B
      fmt.Println("[goroutine B] : exit 为 true，[goroutine B] 结束")
      return
    }
  }
}

func main() {
  var exit *bool = new(bool) // 用于控制 goroutine 退出
  
  wg.Add(2)
  go fnA(exit)
  go fnB(exit)
  wg.Wait()
  
  fmt.Println("[main goroutine] : [goroutine A] 和 [goroutine B] 结束，本程序结束")
}
```

### 方法2：通过 channel

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

var wg sync.WaitGroup

// goroutine A 中执行的任务
func fnA(exitChan chan<- bool) {
  defer wg.Done()
  
  for {
    select {
    case <-time.After(time.Duration(5) * time.Second): // 5秒后让 flag 为 false
      fmt.Println("[goroutine A] : exit 设置为 true， [goroutine B] 将要结束")
      exitChan <- true
      return
    }
  }
}

// goroutine B 中执行的任务
func fnB(exitChan <-chan bool) {
  defer wg.Done()
  
  // 创建一个间隔为500ms的定时器
  t := time.Tick(time.Duration(500) * time.Millisecond)
  
  for {
    select {
      case <-t:
        fmt.Println("This is goroutine B...") // 每500ms输出一次
      case <-exitChan:
        fmt.Println("[goroutine B] : exit 为 true，[goroutine B] 结束")
        return
    }
  }
}

func main() {
  var exitChan chan bool = make(chan bool) // 用于控制 goroutine 退出
  
  wg.Add(2)
  go fnA(exitChan)
  go fnB(exitChan)
  wg.Wait()
  
  fmt.Println("[main goroutine] : [goroutine A] 和 [goroutine B] 结束，本程序结束")
}
```

### 上诉方法的缺点？

无论是通过`指针`还是通过`channel`，都可以实现在一个`goroutine`中控制其他`goroutine`的效果

问题是写法不统一，这不利于用户的使用和`Go`社区生态的发展

比如一个开源库A使用的是`指针`的方式，而另外一个开源库B使用的是`channel`的方式，那么用户使用库前都要去研究这个库是采用了哪种处理方式，而且一个项目中如果导入多个库，则可能发生同时存在两种写法的情况，这很不利于项目的编写和维护

解决方法：`Go`官方推出`context`这个专门用于控制管理`goroutine`的库，制定了规则，统一了写法



## context概述

`Go1.7`时加入了`context`库，这个库专门用于处理多个`goroutine`之间的操作，包括

* 取消信号
* 截至时间
* 超时时间
* 请求域数据 [key-value]

等等操作

这些操作成为了并发控制和超时控制的标准写法



## Context接口

### 接口作用

`Context `接口类型是 `context` 包的核心类型

如果想要使用 `context` 包中的方法，则需要使用到实现了 `Context` 接口的对象

### 接口定义

`context.Context` 是一个接口，定义了下列的方法

```go
type Context interface {
  Value(key interface{}) interface{}
  Done() <-chan struct{}
  Deadline() (deadline time.Time, ok bool) 
  Err() error
}
```

#### Deadline

返回`Context`完成工作的截止时间（deadline）和是否超过截至时间

#### Value

返回`Context`中键对应的值

对于同一个`Context`来说，多次调用`Value`并传入相同的`Key`会返回相同的结果

该方法仅用于传递跨API和进程间跟请求域的数据

#### Done

返回一个`channel struct{}`对象，这个`channel`会在当前工作完成后或者`Context`被取消时关闭

多次调用Done方法返回的是同一个`channel`对象

#### Err

返回当前`Context`结束的原因，它只会在`Done`返回的`channel`由于取消关闭时才会返回非空的值

* 如果当前`Context`被取消就会返回`Canceled`错误

* 如果当前`Context`超时取消就会返回`DeadlineExceeded`错误

### 接口类型使用注意事项

1. 推荐以参数的方式显示传递 `Context` 对象，而不是使用全局变量或者函数返回的方式
2. 以 `Context` 作为参数的函数方法，应该把 `Context` 作为第一个参数
3. 一个函数需要 `Context` 参数时，如果不知道传递什么，不要传递`nil`而是传递 `context.TODO()`
4. `Context`是线程安全的，可以放心的在多个`goroutine`中传递
5. `Context`的`Value`相关方法应该传递请求域的必要数据，不应该用于传递可选参数



## With系列函数

`context` 包定义了4个 `With` 系列函数

### WithCancel

#### 方法签名

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

#### 功能作用

`WithCancel`返回带有`Done`返回通道的父节点的副本

当调用返回的`cancel`函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel



### WithTimeout

#### 方法签名

```go

```



### WithDeadline

#### 方法签名

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

返回`parent context`的副本，并将`context`的`Deadline`调整为不迟于参数`d`

如果`parent context`的`deadline`已经早于参数`d`，则`WithDeadline(parent, d)`在等同于`parent context`。当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。



### WithValue

## TODO和Background

`context`包内置了两个函数`context.TODO()`和`context.Background()`

这两个方法都会返回一个实现了`Context`接口的`emptyCtx`结构类型

`emptyCtx` 是一个没有设置过期时间，没有携带任何值，不可取消的 `Context` 对象

### context.TODO

当目前还不知道需要怎样的`Context`类型时，可以使用`context.TODO()`来暂时占位

```go
ctx : = context.TODO()
```

### context.Background

一般作为`Context`这个树结构的最顶层的`Context`，也就是根`Context`

主要用于`main`函数、初始化以及测试代码中

```go
ctx := context.Background()
```



尽管ctx会自动过期，但在任何情况下主动调用它的cancel函数都是很好的实践

如果不这样做，可能会使上下文及其父类存活时间超过必要的时间

