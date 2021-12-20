// @Title 程序员天地文章抓取
// @Description 抓取业务逻辑处理
// @Author shigx 2021/11/30 11:39 下午
package cxytiandi

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/rediskey"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/redis"
	"gspider/internal/pkg/spider"
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
		Head: func(ctx *spider.Context) error {
			return ctx.VisitForNext("http://cxytiandi.com/article")
		},
		Nodes: map[int]*spider.Node{
			0: step1,
			1: step2,
		},
	},
}

// 列表爬取
var step1 = &spider.Node{
	OnRequest: func(ctx *spider.Context, req *spider.Request) {
		log.Infof("Visiting %s", req.URL.String())
	},
	OnError: func(ctx *spider.Context, res *spider.Response, err error) error {
		log.Errorf("Visiting failed! url:%s, err:%s", res.Request.URL.String(), err.Error())
		// 出错时重试三次
		return Retry(ctx, 3)
	},
	OnHTML: map[string]func(ctx *spider.Context, el *spider.HTMLElement) error{
		`.articleDiv`: func(ctx *spider.Context, el *spider.HTMLElement) error {
			link := el.ChildAttr("a", "href")
			if link == "" {
				return nil
			}

			title := el.ChildText("a")
			visit := el.ChildText("font")
			ctx.PutReqContextValue("title", title)
			ctx.PutReqContextValue("visit", visit)

			return ctx.VisitForNextWithContext(link)
		},
	},
}

// 文章内容抓取
var step2 = &spider.Node{
	OnRequest: func(ctx *spider.Context, req *spider.Request) {
		log.Infof("Visiting %s", req.URL.String())
	},
	OnError: func(ctx *spider.Context, res *spider.Response, err error) error {
		log.Errorf("Visiting failed! url:%s, err:%s", res.Request.URL.String(), err.Error())
		// 出错时重试三次
		return Retry(ctx, 3)
	},
	OnHTML: map[string]func(ctx *spider.Context, el *spider.HTMLElement) error{
		`.homepage`: func(ctx *spider.Context, el *spider.HTMLElement) error {
			var author string
			var cdate string
			el.ForEach(".post-meta span", func(i int, element *spider.HTMLElement) {
				switch i {
				case 0:
					author = element.Text
				case 1:
					cdate = element.Text
				}
			})

			content := el.ChildText("#article_content_str")
			article := entity.Cxytiandi{
				Title:     ctx.GetReqContextValue("title"),
				Cdate:     cdate,
				Visit:     ctx.GetReqContextValue("visit"),
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
func Retry(ctx *spider.Context, count int) error {
	req := ctx.GetRequest()
	key := fmt.Sprintf("err_req_%s", req.URL.String())

	var et int
	if errCount := ctx.GetAnyReqContextValue(key); errCount != nil {
		et = errCount.(int)
		if et >= count {
			return fmt.Errorf("exceed %d counts", count)
		}
	}
	log.Infof("errCount:%d, we wil retry url:%s, after 1 second", et+1, req.URL.String())
	time.Sleep(time.Second)
	ctx.PutReqContextValue(key, et+1)
	ctx.Retry()

	return nil
}
