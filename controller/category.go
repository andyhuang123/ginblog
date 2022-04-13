package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryView(c *gin.Context) {

	c.HTML(http.StatusOK, "category.html", gin.H{
		"active": "category",
	})
}
