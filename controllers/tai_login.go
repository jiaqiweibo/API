package controllers

import (
	"WEB/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type LoginController struct {
	beego.Controller
}

type UserV2 struct {
	Username string
	Password string
}

/*数据库UserInfo 对应 user_info*/
type UserInfo struct {
	Id       int
	UserName string // 数据表中的别名 `json:"username"`
	PassWord string
}

func (this *LoginController) GetFunc() {
	/*测试数据库使用*/

	//注册驱动
	//orm.RegisterDriver("mysql", orm.DR_MySQL)

	// 设置默认数据库
	//mysql用户：root ，密码：zxxx ， 数据库名称：test ， 数据库别名：default  30:设置数据库的最大空闲连接。 40:设置数据库的最大数据库连接
	orm.RegisterDataBase("default", "mysql", "root:root@/sys?charset=utf8", 30, 40)
	// 注册定义的 model
	orm.RegisterModel(new(UserInfo))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))
	o := orm.NewOrm()
	//第一种:下面是插入数据
	// user3 := UserInfo{UserName: "aaaa", PassWord: "123456"}
	// id1, err := o.Insert(&user3)
	// if err != nil {
	// 	this.Ctx.WriteString("mysql插入数据id:" + ",err:" + err.Error())
	// }
	// fmt.Println("sssss:", id1)
	var user_1 models.UserDB
	user_1.Id = 1
	user_1.Age = 10
	fmt.Println("aa", user_1)

	//下面是更新
	//user := UserInfo{Username:"zhangsan", Password:"123456"}
	//user.Id = 1
	//o.Update(&user)

	//下面是读取
	user_r := UserInfo{Id: 1}
	o.Read(&user_r)
	/*第二种*/
	//var maps []orm.Params
	var maps []UserInfo
	// num, err := o.Raw("select * from user_info").Values(&maps)
	o.Raw("select * from user_info").QueryRows(&maps)
	this.Ctx.WriteString(fmt.Sprintf("user info:%v \n", maps))
	// if err != nil {
	// 	return
	// }
	// for i, term := range maps {
	// 	fmt.Println("num:", num, "i=", i, ",term=", term["user_name"])
	// }
	//采用queryBuilder方式进行读取  最常用的方式
	var users []UserInfo
	//var users []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info").OrderBy("id").Desc().Limit(3)
	sql := qb.String()
	o.Raw(sql).QueryRows(&users)
	this.Ctx.WriteString(fmt.Sprintf("sql语句:%v \n", users))

	username := this.Ctx.GetCookie("Username")
	password := this.Ctx.GetCookie("Password")
	user := this.GetSession("Username")
	pass := this.GetSession("Password")
	if userstring, ok := user.(string); ok && userstring != "" {
		this.Ctx.WriteString("uaacser:" + user.(string) + ",pass:" + pass.(string))
	}
	if username == "" || password == "" {
		this.Ctx.WriteString(`<html><form action="http://127.0.0.1:8081/login" method="post">
							<input type="text" name="Username"/>
							<input type="password" name="Password"/>
							<input type="submit" value="提交"/>
					   </form></html>`)
		this.Ctx.WriteString("出差username:" + username + ",出差password:" + password)
	} else {
		this.Ctx.WriteString("aaausername:" + username + ",啊啊password:" + password)
	}

}

func (this *LoginController) PostFunc() {
	// var ob models.Ob
	u := UserV2{}
	if err := this.ParseForm(&u); err != nil {
		this.Ctx.WriteString("err :" + err.Error())
		return
	}
	this.Ctx.SetCookie("Username", u.Username, 150, "/")
	this.Ctx.SetCookie("Password", u.Password, 150, "/")
	this.SetSession("Username", u.Username)
	this.SetSession("Password", u.Password)
	this.Ctx.WriteString("aaausername:" + u.Username + ",aaapassword:" + u.Password)

}
