package controllers

import (
	"WEB/models"
	"WEB/models/myredis"
	"fmt"
)

type MainController struct {
	BaseController
}

func (c *MainController) LoginFunc() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	//c.Ctx.WriteString(fmt.Sprintf("接口:%s \n", c.Ctx.Request.Method))
	fmt.Sprintf("接口:%s \n", c.Ctx.Request.Method)
	user := c.GetSession("Username")
	pass := c.GetSession("Password")
	if userstring, ok := user.(string); ok && userstring != "" {
		if passstring, ok := pass.(string); ok && passstring != "" {
			// c.Ctx.WriteString("uaacser:" + user.(string) + ",pass:" + pass.(string))
			c.Prompt("登录成功", "index")
		} else {
			c.Prompt("密码错误请重新输入", "")
		}
	}

	c.TplName = "login.html"
}
func (c *MainController) PostFunc() {
	username := c.GetString("userEntity.userCode")
	password := c.GetString("userEntity.password")

	if username == "" || password == "" {
		// fmt.Printf("v1 type:%s\n", username)
		// fmt.Printf("v2 type:%s\n", password)
		//c.Ctx.WriteString(fmt.Printf("v1 type:%T\n", username))
		//c.Ctx.WriteString(fmt.Printf("v2 type:%T\n", password))
		//c.Ctx.WriteString("aaausername:" + username + ",aaapassword:" + password)
		c.Prompt("账号不存在", "")
	} else {
		/*测试key，value*/
		_ = myredis.SetToDB(username, password)
		// fmt.Printf("v1 type:%s\n", isSuccess)
		_, _ = myredis.GetFromDB(username)
		// fmt.Printf("v1 type:%s\n", isSucces)
		/*测试批量key，value*/
		// myredis.MsetToDB()
		// myredis.MgetToDB()
		myredis.SetJsonToDB(username, "")
		myredis.GetJsonFromDB(username)
		if models.LoginSubmit(username, password) {

			// u := UserV2{}
			// if err := c.ParseForm(&u); err != nil {
			// 	c.Ctx.WriteString("err :" + err.Error())
			// 	return
			// }
			//c.Ctx.SetCookie("Username", u.Username, 150, "/")
			//c.Ctx.SetCookie("Password", u.Password, 150, "/")
			c.SetSession("Username", username)
			c.SetSession("Password", password)
			//c.Ctx.WriteString("aaausername:" + u.Username + ",aaapassword:" + u.Password)
			c.Prompt("登录成功", "index")
		} else {
			c.Prompt("登录失败", "")
		}

	}
	// var ob models.Ob

}
