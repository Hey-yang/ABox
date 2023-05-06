package main

import (
	"Por03/Test"
	"Test360.com/Por04/Test/Test01"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("你好")
	Test.NetTest()
	Test.NewTest[int32](123)
	Test.NewTest[string]("123")
	Test01.ErrString()
	return
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/pingPost", func(spost *gin.Context) {
		params := spost.Params
		fmt.Println(params)
		spost.JSON(200, gin.H{
			"messgae": "pingPost",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
