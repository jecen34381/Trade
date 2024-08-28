package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Username string `orm:"size(100)"`
	Password string `orm:"size(100)"`
}

func init() {
	orm.RegisterModel(new(User))
}