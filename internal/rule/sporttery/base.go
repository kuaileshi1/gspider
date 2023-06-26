package sporttery

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"gspider/internal/pkg/redis"
	"gspider/internal/pkg/spider"
	"gspider/internal/pkg/utils"
	"time"
)

func Init() {
	spider.Register(ruleScore)
	spider.Register(ruleSpfMatch)
}

func Retry(res *colly.Response, count int) error {
	key := fmt.Sprintf("err_req_%s", utils.Md5(res.Request.URL.String()))

	var et int
	var redisClient = redis.GetClient()
	et, err := redisClient.Get(context.Background(), key).Int()
	if err != redis.Nil && err != nil {
		log.Errorf("get redis key:%s err:%s", key, err.Error())
		return err
	}

	if et >= count {
		return fmt.Errorf("exceed %d counts", count)
	}

	log.Infof("errCount:%d, we will retry url:%s, after 1 second", et+1, res.Request.URL.String())
	time.Sleep(time.Second)

	redisClient.Incr(context.Background(), key)
	redisClient.Expire(context.Background(), key, time.Hour)

	res.Request.Retry()

	return nil
}
