// @Title 任务路由文件
// @Description 该位置可以定义一些和爬虫无关的任务
// @Author shigx 2021/11/29 10:49 上午
package router

import (
	"gspider/internal/controller/job"
	"gspider/internal/pkg/cron"
)

// @Description 任务脚本路由入口
// @Auth shigx
// @Date 2021/11/29 10:51 上午
// @param
// @return
func InitJobRouter() *cron.Job {
	c := cron.NewJob()

	// web页面操作任务通过异步任务处理
	c.Cron.AddFunc("*/2 * * * * *", job.MangeTask)

	c.Cron.AddFunc("*/2 * * * * *", job.CxyOutToMysql)
	c.Cron.AddFunc("*/2 * * * * *", job.WangyiTechOutToMysql)

	return c
}
