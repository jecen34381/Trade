package controllers

import "github.com/beego/beego/v2/server/web"

type ReminderController struct {
	web.Controller
}

func (c *ReminderController) Get() {
	c.Ctx.WriteString("重要提醒!")
	//c.TplName = "index.html"
}