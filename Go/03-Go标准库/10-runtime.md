---
author: "kuonz"
draft: true
title: "runtime"
date: 2020-04-05
categories: ["Go标准库"]
---

## Caller

用于获取执行时的堆栈信息

```go
pc, file, line, ok := runtime.Caller(层数) // 堆栈第n层，如果为0，则表示当前层
```



## FuncForPC



runtime.NumCPU()



runtime.Gosched()