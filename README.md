# gspider
go 实现的爬虫框架，提供web管理页面。主要基于gin+colly开发，web和脚本支持独立部署。

### 项目启动

```bash
0、将gspider.sql导入到mysql，修改configs/app-ga.yaml配置信息
1、本项目web管理页面基于vue-admin-template开发，项目地址：https://github.com/kuaileshi1/gspider-vue。
vue生成文件直接copy到internal/dist目录。
2、$ cd internal/router
3、$ parckr2 build
4、$ go build -o cmd/web/main cmd/web/main.go  # web服务编译
5、$ go build -o cmd/job/main cmd/job/main.go  # 脚本服务编译
6、$ ./cmd/web/main -conf=configs/app-ga.yaml  # 启动web服务
7、$ ./cmd/job/main -conf=configs/app-ga.yaml  # 启动脚本服务

```
项目构建和部署可以根据自己业务实际情况进行操作。由于项目架构按照微服务的形式进行拆分，目前跨应用操作逻辑借助redis进行交互，后期可以基于grpc实现。

### 项目目录结构
```
.
├── api                             web页面接口文档
├── cmd                             应用入口文件
│   ├── job
│   │   └── main.go
│   └── web
│       └── main.go
├── configs                         配置文件
│   └── app-ga.yaml
├── go.mod
├── go.sum
├── gspider.sql                     数据库sql文件
├── internal                        核心业务目录
│   ├── constant                    项目常量目录
│   │   ├── rediskey                redis key定义
│   │   │   └── redis.go
│   │   └── task                    任务常量定义
│   │       └── task.go
│   ├── controller                  控制器层
│   │   ├── job                     脚本任务
│   │   │   ├── cxytiandi.go
│   │   │   ├── task.go
│   │   │   └── wangyi.go
│   │   └── web                     web接口
│   │       ├── demo.go
│   │       ├── task.go
│   │       └── user.go
│   ├── core                        应用核心目录
│   │   ├── application.go
│   │   ├── configure.go
│   │   └── server.go
│   ├── dist                        前端vue生成文件
│   │   ├── favicon.ico
│   │   ├── gspider.jpeg
│   │   ├── index.html
│   │   └── static
│   │       ├── css
│   │       ├── fonts
│   │       ├── img
│   │       └── js
│   ├── middleware                  jin框架中间件
│   │   ├── cors
│   │   │   └── cors.go
│   │   └── jwt
│   │       └── gin_jwt.go
│   ├── model                       应用model层
│   │   ├── cxytiandi_model.go
│   │   ├── entity                  表对应实体
│   │   │   ├── cxytiandi.go
│   │   │   ├── task.go
│   │   │   ├── user.go
│   │   │   └── wangyi_tech_news.go
│   │   ├── task_model.go
│   │   ├── user_model.go
│   │   └── wangyi_tech_news_model.go
│   ├── pkg                        引用外部组件封装
│   │   ├── cron
│   │   │   └── cron.go
│   │   ├── redis
│   │   │   └── client.go
│   │   ├── response
│   │   │   ├── code.go
│   │   │   └── response.go
│   │   └── spider                  爬虫核心业务
│   │       ├── colly.go
│   │       ├── context.go
│   │       ├── helper.go
│   │       ├── spider.go
│   │       ├── task.go
│   │       ├── task_config.go
│   │       └── task_rule.go
│   ├── router                      路由层
│   │   ├── job_route.go
│   │   ├── packrd
│   │   │   └── packed-packr.go
│   │   ├── route.go
│   │   └── router-packr.go
│   ├── rule                        自定义抓取规则逻辑
│   │   ├── cxytiandi
│   │   │   └── cxytiandi.go
│   │   ├── rule.go
│   │   └── wangyi
│   │       └── tech_news.go
│   └── service                     应用service层
│       ├── jobservice
│       │   ├── cron_service.go
│       │   └── task_service.go
│       └── webservice
│           └── task_service.go
├── logs                            日志目录
└── pkg                             外部组件引用目录
```
