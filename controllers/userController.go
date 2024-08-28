package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"jcd-web/models"
)

type UserController struct {
	web.Controller
}

func (c *UserController) Get() {
	fmt.Println("查询USER数据.")
	o := orm.NewOrm()
	var users []models.User
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		c.Ctx.WriteString("Error: " + err.Error())
		fmt.Println("查询失败：", err.Error())
		return
	}
	c.Data["json"] = users
	c.ServeJSON()
}