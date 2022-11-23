package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rbsilmann/api-whatsup/controller"
)

func StartServer() {
	r := gin.Default()
	r.GET("/", controller.GetExample)
	r.GET("/:name", controller.GetEndpoint)
	r.Run(":9098")
}
