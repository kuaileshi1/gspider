// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/13 11:40 下午
package webservice

import (
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/task"
	"gspider/internal/model"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/rpc"
)

// 任务service
type TaskService struct {
	taskModel *model.TaskModel
}

// @Description 实例化task service
// @Auth shigx
// @Date 2021/12/13 11:42 下午
// @param
// @return
func NewTaskService() *TaskService {
	taskModel := model.NewTaskModel("default")
	return &TaskService{
		taskModel: taskModel,
	}
}

// @Description 获取任务列表
// @Auth shigx
// @Date 2021/12/13 11:58 下午
// @param
// @return
func (s *TaskService) GetList(size int, offset int) ([]entity.Task, int64) {
	count, err := s.taskModel.Count()
	if err != nil || count == 0 {
		return nil, count
	}

	list, err := s.taskModel.GetList(size, offset)
	if err != nil {
		return nil, 0
	}

	return list, count
}

// @Description 函数的详细描述
// @Auth shigx
// @Date 2021/12/17 1:52 下午
// @param
// @return
func (s *TaskService) ChangeStatus(taskId int, taskStatus task.Status) bool {
	err := s.taskModel.Updates("id = ?", []interface{}{taskId}, map[string]interface{}{"status": taskStatus})
	if err != nil {
		log.Errorf("update task status error, id: %d, status: %d, err:%v", taskId, taskStatus, err)
		return false
	}

	// rpc调用开启和关闭任务
	if taskStatus == task.StatusStopped || taskStatus == task.StatusRunning {
		client := rpc.NewClient()
		if client == nil {
			return false
		}
		defer client.Close()

		if taskStatus == task.StatusRunning {
			if res := client.Start(taskId); res == false {
				return false
			}
		} else {
			if res := client.Stop(taskId); res == false {
				return false
			}
		}

	}

	return true
}

// @Description 根据任务ID查询任务
// @Auth shigx
// @Date 2021/12/17 2:42 下午
// @param
// @return
func (s *TaskService) GetTaskById(taskId int) (ret *entity.Task, err error) {
	return s.taskModel.GetOneById(taskId)
}

// @Description 创建或更新任务
// @Auth shigx
// @Date 2021/12/19 1:05 上午
// @param
// @return
func (s *TaskService) SaveTask(task *entity.Task) error {
	return s.taskModel.Save(task)
}

// @Description 删除任务
// @Auth shigx
// @Date 2021/12/20 10:40 上午
// @param
// @return
func (s *TaskService) DeleteTask(id int) error {
	return s.taskModel.Delete(id)
}
