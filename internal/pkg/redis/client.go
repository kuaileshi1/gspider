// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/1 3:28 下午
package redis

import (
	"github.com/kuaileshi1/redis"
)

var (
	Client = new(client)
	Nil    = redis.Nil
)

// redis客户端
type client struct {
	defaultClient redis.SplitClient
}

// @Description 初始化
// @Auth shigx
// @Date 2021/12/1 3:37 下午
// @param
// @return
func (c *client) Init() (err error) {
	return c.getClient()
}

// @Description 获取redis链接
// @Auth shigx
// @Date 2021/12/1 3:43 下午
// @param
// @return
func (c *client) getClient() error {
	var err error
	Client.defaultClient, err = redis.GetClient("default")
	if err != nil {
		// 出错重连一次
		Client.defaultClient, err = redis.GetClient("default")
		if err != nil {
			return err
		}
	}

	return nil
}

// @Description 获取redis链接
// @Auth shigx
// @Date 2021/12/1 3:43 下午
// @param
// @return
func GetClient() redis.SplitClient {
	return Client.defaultClient
}
