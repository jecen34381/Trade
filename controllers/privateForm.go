package controllers

import "github.com/beego/beego/v2/server/web"

type PrivateFormController struct {
	web.Controller
}

func (c *PrivateFormController) Get() {
	//c.Ctx.WriteString("Hello, Beego!")
	c.TplName = "privateForm.html"
}