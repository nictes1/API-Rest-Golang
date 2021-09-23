package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nictes1/API-Rest-Golang/controller"
	"github.com/nictes1/API-Rest-Golang/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})
	server.Run(":8080")
}
