---
author: "kuonz"
draft: false
title: "Go操作redis"
date: 2020-04-05
categories: ["Go数据库编程"]
---
  
## 驱动

`Go` 的 `database/sql` 仅支持 `SQL` 语句和关系型数据库，而不支持 `Redis` 在内的 `NoSQL` 

所以`Go`使用 `Redis` 的方法可能因为驱动不同而不同

### redis驱动介绍

`Go` 中的 `redis` 驱动常用的有下列两种

* redigo
* go-redis

本篇文章采用 `go-redis` 作为驱动



## 连接

### 普通连接

```go
package main

import (
  "fmt"
  "github.com/go-redis/redis"
)

func main() {
  var rdb *redis.Client = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379", // redis服务器地址
    Password: "", // redis服务器密码
    DB:       0,  // 连接到的数据集
  })

  _, err = rdb.Ping().Result() // 测试连接是否成功
  
  if err != nil {
    panic(err)
  }
}
```

### 哨兵模式连接

```go
package main

import (
  "fmt"
  "github.com/go-redis/redis"
)

func main()(err error){
  rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "master", // master名称
    SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"}, // 集群地址列表
  })
  
  _, err = rdb.Ping().Result() // 测试连接是否成功
  
  if err != nil {
    panic(err)
  }
}
```

### 集群模式连接

```go
package main

import (
  "fmt"
  "github.com/go-redis/redis"
)

func main() {
  rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
  })
  
  _, err = rdb.Ping().Result() // 测试连接是否成功
  
  if err != nil {
    panic(err)
  }
}
```



## 使用

### set

```go
func redisSetExample() {
  
  err := rdb.Set("score", 100, 0).Err()
  
  if err != nil {
    fmt.Printf("set score failed, err:%v\n", err)
    return
  }
}
```

### get

```go
func redisGetExample() {
  val, err := rdb.Get("score").Result()
  
  if err != nil {
    fmt.Printf("get score failed, err:%v\n", err)
    return
  }
  
  fmt.Println("score", val)

  val2, err := rdb.Get("name").Result()
  
  if err == redis.Nil {
    fmt.Println("name does not exist")
  } else if err != nil {
    fmt.Printf("get name failed, err:%v\n", err)
    return
  } else {
    fmt.Println("name", val2)
  }
}
```

### zset

```go
func redisZSetExample() {
  
  zsetKey := "language_rank"
  
  languages := []*redis.Z{
    &redis.Z{Score: 90.0, Member: "Golang"},
    &redis.Z{Score: 98.0, Member: "Java"},
    &redis.Z{Score: 95.0, Member: "Python"},
    &redis.Z{Score: 97.0, Member: "JavaScript"},
    &redis.Z{Score: 99.0, Member: "C/C++"},
  }
  
  // ZADD
  num, err := rdb.ZAdd(zsetKey, languages...).Result()
  if err != nil {
    fmt.Printf("zadd failed, err:%v\n", err)
    return
  }
  fmt.Printf("zadd %d succ.\n", num)

  // 把Golang的分数加10
  newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
  if err != nil {
    fmt.Printf("zincrby failed, err:%v\n", err)
    return
  }
  fmt.Printf("Golang's score is %f now.\n", newScore)

  // 取分数最高的3个
  ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
  if err != nil {
    fmt.Printf("zrevrange failed, err:%v\n", err)
    return
  }
  for _, z := range ret {
    fmt.Println(z.Member, z.Score)
  }

  // 取95~100分的
  op := &redis.ZRangeBy{
    Min: "95",
    Max: "100",
  }
  ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
  if err != nil {
    fmt.Printf("zrangebyscore failed, err:%v\n", err)
    return
  }
  
  for _, z := range ret {
    fmt.Println(z.Member, z.Score)
  }
}
```

