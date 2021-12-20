// @Title 任务处理服务
// @Description 任务服务类
// @Author shigx 2021/11/26 10:36 下午
package jobservice

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gspider/internal/constant/task"
	"gspider/internal/model"
	"gspider/internal/model/entity"
)

// 任务对应状态chan
var retCh = make(chan entity.TIS, 1)

// @Description 获取chan
// @Auth shigx
// @Date 2021/12/2 3:41 下午
// @param
// @return
func GetTISChan() chan entity.TIS {
	return retCh
}

// 任务service
type TaskService struct {
	taskModel *model.TaskModel
}

// @Description 实例化service
// @Auth shigx
// @Date 2021/11/26 11:19 下午
// @param
// @return
func NewTaskService() *TaskService {
	taskModel := model.NewTaskModel("default")
	return &TaskService{
		taskModel: taskModel,
	}
}

// @Description 检查任务状态并添加到cron
// @Auth shigx
// @Date 2021/11/26 11:38 下午
// @param
// @return
func (s *TaskService) CheckTask() {
	tasks, err := s.taskModel.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("Not fond task, service return")
			return
		}
		log.Errorf("Get task from db error:%v", err)
		return
	}

	for _, val := range tasks {
		if val.Status == task.StatusRunning {
			updates := map[string]interface{}{"status": task.StatusUnExceptedExited}
			err := s.taskModel.Updates("id = ?", []interface{}{val.ID}, updates)
			if err != nil {
				log.Errorf("Update task status err, id:%d, err:%v", val.ID, err)
				continue
			}
		}
		if (val.Status == task.StatusCompleted || val.Status == task.StatusRunning ||
			val.Status == task.StatusUnExceptedExited) && val.CronSpec != "" {
			if err := s.CreateCronTask(val); err != nil {
				log.Errorf("CreateCronTask failed! err:%v", err)
				continue
			}
		}
	}
}

// @Description 创建任务
// @Auth shigx
// @Date 2021/11/29 1:58 下午
// @param
// @return
func (s *TaskService) CreateCronTask(task entity.Task) error {
	cs, err := NewCronService(task.ID, task.CronSpec, GetTISChan())
	if err != nil {
		return err
	}

	return cs.Start()
}

// @Description 更新任务执行状态
// @Auth shigx
// @Date 2021/12/1 10:18 上午
// @param
// @return
func (s *TaskService) ManageTaskStatus() {
	for {
		select {
		case ch := <-retCh:
			taskInfo, err := s.taskModel.GetOneById(ch.ID)
			if err != nil {
				log.Errorf("query model task err :%v", err)
				break
			}
			// 状态已停止就不操作了
			if taskInfo.Status == task.StatusStopped {
				break
			}
			taskInfo.Status = ch.Status
			if ch.Status == task.StatusCompleted {
				taskInfo.Counts += 1
			}
			update := map[string]interface{}{
				"status": taskInfo.Status,
				"counts": taskInfo.Counts,
			}
			if err := s.taskModel.Updates("id = ?", []interface{}{taskInfo.ID}, update); err != nil {
				log.Errorf("update task status err:%v", err)
				break
			}
		}
	}
}

// @Description 根据任务ID查询任务
// @Auth shigx
// @Date 2021/12/17 2:42 下午
// @param
// @return
func (s *TaskService) GetTaskById(taskId int) (ret *entity.Task, err error) {
	return s.taskModel.GetOneById(taskId)
}
