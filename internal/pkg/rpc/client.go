// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/26 6:16 下午
package rpc

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gspider/api/proto/pb"
	"gspider/internal/core"
	"time"
)

type Client struct {
	conn *grpc.ClientConn
	cli  pb.TaskClient
}

// @Description 获取rpc client
// @Auth shigx
// @Date 2021/12/26 6:52 下午
// @param
// @return
func NewClient() *Client {
	rpcConfig := core.App.Config.Rpc
	port := fmt.Sprintf(":%d", rpcConfig.Port)
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("grpc dial err:%v", err)
		return nil
	}

	client := pb.NewTaskClient(conn)
	return &Client{
		conn: conn,
		cli:  client,
	}
}

// @Description 开启任务
// @Auth shigx
// @Date 2021/12/26 8:17 下午
// @param
// @return
func (c *Client) Start(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := c.cli.Start(ctx, &pb.Request{Request: int64(id)})
	if err != nil {
		log.Errorf("rpc client start err:%v", err)
	}
	if res.Code == "ok" {
		return true
	}

	return false
}

// @Description 停止任务
// @Auth shigx
// @Date 2021/12/26 8:33 下午
// @param
// @return
func (c *Client) Stop(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := c.cli.Stop(ctx, &pb.Request{Request: int64(id)})
	if err != nil {
		log.Errorf("rpc client start err:%v", err)
	}
	if res.Code == "ok" {
		return true
	}

	return false
}

// @Description 关闭rpc链接
// @Auth shigx
// @Date 2021/12/26 8:26 下午
// @param
// @return
func (c *Client) Close() {
	c.conn.Close()
}
