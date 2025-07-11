package main

import (
	"form/gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	routers.InitRouters(r)
	//r.GET("/ping", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.POST("/ping", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"message": "post pong",
	//	})
	//})
	r.Run(":8080")
}
