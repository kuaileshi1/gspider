// @Title 规则初始化文件
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/1 9:34 上午
package rule

import (
	"gspider/internal/rule/cxytiandi"
	"gspider/internal/rule/wangyi"
)

// @Description 初始化抓取规则
// @Auth shigx
// @Date 2021/12/2 3:37 下午
// @param
// @return
func init() {
	cxytiandi.Init()
	wangyi.Init()
}
