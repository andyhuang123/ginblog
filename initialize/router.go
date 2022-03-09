package initialize

import (
	"gin-blog/news/middlewares"
	"gin-blog/news/router/api"
	"gin-blog/news/router/web"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	Router := gin.Default()
	//模板解析,静态资源
	Router.LoadHTMLGlob("view/**/**")
	Router.StaticFile("/favicon.ico", "./public/favicon.ico")
	Router.Static("/statics","./public/statics")

	// 路由分组
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 设置跨域中间件
	Router.Use(middlewares.Cors())

	Router.GET("/ws",api.WsClient)

	ApiGroup := Router.Group("/v1/")
	api.UserRouter(ApiGroup)
	api.InitBaseRouter(ApiGroup)


	ViewGroup := Router.Group("/")
	web.WebUserRouter(ViewGroup)



	//注册pprof 相关路由
	pprof.Register(Router)

	return Router
}

