---
author: "kuonz"
draft: true
title: "time"
date: 2020-04-05
categories: ["Go标准库"]
---

```go
now := time.Now() // 当前时间对象

now.Date()
now.Year()
now.Month()
now.Day()
now.Hour()
now.Minute()
now.Second()
```

```go
// 时间戳
now := time.Now()
timestamp1 = now.Unix() // 毫秒时间戳
timestamp2 = now.UnixNano() // 纳秒时间戳
```

```go
// 使用 time.Unix将时间戳转为 Time 类型对象
now := time.Now()
timeStamp = now.UnixNano()
before := time.Unix(timeStamp, 0)
```

```go
// 时间间隔
time.Nanosecond
time.Microsecond
time.Millisecond
time.Second
time.Minute
time.Hour
```

```go
// 时间操作
fmt.Println(now.Add(Time.Hour * 24))
fmt.Println(now.Sub(Time.Hour * 24))
fmt.Println(now.Equal(Time.Hour * 24))
fmt.Println(now.After(Time.Hour * 24)) // bool
fmt.Println(now.Before(Time.Hour * 24)) // bool
```

```go
// 定时器：本质上是一个channel
timer := time.Tick(time.Second)
for t := range timer {
  fmt.Println(t)
}
```

```go
// 时间格式化
// 公式：2006 01 02 15 04 05 000
// 口诀：2016 1 2 3 4 5 000 2006年1月2号下午3点4分5秒000毫秒
查考博客
// 如果需要使用12小时制，则要将 15换为 03， 且加上 PM
```

```go
time.Parse
time.LoadLocation("Asia/Shanghai")
time.ParseInLocation
time.Sleep()
```

time.Tick

time.After

time.Since