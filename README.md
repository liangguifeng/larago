# larago

仿照laravel优雅的写法编写的Go语言开发骨架，方便开箱即用。

# 开发缘由

因为目前go的框架太多了，而我是由PHP转向golang的，对很多golang的语法并不熟悉，之前使用laravel比较多，所以自己按照laravel的设计模式开发了一个go骨架，为了方便自己日后的开发和项目的管理

# 使用方法
1. 复制配置文件，并修改.env文件内的相关配置
```shell
cp .env.example .env
```

2. 安装依赖
```shell
go mod vendor
```

3. 运行项目
```shell
go run index.go
```

4. 访问控制台输出的网址
- 默认是 [http://localhost:3000](http://localhost:3000)

## 功能列表
- [x] gohtml模板
- [x] 中间件
- [x] 路由组
- [x] JWT认证
- [x] config配置文件自动加载
- [x] .env文件
- [x] 日志
- [x] 存储文件
- [x] go mod管理依赖
- [x] 分页处理
- [x] gorm集成
- [x] gin集成