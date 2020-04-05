---
author: "kuonz"
draft: false
title: "flag"
date: 2020-04-05
categories: ["Go标准库"]
---
  
## flag包的作用

`flag`包能够解析命令行的参数，`flag`包使得开发命令行工具更为简单



## flag包支持解析的数据类型

| 支持数据类型 | 合法值说明                                 |
| ------------ | ------------------------------------------ |
| bool         | true/t/1TRUE/True 和 false/f/0/FALSE/False |
| int          | 整数，也可以是八进制，十六进制             |
| int64        | 整数，也可以是八进制，十六进制             |
| uint         | 正整数，也可以是八进制，十六进制           |
| uint64       | 正整数，也可以是八进制，十六进制           |
| float        | 合法的浮点数                               |
| float64      | 合法的浮点数                               |
| string       | 字符串，不用加双引号                       |
| duration     | 合法单位有：ns/us/ms/s/m/h                 |



## flag.[Type]

`flag.[Type]` 能够指定`flag`名，默认值和提示信息，并返回获取到的值的指针类型

### 基本格式

```go
flag.[Type](flag名, 默认值, 帮助信息) *[Type]
```

### 示例代码

```go
name := flag.String("name", "张三", "姓名")
age := flag.Int("age", 18, "年龄")
female := flag.Bool("female", false, "是否是女性")
delay := flag.Duration("duration", 0, "时间间隔")

flag.Parse() // 之后解释作用

// 注意：name, age, female, delay 都是指针类型
fmt.Printf("%T %T %T %T\n", name, age, female, delay) 
// *string *int *bool *time.Duration
```



## flag.[Type]Var

`flag.[Type]Var` 能够指定`flag`名，默认值和提示信息，并把获取到的赋值赋值到指定的变量中

### 基本格式

```go
flag.[Type]Var(变量地址, flag名, 默认值, 帮助信息) *[Type]
```

### 示例代码

```go
var name string
var age int
var female bool
var delay time.Duration

flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&female, "female", false, "是否是女性")
flag.DurationVar(&delay, "duration", 0, "时间间隔")

flag.Parse() // 之后解释作用

fmt.Println(name, age, female, delay)
```



## flag.Parse

通过以上两种方法定义好命令行`flag`参数后，需要通过调用`flag.Parse()`来对命令行参数进行解析

支持的命令行参数格式有以下几种：

| 格式                                  | 例子                       |
| ------------------------------------- | -------------------------- |
| -flag xxx【一个 `-`，使用空格分隔】   | program_name -name golang  |
| -flag=xxx【一个 `-`，使用 `=` 分隔】  | program_name -name=golang  |
| --flag xxx【两个 `-`，使用空格分隔】  | program_name --name golang |
| --flag=xxx【两个 `-`，使用 `=` 分隔】 | program_name --name=golang |

`flag`解析在第一个`非flag`参数之前停止，或者在终止符”–“之后停止



## flag包其他函数

| 函数         | 说明                                                         |
| ------------ | ------------------------------------------------------------ |
| flag.NFlag() | 返回使用的命令行参数个数                                     |
| flag.NArg()  | 返回命令行参数后的其他参数个数                               |
| flag.Args()  | 返回命令行参数后的其他参数，以[]string类型<br />与`os.Args`区别：第一个元素不是程序名称 |

### 示例代码

```go
func main() {
  var name string
  var age int
  var female bool
  var delay time.Duration
  
  flag.StringVar(&name, "name", "golang", "姓名")
  flag.IntVar(&age, "age", 10, "年龄")
  flag.BoolVar(&female "female", false, "是否是女性")
  flag.DurationVar(&delay, "duration", 0, "延迟的时间间隔")

  flag.Parse() //解析命令行参数
  
  fmt.Println(name, age, female, delay)
  
  fmt.Println(flag.NArg()) //返回命令行参数后的其他参数个数
  fmt.Println(flag.Args()) //返回命令行参数后的其他参数
  fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
```

### 使用

```shell
$ ./program_name -name golang -age=10 --female false --duration=1h30m a b c
```

### 输出

```shell
golang 10 false 1h30m
4
[a b c]
3
```



## 查看帮助信息

在使用`type.[Type]`和`flag.[Type]Var`时会指定帮助信息

通过`-h` 或者 `--h` 或者  `--help`或者 `-help` 就能查看帮助信息

### 使用

```shell
$ program_name -h
$ program_name --h
$ program_name -help
$ program_name --help
```

### 输出

```shell
Usage of C:\Users\idx00\Desktop\main.exe:
  -age int
        年龄 (default 18)
  -duration duration
        时间间隔
  -female
        是否是女性
  -name string
        姓名 (default "张三")
```

### 注意

如果想要使用帮助信息，则必须进行`flag.Parse`操作