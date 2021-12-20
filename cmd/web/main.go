// @Title 项目入口文件
// @Description 网站启动入口文件
// @Author shigx 2021/11/24 6:25 下午
package main

import (
	log "github.com/sirupsen/logrus"
	"gspider/internal/core"
	"gspider/internal/pkg/redis"
	"gspider/internal/router"
	_ "gspider/internal/rule"
	"time"
)

func main() {
	app := core.NewApp().
		SetName("gspider-web").
		SetVersion("v1.0.0").
		SetPath("/opt/case/gspider").
		Init()

	Init(app)

	app.Start()
}

// @Description 入口初始化
// @Auth shigx
// @Date 2021/12/2 2:36 下午
// @param
// @return
func Init(app core.AppContract) {
	time.Local, _ = time.LoadLocation("Asia/Shanghai")

	app.SetServerHttpHandler(router.InitRouter())

	err := redis.Client.Init()
	if err != nil {
		log.Errorf("Init:redis init failed,err:%v", err)
		panic(err)
	}

	// 开启debug模式
	app.SetDebug()
}
