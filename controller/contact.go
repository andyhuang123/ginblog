package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContactView(c *gin.Context) {

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "hello gin ",
	})
}
