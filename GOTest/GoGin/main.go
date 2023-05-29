package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("111")
	getGinServer()
}

func getGinServer() {
	//创建一个服务
	ginServer := gin.Default()

	//响应页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		// context.JSON(200, gin.H{
		// 	"msg": "这是go后台传来的"}) //返回json数据
		context.HTML(200, "index.html", gin.H{
			"msg": "这是go后台传来的",
		})
	})

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,world GET"})
	})

	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,world POST"})
	})
	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,world PUT"})
	})
	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,world DELETE"})
	})
	//服务器端口
	ginServer.Run(":8082")
}
