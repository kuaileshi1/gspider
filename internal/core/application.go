// @Title 应用初始化文件
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/24 10:51 下午
package core

import (
	"github.com/kuaileshi1/dbable"
	"github.com/kuaileshi1/redis"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"gspider/internal/pkg/cron"
	"net/http"
	"path"
	"time"
)

const (
	defaultName    = "gspider"
	defaultVersion = "v1.0.0"
	defaultPath    = "/opt/case/gspider"
)

// @Description 定义应用接口
// @Auth shigx
// @Date 2021/12/2 2:37 下午
// @param
// @return
type AppContract interface {
	InitContract
	SetServerHttpHandler(httpHandler http.Handler) AppContract
	SetServerCliHandler(jobHandler *cron.Job) AppContract
	AppConfig(key string) string
	SetDebug()
	Start()
}

// 服务初始化接口
type InitContract interface {
	Init() AppContract
	SetName(name string) AppContract
	SetVersion(version string) AppContract
	SetPath(path string) AppContract
}

var (
	App           *Application
	isCommandLine bool
)

// @Description AppContract接口的实现
// @Auth shigx
// @Date 2021/12/2 2:38 下午
// @param
// @return
type Application struct {
	Name       string
	Version    string
	Path       string
	Config     *Config
	Debug      bool
	CliHandler *cron.Job // 命令行模式下handler
}

// @Description 构造Application
// @Auth shigx
// @Date 2021/11/25 4:48 下午
// @param
// @return
func NewApp() AppContract {
	App = &Application{
		Name:    defaultName,
		Version: defaultVersion,
		Path:    defaultPath,
	}
	return App
}

// @Description 应用初始化
// @Auth shigx
// @Date 2021/11/25 3:46 下午
// @param
// @return
func (app *Application) Init() AppContract {
	app.Config = InitConfig()

	app.initLog()   // 初始化日志
	app.initMysql() // 初始化mysql
	app.initRedis() // 初始化redis

	return App
}

// @Description 设置应用名称
// @Auth shigx
// @Date 2021/11/25 3:46 下午
// @param
// @return
func (app *Application) SetName(name string) AppContract {
	app.Name = name

	return app
}

// @Description 设置应用版本
// @Auth shigx
// @Date 2021/11/25 3:46 下午
// @param
// @return
func (app *Application) SetVersion(version string) AppContract {
	app.Version = version
	return app
}

// @Description 设置应用路径
// @Auth shigx
// @Date 2021/11/25 3:45 下午
// @param
// @return
func (app *Application) SetPath(path string) AppContract {
	app.Path = path
	return app
}

// @Description 设置web httpHandler
// @Auth shigx
// @Date 2021/11/25 3:45 下午
// @param
// @return
func (app *Application) SetServerHttpHandler(httpHandler http.Handler) AppContract {
	app.Config.Server.HttpHandler = httpHandler
	isCommandLine = false
	return app
}

// @Description 设置任务jobHandler
// @Auth shigx
// @Date 2021/12/2 2:39 下午
// @param
// @return
func (app *Application) SetServerCliHandler(jobHandler *cron.Job) AppContract {
	app.CliHandler = jobHandler
	isCommandLine = true

	return app
}

// @Description 程序启动入口
// @Auth shigx
// @Date 2021/11/25 2:55 下午
// @param
// @return
func (app *Application) Start() {

	if isCommandLine {
		go app.startRpc()
		app.CliHandler.Run()
		return
	}

	app.startWeb()
}

// @Description 应用配置查询
// @Auth shigx
// @Date 2021/11/25 3:07 下午
// @param
// @return
func (app *Application) AppConfig(key string) string {
	var value string
	if v, ok := app.Config.App[key]; ok {
		value = v.(string)
	}

	return value
}

// @Description 开启debug模式
// @Auth shigx
// @Date 2021/11/26 9:18 上午
// @param
// @return
func (app *Application) SetDebug() {
	app.Debug = true
	dbable.SetDebugOn()
	log.SetLevel(log.DebugLevel)
}

// @Description 初始化日志
// @Auth shigx
// @Date 2021/11/25 3:01 下午
// @param fileName string 日志文件名前缀
// @return
func (app *Application) initLog() {
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	log.SetReportCaller(true)

	fullPath := path.Join(app.AppConfig("log_path"), app.Name)
	writer, err := rotatelogs.New(
		fullPath+"_%Y%m%d.log",
		rotatelogs.WithLinkName(fullPath),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24*7),     // 日志最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatalf("Init log failed error:%v", err)
	}
	log.SetOutput(writer)
}

// @Description 初始化mysql
// @Auth shigx
// @Date 2021/11/25 5:02 下午
// @param
// @return
func (app *Application) initMysql() {
	for instance, config := range App.Config.Mysql {
		dbable.Init(config, instance)
	}
}

// @Description 初始化redis
// @Auth shigx
// @Date 2021/12/1 3:56 下午
// @param
// @return
func (app *Application) initRedis() {
	for instance, config := range App.Config.Redis {
		if client, err := redis.NewClient(config); err == nil {
			redis.PutClient(instance, client)
		}
	}
}

// @Description 启动web服务
// @Auth shigx
// @Date 2021/11/25 3:33 下午
// @param
// @return
func (app *Application) startWeb() {
	if err := StartServer(app.Config.Server); err != nil {
		log.Fatal("ListenAndServe Failed: ", err)
	}
}

// @Description 启动rpc服务
// @Auth shigx
// @Date 2021/12/26 4:03 下午
// @param
// @return
func (app *Application) startRpc() {
	if err := StartRpcServer(app.Config.Rpc); err != nil {
		log.Fatal("ListenAndRpcServe Failed: ", err)
	}
}
