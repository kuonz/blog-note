---
author: "kuonz"
draft: false
title: "Redis下载与安装"
date: 2020-04-05
categories: ["Redis介绍、下载与安装"]
---
  
> 由于 Redis 对 Windows 平台支持并不好，且 Redis 一般只用于 Linux 服务器上，所以本文只介绍如何在 Linux 中使用 Redis

## 普通方式

### 下载

访问 [Redis官方网站](https://redis.io/download)，挑选合适的版本进行下载

![](/02-Redis下载与安装-images/image-20200322153536217.png)

### 安装

在下载完毕后，执行下列命令进行解压

```shell
tar -zxvf 文件名.tar.gz
```

然后把解压后的文件夹放置到合适的路径下，执行下列的命令进行编译和安装

```shell
make && make install [destdir=安装目录]
```



## Docker方式