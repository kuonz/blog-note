---
author: "kuonz"
draft: false
title: "fmt"
date: 2020-04-05
categories: ["Go标准库"]
---
  
## 前置知识

### 占位符

| 占位符 | 含义                                             |
| ------ | ------------------------------------------------ |
| %T     | 查看类型                                         |
| %v     | 查看值（万能的）                                 |
| %+v    | 查看值，但输出结构体时，会添加字段名             |
| %#v    | 查看值，会带上更详细信息                         |
| %c     | 字符                                             |
| %s     | 字符串                                           |
| %q     | 会使用双引号括起来，必要时会采用安全的转义表示   |
| %b     | 2进制，包括整数和浮点数                          |
| %o     | 8进制                                            |
| %d     | 10进制                                           |
| %x     | 16进制（其中字母都是小写）                       |
| %X     | 16进制（其中字母都是大写）                       |
| %p     | 打印指针，会加上前导0x                           |
| %e     | 科学技术法，如-123.456e+78                       |
| %E     | 科学技术法，如-123.456E+78                       |
| %f     | 有小数部分但无指令部分，如123.456                |
| %F     | 等价于%f                                         |
| %g     | 根据实际情况采用%e或%f格式（更加简洁，更加清晰） |
| %G     | 根据实际情况采用%E或%F格式（更加简洁，更加清晰） |
| %n.m   | 宽度为n，小数保留m位                             |
| %+     | 总是输出正负号                                   |
| %-     | 在右边填空白而不是默认的左边填空白               |
| %0     | 使用0而不是空格来填充                            |

### 转义字符

| 转义字符 | 含义                             |
| -------- | -------------------------------- |
| \n       | 换行（直接跳到下一行的同列位置） |
| \t       | 制表符                           |
| \r       | 回车符（返回行首）               |
| \\'      | 一个单引号                       |
| \\"      | 一个双引号                       |
| \\\\     | 一个反斜杠                       |
| %%       | 一个百分号                       |



## Print系列

### Print

打印，不自动换行

```go
fmt.Print("Hello")
fmt.Print("Hello\n")
```

可以传入多个参数，打印时会用`  ,`分隔隔

```go
fmt.Print("Hello", "World")
```

### Println

打印，会自动换行

```go
fmt.Println("Hello")
```

可以传入多个参数，打印时会用`  ,`分隔隔

```go
fmt.Print("Hello", "World")
```

### Printf

可以使用占位符对字符串进行格式化操作


```go
var a1 int = 123
var a2 = struct {
  Name string
  Age int
} {
  Name: "golang",
  Age: 10,
}
var a3 rune = '中'
var a4 string ="golang"
var a5 int = 100
var a6 int = 0777
var a7 int = 0xFF6666
var a8 float64 = 3.1415926

fmt.Printf("%T", a1)
fmt.Printf("%v", a2)
fmt.Printf("%+v", a2)
fmt.Printf("%#v", a2)
fmt.Printf("%c", a3)
fmt.Printf("%s", a4)
fmt.Printf("%b", a5)
fmt.Printf("%o", a6)
fmt.Printf("%d", a5)
fmt.Printf("%x", a7)
fmt.Printf("%X", a7)
fmt.Printf("%b", a8)
fmt.Printf("%e", a8)
fmt.Printf("%E", a8)
fmt.Printf("%f", a8)
fmt.Printf("%F", a8)
fmt.Printf("%g", a8)
fmt.Printf("%G", a8)
```



## Scan系列

### Scan

功能：从标准输入中扫描文本，读取由空白符（换行符也视为空白符）分隔的值，并保存到本函数的参数中

返回值：读取中的数据个数和遇到的任何错误

停止的标志：读取到和参数个数相同的值

```go
var s string
fmt.Scanf(&s)
fmt.Println(s)
```

### Scanf

可以使用占位符的`Scan`语句

停止的标志：按照占位符语句读取

```go
var (
  name string
  age int
  class string
)

fmt.Scanf("%s %d %s\n", &name, &age, &class) // 需要手动加换行符
fmt.Println(name, age, class)
```

### Scanln

停止的标志：扫描到换行结束

```go
var (
  name string
  age int
  class string
)

fmt.Scanf("%s %d %s", &name, &age, &class) // 不需要手动加换行符
fmt.Println(name, age, class)
```



## 其他常用方法

### Sprintf

使用占位符完成字符串拼接

返回值是`string`类型

```go
s1 := "Hello"
s2 := "World"

s3 := fmt.Sprintf("%s %s", s1, s2)
fmt.Println(s3) // Hello World
```

### Errorf

使用占位符创建新的`error`类型变量

返回值是`error`类型

```go
msg := "这是我自定义的error"
code := 1000

err := fmt.Errorf("msg: %s，code: %d", msg, code)

fmt.Println(err.Error()) // msg: 这是我自定义的error，code: 1000
fmt.Printf("type: %T\n", err) // type: *errors.errorString
```