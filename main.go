package main

import (
	"SimpleGOServer/service"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.New()
	service.InitDictionary()

	ginRouter.GET("/time", service.GetTimeEpoch)
	ginRouter.GET("/datetime", service.GetDateTime)
	ginRouter.GET("/localtime/:timezone", service.GetTimeTZ)

	ginRouter.Run()
}
