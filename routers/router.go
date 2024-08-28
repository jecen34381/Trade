package routers

import (
	"github.com/beego/beego/v2/server/web"
	"jcd-web/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/jcd/market", &controllers.MarketController{})
	web.Router("/jcd/announcement", &controllers.AnnounController{})
	web.Router("/jcd/reminder", &controllers.ReminderController{})
	web.Router("/jcd/private", &controllers.PrivateFormController{})
	web.Router("/jcd/front", &controllers.FrontController{})

	web.Router("/jcd/sql/user", &controllers.UserController{})
}