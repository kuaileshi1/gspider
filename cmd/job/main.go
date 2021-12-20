// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/24 6:26 下午
package main

import (
	log "github.com/sirupsen/logrus"
	"gspider/internal/core"
	"gspider/internal/pkg/redis"
	"gspider/internal/router"
	_ "gspider/internal/rule"
	"gspider/internal/service/jobservice"
	"time"
)

var app core.AppContract

func init() {
	app = core.NewApp().
		SetName("gspider-job").
		SetVersion("v1.0.0").
		SetPath("/opt/case/gspider").
		Init()
}

func main() {
	time.Local, _ = time.LoadLocation("Asia/Shanghai")
	app.SetDebug()
	app.SetServerCliHandler(router.InitJobRouter())

	err := redis.Client.Init()
	if err != nil {
		log.Errorf("Init:redis init failed,err:%v", err)
		panic(err)
	}

	taskService := jobservice.NewTaskService()
	go taskService.CheckTask()
	go taskService.ManageTaskStatus()

	app.Start()
}
