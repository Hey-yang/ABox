package goginserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `json:"user" binding:"required"`
	Pssword string `json:"password" binding:"required"`
}

func GetGinServerSan() *gin.Engine {
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

func GetGinServerEr() *gin.Engine {
	// 初始化 Gin 框架默认实例，该实例包含了路由、中间件以及配置信息
	r := gin.Default()
	// 路由组1 ，处理GET请求
	v1 := r.Group("api/v1")
	// {} 是书写规范
	{
		v1.GET("/register", register)
		v1.GET("/login/:id", login)
	}
	v2 := r.Group("api/v2")
	{
		v2.POST("/register", register)
		v2.POST("/login", login)
	}
	return r
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	name1 := c.Query("name")
	name2 := c.Param("id")
	fmt.Println(c.Params)
	c.String(http.StatusOK, fmt.Sprintf("hello %s %s %s\n", name, name1, name2))
}

func register(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

// 通过字典模拟 DB
var db = make(map[string]string)

func GetGinServerYi() *gin.Engine {
	// 初始化 Gin 框架默认实例，该实例包含了路由、中间件以及配置信息
	r := gin.Default()

	// Ping 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 获取用户数据路由
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")

		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// 需要 HTTP 基本授权认证的子路由群组设置   foo:bar@localhost:8082/admin
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // 用户名:foo 密码:bar
		"manu": "123", // 用户名:manu 密码:123
	}))

	// 保存用户信息路由
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// 解析并验证 JSON 格式请求数据
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

// 简单的使用
func GetGinServer() *gin.Engine {
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
		context.JSON(http.StatusOK, gin.H{"msg": "hello,world GET"})
	})

	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "hello,world POST"})
	})
	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "hello,world PUT"})
	})
	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "hello,world DELETE"})
	})
	return ginServer
}
