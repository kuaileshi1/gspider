// @Title 程序员天地文章抓取
// @Description 抓取业务逻辑处理
// @Author shigx 2021/11/30 11:39 下午
package cxytiandi

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
	"time"
)

// @Description 注册抓取规则
// @Auth shigx
// @Date 2021/12/2 3:35 下午
// @param
// @return
func Init() {
	spider.Register(rule)
}

// 规则定义
var rule = &spider.TaskRule{
	Name:        "程序员天地",
	Description: "技术文章抓取",
	Rule: &spider.Rule{
		Url: "http://cxytiandi.com/article",
		Nodes: map[int]*spider.Node{
			0: step1,
			1: step2,
		},
	},
}

// 列表爬取
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
		`.articleDiv`: func(el *colly.HTMLElement, nextC *colly.Collector) error {
			link := el.ChildAttr("a", "href")
			if link == "" {
				return nil
			}

			title := el.ChildText("a")
			visit := el.ChildText("font")
			ctx := colly.NewContext()
			ctx.Put("title", title)
			ctx.Put("visit", visit)

			return nextC.Request("GET", el.Request.AbsoluteURL(link), nil, ctx, nil)
		},
	},
}

// 文章内容抓取
var step2 = &spider.Node{
	OnRequest: func(req *colly.Request) {
		log.Infof("Visiting %s", req.URL.String())
	},
	OnError: func(res *colly.Response, err error) error {
		log.Errorf("Visiting failed! url:%s, err:%s", res.Request.URL.String(), err.Error())
		// 出错时重试三次
		return Retry(res, 3)
	},
	OnHTML: map[string]func(el *colly.HTMLElement, nextC *colly.Collector) error{
		`.homepage`: func(el *colly.HTMLElement, nextC *colly.Collector) error {
			var author string
			var cdate string
			el.ForEach(".post-meta span", func(i int, element *colly.HTMLElement) {
				switch i {
				case 0:
					author = element.Text
				case 1:
					cdate = element.Text
				}
			})

			content := el.ChildText("#article_content_str")
			article := entity.Cxytiandi{
				Title:     el.Request.Ctx.Get("title"),
				Cdate:     cdate,
				Visit:     el.Request.Ctx.Get("visit"),
				Author:    author,
				Content:   content,
				CreatedAt: time.Now(),
			}
			jsonData, err := json.Marshal(article)
			if err != nil {
				return err
			}

			if _, err := redis.GetClient().RPush(context.Background(), rediskey.CxytiandiOutKey, jsonData).Result(); err != nil {
				return err
			}
			redis.GetClient().Expire(context.Background(), rediskey.CxytiandiOutKey, time.Hour*24)

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
