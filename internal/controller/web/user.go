// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/13 5:00 下午
package web

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gspider/internal/pkg/response"
	"strings"
)

type User struct {
}

// @Description 获取登录信息
// @Auth shigx
// @Date 2021/12/13 6:12 下午
// @param
// @return
func (w *User) GetUserInfo(c *gin.Context) interface{} {
	v, ok := c.Get(jwt.IdentityKey)
	if !ok {
		return response.Error(c, response.FailureCode, response.GetMessage(response.FailureCode))
	}

	data := v.(map[string]interface{})
	roles := data["roles"].(string)
	data["roles"] = strings.Split(roles, ",")

	return response.Success(c, data)
}

// @Description 退出操作
// @Auth shigx
// @Date 2021/12/13 6:13 下午
// @param
// @return
func (w *User) Logout(c *gin.Context) interface{} {
	return response.Success(c)
}
