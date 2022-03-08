package initialize

import (
	"gin-blog/news/middlewares"
	"gin-blog/news/router"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//模板解析
	Router.LoadHTMLGlob("view/**/**")

	// 路由分组
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 设置跨域中间件
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	ViewGroup := Router.Group("/")
	router.WebUserRouter(ViewGroup)

	//注册pprof 相关路由
	pprof.Register(Router)

	return Router
}

