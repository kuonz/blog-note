---
author: "kuonz"
draft: false
title: "Go-Regexp"
date: 2020-04-05
categories: ["Go正则表达式"]
---
  
## Regexp包

在 `Go` 中，如果想要使用正则表达式，则需要使用到内置的 `Regexp` 包



## 创建Regexp对象

| 方法签名                                   | 注意事项                                                     |
| ------------------------------------------ | ------------------------------------------------------------ |
| Compile(str string) (*Regexp, error)       | 采用最左最短方式搜索<br />如果传入的正则表达式错误，则会返回`error` |
| MustCompile(str string) *Regexp            | 采用最左最短方式搜索<br />如果传入的正则表达式解析错误，会引发`panic` |
| CompilePOSIX(str string)  (*Regexp, error) | 采用最左最长方式搜索<br />采用`POSIX` 语法<br />不支持：`\d、\D、\s、\S、\w、\W` |
| MustCompilePOSIX(str string) *Regexp       | 采用最左最长方式搜索<br />采用`POSIX` 语法<br />不支持：`\d、\D、\s、\S、\w、\W` |

### 示例代码

```go
package main

import (
  "regexp"
)

func main() {
  reg1, err := regexp.Compile(`[a-z]+`)
  reg2 := regexp.MustCompile(`[a-z]+`)
  reg3, err := regexp.CompilePOSIX(`[a-z]+`)
  reg4 := regexp.MustCompilePOSIX(`[a-z]+`)
}
```



## 验证字符串

| 方法签名                                                     | 说明                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Match(pattern string, b []byte) (matched bool, err error)    | 接收字节数组判断是否符合正则表达式<br />matched：符合为true，不符合为false<br />err：查找过程中遇到的任何错误 |
| MatchReader(pattern string, r io.RuneReader) (matched bool, err error) | 从Reader接收数据判断是否符合正则表达式<br />matched：符合为true，不符合为false<br />err：查找过程中遇到的任何错误 |
| MatchString(pattern string, s string) (matched bool, err error) | 接收字符串判断是否符合正则表达式<br />matched：符合为true，不符合为false<br />err：查找过程中遇到的任何错误 |

### 示例代码

```go
package main

import (
  "regexp"
)

func main() {  
  regexp.Match(`^[a-z]+$`, []byte("abcdef")) // true
  regexp.Match(`^[a-z]+$`, []byte("123456")) // false
  
  regexp.MatchString(`^[a-z]+$`, "abcdef") // true
  regexp.MatchString(`^[a-z]+$`, "123456") // false
  
  reader1 := bytes.NewReader([]byte("abcdef"))
  reader2 := bytes.NewReader([]byte("123456"))
  regexp.MatchReader(`^[a-z]+$`, reader1) // true
  regexp.MatchReader(`^[a-z]+$`, reader2) // false
}
```



## 替换字符串

| 方法签名                                                     | 说明                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| ReplaceAll(src, repl []byte) []byte                          | 在 `src` 中搜索匹配项，并替换为 `repl` 指定的内容            |
| ReplaceAllString(src, repl string) string                    | 在 `src` 中搜索匹配项，并替换为 `repl` 指定的内容            |
| ReplaceAllLiteral(src, repl []byte) []byte                   | 在 `src` 中搜索匹配项，并替换为 `repl` 指定的内容<br />如果 `repl` 中有分组引用符`$`（如`$1`、`$name`），则将`$` 当普通字符处理 |
| ReplaceAllLiteralString(src, repl string) string             | 在 `src` 中搜索匹配项，并替换为 `repl` 指定的内容<br />如果 `repl` 中有分组引用符`$`（如`$1`、`$name`），则将`$` 当普通字符处理 |
| ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte  | 在 `src` 中搜索匹配项，然后将匹配的内容经过 `repl` 处理后，替换 `src` 中的匹配项<br />如果 `repl` 中有分组引用符`$`（如`$1`、`$name`），则将`$` 当普通字符处理 |
| ReplaceAllStringFunc(src string, repl func(string) string) string | 在 `src` 中搜索匹配项，然后将匹配的内容经过 `repl` 处理后，替换 `src` 中的匹配项<br />如果 `repl` 中有分组引用符`$`（如`$1`、`$name`），则将`$` 当普通字符处理 |



## 提取字符串

使用`Regexp`对象的方法进行提取字符串

| 常用方法签名                                        | 说明                                 |
| --------------------------------------------------- | ------------------------------------ |
| Find(b []byte) []byte                               | 返回第一个匹配的内容                 |
| FindString(s string) string                         | 返回第一个匹配的内容                 |
| FindAll(b []byte, n int) \[][]byte                  | 返回所有匹配的内容                   |
| FindAllString(s string, n int) []string             | 返回所有匹配的内容                   |
| FindIndex(b []byte) []int                           | 返回第一个匹配的位置                 |
| FindStringIndex(s string) loc []int                 | 返回第一个匹配的位置                 |
| FindReaderIndex(r io.RuneReader) (loc []int)        | 返回第一个匹配的位置                 |
| FindAllIndex(b []byte, n int) [][]int               | 并返回所有匹配的位置[首下标，尾下标] |
| FindAllStringIndex(s string, n int) [][]int         | 并返回所有匹配的位置[首下标，尾下标] |
| FindSubmatch(b []byte) [][]byte                     |                                      |
| FindStringSubmatch(s string) []string               |                                      |
| FindAllSubmatch(b []byte, n int) [][][]byte         |                                      |
| FindAllStringSubmatch(s string, n int) [][]string   |                                      |
| FindSubmatchIndex(b []byte) []int                   |                                      |
| FindStringSubmatchIndex(s string) []int             |                                      |
| FindReaderSubmatchIndex(r io.RuneReader) []int      |                                      |
| FindAllSubmatchIndex(b []byte, n int) [][]int       |                                      |
| FindAllStringSubmatchIndex(s string, n int) [][]int |                                      |