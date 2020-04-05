---
author: "kuonz"
draft: false
title: "Go读写csv文件"
date: 2020-04-05
categories: ["Go常用操作"]
---
  
## csv文件写入

类似于 `bufio` 的用法

```go
package main

import (
  "encoding/csv"
  "os"
)

func main() {
  file, err := os.Create("test.csv") // 创建文件
  // file, err := os.OpenFile("./test.csv", os.O_APPEND|os.O_CREATE, 0666) // 创建文件
  
  if err != nil { // 如果创建文件失败，则抛出异常
    panic(err)
  }
  
  defer file.Close() // 延迟关闭
  
  writer := csv.NewWriter(file)
  
  writer.Write([]string{"123", "456", "789"})
  
  writer.Flush() // 刷新缓冲区内容
}
```



## csv文件读取

类似于 `bufio` 的用法

`FieldsPerRecord` 用于规定每条记录的字段数，如果字段数不满足，则报错，如果设置为`-1`，则不限制

```go
package main

import (
  "encoding/csv"
  "os"
)

func main() {
  file, err := os.Open("./test.csv") // 打开文件
  
  if err != nil {
    panic(err) // 如果打错错误，则panic
  }
  
  defer file.Close() // 延时关闭文件
  
  reader := csv.NewReader(file)
  
  reader.FieldsPerRecord = -1 // 设置不限制每条记录的字段数目
  
  record, err := reader.ReadAll()
  
  if err != nil {
    panic(err)  
  }
  
  for _, item := range record {
    fmt.Printf("%s ", item[0])
    fmt.Printf("%s ", item[1])
    fmt.Printf("%s ", item[2])
  }
}
```





