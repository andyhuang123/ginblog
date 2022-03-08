package router

import (
	"gin-blog/news/controller"
	"github.com/gin-gonic/gin"
)

// InitBaseRouter 图形验证码的路由
func InitBaseRouter(Router *gin.RouterGroup) {

	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", controller.GetCaptcha)
	}

}
