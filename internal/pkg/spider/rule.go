// @Title 爬虫任务规则文件
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/30 9:14 上午
package spider

import (
	"errors"
	"github.com/gocolly/colly/v2"
)

// 规则哈希map
var rules = make(map[string]*TaskRule)

// @Description 注册规则
// @Auth shigx
// @Date 2021/12/2 11:02 下午
// @param
// @return
func Register(rule *TaskRule) {
	if err := checkRule(rule); err != nil {
		panic(err)
	}

	rules[rule.Name] = rule
}

// @Description 根据规则名称获取规则
// @Auth shigx
// @Date 2021/12/2 11:01 下午
// @param
// @return
func GetTaskRule(ruleName string) (*TaskRule, error) {
	if rule, ok := rules[ruleName]; ok {
		return rule, nil
	}

	return nil, errors.New("task rule not exist")
}

// @Description 获取所有规则名称
// @Auth shigx
// @Date 2021/12/2 11:01 下午
// @param
// @return
func GetTaskRuleKeys() []string {
	keys := make([]string, 0, len(rules))
	for k := range rules {
		keys = append(keys, k)
	}

	return keys
}

// 任务规则
type TaskRule struct {
	Name                   string
	Description            string
	DisableCookies         bool // 是否禁用cookie
	AllowURLRevisit        bool // 是否允许重复抓取
	IgnoreRobotsTxt        bool
	InsecureSkipVerify     bool
	ParseHTTPErrorResponse bool
	Rule                   *Rule
}

// 回调具体实现
type Rule struct {
	EnterFun func(c *colly.Collector) error // 入口函数
	Nodes    map[int]*Node                  // 节点列表
}

// 页面回调函数
type Node struct {
	OnRequest  func(request *colly.Request)
	OnError    func(response *colly.Response, e error) error
	OnResponse func(response *colly.Response, nextC *colly.Collector) error
	OnHTML     map[string]func(e *colly.HTMLElement, nextC *colly.Collector) error
	OnXML      map[string]func(e *colly.XMLElement, nextC *colly.Collector) error
	OnScraped  func(response *colly.Response) error
}

// @Description 规则检查
// @Auth shigx
// @Date 2021/12/2 10:58 下午
// @param
// @return
func checkRule(rule *TaskRule) error {
	if rule == nil || rule.Rule == nil {
		return errors.New("task rule is nil")
	}
	if rule.Name == "" {
		return errors.New("task rule name is empty")
	}
	if rule.Rule.EnterFun == nil {
		return errors.New("task rule enterFun is nil")
	}
	if len(rule.Rule.Nodes) == 0 {
		return errors.New("task rule nodes len is invalid")
	}
	for i := 0; i < len(rule.Rule.Nodes); i++ {
		if _, ok := rule.Rule.Nodes[i]; !ok {
			return errors.New("task rule nodes key should start from 0 and monotonically increasing")
		}
	}
	if _, ok := rules[rule.Name]; ok {
		return errors.New("task rule name duplicated")
	}

	return nil
}
