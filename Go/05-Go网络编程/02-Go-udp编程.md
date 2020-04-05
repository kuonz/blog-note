---
author: "kuonz"
draft: false
title: "Go-udp编程"
date: 2020-04-05
categories: ["Go网络编程"]
---
  
## UDP服务端

### 流程

1. 监听套接字（IP+端口）
2. 接收/发送数据

### 示例代码

```go
func main() {
  // 监听套接字
  listen, err := net.ListenUDP("udp", &net.UDPAddr{
    IP:   net.IPv4(0, 0, 0, 0),
    Port: 30000,
  })
  
  if err != nil {
    fmt.Println("listen failed, err:", err)
    return
  }
  
  defer listen.Close() // 延时关闭服务器
  
  for {
    var data [1024]byte // 定义接收数据的缓冲区
    
    n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
    
    if err != nil {
      fmt.Println("read udp failed, err:", err)
      continue
    }
    
    // 处理接收到的数据
    fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n) 
    
    _, err = listen.WriteToUDP(data[:n], addr) // 发送数据
    
    if err != nil {
      fmt.Println("write to udp failed, err:", err)
      continue
    }
  }
}
```



## UDP客户端

### 流程

1. 创建socket
2. 接收/发送数据

### 示例代码

```go
func main() {
  // 创建 socket
  socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
    IP:   net.IPv4(0, 0, 0, 0),
    Port: 30000,
  })
  
  if err != nil {
    fmt.Println("连接服务端失败，err:", err)
    return
  }
  
  defer socket.Close() // 延时关闭socket
  
  sendData := []byte("Hello server") // 定义要发送的数据
  
  _, err = socket.Write(sendData) // 发送数据
  
  if err != nil {
    fmt.Println("发送数据失败，err:", err)
    return
  }
  
  data := make([]byte, 4096) // 定义接收数据缓冲区
  
  n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
  
  if err != nil {
    fmt.Println("接收数据失败，err:", err)
    return
  }
  
   // 处理接收到的数据
  fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
```

