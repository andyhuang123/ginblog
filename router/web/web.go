package web

import (
	"gin-blog/news/controller"
	"github.com/gin-gonic/gin"
)

func WebUserRouter(Router *gin.RouterGroup) {


	index := Router.Group("/")
	{
		index.Any("", controller.HomeView)
	}

	UserRouter := Router.Group("home")
	{
		UserRouter.GET("index", controller.HomeView)
	}
}

