// @Title api接口返回统一处理
// @Description 接口成功失败返回封装
// @Author shigx 2021/11/26 5:15 下午
package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// 返回结构体
type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type HandlerFunc func(ctx *gin.Context) interface{}

// @Description 统一返回处理
// @Auth shigx
// @Date 2021/11/26 5:24 下午
// @param
// @return
func Wrap(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		res := handlerFunc(c)
		switch repType := res.(type) {
		case string:
			c.String(http.StatusOK, res.(string))
		case Resp:
			c.JSON(http.StatusOK, res.(Resp))
		default:
			log.Error(fmt.Sprintf("Response type not support. Got:[%T]", repType))
		}

		return
	}
}

// @Description 成功处理
// @Auth shigx
// @Date 2021/11/26 5:58 下午
// @param
// @return
func Success(ctx *gin.Context, data ...interface{}) Resp {
	return resp(ctx, SuccessCode, GetMessage(SuccessCode), data...)
}

// @Description 失败返回处理
// @Auth shigx
// @Date 2021/11/26 6:00 下午
// @param
// @return
func Error(ctx *gin.Context, code int, msg string, data ...interface{}) Resp {
	return resp(ctx, code, msg, data...)
}

// @Description 当状态码不确定是成功还是失败时调用，只是为了在调用时看着舒服些
// @Auth shigx
// @Date 2021/11/26 6:05 下午
// @param
// @return
func Common(ctx *gin.Context, code int, msg string, data ...interface{}) Resp {
	return resp(ctx, code, msg, data...)
}

// @Description 返回值处理
// @Auth shigx
// @Date 2021/11/26 5:56 下午
// @param
// @return
func resp(ctx *gin.Context, code int, msg string, data ...interface{}) Resp {
	resp := Resp{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	if len(data) == 1 {
		resp.Data = data[0]
	}
	// TODO 这里可以对返回做一些特殊处理，如加密。

	return resp
}
