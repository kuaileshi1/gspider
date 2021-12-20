// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/12/12 10:30 下午
package jwt

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gspider/internal/model"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/response"
	"log"
	"time"
)

func GinJwtInt() (authMiddleware *jwt.GinJWTMiddleware) {
	type login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte("gspider"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "identity",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 登录期间回调函数
			if v, ok := data.(*entity.User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var param login
			if err := c.ShouldBind(&param); err != nil {
				return nil, errors.New("用户名或密码不能为空")
			}
			user, err := model.NewUserModel("default").IsValidUser(param.Username, param.Password)
			if err != nil {
				return nil, errors.New("用户名或密码错误")
			}
			return user, nil
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, response.Resp{
				Code: code,
				Data: nil,
				Msg:  message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, response.Resp{
				Code: code,
				Data: map[string]string{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				},
				Msg: response.GetMessage(code),
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie:jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		SendCookie:    true,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return
}
