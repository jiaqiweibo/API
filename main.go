package main

import (
	"WEB/models"
	_ "WEB/routers"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//引入数据模型
func init() {
	// 注册数据库
	models.RegisterDB()

}

func main() {
	// 开启 ORM 调试模式
	fmt.Println(strings.Index("widuu", "uu")) //0
	orm.Debug = false
	// 设置级别
	// LevelEmergency
	// LevelAlert
	// LevelCritical
	// LevelError
	// LevelWarning
	// LevelNotice
	// LevelInformational
	// LevelDebug
	beego.SetLevel(beego.LevelNotice)
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.SetLogFuncCall(true)
	// 这个默认情况就会同时输出到两个地方，一个 console，一个 file，如果只想输出到文件，就需要调用删除操作：
	//beego.BeeLogger.DelLogger("console")

	//beego.Informational("aaaa")
	// 自动建表
	//orm.RunSyncdb("default", true, true)

	beego.Run()
	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
}
