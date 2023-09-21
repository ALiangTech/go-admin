# go-admin

给admin-template 提供服务端支持 


## 项目结构说明

```go
go-admin/
    ├── main.go
    ├── api/  //包含了你的API的不同版本，每个版本有自己的路由、控制器、中间件和模型。
    │   ├── v1/ // 不同版本的API。
    │   │   ├── routes/ // 每个版本的路由定义。
    │   │   │   ├── user_routes.go
    │   │   │   └── ...
    │   │   ├── controllers/ // 控制器，用于处理HTTP请求和响应。
    │   │   │   ├── user_controller.go
    │   │   │   └── ...
    │   │   ├── middlewares/ // 中间件，用于处理请求前、请求后的逻辑。
    │   │   │   ├── auth_middleware.go
    │   │   │   └── ...
    │   │   └── models/ // 数据模型，用于与数据库交互。
    │   │       ├── user.go
    │   │       └── ...
    │   ├── v2/
    │   │   ├── routes/
    │   │   ├── controllers/
    │   │   ├── middlewares/
    │   │   └── models/
    │   └── ...
    ├── config/ // 包含应用程序的配置文件，如数据库配置等。
    │   ├── config.go
    │   └── database.go
    ├── utils/ // 包含一些通用的工具函数和帮助函数。
    ├── middleware/ // 自定义中间件函数。
    ├── services/ // 包含一些业务逻辑的服务层。
    ├── tests/ // 包含单元测试和集成测试。
    ├── static/ // 静态文件，如CSS、JavaScript文件等。
    ├── templates/ // 模板文件，如果你的应用使用模板引擎来渲染视图。
    ├── migrations/ // 数据库迁移文件，用于管理数据库结构变化。
    └── README.md
```


## 能力
基础CRUD
权限


## 服务端单一逻辑

请求 => 处理请求 
响应 <= 处理请求

## 针对响应 不会提示任何服务端的错误信息

举例来说 content-type 传递的类型 不对 不会通过接口告知前端

如果需要服务端校验呢？新开校验接口给前端

对前端来说 无非就是有数据  和 无数据