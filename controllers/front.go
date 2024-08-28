package controllers

import "github.com/beego/beego/v2/server/web"

type FrontController struct {
	web.Controller
}

func (c *FrontController) Get() {
	c.TplName = "front.html"
}