---
author: "kuonz"
draft: false
title: "template"
date: 2020-04-05
categories: ["Go标准库"]
---
  
## 概述

`GO` 中的模板引擎分为两个包 `text/template` 和 `html/template`

| 模板包        | 说明                     |
| ------------- | ------------------------ |
| html/template | 专门用于生成 `HTML` 文档 |
| text/template | 可以生成其他格式的文档   |



## 模板引擎使用流程

1. 定义模板文件
2. 解析模板文件
3. 进行模板渲染



## 定义模板文件

### 文件要求

| 要求   | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| 后缀名 | 模板文件后缀名并没有硬性要求，但一般约定为 `.tmpl` 或 `.tpl` |
| 编码   | 必须是 `UTF-8` 编码                                          |

### 模板语法标识符

模板语法都需要包含在`{{`和`}}`中间

标识符可以通过`Template`对象的`Delims`来自定义符号，从而避免与别的框架发生冲突

```go
template.New("index.tmpl").Delims("{[","]}") // 此时不再是 {{}}，而是 {[]}
```

### 注释

在`{{ }}`中使用`/**/`

```go
{{ /* 注释 */ }}
```

注释是可以换行的

```go
{{/*
  这是
  多行
  注释
*/}}
```

### 数据

在 `{{ }}` 中用 `.`表示传递给模板的数据

```go
// 模板文件：index.tmpl
<h1> Hello {{ . }} </h1>

// 解析文件：index.go
t, _ := template.ParseFiles("index.tmpl")
t.Execute(w, "Golang")

// 渲染后文件：index.html
<h1>Hello Golang</h1> 
```

如果传入的是结构体对象，则可以使用 `{{ .字段名 }}` 来获取结构体字段的值

```go
// 模板文件：index.tmpl
<h1> Name: {{ .Name }} </h1>
<h1> Age: {{ .Age }} </h1>

// 解析文件：index.go
type Person struct {
  Name string
  Age int
}

p := Person{
  Name: "Alice",
  Age: 20,
}

t, _ := template.ParseFiles("index.tmpl")
t.Execute(w, p)

// 渲染后文件：index.html
<h1> Name: Alice </h1>
<h1> Age: 20 </h1>
```

如果传入的是映射对象，则可以使用 `{{ .键名 }}` 来获取键相对应的值

```go
// 模板文件：index.tmpl
<h1> Name: {{ .name }} </h1>
<h1> Country: {{ .country }} </h1>

// 解析文件：index.go
var m map[string]string = make(map[string]string)

m[name] = "张三"
m[country] = "中国"

t, _ := template.ParseFiles("index.tmpl")
t.Execute(w, m)

// 渲染后文件：index.html
<h1> Name: 张三 </h1>
<h1> Country: 中国 </h1>
```

### with

使用`{{ with 数据 }} {{end}}`创建局部作用域，类似 `JavaScript` 的 `with` 语句

```go
{{ with .m }}
{{ .name }}
{{ .age }}
{{ .country }}
{{ end }}

等价于

{{ .m.name }}
{{ .m.age }}
{{ .m.country }}
```

### 变量

在模板中可以声明变量 `$变量名`，从而保存传入到模板中的数据或其他语句生成的结果

```go
// 模板文件：index.tmpl
<h1> Name: {{ $name := .name }} </h1>
<h1> Name: {{ $name }} </h1>

// 解析文件：index.go
var m map[string]string = make(map[string]string)

m[name] = "张三"

t, _ := template.ParseFiles("index.tmpl")
t.Execute(w, m)

// 渲染后文件：index.html
<h1> Name: 张三 </h1>
<h1> Name: 张三 </h1>
```

### 比较函数

| 比较函数     | 说明                                  |
| ------------ | ------------------------------------- |
| lt arg1 arg2 | arg1 < arg2 返回 true，否则返回false  |
| gt arg1 arg2 | arg1 > arg2 返回 true，否则返回false  |
| eq arg1 arg2 | arg1 = arg2 返回 true，否则返回false  |
| nq arg1 arg2 | arg1 != arg2 返回 true，否则返回false |
| le arg1 arg2 | arg1 <= arg2 返回 true，否则返回false |
| ge arg1 arg2 | arg1 >= arg2 返回 true，否则返回false |

### 条件判断

使用 `{{if 条件}} {{else if 条件}} {{else}} {{end}}` 来实现条件判断

```go
{{if lt $age 22}}
小于22岁
{{else if lt $age 35}}
大于等于22岁，小于35岁
{{else}}
大于等于35岁
{{end}}
```

### 遍历

使用 `{{range $下标, $项 := .目标数据}} {{ else }} {{end}}` 来实现遍历

```go
{{ range $idx, $hobby := .hobby }}
<p>Index: $idx, Hobby: $hobby</p>
{{ else }}
hobby为空
{{ end }}
```

### 预定义函数

模板引擎中定义了一些常用的函数，称为预定义函数

| 预定义函数 | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| and        | "and x y"等价于"if x then y else x                           |
| or         | "or x y"等价于"if x then x else y"                           |
| not        | 返回它的单个参数的布尔值的否定                               |
| len        | 返回它的参数的整数类型长度                                   |
| index      | 执行结果为第一个参数以剩下的参数为索引/键指向的值            |
| print      | 即fmt.Sprint                                                 |
| printf     | 即fmt.Sprintf                                                |
| println    | 即fmt.Sprintln                                               |
| js         | 返回与其参数的文本表示形式等效的转义JavaScript               |
| call       | 执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数 |

### 自定义函数

通过 `Template `对象的 `Funcs` 函数来绑定自定义函数

```go
func sayHello(w http.ResponseWriter, r *http.Request) {
  // 自定义一个显示=当前时间的模板函数
  currentTime := func() string {
    return time.Now().UnixNano()
  }
  // 自定义一个显示参数的模板函数
  showArgument := func(arg string) string {
    return arg
  }
  
  // 创建模板对象
  t := template.New("index.tmpl")
  // 绑定自定义函数
  t.Funcs(template.FuncMap{
    "currentTime": currentTime,
    "showArgument": showArgument,
  })
  // 解析模板文件：名字需要对应上
  t.ParseFiles("index.tmpl")

  // 定义数据
  p := Person{
    Name:   "Alice",
    Age:    20,
  }
  
  // 使用 p 渲染模板，并将结果写入w
  t.Execute(w, p)
}
```

自定义函数在模板中的使用

```html
{{ currentTime }}
{{ showArgument .Name }}
```

### 模板嵌套

概念：复制别的模板内容到当前模板中

嵌套的模板分为两类：

1. 通过`define`定义模板

```html
{{ define "ol.tmpl"}}
<ol>
  <li>吃饭</li>
  <li>睡觉</li>
  <li>写代码</li>
</ol>
{{end}}
```

2. 单独的模板文件

```html
<!-- ul.tmpl -->
<ul>
  <li>注释</li>
  <li>日志</li>
  <li>测试</li>
</ul>
```

模板嵌套操作

```go
{{ define "ol.tmpl"}}
<ol>
  <li>吃饭</li>
  <li>睡觉</li>
  <li>写代码</li>
</ol>
{{end}}

<h1>测试嵌套template语法</h1>
<hr />
{{template "ul.tmpl"}} {{/* 嵌套单独的模板文件 */}}
<hr />
{{template "ol.tmpl"}} {{/* 嵌套define定义的模板 */}}
```

渲染后的`HTML`文件

```html
<h1>测试嵌套template语法</h1>
<hr />
<ul>
  <li>注释</li>
  <li>日志</li>
  <li>测试</li>
</ul>
<hr />
<ol>
  <li>吃饭</li>
  <li>睡觉</li>
  <li>写代码</li>
</ol>
```

### 模板继承

概念：给当前模板留出空间，可以自定义留出空间中显示的内容

`{{block "block_name" data}} {{end}}` 实现模板继承

父模板：

```html
<!-- base.tmpl -->
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <title>Go Templates</title>
</head>
<body>
<div class="container-fluid">
  {{block "content" . }}{{end}}
</div>
</body>
</html>
```

子模板：

```html
<!-- index.tmpl -->
{{template "base.tmpl"}} {{/* 通过模板嵌套把父模板复制过来 */}}

{{define "content"}} {{/* 自定义父模板中留出空间的内容 */}}
  <div>Hello world!</div>
  <div>{{ . }}</div>
{{end}}
```

解析、渲染过程

```go
t := template.New("index.tmpl")
t.ParseGlob("./*.tmpl")
t.ExecuteTemplate(w, "index.tmpl", "张三")
```

渲染后结果：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <title>Go Templates</title>
</head>
<body>
<div class="container-fluid">
  <div>Hello world!</div>
  <div>张三</div>
</div>
</body>
</html>
```

模板嵌套和模板继承的区别：

* 模板嵌套：复制别的模板内容到当前模板中
* 模板继承：给当前模板留出空间，可以自定义留出空间中显示的内容，通常需要配合模板嵌套使用

### 移除空格

在使用模板语法时，可能会引入一些空格或换行符，这会造成渲染后效果不一样的情况，所以需要去除空白字符

使用 `{{- -}}` 来去除空白字符

```go
{{- .Name -}}
```

### pipeline 管道操作

使用`|`进行管道操作（pipeline）

`|` 前的命令执行结果会作为参数传递给 `|` 后的命令，类似于 `Linux` 中的管道操作 



## 解析模板文件

定义好模板文件后，创建模板对象，模板对象解析模板文件或模板字符串

### New

使用场景：创建模板对象

函数签名：

```go
func New(name string) *Template
```

示例代码：

```go
t := template.New("template_name")
```

### Parse

使用场景：解析模板字符串

函数签名：

```go
func (t *Template)Parse(src string) (*Template, error)  // template 对象的方法
```

示例代码：

```go
src := "<h1>Hello World</h1>"
t := template.New("template_name")
t, err := t.Parse(src)
```

### ParseFiles

使用场景：解析模板文件

函数签名：

```go
func (t *Template) ParseFiles(filenames ...string) (*Template, errror) // Template 对象的方法
func ParseFiles(filenames ...string) (*Template, errror) // template 包的函数，语法糖函数，无需事先创建模板对象
```

注意事项：模板对象的名称必须出现在解析的模板文件名中

示例代码：

```go
// 写法1
t := template.New("index.tmpl") // 模板对象的名称必须出现在解析的模板文件名中，此处是 index.tmpl
t, err := t.ParseFiles("./index.tmpl", "./user.tmpl", "./login.tmpl")

// 写法2 - 语法糖写法，不用事先创建 Template 对象
t, err := template.ParseFiles("./index.tmpl", "./user.tmpl", "./login.tmpl")
```

### ParseGlob

使用场景：用于解析多个模板文件时，不写具体名字，可以写正则表达式进行匹配

函数签名：

```go
func (t *Template) ParseFiles(pattern string) (*Template, errror) // Template 对象的方法
func ParseFiles(pattern string) (*Template, errror) // template 包的函数，语法糖函数，无需事先创建模板对象
```

示例代码：

```go
// 写法1
t := template.New("index.tmpl") // 模板对象的名称必须出现在解析的模板文件名中，此处是 index.tmpl
t, err := t.ParseGlob("*.tmpl")

// 写法2 - 语法糖写法，不用事先创建 Template 对象
t, err := template.ParseGlob("*.tmpl")
```



## 进行模板渲染

解析模板文件，得到模板对象后，可以使用下列模板对象的方法进行模板的渲染

### Execute

使用场景：模板对象只包含一个模板文件的解析内容

函数签名：

```go
func (t *Template) Execute(w io.Writer, data interface{}) error
```

示例代码：

```go
func handleIndex(w http.ResponseWriter, r *http.Request) {
  err := t.Execute(w, "数据")
}
```

### ExecuteTemplate

使用场景：模板对象包含了多个模板文件的解析内容，使用时需要指定使用哪个模板文件的解析内容

函数签名：

```go
func (t *Template) ExecuteTemplate(w io.Writer, name string, data interface{}) error
```

示例代码：

```go
func handleIndex(w http.ResponseWriter, r *http.Request) {
  err := t.ExecuteTemplate(w, "index.tmpl", "首页数据")
}

func handleUesr(w http.ResponseWriter, r *http.Request) {
  err := t.ExecuteTemplate(w, "user.tmpl", "用户数据")
}
```



## text/template和html/template区别

两者的使用方法基本一致，区别在于`html/template`会对一些风险内容进行转义，以此防范跨站脚本攻击(xss)

跨站脚本攻击(xss)：有时网站需要接收用户提供的文本，并加入到网站HTML中（比如评论功能），那么用户就有可能提交一段可以执行恶意`JavaScript`的`HTML`文本，从而攻击网站，例子如下

模板文件：

```go
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Hello</title>
</head>
<body>
  用户的评论是：{{ . }}
</body>
</html>
```

用户提交的文本：

```html
<script>
  alert("这是恶意攻击")
</script>
```

如果此时将用户提交的文本直接加入到网站的`HTML`中，则被恶意攻击了：其他用户打开网站时都会执行`alert("这是恶意攻击")`这段`JavaScript`代码

解决方法：为了防范跨站脚本攻击，需要对一些敏感的`HTML`标签进行转义，即它转义后加入网站的`HTML`时，保证是安全的

而`html/template`就对这些敏感的`HTML`标签进行了转义，不需要程序员手动转义