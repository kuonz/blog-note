---
author: "kuonz"
draft: false
title: "Go-tcp编程"
date: 2020-04-05
categories: ["Go网络编程"]
---
  
## TCP服务端

### 流程

1. 监听套接字（IP+端口），等待客户端连接
2. 接受客户端连接
3. 创建`goroutine`处理客户端连接

### 示例代码

```go
func main() {
  // 监听套接字
  listen, err := net.Listen("tcp", "127.0.0.1:20000")
  
  if err != nil {
    fmt.Println("listen failed, err:", err)
    return
  }
  
  for { // 等待客户端连接
    conn, err := listen.Accept() // 接受连接
    
    if err != nil {
      fmt.Println("accept failed, err:", err)
      continue
    }
    
    go process(conn) // 启动一个goroutine处理连接
  }
}

// 处理连接的函数
func process(conn net.Conn) {
  
  defer conn.Close() // 延时关闭连接
  
  for { // 不断读取客户端的数据
    
    var buf [128]byte // 创建缓存区
    
    n, err := conn.Read(buf[:]) // 读取数据
    
    if err == io.EOF { // 判断是否读取完毕
      break
    }
    
    if err != nil {
      fmt.Println("read from client failed, err:", err)
      break
    }
    
    recvStr := string(buf[:n]) // 处理读取到的数据
    
    fmt.Println("收到client端发来的数据：", recvStr)
    
    conn.Write([]byte(recvStr)) // 发送数据
  }
}
```



## TCP客户端

### 流程

1. 建立与服务端的连接
2. 进行数据的发送和接收
3. 关闭连接

### 示例代码

```go
func main() {
  conn, err := net.Dial("tcp", "127.0.0.1:20000") // 建立与客户端的连接
  
  if err != nil {
    fmt.Println("err :", err)
    return
  }
  
  defer conn.Close() // 延时关闭连接
  
  inputReader := bufio.NewReader(os.Stdin) // 使用 bufio.Reader 获取标准输入
  
  for {
    input, _ := inputReader.ReadString('\n') // 读取用户输入
    inputInfo := strings.Trim(input, "\r\n") // 处理用户输入
    
    if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
      return
    }
    
    _, err = conn.Write([]byte(inputInfo)) // 发送数据
    
    if err != nil {
      return
    }
    
    buf := [512]byte{} // 创建读取缓冲区
    
    n, err := conn.Read(buf[:])
    
    if err != nil {
      fmt.Println("recv failed, err:", err)
      return
    }
    fmt.Println(string(buf[:n]))
  }
}
```



## TCP粘包

###  粘包示例代码

#### 服务器端

```go
func main() {

  listen, err := net.Listen("tcp", "127.0.0.1:30000") // 监听端口
  
  if err != nil {
    fmt.Println("listen failed, err:", err)
    return
  }
  
  defer listen.Close() // 延时关闭
  
  for { // 等待客户端连接
    
    conn, err := listen.Accept() // 接受连接
    
    if err != nil {
      fmt.Println("accept failed, err:", err)
      continue
    }
    
    go process(conn) // 启动一个goroutine去处理连接
  }
}

func process(conn net.Conn) {
  
  defer conn.Close() // 延时关闭
  
  var buf [1024]byte // 创建读取缓存区
  
  for {
    n, err := conn.Read(buf[:])
    
    if err == io.EOF { // 判断是否读取完毕
      break
    }
    
    if err != nil {
      fmt.Println("read from client failed, err:", err)
      break
    }
    
    recvStr := string(buf[:n]) // 处理读取到的数据
    fmt.Println("收到client发来的数据：", recvStr)
  }
}
```

#### 客户端

```go
func main() {
  
  conn, err := net.Dial("tcp", "127.0.0.1:30000") // 建立连接
  
  if err != nil {
    fmt.Println("dial failed, err", err)
    return
  }
  
  defer conn.Close() // 延时关闭连接
  
  for i := 0; i < 20; i++ { // 连续发送数据
    msg := `Hello, Hello. How are you?`
    conn.Write([]byte(msg))
  }
}
```

#### 效果

并不是预期20条`Hello, Hello. How are you?`，而是数据部分混合在一起了，这种现象称为`tcp粘包`



### 发生粘包原因

粘包是因为tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的接收和发送

粘包可发生在发送端也可发生在接收端

#### 发生在发送端

Nagle算法可能会造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法，简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去

#### 发生在接收端

接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据

当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据



### 解决粘包方案

出现粘包的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作

#### 封包

封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容

包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包

#### 解包

解包就是根据包头长度固定以及包头中含有包体长度的变量来拆分出一个完整的数据包的过程



### 解决方案示例代码

#### 自定义协议

```go
package proto

import (
  "bufio"
  "bytes"
  "encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
  
  // 读取消息的长度，转换成int32类型（占4个字节）
  var length = int32(len(message))
  var pkg = new(bytes.Buffer)
  
  // 创建消息头
  err := binary.Write(pkg, binary.LittleEndian, length)
  if err != nil {
    return nil, err
  }
  
  // 创建消息体
  err = binary.Write(pkg, binary.LittleEndian, []byte(message))
  if err != nil {
    return nil, err
  }
  
  return pkg.Bytes(), nil // 返回字节数组形式的封包
}

// Decode 解码消息
func Decode(conn *bufio.Reader) (string, error) {
  
  // 读取消息的长度
  lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
  lengthBuff := bytes.NewBuffer(lengthByte)
  
  var length int32
  err := binary.Read(lengthBuff, binary.LittleEndian, &length)
  if err != nil {
    return "", err
  }
  
  // Buffered返回缓冲中现有的可读取的字节数。
  if int32(conn.Buffered()) < length+4 {
    return "", err
  }

  // 读取真正的消息数据
  pack := make([]byte, int(4+length))
  _, err = conn.Read(pack)
  
  if err != nil {
    return "", err
  }
  
  return string(pack[4:]), nil // 返回读取到实际信息
}
```

#### 服务端

```go
func main() {

  listen, err := net.Listen("tcp", "127.0.0.1:30000") // 监听套接字
  
  if err != nil {
    fmt.Println("listen failed, err:", err)
    return
  }
  
  defer listen.Close() // 延时关闭监听
  
  for { // 等待用户的连接
    
    conn, err := listen.Accept() // 接受用户连接
    
    if err != nil {
      fmt.Println("accept failed, err:", err)
      continue
    }
    
    go process(conn) // 启动一个goroutine去处理连接
  }
}

func process(conn net.Conn) {
  
  defer conn.Close() // 延时关闭连接
  
  reader := bufio.NewReader(conn)
  
  for {
    msg, err := proto.Decode(reader)
    if err == io.EOF {
      return
    }
    if err != nil {
      fmt.Println("decode msg failed, err:", err)
      return
    }
    fmt.Println("收到client发来的数据：", msg)
  }
}
```

#### 客户端

```go
func main() {
  
  conn, err := net.Dial("tcp", "127.0.0.1:30000") // 建立连接
  
  if err != nil {
    fmt.Println("dial failed, err", err)
    return
  }
  
  defer conn.Close() // 延时关闭连接
  
  for i := 0; i < 20; i++ { // 发送20条数据
    
    msg := `Hello, Hello. How are you?` // 包体内容
    
    data, err := proto.Encode(msg) // 进行封包操作
    
    if err != nil {
      fmt.Println("encode msg failed, err:", err)
      return
    }
    
    conn.Write(data) // 发送封包
  }
}
```

