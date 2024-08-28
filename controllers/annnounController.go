package controllers

import "github.com/beego/beego/v2/server/web"

type AnnounController struct {
	web.Controller
}

func (c *AnnounController) Get() {
	//c.Ctx.WriteString("最新公告！")
	c.TplName = "annon.html"
}