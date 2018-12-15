package controllers

import (
	"github.com/astaxie/beego"
)

type IntroduceController struct {
	beego.Controller
}

func (c *IntroduceController) GetFunc() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "introduce.html"

}
func (c *IntroduceController) PostFunc() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "introduce.html"

}
