package controllers

import (
	"WEB/models"
	"fmt"

	"github.com/astaxie/beego"
)

type HouseEditController struct {
	BaseController
}

func (c *HouseEditController) GetFunc() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "house_edit.html"

}
func (c *HouseEditController) PostFunc() {
	fmt.Printf("aaaaaaaaaaaaaaaaaaaaacontrollerName:%s\nactionName:%s\n", c.controllerName, c.actionName)
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	models.AddUser("vv", "vv")
	// models.UpdateUser("vv", "啊啊")
	//models.DeleteUserxvv")
	beego.Notice("aaaa擦擦擦")
	models.PoolRedis()

	c.TplName = "house_edit.html"

}
