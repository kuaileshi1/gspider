// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/13 11:27 下午
package web

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/task"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/response"
	"gspider/internal/pkg/spider"
	"gspider/internal/service/webservice"
	"strconv"
)

// 列表参数
type listParam struct {
	Size   int `json:"size" form:"size"`
	Offset int `json:"offset" form:"offset"`
}

// 列表返回
type listResp struct {
	List  []entity.Task `json:"list"`
	Total int64         `json:"total"`
}

// 保存数据参数
type saveTaskReq struct {
	*entity.Task
}

// 保存数据返回
type saveTaskResp struct {
	ID int `json:"id"`
}

type Task struct {
	taskService *webservice.TaskService
}

// @Description 实例化
// @Auth shigx
// @Date 2021/12/17 1:28 下午
// @param
// @return
func NewTask() *Task {
	return &Task{
		taskService: webservice.NewTaskService(),
	}
}

// @Description 任务列表
// @Auth shigx
// @Date 2021/12/13 11:37 下午
// @param
// @return
func (w *Task) List(c *gin.Context) interface{} {
	var param listParam
	if err := c.BindQuery(&param); err != nil {
		return response.Error(c, response.ParamsErrorCode, response.GetMessage(response.ParamsErrorCode))
	}
	list, count := w.taskService.GetList(param.Size, param.Offset)

	data := listResp{
		List:  list,
		Total: count,
	}

	return response.Success(c, data)
}

// @Description 创建和更新数据
// @Auth shigx
// @Date 2021/12/19 1:11 上午
// @param
// @return
func (w *Task) Save(c *gin.Context) interface{} {
	var param saveTaskReq
	if err := c.BindJSON(&param); err != nil {
		return response.Error(c, response.ParamsErrorCode, response.GetMessage(response.ParamsErrorCode))
	}
	task := param.Task
	err := w.taskService.SaveTask(task)
	if err != nil {
		return response.Error(c, response.FailureCode, response.GetMessage(response.FailureCode))
	}

	return response.Success(c, saveTaskResp{ID: task.ID})
}

// @Description 查询任务
// @Auth shigx
// @Date 2021/12/17 5:22 下午
// @param
// @return
func (w *Task) GetTask(c *gin.Context) interface{} {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return response.Error(c, response.ParamsErrorCode, response.GetMessage(response.ParamsErrorCode))
	}

	taskInfo, err := w.taskService.GetTaskById(id)
	if err != nil {
		log.Errorf("GetTask from db failed, id:%d, err:%v", id, err)
		return response.Error(c, response.ParamsErrorCode, response.GetMessage(response.ParamsErrorCode))
	}

	return response.Success(c, taskInfo)
}

// @Description 停止任务
// @Auth shigx
// @Date 2021/12/17 1:56 下午
// @param
// @return
func (w *Task) Stop(c *gin.Context) interface{} {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return response.Error(c, response.ParamsErrorCode, response.GetMessage(response.ParamsErrorCode))
	}

	if res := w.taskService.ChangeStatus(id, task.StatusStopped); !res {
		return response.Error(c, response.FailureCode, response.GetMessage(response.FailureCode))
	}

	return response.Success(c)

}

// @Description 开启任务
// @Auth shigx
// @Date 2021/12/17 3:39 下午
// @param
// @return
func (w *Task) Start(c *gin.Context) interface{} {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return response.Error(c, response.ParamsErrorCode, response.GetMessage(response.ParamsErrorCode))
	}

	if res := w.taskService.ChangeStatus(id, task.StatusRunning); !res {
		return response.Error(c, response.FailureCode, response.GetMessage(response.FailureCode))
	}

	return response.Success(c)
}

// @Description 删除任务
// @Auth shigx
// @Date 2021/12/20 10:32 上午
// @param
// @return
func (w *Task) Delete(c *gin.Context) interface{} {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return response.Error(c, response.ParamsErrorCode, response.GetMessage(response.ParamsErrorCode))
	}

	if err := w.taskService.DeleteTask(id); err != nil {
		return response.Error(c, response.FailureCode, response.GetMessage(response.FailureCode))
	}

	return response.Success(c)
}

// @Description 获取抓取规则信息
// @Auth shigx
// @Date 2021/12/17 5:08 下午
// @param
// @return
func (w *Task) Rules(c *gin.Context) interface{} {
	keys := spider.GetTaskRuleKeys()
	return response.Success(c, keys)
}
