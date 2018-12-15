package controllers

import (
	"github.com/astaxie/beego"
)

type HouseListController struct {
	beego.Controller
}

func (c *HouseListController) GetFunc() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "house_list.html"

}
func (c *HouseListController) PostFunc() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "house_list.html"

}
