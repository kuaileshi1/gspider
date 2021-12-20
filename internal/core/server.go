// @Title Web Server 处理
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/25 2:24 下午
package core

import (
	"errors"
	"fmt"
	"github.com/fvbock/endless"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// 服务启动相关配置
type ServerConfig struct {
	Port           int
	PidDir         string
	ReadTimeOut    time.Duration
	WriteTimeOut   time.Duration
	MaxHeaderBytes int
	HttpHandler    http.Handler
}

// @Description 服务启动参数校验并设置默认值
// @Auth shigx
// @Date 2021/11/25 2:36 下午
// @param
// @return
func (sc *ServerConfig) initServerConfig() {
	if sc.Port == 0 {
		sc.Port = 80
	}
	if sc.PidDir == "" {
		sc.PidDir = filepath.Dir(os.Args[0])
	}
	if sc.ReadTimeOut == 0 {
		sc.ReadTimeOut = 60
	}
	if sc.WriteTimeOut == 0 {
		sc.WriteTimeOut = 60
	}
	if sc.MaxHeaderBytes == 0 {
		sc.MaxHeaderBytes = 1 << 20
	}
}

// @Description 启动Web服务
// @Auth shigx
// @Date 2021/11/25 2:51 下午
// @param config ServerConfig 服务配置
// @return error
func StartServer(config *ServerConfig) error {
	endless.DefaultReadTimeOut = config.ReadTimeOut * time.Second
	endless.DefaultWriteTimeOut = config.WriteTimeOut * time.Second
	endless.DefaultMaxHeaderBytes = config.MaxHeaderBytes
	endPoint := fmt.Sprintf(":%d", config.Port)

	server := endless.NewServer(endPoint, config.HttpHandler)
	server.BeforeBegin = func(add string) {
		f, err := os.OpenFile(filepath.Join(config.PidDir, "/PID"), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
		defer f.Close()
		if err != nil {
			panic(fmt.Sprintf("Fail to OpenFile: %v", err))
		}
		fmt.Fprintln(f, syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		return errors.New(fmt.Sprintf("Server err: %v", err))
	}

	return nil
}
