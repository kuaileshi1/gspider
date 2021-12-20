// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/26 6:46 下午
package entity

import (
	"time"
)

// 抓取内容存入表
type Cxytiandi struct {
	ID        int
	Title     string
	Cdate     string
	Visit     string
	Author    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*Cxytiandi) TableName() string {
	return "cxytiandi"
}
