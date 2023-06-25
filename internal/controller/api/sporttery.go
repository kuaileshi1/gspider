// Package api
// @Description: 体育接口
// @Auth shigx 2023-04-28 10:36:06
package api

import (
	"github.com/gin-gonic/gin"
	"gspider/internal/pkg/response"
	"gspider/internal/service/apiservice"
)

type Sporttery struct {
	sportteryService *apiservice.SportteryService
}

func NewSporttery() *Sporttery {
	return &Sporttery{
		sportteryService: apiservice.NewSportteryService(),
	}
}

// @Description 任务列表
// @Auth shigx
// @Date 2021/12/13 11:37 下午
// @param
// @return
func (a Sporttery) JczqScoreList(c *gin.Context) interface{} {
	list := a.sportteryService.JczqScoreList()

	data := map[string]interface{}{
		"list": list,
	}

	return response.Success(c, data)
}
