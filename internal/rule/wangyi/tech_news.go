// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/3 11:36 上午
package wangyi

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/rediskey"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/redis"
	"gspider/internal/pkg/spider"
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
		Head: func(ctx *spider.Context) error {
			return ctx.VisitForNext("http://tech.163.com/gd/")
		},
		Nodes: map[int]*spider.Node{
			0: step1,
			1: step2,
		},
	},
}

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
		`.bigsize`: func(ctx *spider.Context, el *spider.HTMLElement) error {
			link := el.ChildAttr("a", "href")
			if link == "" {
				return nil
			}

			title := el.ChildText("a")
			ctx.PutReqContextValue("title", title)

			return ctx.VisitForNextWithContext(link)
		},
	},
}

var step2 = &spider.Node{
	OnRequest: func(ctx *spider.Context, req *spider.Request) {
		log.Infof("Visiting %s", req.URL.String())
	},
	OnHTML: map[string]func(ctx *spider.Context, el *spider.HTMLElement) error{
		`#container`: func(ctx *spider.Context, el *spider.HTMLElement) error {

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
				Title:     ctx.GetReqContextValue("title"),
				Cdate:     cdate,
				Content:   content,
				CreatedAt: time.Now(),
			}
			jsonData, err := json.Marshal(article)
			if err != nil {
				return err
			}

			if _, err := redis.GetClient().RPush(context.Background(), rediskey.WangyiTechOutKey, jsonData).Result(); err != nil {
				return err
			}
			redis.GetClient().Expire(context.Background(), rediskey.WangyiTechOutKey, time.Hour*24)

			return nil
		},
	},
}

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
