---
author: "kuonz"
draft: false
title: "Go操作mysql"
date: 2020-04-05
categories: ["Go数据库编程"]
---
  
## 驱动与连接

`Go` 中的 `database/sql` 包提供了 `SQL` 数据库的泛用接口，并不提供具体的数据库驱动

使用 `database/sql` 包时必须注入数据库驱动。

### MySQL驱动

```go
import _ "github.com/go-sql-driver/mysql"
```

### 连接MySQL

使用 `sql.Open` 函数进行连接

`sql.Open` 返回的 `sql.DB`对象可以安全地被多个 `goroutine` 并发使用，并且维护自己的空闲连接池

因此，`sql.Open` 函数应该仅被调用一次，且很少需要关闭返回的 `DB` 对象

```go
package main

import (
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

func main() {
  db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
  if err != nil {
    panic(err)
  }
  defer db.Close()
}
```

### 测试连接

如果要检查数据源的名称是否真实有效，应该调用 `Ping` 方法

```go
package main

import (
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

func main() {
  db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
  if err != nil {
    panic(err)
  }
  defer db.Close()
  
  err = db.Ping() // 测试连接是否成功
  if err != nil {
    return err
  }
  return nil
}
```



## 最大连接数与最大空闲数

`database/sql` 内部维护有一个并发安全的连接池，所以可以设置其最大连接数和最大空闲数

### SetMaxOpenConns

```go
func (db *DB) SetMaxOpenConns(n int)
```

`SetMaxOpenConns`设置与数据库建立连接的最大数目

如果 `n` 小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制

如果 `n` 小于 0，不会限制最大开启连接数，默认为0（无限制）

```go
db.SetMaxOpenConns(10)
db,SetMaxOpenConns(-1)
```

### SetMaxIdleConns

```go
func (db *DB) SetMaxIdleConns(n int)
```

`SetMaxIdleConns`设置连接池中的最大闲置连接数

如果 `n` 大于最大开启连接数，则最大闲置连接数会减小到匹配最大开启连接数的限制

如果 `n` 小于 0，则不会保留闲置连接

```go
db.SetMaxIdleConns(10)
db.SetMaxIdleConns(-1)
```



## 增删改查

### 方法介绍

#### Exec

增删改都使用`Exec`方法

```go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
```

`Exec` 参数 `args` 表示 `exec` 中的占位参数

`Exec` 返回的 `Result` 是已执行的 `SQL` 命令的结果集

`Exec` 的结果集不会拥有数据库的连接，即不用释放

#### Query

查询操作使用`Query` 方法

```go
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```

`Query` 参数 `args` 表示 `query` 中的占位参数

`Query` 返回的 `Result` 是查询结果集

`Query` 的结果集会拥有数据库的连接，即用完需要释放

### 插入数据【增】

```go
func insertDemo() {
  sqlStr := "INSERT INTO user(name, age) VALUES (?,?)"
  
  ret, err := db.Exec(sqlStr, "王五", 38)
  
  if err != nil {
    fmt.Printf("insert failed, err:%v\n", err)
    return
  }
  
  theID, err := ret.LastInsertId() // 新插入数据的id
  
  if err != nil {
    fmt.Printf("get lastinsert ID failed, err:%v\n", err)
    return
  }
  
  fmt.Printf("insert success, the id is %d.\n", theID)
}
```

### 删除数据【删】

```go
func deleteDemo() {
  sqlStr := "delete from user where id = ?"
  
  ret, err := db.Exec(sqlStr, 3)
  
  if err != nil {
    fmt.Printf("delete failed, err:%v\n", err)
    return
  }
  
  n, err := ret.RowsAffected() // 操作影响的行数
  
  if err != nil {
    fmt.Printf("get RowsAffected failed, err:%v\n", err)
    return
  }
  
  fmt.Printf("delete success, affected rows:%d\n", n)
}
```

### 更新数据【改】

```go
func updateDemo() {
  sqlStr := "update user set age=? where id = ?"
  
  ret, err := db.Exec(sqlStr, 39, 3)
  
  if err != nil {
    fmt.Printf("update failed, err:%v\n", err)
    return
  }
  
  n, err := ret.RowsAffected() // 操作影响的行数
  
  if err != nil {
    fmt.Printf("get RowsAffected failed, err:%v\n", err)
    return
  }
  
  fmt.Printf("update success, affected rows:%d\n", n)
}
```

### 查询数据【查】

#### 单行数据查询

```go
func querySingleRowDemo() {
  sqlStr := "select id, name, age from user where id=?"
  
  var u user
  rows, err := db.QueryRow(sqlStr, 1)
  
  if err != nil {
    fmt.Printf("query failed, err:%v\n", err)
    return
  }
  
  defer rows.Close() // 非常重要：释放连接
  
  rows.Scan(&u.id, &u.name, &u.age) // 读取数据
  
  if err != nil {
    fmt.Printf("scan failed, err:%v\n", err)
    return
  }
  
  fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
```

#### 多行数据查询

```go
func queryMultiRowDemo() {
  sqlStr := "select id, name, age from user where id > ?"
  
  rows, err := db.Query(sqlStr, 0)
  
  if err != nil {
    fmt.Printf("query failed, err:%v\n", err)
    return
  }
  
  defer rows.Close() // 非常重要：释放连接

  // 循环读取结果集中的数据
  for rows.Next() {
    var u user
    err := rows.Scan(&u.id, &u.name, &u.age)
    if err != nil {
      fmt.Printf("scan failed, err:%v\n", err)
      return
    }
    fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
  }
}
```



## SQL语句预处理

### 预处理是什么

#### 普通SQL语句执行过程：

1. `Go` 程序对 `SQL` 语句进行占位符替换得到完整的 `SQL` 语句
2. `Go` 程序发送完整 `SQL` 语句到 `MySQL` 服务端
3. `MySQL` 服务端执行完整的 `SQL` 语句并将结果返回给 `Go` 程序

#### 预处理执行过程：

1. `Go` 程序把`SQL`语句分成两部分：命令部分与数据部分
2. `Go` 程序先把命令部分发送给 `MySQL` 服务端，`MySQL` 服务端进行 `SQL` 预处理（比如编译等操作）
3. `Go` 程序然后把数据部分发送给 `MySQL` 服务端，`MySQL` 服务端对 `SQL` 语句进行占位符替换
4. `MySQL` 服务端执行完整的 `SQL` 语句并将结果返回给 `Go` 程序

### 预处理的优点

1. 提高性能：避免 `MySQL` 反复对相同 `SQL` 语句进行编译，即一次编译多次运行
2. 避免 `SQL` 注入

### Go中使用预处理

#### 方法签名

```go
func (db *DB) Prepare(query string) (*Stmt, error)
```

#### stmt使用

`stmt` 对象具有 `Exec` 或 `Query` 方法，即使用与原来的增删改查方法一致

`stmt` 对象使用完毕后需要释放连接

#### 示例代码

```go
func prepareInsertDemo() {
  sqlStr := "insert into user(name, age) values (?,?)"
  
  stmt, err := db.Prepare(sqlStr) // 进行预处理
  
  if err != nil {
    fmt.Printf("prepare failed, err:%v\n", err)
    return
  }
  
  defer stmt.Close() // 释放连接
  
  _, err = stmt.Exec("zhangsan", 18)
  if err != nil {
    fmt.Printf("insert failed, err:%v\n", err)
    return
  }
  
  _, err = stmt.Exec("lisi", 20)
  if err != nil {
    fmt.Printf("insert failed, err:%v\n", err)
    return
  }
  
  fmt.Println("insert success.")
}
```

```go
func prepareQueryDemo() {
  sqlStr := "select id, name, age from user where id > ?"
  
  stmt, err := db.Prepare(sqlStr)
  if err != nil {
    fmt.Printf("prepare failed, err:%v\n", err)
    return
  }
  
  defer stmt.Close() // 释放连接
  
  rows, err := stmt.Query(0)
  if err != nil {
    fmt.Printf("query failed, err:%v\n", err)
    return
  }
  
  defer rows.Close() // 释放连接
  
  // 循环读取结果集中的数据
  for rows.Next() {
    var u user
    err := rows.Scan(&u.id, &u.name, &u.age)
    if err != nil {
      fmt.Printf("scan failed, err:%v\n", err)
      return
    }
    fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
  }
}
```



## 事务

`MySQL` 中只有使用了 `Innodb` 数据库引擎的数据库或表才支持事务

`Go`中使用下列三个方法进行事务操作

| 方法签名                            | 说明     |
| ----------------------------------- | -------- |
| func (db *DB) Begin() (\*Tx, error) | 开始事务 |
| func (tx *Tx) Commit() error        | 提交事务 |
| func (tx *Tx) Rollback() error      | 回滚事务 |

### 示例代码

```go
func transactionDemo() {
  tx, err := db.Begin() // 开启事务
  
  if err != nil {
    if tx != nil {
      tx.Rollback() // 事务回滚
    }
    fmt.Printf("begin trans failed, err:%v\n", err)
    return
  }
  
  sqlStr1 := "Update user set age = 30 where id = ?"
  
  _, err = tx.Exec(sqlStr1, 2)
  
  if err != nil {
    tx.Rollback() // 事务回滚
    fmt.Printf("exec sql1 failed, err:%v\n", err)
    return
  }
  
  sqlStr2 := "Update user set age = 40 where id = ?"
  
  _, err = tx.Exec(sqlStr2, 4)
  
  if err != nil {
    tx.Rollback() // 事务回滚
    fmt.Printf("exec sql2 failed, err:%v\n", err)
    return
  }
  
  err = tx.Commit() // 事务提交
  
  if err != nil {
    tx.Rollback() // 回滚事务
    fmt.Printf("commit failed, err:%v\n", err)
    return
  }
  
  fmt.Println("exec trans success!")
}
```



## 总结

### 需要释放连接的东西

1. db对象
2. stmt预处理对象
3. 查询结果集对象