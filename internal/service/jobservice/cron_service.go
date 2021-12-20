// @Title 定时任务服务类
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/28 5:32 下午
package jobservice

import (
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	constant_task "gspider/internal/constant/task"
	"gspider/internal/core"
	"gspider/internal/model"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/spider"
	"sync"
)

// 用于存储已启动任务
var cronTaskMap = &sync.Map{}

// 定时任务服务
type CronService struct {
	taskID   int
	cronSpec string
	entityID cron.EntryID
	retCh    chan<- entity.TIS
}

// @Description 实例化cron
// @Auth shigx
// @Date 2021/11/28 5:50 下午
// @param
// @return
func NewCronService(taskID int, cronSpec string, retCh chan<- entity.TIS) (*CronService, error) {
	cs := &CronService{
		taskID:   taskID,
		cronSpec: cronSpec,
		retCh:    retCh,
	}
	if err := AddCronService(cs); err != nil {
		return nil, err
	}

	return cs, nil
}

// @Description 任务执行方法
// @Auth shigx
// @Date 2021/11/30 11:19 下午
// @param
// @return
func (cs *CronService) Run() {
	taskModel := model.NewTaskModel("default")
	task, err := taskModel.GetOneById(cs.taskID)
	if err != nil {
		log.Errorf("run cron task failed, query task err:%v", err)
		return
	}
	if task.Status != constant_task.StatusCompleted &&
		task.Status != constant_task.StatusUnExceptedExited &&
		task.Status != constant_task.StatusRunning {
		log.Errorf("run cron task failed, status:%d", task.Status)
		return
	}
	spiderTask, err := spider.GetSpiderTask(task)
	if err != nil {
		log.Errorf("run cron task failed, err:%v", err)
		return
	}
	s := spider.New(spiderTask, cs.retCh)
	if err := s.Run(); err != nil {
		log.Errorf("run cron task failed, err:%v", err)
		cs.retCh <- entity.TIS{ID: task.ID, Status: constant_task.StatusUnExceptedExited}
		return
	}

	cs.retCh <- entity.TIS{ID: task.ID, Status: constant_task.StatusRunning}
}

// @Description 添加任务到定时管理器
// @Auth shigx
// @Date 2021/11/30 11:19 下午
// @param
// @return
func (cs *CronService) Start() error {
	entityID, err := core.App.CliHandler.Cron.AddJob(cs.cronSpec, cs)
	if err != nil {
		return err
	}
	cs.entityID = entityID

	return nil
}

// @Description 将任务从管理器中移除
// @Auth shigx
// @Date 2021/11/30 11:20 下午
// @param
// @return
func (cs *CronService) Stop() {
	core.App.CliHandler.Cron.Remove(cs.entityID)
	cronTaskMap.Delete(cs.taskID)
}

// @Description 添加任务service
// @Auth shigx
// @Date 2021/11/28 5:47 下午
// @param
// @return
func AddCronService(cs *CronService) error {
	if _, loaded := cronTaskMap.LoadOrStore(cs.taskID, cs); loaded {
		return fmt.Errorf("add cron task failed, taskID:%d", cs.taskID)
	}

	return nil
}

// @Description 获取任务service
// @Auth shigx
// @Date 2021/11/28 5:56 下午
// @param
// @return
func GetCronService(taskID int) *CronService {
	cs, ok := cronTaskMap.Load(taskID)
	if !ok {
		return nil
	}

	return cs.(*CronService)
}
