package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

type Usera struct {
	Username string
	Password string
}

func (this *HomeController) GetFunc() {
	username := this.GetString("Username")
	password := this.GetString("Password")
	if username == "" || password == "" {
		this.Ctx.WriteString(`<html><form action="http://127.0.0.1:8081/user" method="post">
							<input type="text" name="Username"/>
							<input type="password" name="Password"/>
							<input type="submit" value="提交"/>
					   </form></html>`)
		this.Ctx.WriteString("username:" + username + ",password:" + password)
	} else {
		this.Ctx.WriteString("username:" + username + ",password:" + password)
	}
	// this.Ctx.WriteString("<html>" + id + "<br/>")
	user := this.GetSession("Username")
	pass := this.GetSession("Password")
	if userstring, ok := user.(string); ok && userstring != "" {
		this.Ctx.WriteString("uaacser:" + user.(string) + ",pass:" + pass.(string))
	}

}
func (this *HomeController) PostFunc() {

	// var ob models.Ob
	id := this.GetString("id")
	if id == "" {
		fmt.Printf("err:id is", id)
	}
	fmt.Println(id)
	u := Usera{}
	if err := this.ParseForm(&u); err != nil {
		this.Ctx.WriteString("err :" + err.Error())
	}
	this.Ctx.WriteString("usebbrname:" + u.Username + ",pabbssword:" + u.Password)
}
