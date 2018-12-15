package routers

import (
	"WEB/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:LoginFunc;post:PostFunc")

	beego.Router("/index", &controllers.IndexController{}, "get:GetFunc;post:PostFunc")
	beego.Router("/introduce", &controllers.IntroduceController{}, "get:GetFunc;post:PostFunc")
	beego.Router("/house_list.html", &controllers.HouseListController{}, "get:GetFunc;post:PostFunc")
	beego.Router("/house_edit.html", &controllers.HouseEditController{}, "get:GetFunc;post:PostFunc")
	beego.Router("/loupanchart.html", &controllers.LoupanchartController{}, "get:GetFunc;post:PostFunc")

	beego.Router("/user", &controllers.HomeController{}, "get:GetFunc;post:PostFunc")
	beego.Router("/login", &controllers.LoginController{}, "get:GetFunc;post:PostFunc")
}
