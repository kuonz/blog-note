---
author: "kuonz"
draft: false
title: "Git配置与工作区域"
date: 2020-04-05
categories: ["Git"]
---
  
## Git配置

### Git配置分为三种级别

| 分类     | 说明                                   | 配置写法            | 配置存放路径    |
| -------- | -------------------------------------- | ------------------- | --------------- |
| 系统级别 | 同计算机所有账号的仓库都使用这份配置   | git config --system | /etc/.gitconfig |
| 全局级别 | 同系统账号下的不同的仓库都使用这份配置 | git config --global | ~/gitconfig     |
| 仓库级别 | 只有当前仓库使用这份配置               | git config --local  | ./.git/config   |

### 必须配置

Git在使用前必须要配置使用者的信息，包括用户名和邮箱，配置方法如下

```shell
git config --system user.name <用户名>
git config --system user.email <邮箱>
```

示例

```shell
git config --system user.name alice
git config --system user.email alice@outlook.com
```

### 查看所有的配置项

```shell
git config -l
```



## 其他配置

### ssh公钥

#### ssh公钥生成命令如下

```shell
ssh-keygen -t rsa -C <邮箱地址>
# -t： 表示加密类型
# rsa： 使用rsa加密方式
# -C： 表示注释
# <邮箱地址>： 最好与Git配置中user.email相同
```

生成私钥：**id_rsa**，生成公钥：**id_rsa.pub**

#### ssh公钥存放位置

* Windows: `C:/User/<UserName>/.ssh/id_rsa.pub`

* Linux: `~/.ssh/id_rsa.pub`

### .gitignore

作用：用于忽略追踪，即在**.gitignore**文件中注明的某个文件或某类文件将被git所忽略

位置：**.gitignore**文件与**.git**放在同一个目录下

切记：不要管理敏感信息

自定义匹配规则

| 规则           | 示例                                                         |
| -------------- | ------------------------------------------------------------ |
| 匹配文件名     | a.js // 不管理a.js<br/>b.go // 不管理c.go<br/>.gitignore // 不管理.gitignore |
| 匹配文件名后缀 | * .js // 不管理js源码 *.go // 不管理go源码                   |
| 匹配文件夹     | files/<br/>node_modules/                                     |
| 取反，使用 !   | files/<br/>!files/a.js # files中除了 a.js 外都不进行管理     |



## Git工作区域

Git本地有三个工作区域：

* 工作区（Working Directory）
* 暂存区(Stage/Index)
* 本地库(Repository)

此外还有远程仓库，关系图如下

![](/03-Git配置与工作区域-images/image-20200324051731618.png)