// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/30 10:22 下午
package spider

import (
	"gspider/internal/model/entity"
	"regexp"
	"strings"
	"time"
)

// @Description 获取爬虫任务配置
// @Auth shigx
// @Date 2021/12/2 10:47 下午
// @param
// @return
func GetSpiderTask(task *entity.Task) (*Task, error) {
	rule, err := GetTaskRule(task.TaskRuleName)
	if err != nil {
		return nil, err
	}
	var optAllowedDomains []string
	if task.OptAllowedDomains != "" {
		optAllowedDomains = strings.Split(task.OptAllowedDomains, ",")
	}
	var urlFiltersReg []*regexp.Regexp
	if task.OptUrlFilters != "" {
		urlFilters := strings.Split(task.OptUrlFilters, ",")
		for _, v := range urlFilters {
			reg, err := regexp.Compile(v)
			if err != nil {
				return nil, err
			}
			urlFiltersReg = append(urlFiltersReg, reg)
		}
	}

	config := TaskConfig{
		Option: Option{
			UserAgent:              task.OptUserAgent,
			MaxDepth:               task.OptMaxDepth,
			AllowedDomains:         optAllowedDomains,
			URLFilters:             urlFiltersReg,
			AllowURLRevisit:        rule.AllowURLRevisit,
			MaxBodySize:            task.OptMaxBodySize,
			IgnoreRobotsTxt:        rule.IgnoreRobotsTxt,
			InsecureSkipVerify:     rule.InsecureSkipVerify,
			ParseHTTPErrorResponse: rule.ParseHTTPErrorResponse,
			DisableCookies:         rule.DisableCookies,
		},
		Limit: Limit{
			Enable:      task.LimitEnable,
			DomainGlob:  task.LimitDomainGlob,
			Delay:       time.Duration(task.LimitDelay) * time.Millisecond,
			RandomDelay: time.Duration(task.LimitRandomDelay) * time.Millisecond,
			Parallelism: task.LimitParallelism,
		},
	}

	if task.OptRequestTimeout > 0 {
		config.Option.RequestTimeout = time.Duration(task.OptRequestTimeout) * time.Second
	}
	if urls := strings.TrimSpace(task.ProxyUrls); len(urls) > 0 {
		config.ProxyURLs = strings.Split(urls, ",")
	}

	return NewTask(task.ID, *rule, config), nil
}
