// @Title 定时任务封装
// @Description 对定时任务组件进行封装操作
// @Author shigx 2021/11/29 10:23 上午
package cron

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

type EntryID cron.EntryID

// Job结构体
type Job struct {
	cron *cron.Cron
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
		cron: cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(logger))),
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
	j.cron.Run()
}

// @Description 添加任务
// @Auth shigx
// @Date 2021/12/29 5:33 下午
// @param
// @return
func (j *Job) AddFunc(spec string, cmd func()) (EntryID, error) {
	id, err := j.cron.AddFunc(spec, cmd)

	return EntryID(id), err
}

// @Description 添加任务
// @Auth shigx
// @Date 2021/12/29 5:39 下午
// @param
// @return
func (j *Job) AddJob(spec string, cmd cron.Job) (EntryID, error) {
	id, err := j.cron.AddJob(spec, cmd)

	return EntryID(id), err
}

// @Description 移除定时任务
// @Auth shigx
// @Date 2021/12/29 5:42 下午
// @param
// @return
func (j *Job) Remove(id EntryID) {
	j.cron.Remove(cron.EntryID(id))
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
