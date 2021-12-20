// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/26 6:10 下午
package web

import (
	"github.com/gin-gonic/gin"
	"gspider/internal/pkg/response"
)

type Demo struct {
}

// @Description Demo页面
// @Auth shigx
// @Date 2021/11/26 6:12 下午
// @param
// @return
func (c *Demo) Demo(ctx *gin.Context) interface{} {
	// return response.Error(ctx, response.FailureCode, response.GetMessage(response.FailureCode))
	return response.Success(ctx, "Hello")
}
