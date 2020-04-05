---
author: "kuonz"
draft: false
title: "atomic"
date: 2020-04-05
categories: ["Go标准库"]
---
  
### 使用场景

由于`加锁/解锁`操作涉及到操作系统`用户态`和`内核态`的转换操作，耗费资源大

所以`Go`提供了`atomic`包，该包封装了基础数据类型的原子操作，即能够在不加锁和解锁（不切换到内核态，仅仅在用户态）的情况下，保证基础数据类型的并发安全



### 读取操作

```go
func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
func LoadUint32(addr *uint32) (val uint32)
func LoadUint64(addr *uint64) (val uint64)
func LoadUintptr(addr *uintptr) (val uintptr)
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
```



### 写入操作

```go
func StoreInt32(addr *int32, val int32)
func StoreInt64(addr *int64, val int64)
func StoreUint32(addr *uint32, val uint32)
func StoreUint64(addr *uint64, val uint64)
func StoreUintptr(addr *uintptr, val uintptr)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
```

### 修改操作

```go
func AddInt32(addr *int32, delta int32) (new int32)
func AddInt64(addr *int64, delta int64) (new int64)
func AddUint32(addr *uint32, delta uint32) (new uint32)
func AddUint64(addr *uint64, delta uint64) (new uint64)
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
```



### 交换操作

```go
func SwapInt32(addr *int32, new int32) (old int32)
func SwapInt64(addr *int64, new int64) (old int64)
func SwapUint32(addr *uint32, new uint32) (old uint32)
func SwapUint64(addr *uint64, new uint64) (old uint64)
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
```



### 比较并交换操作

```go
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
```



### 使用例子

```go
package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

var globalVariable int32 // 全局变量，用于模拟共享内存
var wg sync.WaitGroup

func fn() {
  for i := 0; i < 10000; i++ {
    atomic.AddInt32(&globalVariable, 1) // 使用原子操作
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