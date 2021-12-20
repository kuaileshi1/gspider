// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/30 9:14 上午
package spider

import (
	"context"
	"fmt"
	"sync"
)

// 爬虫任务结构体
type Task struct {
	ID int
	TaskRule
	TaskConfig
}

// @Description 实例化任务
// @Auth shigx
// @Date 2021/12/2 11:03 下午
// @param
// @return
func NewTask(id int, rule TaskRule, config TaskConfig) *Task {
	return &Task{
		ID:         id,
		TaskRule:   rule,
		TaskConfig: config,
	}
}

var (
	ctlMu  = &sync.RWMutex{}
	ctlMap = make(map[int]context.CancelFunc) // 上下文取消函数
)

// @Description 执行上下文取消
// @Auth shigx
// @Date 2021/12/2 11:08 下午
// @param
// @return
func CancelTask(taskID int) bool {
	ctlMu.Lock()
	defer ctlMu.Unlock()

	cancel, ok := ctlMap[taskID]
	if !ok {
		return false
	}
	cancel()
	delete(ctlMap, taskID)

	return true
}

// @Description 添加取消函数
// @Auth shigx
// @Date 2021/12/2 11:09 下午
// @param
// @return
func addTaskCtrl(taskID int, cancelFunc context.CancelFunc) error {
	ctlMu.Lock()
	defer ctlMu.Unlock()

	if _, ok := ctlMap[taskID]; ok {
		return fmt.Errorf("duplicate taskID:%d", taskID)
	}

	ctlMap[taskID] = cancelFunc

	return nil
}
