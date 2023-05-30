package routers

import (
	"quickstart/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/user/GetUser", &controllers.UserController{}, "*:GetUser")

	// 这个会自动匹配相应的接口
	beego.AutoRouter(&controllers.UserController{})
}
