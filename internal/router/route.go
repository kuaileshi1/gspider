// @Title 路由处理
// @Description web页面路由处理逻辑
// @Author shigx 2021/11/25 3:36 下午
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"gspider/internal/controller/api"
	"gspider/internal/controller/web"
	"gspider/internal/core"
	"gspider/internal/middleware/cors"
	"gspider/internal/middleware/jwt"
	"gspider/internal/pkg/response"
)

func InitRouter() *gin.Engine {
	if mode := core.App.AppConfig("mode"); mode != "" {
		gin.SetMode(mode)
	}
	route := gin.New()
	route.Use(gin.Recovery())

	route.Use(cors.Cors())

	route.NoRoute(response.Wrap(func(c *gin.Context) interface{} {
		return response.Error(c, response.NotFoundCode, response.GetMessage(response.NotFoundCode))
	}))

	authMiddleware := jwt.GinJwtInt()

	route.POST("/login", authMiddleware.LoginHandler)

	webApi := route.Group("/webapi")
	webApi.Use(authMiddleware.MiddlewareFunc())
	{
		// 重新刷新token
		webApi.GET("/refresh_token", authMiddleware.RefreshHandler)

		user := new(web.User)
		webApi.GET("/user_info", response.Wrap(user.GetUserInfo))
		webApi.POST("logout", response.Wrap(user.Logout))

		task := web.NewTask()
		webApi.GET("/tasks", response.Wrap(task.List))
		webApi.POST("/tasks", response.Wrap(task.Save))
		webApi.GET("/tasks/:id", response.Wrap(task.GetTask))
		webApi.GET("/tasks/:id/stop", response.Wrap(task.Stop))
		webApi.GET("/tasks/:id/start", response.Wrap(task.Start))
		webApi.GET("/tasks/:id/delete", response.Wrap(task.Delete))
		webApi.GET("/tasks/rules", response.Wrap(task.Rules))

		demo := new(web.Demo)
		webApi.GET("/ping", response.Wrap(demo.Demo))
	}

	box := packr.New("dist box", "../dist")
	box1 := packr.New("static box", "../dist/static")
	route.StaticFS("/admin", box)
	route.StaticFS("/static", box1)

	// 对外接口定义
	apiGroup := route.Group("/api")
	sporttery := api.NewSporttery()
	apiGroup.GET("/sporttery/jczq_score_list", response.Wrap(sporttery.JczqScoreList))

	return route
}
