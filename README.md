### backend-admin
后台管理

#### 功能特征
- 基于 GIN WEB API 框架
- 基于Casbin的 RBAC 访问控制模型
- JWT 认证
- 支持 Swagger 文档(基于swaggo)
- 基于 GORM 的数据库存储，可扩展多种类型数据库
- 配置文件简单的模型映射，快速能够得到想要的配置
- 支持平滑升级
#### 接入准备

* 安装命令
```
    //cd bkAdmin文件夹
    go install
    //安装swag 
    go install github.com/swaggo/swag/cmd/swag
   
```

* 创建数据表结构,支持postgres和mysql
```
   //默认账号:admin 密码:123456
   bkAdmin migrate  
```

* 执行命令
```
    //生成swag文档
    swag init -g cmd/bkAdmin/main.go
    //运行http服务
    bkAdmin run 
```
