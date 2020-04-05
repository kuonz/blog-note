---
author: "kuonz"
draft: false
title: "Docker私有仓库"
date: 2020-04-05
categories: ["Docker"]
---
  
## 搭建私有仓库

1. 拉取私有仓库镜像

   ```shell
   docker pull registry
   ```

2. 启动私有仓库容器

   ```shell
   docker run -id --name registry -p 宿主机端口:5000 registry
   ```

3. 检查私有仓库是否启动成功

   ```shell
   浏览器访问 http://私有仓库服务器ip:端口号/v2/_catalog
   ```

4. 修改`daemon.json`配置

   ```shell
   sudo vim /etc/docker/daemon.json
   # 加上下面一行
   {
     "insecure-registries": ["私有仓库服务ip:端口号"]
   }
   ```

5. 重启Docker服务

   ```shell
   docker start registry
   ```

   

## 上传与拉取私有仓库

### 上传镜像到私有仓库

1. 为上传镜像打标签

   ```shell
   docker tag 要上传的镜像名:版本号 私有仓库服务器ip:端口号/镜像名:版本号
   ```

2. 上传标记的镜像

   ```shell
   docker push 私有仓库服务器ip:端口号/镜像名:版本号
   ```


### 从私有仓库拉镜像

```shell
docker pull 私有仓库服务器ip:端口号/镜像名:版本号
```