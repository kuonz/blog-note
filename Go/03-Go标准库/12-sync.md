---
author: "kuonz"
draft: false
title: "sync"
date: 2020-04-05
categories: ["Go标准库"]
---
  
## sync.WaitGroup

### 功能

`sync.WaitGroup` 用于实现 `goroutine` 任务的同步

### 原理

`sync.WaitGroup` 内维护了一个计数器，用来记录当前`goroutine`执行的任务的数量

### 方法

`sync.WaitGroup` 有三个 `API`，见下表

| API            | 功能                                   |
| -------------- | -------------------------------------- |
| Add(delta int) | 令内部计数器加 delta                   |
| Done()         | 令内部计数器减 1                       |
| Wait()         | 阻塞当前`goroutine`，直到内部计数器为0 |

### 示例代码

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

var wg sync.WaitGroup

func main() {
  wg.Add(2) // 令wg内部计数器加2，表示有两个正在执行任务的goroutine

  // 启动一个goroutine执行job1
  go func() {
    time.Sleep(2 * time.Second)
    fmt.Println("This is job1")
    wg.Done() // 令wg内部计数器减1，表示有一个任务已经完成
  }()

  // 启动一个goroutine执行job2
  go func() {
    time.Sleep(5 * time.Second)
    fmt.Println("This is job2")
    wg.Done() // 令wg内部计数器减1，表示有一个任务已经完成
  }()

  fmt.Println("等待其他任务完成...")
  wg.Wait() // 等待wg内部计数器清零，即等待所有goroutine任务完成
}
```



## sync.Once

### 使用场景

在高并发情况下，某些操作只能做一次或者只要做一次，为了避免多个`goroutine`执行多次这些操作，则需要使用`sync.Once`

### 实现方案

通过 `sync.Once` 来实现

```go
// 创建 sync.Once 对象
var once sync.Once
// 实现只执行一次的效果
once.Do(func() {})
```

### 示例代码

当有很多个goroutine使用了同一个channel，那么如果要关闭这个channel，则只能关闭一次，如果关闭多次，就会发生`panic`

```go
package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

var once sync.Once
var wg sync.WaitGroup

func fn(id int, ch chan int) {
  randNum := rand.Intn(10) + 1 // 获取[1,10]的随机数

  time.Sleep(time.Duration(randNum) * time.Second) // 随机睡1-10秒

  fmt.Printf("goroutine %d 睡眠了 %d 秒\n", id, randNum)
  
  // 睡完后关闭ch
  once.Do(func() { // 使用 sync.Once 避免被多次关闭
    fmt.Printf("goroutine %d 关闭了channel\n", id)
    close(ch)
  })

  wg.Done() // 标记当前 goroutine 完成
}

func main() {
  var ch chan int = make(chan int, 10)

  wg.Add(2)
  go fn(1, ch)
  go fn(2, ch)

  fmt.Println("等待全部 goroutine 完成")
  wg.Wait() // 等待全部 goroutine 完成
  fmt.Println("全部 goroutine 执行完毕")
}

```

#### 实现并发安全的单例模式

```go
package singleton

import "sync"

var once sync.Once

type singleton struct {}
var instance *singleton

func GetInstance() *singleton {
  once.Do(func() {
    instance = &singleton{}
  })
  return instance
}
```



## sync.Mutex

### 使用场景

多个`goroutine`之间的通信除了可以通过管道`channel`实现之外，还可以使用`共享内存`来实现

多个`goroutine`在对共享内存进行操作时，可能会发生数据竞争

### 竞争例子

```go
package main

import (
  "fmt"
  "sync"
)

var globalVariable int // 全局变量，用于模拟共享内存
var wg sync.WaitGroup

func fn() {
  for i := 0; i < 10000; i++ {
    globalVariable = globalVariable + 1
  }
  wg.Done()
}

func main() {
  wg.Add(2)
  go fn()
  go fn()
  wg.Wait()
  fmt.Println(globalVariable) // 结果并不是预期的 20000
}
```

在上述例子中，`globalVariable`的值并不是预期的`20000`，这是因为`globalVariable = globalVariable + 1` 这句代码并不是原子操作，比如当`globalVariable`为10时，两个`goroutine` 同时执行到`globalVariable = globalVariable + 1` 这句代码，即同时让`globalVariable`的值为11，即本来应该`+2`的操作被合并为`+1`了，这就发生了数据竞争的问题

### 解决方案

在执行对共享之间的操作之前，需要实现除了当前`goroutine`能够操作该共享资源外，其他的`goroutine`都无法操作该共享资源的效果，实现这个效果的操作称为`加锁`，在完成对共享资源的操作后需要把锁撤离，这个操作称为`解锁`

### 实现方案

`Go`中实现这种加锁，解锁操作的类型称为互斥锁，由`sync.Mutex`实现

```go
// 创建sync.Mutex对象
var mutexLock sync.Mutex
// 加锁操作
mutexLock.Lock()
// 解锁操作
mutexLock.Unlock()
```

### 解决竞争例子

```go
package main

import (
  "fmt"
  "sync"
)

var globalVariable int // 全局变量，用于模拟共享内存
var wg sync.WaitGroup
var mutexLock sync.Mutex // 创建互斥锁对象

func fn() {
  for i := 0; i < 10000; i++ {
    mutexLock.Lock() // 对共享资源进行操作前加锁
    globalVariable = globalVariable + 1
    mutexLock.Unlock() // 对共享资源操作完后进行解锁
  }
  wg.Done()
}

func main() {
  wg.Add(2)
  go fn()
  go fn()
  wg.Wait()
  fmt.Println(globalVariable) // 结果是预期的 20000
}
```



## sync.RWMutex

### 使用场景

互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源而不涉及资源的修改时是没有必要加锁的，这种场景下使用读写锁的性能会更好

### 实现方案

`Go`中读写锁由 `sync.RWMutex `实现

```go
// 创建读写锁对象
var rwmutexLock sync.RWMutex
// 加写锁
rwmutexLock.Lock()
// 解写锁
rwmutexLock.Unlock()
// 加读锁
rwmutexLock.RLock()
// 解读锁
rwmutexLock.RUnlock()
```

### 读写锁分类

读写锁具体分为两种锁：读锁和写锁

* 当一个`goroutine`获取了读锁
  * 其他的`goroutine`要获取读锁：能够不阻塞地获得读锁
  * 其他的`goroutine`要获取写锁：阻塞，直到全部读锁释放，然后获得写锁

* 当一个`goroutine`获取了写锁
  * 其他的`goroutine`无论是获取读锁还是写锁都会等待该写锁释放

#### 使用互斥锁耗时

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

var globalVariable int // 模拟共享内存
var mutexLock sync.Mutex
var wg sync.WaitGroup

func read() { // 模拟读操作
  mutexLock.Lock()
  time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
  mutexLock.Unlock()

  wg.Done()
}

func write() { // 模拟写操作
  mutexLock.Lock()
  globalVariable = globalVariable + 1
  time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
  mutexLock.Unlock()

  wg.Done()
}

func main() {
  start := time.Now()

  // 10次写操作
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go write()
  }

  // 1000次读操作，即读多写少场景
  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go read()
  }

  wg.Wait()
  end := time.Now()
  fmt.Printf("使用互斥锁耗时：%v", end.Sub(start))
}
```

上述例子结果为 `1.8370537s`

#### 使用读写锁耗时

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

var globalVariable int // 模拟共享内存
var mutexLock sync.RWMutex
var wg sync.WaitGroup

func read() { // 模拟读操作
  mutexLock.RLock()
  time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
  mutexLock.RUnlock()

  wg.Done()
}

func write() { // 模拟写操作
  mutexLock.Lock()
  globalVariable = globalVariable + 1
  time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
  mutexLock.Unlock()

  wg.Done()
}

func main() {
  start := time.Now()

  // 10次写操作
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go write()
  }

  // 1000次读操作，即读多写少场景
  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go read()
  }

  wg.Wait()
  end := time.Now()
  fmt.Printf("使用互斥锁耗时：%v", end.Sub(start))
}
```

上述例子结果为 `111.6956ms`



## sync.Map

### 使用场景

`Go`内置的`map`类型不是并发安全，如果想要使用并发安全的`map`，则可以考虑进行加锁，但加锁/解锁涉及到用户态和内核态的切换，太过耗费资源，所以`Go`提供了并发安全的映射类型`sync.Map`

### 实现方案

`sync.Map`是一个结构体类型，并不用像内置的`map`一样需要使用`make`分配内存，而是可以直接使用

由于`Go`中不存在泛型，所以`sync.Map`的键和值的类型都是空接口`interface{}`类型

```go
// 创建 sync.Map 对象
var m sync.Map
// 增 改
m.Store(key, value)
// 删
m.Delete(key)
// 查
m.Load(key)
// 遍历
m.Range( func(key, value interface{}) bool )
```

### 实例代码

```go
package main

import (
  "fmt"
  "sync"
)

func main() {
  var m sync.Map

  fmt.Println("======增======")
  m.Store("one", 1)
  m.Store("two", 2)
  m.Store("three", 3)

  fmt.Println("\n======查======")
  v, _ := m.Load("three")

  fmt.Printf("%v\n", v)

  fmt.Println("\n======改======")
  m.Store("three", 33)

  v, _ = m.Load("three")

  fmt.Printf("%v\n", v)

  fmt.Println("\n======遍历======")
  m.Range(func(key, value interface{}) bool {
    fmt.Printf("key: %v, value: %v\n", key, value)
    return true
  })

  fmt.Println("\n======删======")
  m.Delete("one")

  m.Range(func(key, value interface{}) bool {
    fmt.Printf("key: %v, value: %v\n", key, value)
    return true
  })
}
```



## sync.Pool