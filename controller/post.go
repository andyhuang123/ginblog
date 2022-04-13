package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostView(c *gin.Context){

	c.HTML(http.StatusOK, "post.html", gin.H{
		"title": "hello gin ",
	})
}