// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/30 11:11 下午
package spider

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"gspider/internal/constant/task"
	"gspider/internal/model/entity"
)

// 爬虫主结构体
type Spider struct {
	task  *Task
	retCh chan<- entity.TIS
}

// 实例化
func New(task *Task, retCh chan<- entity.TIS) *Spider {
	return &Spider{
		task:  task,
		retCh: retCh,
	}
}

// @Description 运行
// @Auth shigx
// @Date 2021/12/2 11:45 下午
// @param
// @return
func (s *Spider) Run() error {
	c, err := newCollector(s.task.TaskConfig)
	if err != nil {
		return err
	}
	nodesLen := len(s.task.Rule.Nodes)

	// 根据nodes数量初始化colly对象
	collectors := make([]*colly.Collector, 0, nodesLen)
	for i := 0; i < nodesLen; i++ {
		nextC := c.Clone()
		collectors = append(collectors, nextC)
	}

	// 根据nodes数量初始化带取消的上下文
	ctxCtl, cancel := context.WithCancel(context.Background())
	ctxs := make([]*Context, 0, nodesLen)
	for i := 0; i < nodesLen; i++ {
		var ctx *Context
		if i != nodesLen-1 {
			ctx = newContext(ctxCtl, cancel, s.task, collectors[i], collectors[i+1])
		} else {
			ctx = newContext(ctxCtl, cancel, s.task, collectors[i], nil)
		}
		ctxs = append(ctxs, ctx)

		addCallback(ctx, s.task.Rule.Nodes[i])
	}
	// 请求入口上下文
	headCtx := newContext(ctxCtl, cancel, s.task, c, collectors[0])

	headWrapper := func(ctx *Context) (err error) {
		defer func() {
			if e := recover(); e != nil {
				if v, ok := e.(error); ok {
					err = v
				} else {
					err = fmt.Errorf("%v", e)
				}
			}
		}()

		return s.task.Rule.Head(ctx)
	}
	if err := headWrapper(headCtx); err != nil {
		return err
	}
	if err := addTaskCtrl(s.task.ID, cancel); err != nil {
		return err
	}

	// 开启协程等待回调返回
	go func() {
		for i := 0; i < nodesLen; i++ {
			collectors[i].Wait()
			log.Infof("task:%s %d step completed...", s.task.Name, i+1)
		}
		CancelTask(s.task.ID)
		s.retCh <- entity.TIS{ID: s.task.ID, Status: task.StatusCompleted}

		log.Infof("task:%s run completed...", s.task.Name)
	}()

	return nil
}

func cbDefer(ctx *Context, info string) {
	if e := recover(); e != nil {
		log.Error(info, fmt.Sprintf(", err: %+v", e))
		ctx.ctlCancel()
	}
}

// @Description 添加回调
// @Auth shigx
// @Date 2021/12/2 11:45 下午
// @param
// @return
func addCallback(ctx *Context, node *Node) {
	if node.OnRequest != nil {
		ctx.c.OnRequest(func(request *colly.Request) {
			defer cbDefer(ctx, fmt.Sprintf("OnRequest unexcepted exited, url:%s", request.URL.String()))
			newCtx := ctx.cloneWithReq(request)
			select {
			case <-newCtx.ctlCtx.Done():
				newCtx.Abort()
				return
			default:

			}
			node.OnRequest(newCtx, newRequest(request, newCtx))
		})
	}

	if node.OnError != nil {
		ctx.c.OnError(func(response *colly.Response, e error) {
			defer cbDefer(ctx, fmt.Sprintf("OnError unexcepted exited, url:%s", response.Request.URL.String()))
			newCtx := ctx.cloneWithReq(response.Request)
			select {
			case <-newCtx.ctlCtx.Done():
				log.Warnf("request has ben canceled in OnError, url:%s", newCtx.GetRequest().URL.String())
				return
			default:

			}
			err := node.OnError(newCtx, newResponse(response, newCtx), e)
			if err != nil {
				log.Errorf("node.OnError return err:%v, request url:%s", err, response.Request.URL.String())
			}
		})
	}

	if node.OnResponse != nil {
		ctx.c.OnResponse(func(response *colly.Response) {
			defer cbDefer(ctx, fmt.Sprintf("OnResponse unexecpted exited, url:%s", response.Request.URL.String()))
			newCtx := ctx.cloneWithReq(response.Request)
			select {
			case <-newCtx.ctlCtx.Done():
				log.Warnf("request has been canceled in OnResponse, url:%s", newCtx.GetRequest().URL.String())
				return
			default:

			}
			err := node.OnResponse(newCtx, newResponse(response, newCtx))
			if err != nil {
				log.Errorf("node.OnResponse return err:%+v, request url:%s", err, response.Request.URL.String())
			}
		})
	}

	if node.OnHTML != nil {
		for selector, fn := range node.OnHTML {
			ctx.c.OnHTML(selector, func(element *colly.HTMLElement) {
				defer cbDefer(ctx, fmt.Sprintf("OnHTML unexcepted exited, selector:%s, url:%s", selector, element.Request.URL.String()))
				newCtx := ctx.cloneWithReq(element.Request)
				select {
				case <-newCtx.ctlCtx.Done():
					log.Warnf("request has been canceled in OnHTML, url:%s", newCtx.GetRequest().URL.String())
					return
				default:
				}

				err := fn(newCtx, newHTMLElement(element, newCtx))
				if err != nil {
					log.Errorf("node.OnHTML:%s return err:%+v, request url:%s", selector, err, element.Request.URL.String())
				}
			})
		}
	}

	if node.OnXML != nil {
		for selector, fn := range node.OnXML {
			ctx.c.OnXML(selector, func(element *colly.XMLElement) {
				defer cbDefer(ctx, fmt.Sprintf("OnXML unexcepted exited, selector:%s, url:%s", selector, element.Request.URL.String()))

				newCtx := ctx.cloneWithReq(element.Request)
				select {
				case <-newCtx.ctlCtx.Done():
					log.Warnf("request has been canceled in OnXML, url:%s", newCtx.GetRequest().URL.String())
					return
				default:
				}

				err := fn(newCtx, newXMLElement(element, newCtx))
				if err != nil {
					log.Errorf("node.OnXML:%s return err:%+v, request url:%s", selector, err, element.Request.URL.String())
				}
			})
		}
	}

	if node.OnScraped != nil {
		ctx.c.OnScraped(func(response *colly.Response) {
			defer cbDefer(ctx, fmt.Sprintf("OnScraped unexcepted exited, url:%s", response.Request.URL.String()))

			newCtx := ctx.cloneWithReq(response.Request)
			select {
			case <-newCtx.ctlCtx.Done():
				log.Warnf("request has been canceled in OnScraped, url:%s", newCtx.GetRequest().URL.String())
				return
			default:
			}

			err := node.OnScraped(newCtx, newResponse(response, newCtx))
			if err != nil {
				log.Errorf("node.OnScraped return err:%+v, request url:%s", err, response.Request.URL.String())
			}
		})
	}
}
