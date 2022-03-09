package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HomeView(c *gin.Context){

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin " + strings.ToLower(c.Request.Method) + " method",
	})
}
