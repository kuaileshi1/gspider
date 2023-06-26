// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/30 10:08 上午
package spider

import (
	"crypto/tls"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"net/http"
	"regexp"
	"time"
)

// Task
// @Description: 爬虫任务结构体
type Task struct {
	ID int
	TaskRule
	TaskConfig
}

// NewTask
// @Description: 实例化任务
// @Auth shigx 2023-06-26 15:39:09
// @param id 任务id
// @param rule 任务规则
// @param config 任务配置
// @return *Task 任务实例
func NewTask(id int, rule TaskRule, config TaskConfig) *Task {
	return &Task{
		ID:         id,
		TaskRule:   rule,
		TaskConfig: config,
	}
}

// 任务配置
type TaskConfig struct {
	Option    Option
	Limit     Limit
	ProxyURLs []string
}

// colly Option配置
type Option struct {
	UserAgent              string
	MaxDepth               int
	AllowedDomains         []string
	URLFilters             []*regexp.Regexp
	AllowURLRevisit        bool
	MaxBodySize            int
	IgnoreRobotsTxt        bool
	InsecureSkipVerify     bool
	ParseHTTPErrorResponse bool
	DisableCookies         bool
	RequestTimeout         time.Duration
}

// colly limit配置
type Limit struct {
	Enable       bool
	DomainRegexp string
	DomainGlob   string
	Delay        time.Duration
	RandomDelay  time.Duration
	Parallelism  int
}

// @Description 实例化colly
// @Auth shigx
// @Date 2021/12/2 11:27 下午
// @param
// @return
func newCollector(config TaskConfig) (*colly.Collector, error) {
	opts := make([]colly.CollectorOption, 0)

	opts = append(opts, colly.Async(true))
	if config.Option.MaxDepth > 1 {
		opts = append(opts, colly.MaxDepth(config.Option.MaxDepth))
	}

	if len(config.Option.AllowedDomains) > 0 {
		opts = append(opts, colly.AllowedDomains(config.Option.AllowedDomains...))
	}

	if config.Option.AllowURLRevisit {
		opts = append(opts, colly.AllowURLRevisit())
	}
	if config.Option.IgnoreRobotsTxt {
		opts = append(opts, colly.IgnoreRobotsTxt())
	}
	if config.Option.MaxBodySize > 0 {
		opts = append(opts, colly.MaxBodySize(config.Option.MaxBodySize))
	}
	if config.Option.UserAgent != "" {
		opts = append(opts, colly.UserAgent(config.Option.UserAgent))
	}
	if config.Option.ParseHTTPErrorResponse {
		opts = append(opts, colly.ParseHTTPErrorResponse())
	}
	if len(config.Option.URLFilters) > 0 {
		opts = append(opts, colly.URLFilters(config.Option.URLFilters...))
	}

	c := colly.NewCollector(opts...)
	if config.Option.DisableCookies {
		c.DisableCookies()
	}

	if len(config.ProxyURLs) > 0 {
		rp, err := proxy.RoundRobinProxySwitcher(config.ProxyURLs...)
		if err != nil {
			return nil, err
		}
		c.SetProxyFunc(rp)
	}
	if config.Limit.Enable {
		var limit colly.LimitRule
		if config.Limit.Delay > 0 {
			limit.Delay = config.Limit.Delay
		}
		if config.Limit.DomainGlob != "" {
			limit.DomainGlob = config.Limit.DomainGlob
		} else {
			limit.DomainGlob = "*"
		}

		if config.Limit.DomainRegexp != "" {
			limit.DomainRegexp = config.Limit.DomainRegexp
		}
		if config.Limit.Parallelism > 0 {
			limit.Parallelism = config.Limit.Parallelism
		}
		if config.Limit.RandomDelay > 0 {
			limit.RandomDelay = config.Limit.RandomDelay
		}

		c.Limit(&limit)

	}
	if config.Option.RequestTimeout > 0 {
		c.SetRequestTimeout(config.Option.RequestTimeout)
	}
	if config.Option.InsecureSkipVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.WithTransport(tr)
	}

	return c, nil
}
