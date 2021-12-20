// @Title 定时任务封装
// @Description 对定时任务组件进行封装操作
// @Author shigx 2021/11/29 10:23 上午
package cron

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// Job结构体
type Job struct {
	Cron *cron.Cron
}

// 任务日志结构体
type JobLogger struct {
}

var (
	job     *Job
	logger  *JobLogger
	running bool
)

// @Description 实例化job
// @Auth shigx
// @Date 2021/11/29 10:39 上午
// @param
// @return
func NewJob() *Job {
	if running {
		return job
	}
	job = &Job{
		Cron: cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(logger))),
	}

	return job
}

// @Description 对定时任务Run方法重写
// @Auth shigx
// @Date 2021/12/2 2:31 下午
// @param
// @return
func (j *Job) Run() {
	running = true
	j.Cron.Run()
}

// @Description 实现接口Info方法
// @Auth shigx
// @Date 2021/12/2 2:32 下午
// @param
// @return
func (l *JobLogger) Info(msg string, keysAndValues ...interface{}) {
	log.Info(msg, keysAndValues)
}

// @Description 实现接口Error方法
// @Auth shigx
// @Date 2021/12/2 2:32 下午
// @param
// @return
func (l *JobLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	log.Error(err.Error(), msg, keysAndValues)
}
