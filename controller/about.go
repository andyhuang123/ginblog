package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutView(c *gin.Context) {

	c.HTML(http.StatusOK, "about.html", gin.H{
		"title": "hello gin ",
	})
}
