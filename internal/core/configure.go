// @Title 应用配置处理
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/24 11:02 下午
package core

import (
	"flag"
	"fmt"
	"github.com/kuaileshi1/dbable"
	"github.com/kuaileshi1/redis"
	"github.com/spf13/viper"
	"os"
)

var configFile string

// 配置结构体
type Config struct {
	App    map[string]interface{}
	Server *ServerConfig
	Mysql  map[string]*dbable.MysqlConfig
	Redis  map[string]*redis.ConfigRedis
}

// @Description 初始化配置文件
// @Auth shigx
// @Date 2021/11/25 4:34 下午
// @param
// @return
func InitConfig() *Config {
	parseOSArg()

	viper := viper.New()

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file:%v", err))
	}

	reslut := new(Config)
	if err := viper.Unmarshal(&reslut); err != nil {
		panic(fmt.Sprintf("Viper Unmarshal err:%v", err))
	}
	reslut.Server.initServerConfig()

	return reslut
}

// @Description 解析启动命令配置参数
// @Auth shigx
// @Date 2021/11/25 4:33 下午
// @param
// @return
func parseOSArg() {
	flags := flag.NewFlagSet(App.Name, flag.ExitOnError)
	flags.StringVar(&configFile, "conf", "", "config file with whole path")
	flags.Parse(os.Args[1:])
}
