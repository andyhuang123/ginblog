package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeView(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"active": "home",
	})
}
