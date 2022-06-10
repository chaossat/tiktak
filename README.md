# tiktak
A minimalist back-end framwork built in golang language to imitate tiktok

# 使用方法
先放置配置文件到./config
## windows
EasyStart.bat

EasyStartExe.bat

## linux环境下启动
```shell
./build-start.sh #一键编译并启动所有服务
./start.sh  #一键启动所有服务
./shutdown.sh #一键杀死所有服务
```

# 项目结构
- common
  - 数据库等基础构件包，目前仅数据库
- config
  - 配置文件夹
- controller
  - web处理器函数包
- db
  - 数据库操作函数包
- middleware
  - 鉴权包等中间件包，目前仅jwt
- migration
  - gorm自动迁移
- model
  - 数据库model
- oss
  - oss对象储存操作函数包
- router
  - gin路由函数包
- service
  - 微服务包
- static
  - 静态资源
- tempfile
  - 视频文件本地临时储存位置
- tempimage
  - 视频封面本地临时储存位置
- util
  - 哈希、加密等工具包
