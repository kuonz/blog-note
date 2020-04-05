---
author: "kuonz"
draft: false
title: "Git介绍下载安装"
date: 2020-04-05
categories: ["Git"]
---
  
## Git介绍

### Git是什么

* Git是目前世界上最先进的开源免费的分布式版本控制系统

* Linux的发起者linus为了方便管理的Linux代码，花了2周开发出来的版本管理工具，目的是替代BitKeeper

### Git优点与缺点

| 优点                                   | 缺点                                |
| -------------------------------------- | ----------------------------------- |
| 适合分布式开发，强调个体               | 模式上比SVN更加复杂                 |
| 公共服务器压力和数据量都不会太大       | 权限控制没有SVN方便，代码保密性较差 |
| 速度快、灵活                           |                                     |
| 任意两个开发者之间可以很容易的解决冲突 |                                     |
| 支持离线工作                           |                                     |

### Git官网与源码

官网： [https://git-scm.com/](https://link.zhihu.com/?target=https%3A//git-scm.com/)

源码： [https://github.com/git/git/](https://link.zhihu.com/?target=https%3A//github.com/git/git/)



## Git下载安装

### Windows

* 下载：https://git-scm.com
* 安装：按照默认设置，一直下一步即可完成安装

### Linux

```shell
sudo apt-get install git # Ubuntu系
sudo yum -y install git # Centos系
```

### Mac OS

```shell
brew install git
```

### 检查是否安装成功

```shell
git --version
```

