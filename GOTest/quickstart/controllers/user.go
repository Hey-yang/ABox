package controllers

import beego "github.com/beego/beego/v2/server/web"

type UserController struct {
	beego.Controller
}

type UserInfo struct {
	ID    int64
	Name  string
	Phone string
	Addr  string
}

// func (c *UserController) Get() {
// 	c.Ctx.WriteString("1231312")
// }

// func (c *UserController) Post() {
// 	c.Data["Website"] = "beego.vip"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }

func (c *UserController) GetUser() {
	user := UserInfo{ID: 1, Name: "user", Phone: "12333333", Addr: "123111232123123123"}
	c.Data["json"] = user
	c.ServeJSON()
	// c.Ctx.WriteString("GetUser")
}
