package router

import (
	"connector/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//Get gatewaymassage from manager and generate gateway.json at the same time
	r.POST("/gateway/data", controller.GenerateGatewayJSON)
	r.GET("/nginx_status", controller.GetNginxStatus)
	return r
}
