// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/26 6:46 下午
package entity

import (
	"time"
)

// 用户表
type User struct {
	ID           uint64    `json:"id"`
	Username     string    `json:"username"`
	TaskRuleName string    `json:"task_rule_name"`
	Password     string    `json:"-"`
	Roles        string    `json:"roles"`
	Introduction string    `json:"introduction"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (*User) TableName() string {
	return "user"
}
