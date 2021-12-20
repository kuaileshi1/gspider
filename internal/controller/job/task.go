// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/29 11:09 上午
package job

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/rediskey"
	"gspider/internal/constant/task"
	"gspider/internal/pkg/redis"
	"gspider/internal/service/jobservice"
)

type jsonData struct {
	ID     int
	Status task.Status
}

// @Description 管理任务
// @Auth shigx
// @Date 2021/12/17 3:01 下午
// @param
// @return
func MangeTask() {
	for {
		res, err := redis.GetClient().LPop(context.Background(), rediskey.TaskStatusChangeKey).Bytes()
		if err != nil && !errors.Is(err, redis.Nil) {
			log.Errorf("MangeTask get data from pop failed, data:%v, err:%v", res, err)
			continue
		}
		if res == nil {
			break
		}
		var ret jsonData
		if err := json.Unmarshal(res, &ret); err != nil {
			log.Errorf("MangeTask json unmarshal failed, data:%v, err:%v", res, err)
			continue
		}

		switch ret.Status {
		case task.StatusRunning:
			// 开启任务
			startTask(ret)
			break
		case task.StatusStopped:
			// 停止任务
			cronService := jobservice.GetCronService(ret.ID)
			// 如果任务未启动,则不需要停止
			if cronService == nil {
				break
			}

			cronService.Stop()
			break
		}
	}
}

// @Description 开启定时任务
// @Auth shigx
// @Date 2021/12/17 3:01 下午
// @param
// @return
func startTask(ret jsonData) {
	// 开启任务
	cronService := jobservice.GetCronService(ret.ID)
	// 如果任务已启动,则不需要重新启动
	if cronService != nil {
		return
	}

	service := jobservice.NewTaskService()
	task, err := service.GetTaskById(ret.ID)

	if err != nil {
		log.Errorf("MangeTask get task failed, data:%v, err:%v", ret, err)
		return
	}

	if err := service.CreateCronTask(*task); err != nil {
		log.Errorf("MangeTask CreateCronTask failed, data:%v, err:%v", ret, err)
		return
	}
}
