package goginserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGinServerSi() *gin.Engine {
	// 初始化 Gin 框架默认实例，该实例包含了路由、中间件以及配置信息
	r := gin.Default()
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// new结构体，给结构体分配内存
		login := new(Login)
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&login); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(*login)
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	return r
}
