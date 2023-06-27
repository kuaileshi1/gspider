// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/3 11:36 上午
package wangyi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/rediskey"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/redis"
	"gspider/internal/pkg/spider"
	"gspider/internal/pkg/utils"
	"strings"
	"time"
)

func Init() {
	spider.Register(rule)
}

var rule = &spider.TaskRule{
	Name:        "网易科技新闻",
	Description: "科技新闻抓取",
	Rule: &spider.Rule{
		EnterFun: func(c *colly.Collector) error {
			return c.Visit("http://tech.163.com/gd/")
		},
		Nodes: map[int]*spider.Node{
			0: step1,
			1: step2,
		},
	},
}

var step1 = &spider.Node{
	OnRequest: func(req *colly.Request) {
		log.Infof("Visiting %s", req.URL.String())
	},
	OnError: func(res *colly.Response, err error) error {
		log.Errorf("Visiting failed! url:%s, err:%s", res.Request.URL.String(), err.Error())
		// 出错时重试三次
		return Retry(res, 3)
	},
	OnHTML: map[string]func(el *colly.HTMLElement, nextC *colly.Collector) error{
		`.bigsize`: func(el *colly.HTMLElement, nextC *colly.Collector) error {
			link := el.ChildAttr("a", "href")
			if link == "" {
				return nil
			}

			title := el.ChildText("a")
			ctx := colly.NewContext()
			ctx.Put("title", title)

			return nextC.Request("GET", el.Request.AbsoluteURL(link), nil, ctx, nil)
		},
	},
}

var step2 = &spider.Node{
	OnRequest: func(req *colly.Request) {
		log.Infof("Visiting %s", req.URL.String())
	},
	OnHTML: map[string]func(el *colly.HTMLElement, nextC *colly.Collector) error{
		`#container`: func(el *colly.HTMLElement, nextC *colly.Collector) error {

			var cdate string
			ptime := el.ChildText(".post_info")
			if ptime != "" {
				ptime := strings.Split(ptime, " ")
				if len(ptime) > 0 {
					cdate = ptime[0]
				}
			}
			content := el.ChildText(".post_body")

			article := entity.WangyiTechNews{
				Title:     el.Request.Ctx.Get("title"),
				Cdate:     cdate,
				Content:   content,
				CreatedAt: time.Now(),
			}
			jsonData, err := json.Marshal(article)
			if err != nil {
				return err
			}

			fmt.Println(article)

			if _, err := redis.GetClient().RPush(context.Background(), rediskey.WangyiTechOutKey, jsonData).Result(); err != nil {
				return err
			}
			redis.GetClient().Expire(context.Background(), rediskey.WangyiTechOutKey, time.Hour*24)

			return nil
		},
	},
}

// @Description 错误重试
// @Auth shigx
// @Date 2021/12/3 12:05 上午
// @param
// @return
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
