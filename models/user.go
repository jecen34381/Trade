package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id   int64  `orm:"auto"` // 自增ID
	Name string `orm:"size(100)"`
	Age  int
}

func init() {
	orm.RegisterModel(new(User))
}
