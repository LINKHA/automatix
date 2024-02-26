<p align="center">
<h1 align="center">automatix</h1>
<h6 align="center">golang游戏平台，用于社交和实时游戏和应用的分布式服务器 </h6>
</p>


## 组件

* **gate** - 长连接gate，负责转发消息到内部服务.
* **login** - 登录服务器，管理进入游戏服前的链路.
* **usercenter** - 用户中心，负责账号级别管理.
* **roommamanger** - 房间系统，支持房间匹配.
* **rolemanager** - 角色系统，负责角色级别管理.
* **punishment** - 触发、举报系统.

## 快速开始
### 前提
* **ubuntu** - Ubuntu 22.04以上版本
* **docker** - 需要安装docker
* **golang** - 1.20.6以上版本

### 启动
运行服务器和数据库最快的方式是使用Docker。

1. 生成automatix镜像，使用了modd做热更新:

   ```shell
   docker build -t automatix .
   ```

2. 创建automatix_net网桥:

   ```shell
   docker network create automatix_net
   ```

3. 启动中间件容器:

   ```shell
   docker-compose -f docker-compose-env.yml up -d
   ```

4. 启动automatix容器:

   ```shell
   docker-compose up -d 
   ```

5. 手动设置data权限:

   ```shell
   chmod -R 777 data/
   ```