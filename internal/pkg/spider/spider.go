// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/30 11:11 下午
package spider

import (
	"fmt"
	"github.com/gocolly/colly/v2"
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
	collectors = append(collectors, c)
	for i := 1; i < nodesLen; i++ {
		nextC := c.Clone()
		collectors = append(collectors, nextC)
	}

	// 注册回调函数
	for i := 0; i < nodesLen; i++ {
		var currentC *colly.Collector
		var nextC *colly.Collector
		currentC = collectors[i]
		if i != nodesLen-1 {
			nextC = collectors[i+1]
		} else {
			nextC = nil
		}

		addCallback(currentC, nextC, s.task.Rule.Nodes[i])
	}

	enterWrapper := func(c *colly.Collector) (err error) {
		defer func() {
			if e := recover(); e != nil {
				if v, ok := e.(error); ok {
					err = v
				} else {
					err = fmt.Errorf("%v", e)
				}
			}
		}()

		return s.task.Rule.EnterFun(c)
	}
	if err = enterWrapper(c); err != nil {
		return err
	}

	// 开启协程等待回调返回
	go func() {
		for i := 0; i < nodesLen; i++ {
			collectors[i].Wait()
			log.Infof("task:%s %d step completed...", s.task.Name, i+1)
		}
		s.retCh <- entity.TIS{ID: s.task.ID, Status: task.StatusCompleted}

		log.Infof("task:%s run completed...", s.task.Name)
	}()

	return nil
}

// @Description 添加回调
// @Auth shigx
// @Date 2021/12/2 11:45 下午
// @param
// @return
func addCallback(currentC *colly.Collector, nextC *colly.Collector, node *Node) {
	if node.OnRequest != nil {
		currentC.OnRequest(func(request *colly.Request) {
			node.OnRequest(request)
		})
	}

	if node.OnError != nil {
		currentC.OnError(func(response *colly.Response, e error) {
			err := node.OnError(response, e)
			if err != nil {
				log.Errorf("node.OnError return err:%v, request url:%s", err, response.Request.URL.String())
			}
		})
	}

	if node.OnResponse != nil {
		currentC.OnResponse(func(response *colly.Response) {
			err := node.OnResponse(response, nextC)
			if err != nil {
				log.Errorf("node.OnResponse return err:%+v, request url:%s", err, response.Request.URL.String())
			}
		})
	}

	if node.OnHTML != nil {
		for selector, fn := range node.OnHTML {
			currentC.OnHTML(selector, func(element *colly.HTMLElement) {
				err := fn(element, nextC)
				if err != nil {
					log.Errorf("node.OnHTML:%s return err:%+v, request url:%s", selector, err, element.Request.URL.String())
				}
			})
		}
	}

	if node.OnXML != nil {
		for selector, fn := range node.OnXML {
			currentC.OnXML(selector, func(element *colly.XMLElement) {
				err := fn(element, nextC)
				if err != nil {
					log.Errorf("node.OnXML:%s return err:%+v, request url:%s", selector, err, element.Request.URL.String())
				}
			})
		}
	}

	if node.OnScraped != nil {
		currentC.OnScraped(func(response *colly.Response) {
			err := node.OnScraped(response)
			if err != nil {
				log.Errorf("node.OnScraped return err:%+v, request url:%s", err, response.Request.URL.String())
			}
		})
	}
}
