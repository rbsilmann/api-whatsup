package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEndpoint(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"api says": "hey " + name + ", everything ok?",
	})
	return
}

func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"api says": "send a get request to the endpoint with your name!",
	})
}