package web

import (
	"gin-blog/news/controller"
	"github.com/gin-gonic/gin"
)

func WebUserRouter(Router *gin.RouterGroup) {

	UserRouter := Router.Group("home")
	{
		UserRouter.GET("index", controller.Home)
	}
}

