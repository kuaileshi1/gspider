// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/1 6:01 下午
package job

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/rediskey"
	"gspider/internal/model"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/redis"
)

// @Description 程序员天地抓取内容入mysql
// @Auth shigx
// @Date 2021/12/2 3:24 下午
// @param
// @return
func CxyOutToMysql() {
	cxyModel := model.NewCxytiandiModel("default")

	inserData := make([]entity.Cxytiandi, 0)
	i := 0

	for {
		res, err := redis.GetClient().LPop(context.Background(), rediskey.CxytiandiOutKey).Bytes()
		if err != nil && !errors.Is(err, redis.Nil) {
			log.Errorf("cxytiandi get data from pop failed, err:%v", err)
			continue
		}
		if res == nil {
			break
		}
		var ret entity.Cxytiandi
		if err := json.Unmarshal(res, &ret); err != nil {
			log.Errorf("cxytiandi json unmarshal failed, err:%v", err)
			continue
		}
		if ret.Title != "" {
			inserData = append(inserData, ret)
			i++
		}
		if i >= 100 {
			cxyModel.BatchInsert(inserData)
			inserData = inserData[:0]
		}
	}
	if len(inserData) > 0 {
		cxyModel.BatchInsert(inserData)
	}
}
