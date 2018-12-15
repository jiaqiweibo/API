package controllers

import (
	"github.com/astaxie/beego"
)

type LoupanchartController struct {
	beego.Controller
}

func (c *LoupanchartController) GetFunc() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "loupanchart.html"

}
func (c *LoupanchartController) PostFunc() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "loupanchart.html"

}
