// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/26 6:46 下午
package entity

import (
	"gspider/internal/constant/task"
	"time"
)

// 任务表
type Task struct {
	ID                int         `json:"id"`
	TaskName          string      `json:"task_name"`
	TaskRuleName      string      `json:"task_rule_name"`
	TaskDesc          string      `json:"task_desc"`
	Status            task.Status `json:"status"`
	Counts            int         `json:"counts"`
	CronSpec          string      `json:"cron_spec"`
	OptUserAgent      string      `json:"opt_user_agent"`
	OptMaxDepth       int         `json:"opt_max_depth"`
	OptAllowedDomains string      `json:"opt_allowed_domains"`
	OptUrlFilters     string      `json:"opt_url_filters"`
	OptMaxBodySize    int         `json:"opt_max_body_size"`
	OptRequestTimeout int         `json:"opt_request_timeout"`
	LimitEnable       bool        `json:"limit_enable"`
	LimitDomainRegexp string      `json:"limit_domain_regexp"`
	LimitDomainGlob   string      `json:"limit_domain_glob"`
	LimitDelay        int         `json:"limit_delay"`
	LimitRandomDelay  int         `json:"limit_random_delay"`
	LimitParallelism  int         `json:"limit_parallelism"`
	ProxyUrls         string      `json:"proxy_urls"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func (*Task) TableName() string {
	return "task"
}

// 任务精简结构体
type TIS struct {
	ID     int
	Status task.Status
}
