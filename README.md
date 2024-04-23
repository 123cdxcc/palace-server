# Go gin+websocket聊天室服务端简单实现

练手项目，聊天室一般分为服务端和客户端，本文为服务端实现。

一般会都用到数据库实现，但是为了快和简单，不使用数据库，数据全部都在内存中定义

## 功能

* 注册/注销
* 房间群聊
* 使用jwt权限校验

## 结构

```shell
.
├─cmd // 启动目录
├─configs // 配置文件
├─internal
│  ├─handler // 路由处理
│  ├─logic // 内部逻辑
│  ├─middlewares // 中间件
│  ├─model // 模型
│  └─router // 路由定义
├─pkg
│  ├─constant // 常量
│  ├─errors // 错误
│  ├─log // 日志
│  ├─resp // http返回方法
│  └─utils // 工具方法
└─types // 与前端交互模型定义
```

