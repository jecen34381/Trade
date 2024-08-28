package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	//c.Ctx.WriteString("Hello, Beego!")
	c.TplName = "index.html"
}