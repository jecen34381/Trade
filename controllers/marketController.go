package controllers

import "github.com/beego/beego/v2/server/web"

type MarketController struct {
	web.Controller
}

func (c *MarketController) Get() {
	c.Ctx.WriteString("最新行情！")
	//c.TplName = "index.html"
}