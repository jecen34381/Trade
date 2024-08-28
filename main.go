package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // 使用 MySQL 驱动
	_ "jcd-web/routers"
)

func init() {
	// 注册数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "zhangyonghe:@tcp(47.95.6.83:33061)/jcd?charset=utf8mb4")
	// 自动创建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	// 设置默认端口
	web.BConfig.Listen.HTTPPort = 8088
	web.Run()
}
