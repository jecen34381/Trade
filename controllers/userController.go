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
	// 创建一个新的用户
	o := orm.NewOrm()
	user := models.User{Name: "张三", Age: 30}
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("插入失败:", err.Error())
	}

	// 查询用户
	var users []models.User
	_, err = o.QueryTable("user").All(&users)
	if err != nil {
		fmt.Println("查询失败:", err.Error())
	}

	fmt.Println(users[0].Name)

	c.Ctx.WriteString(users[0].Name)
}