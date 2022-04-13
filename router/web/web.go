package web

import (
	"gin-blog/news/controller"

	"github.com/gin-gonic/gin"
)

func WebUserRouter(Router *gin.RouterGroup) {

	HomeRouter := Router.Group("/")
	{
		HomeRouter.Any("", controller.HomeView)

		HomeRouter.GET("index.html", controller.HomeView)

		HomeRouter.GET("post.html", controller.PostView)

		HomeRouter.GET("category.html", controller.CategoryView)

		HomeRouter.GET("about.html", controller.AboutView)

		HomeRouter.GET("contact.html", controller.ContactView)

	}

}
