// @Title Web Server 处理
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/25 2:24 下午
package core

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gspider/api/proto/pb"
	"gspider/internal/service/rpcservice"
	"net"
)

// RpcServer配置
type RpcServerConfig struct {
	Port int
}

// @Description 服务启动参数校验并设置默认值
// @Auth shigx
// @Date 2021/11/25 2:36 下午
// @param
// @return
func (sc *RpcServerConfig) initServerConfig() {
	if sc.Port == 0 {
		sc.Port = 8888
	}
}

// @Description 启动RPC服务
// @Auth shigx
// @Date 2021/11/25 2:51 下午
// @param config ServerConfig 服务配置
// @return error
func StartRpcServer(config *RpcServerConfig) error {
	server := grpc.NewServer()
	port := fmt.Sprintf(":%d", config.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.New(fmt.Sprintf("Rpc server err:%v", err))
	}
	pb.RegisterTaskServer(server, &rpcservice.TaskRpcService{})
	reflection.Register(server)

	err = server.Serve(lis)
	if err != nil {
		return errors.New(fmt.Sprintf("Rpc server err:%v", err))
	}

	return nil
}
