// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/26 3:48 下午
package rpcservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gspider/api/proto/pb"
	"gspider/internal/service/jobservice"
)

type TaskRpcService struct {
}

// @Description 开启任务
// @Auth shigx
// @Date 2021/12/26 10:31 下午
// @param
// @return
func (s *TaskRpcService) Start(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	taskId := int(r.GetRequest())
	// 开启任务
	cronService := jobservice.GetCronService(taskId)
	// 如果任务已启动,则不需要重新启动
	if cronService != nil {
		return &pb.Response{Code: "ok"}, nil
	}

	service := jobservice.NewTaskService()
	task, err := service.GetTaskById(taskId)

	if err != nil {
		log.Errorf("MangeTask get task failed, data:%v, err:%v", r.GetRequest(), err)
		return &pb.Response{Code: "fail"}, err
	}

	if err := service.CreateCronTask(*task); err != nil {
		log.Errorf("MangeTask CreateCronTask failed, data:%v, err:%v", r.GetRequest(), err)
		return &pb.Response{Code: "fail"}, err
	}
	return &pb.Response{Code: "ok"}, nil
}

// @Description 停止任务
// @Auth shigx
// @Date 2021/12/26 10:34 下午
// @param
// @return
func (s *TaskRpcService) Stop(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	cronService := jobservice.GetCronService(int(r.GetRequest()))
	// 如果任务未启动,则不需要停止
	if cronService == nil {
		return &pb.Response{Code: "ok"}, nil
	}

	cronService.Stop()
	return &pb.Response{Code: "ok"}, nil
}
