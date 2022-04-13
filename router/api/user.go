package api

import (
	"gin-blog/news/controller"
	"gin-blog/news/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {

	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)

		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), controller.GetUserList)

		UserRouter.GET("info", middlewares.ApiSignMIddleware(), controller.Info)

	}
}
