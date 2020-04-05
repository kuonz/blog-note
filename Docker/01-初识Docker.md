---
author: "kuonz"
draft: false
title: "初识Docker"
date: 2020-04-05
categories: ["Docker"]
---
  
## Docker概述

### Docker是什么

1. `Docker`是一个开源的应用容器引擎，解决应用软件跨环境迁移的问题
2. `Docker`诞生于2013年初，基于`Go`语言实现，由`dotCloud`公司（后改名`Docker Inc`）出品

### Docker特点

1. 容器使用沙箱机制，不同的容器之间相互隔离
2. 容器性能开销极低，速度和资源消耗都优于虚拟机

### Docker版本

`Docker`从17.03版本后分为 `CE`（Community Edition）和 `EE` (Enterprise Edition) 两个版本



## Docker安装

### TODO

### Docker下载加速



## Docker与虚拟机区别

容器是在`Linux`系统上运行的一个独立进程，其并与其他容器共享主机的内核，不占用其他任何可执行文件的内存，非常轻量

虚拟机运行的是一个完成的操作系统，通过虚拟机管理程序对主机资源进行虚拟访问，相比之下需要的资源更多

![](/01-初识Docker-images/image-20200327214541526.png)

| 对比项     | Docker             | VM             |
| ---------- | ------------------ | -------------- |
| 启动       | 秒级               | 分钟级         |
| 硬盘使用量 | MB级               | GB级           |
| 性能       | 接近原生           | 弱于原生       |
| 系统支持量 | 单机支持上千个容器 | 单机支持几十个 |