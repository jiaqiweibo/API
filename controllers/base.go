package controllers

import (
	"fmt"

	"WEB/models"
	"strings"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	//o              orm.Ormer
	//isFirst        bool  //测试后发现不是全局变量
}

func (p *BaseController) Prepare() {

	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName) //MainController
	p.actionName = strings.ToLower(actionName)         //post:PostFunc
	fmt.Printf("controllerName:%s\nactionName:%s\n", controllerName, actionName)
	//p.o = orm.NewOrm()
	// if p.isFirst == false {
	// 	p.isFirst = true

	// 	fmt.Printf("vvvvvvvvvvvv:%s\nactionName:%s\n", controllerName, actionName)
	// } else {
	// 	fmt.Printf("ccccccccccccccccccccccontrollerName:%s\nactionName:%s\n", controllerName, actionName)
	// }

	// controllerName, actionName := p.GetControllerAndAction()
	// p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	// p.actionName = strings.ToLower(actionName)
	// p.o = orm.NewOrm()
	if actionName == "PostFunc" {
		//对特定页面请求进行账号验证
		if controllerName == "HouseEditController" && controllerName == "HouseListController" {

			user := p.GetSession("Username")
			pass := p.GetSession("Password")

			if userstring, ok := user.(string); ok && userstring != "" {
				if passstring, ok := pass.(string); ok && passstring != "" {
					if models.LoginSubmit(userstring, passstring) {
						fmt.Printf("登录成功%v,密码：%v", user, pass)
						//p.Prompt("登录成功", "")
					} else {
						p.Prompt("登录失败", "")
					}
				} else {
					p.Prompt("登录失败", "")
				}
			} else {
				fmt.Printf("登录aaassd成功")
				p.Prompt("登录失败", "")
			}
		}
	}

	// //初始化前台页面相关元素
	// if strings.ToLower(p.controllerName) == "blog" {

	// 	p.Data["actionName"] = strings.ToLower(actionName)
	// 	var result []*models.Config
	// 	p.o.QueryTable(new(models.Config).TableName()).All(&result)
	// 	configs := make(map[string]string)
	// 	for _, v := range result {
	// 		configs[v.Name] = v.Value
	// 	}
	// 	p.Data["config"] = configs
	// }

}
func (c *BaseController) Prompt(msg string, url string) {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// //c.TplName = "index.tpl"
	//c.TplName = "login.html"
	if url == "" {
		c.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		c.StopRun() //StopRun对USERSTOPRUN错误产生恐慌，定义后返回函数
	} else {
		c.Redirect(url, 302) //重定向使用状态代码将重定向响应发送到url。
	}
}
