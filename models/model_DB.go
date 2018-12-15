package models

import (
	"fmt"

	// "WEB/models/myredis"

	// "WEB/models/mymysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

var o orm.Ormer

// //初始化
func RegisterDB() {

	//注册数据驱动
	// orm.RegisterDriver("mysql", orm.DR_MySQL)
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//注册数据库 ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:root@/sys?charset=utf8", 30, 40)
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        数据库驱
	// 参数3        对应的连接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)

	//注册模型
	//orm.RegisterModel(new(Users), new(UserDB), new(Article), new(Articletype), new(Comment), new(Profile), new(Post), new(Tag))
	orm.RegisterModel(new(Users), new(UserDB), new(Article), new(Articletype), new(Comment))
	o = orm.NewOrm()
	//orm.RegisterModel(new(User), new(Profile), new(Post), new(Tag))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	//orm.RunSyncdb("default", true, true)
}
func LoginSubmit(username string, pwd string) bool {

	dbhost := beego.AppConfig.String("dbhost")
	fmt.Printf("配置文件字段:%s \n", dbhost)
	// o := orm.NewOrm()
	var userDB UserDB
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_d_b").Where(fmt.Sprintf("name='%s'", username)).And(fmt.Sprintf("pass='%s'", pwd))
	sql := qb.String()
	o.Raw(sql).QueryRow(&userDB)

	//fmt.Printf("sql语句:%v \n", sql)
	//fmt.Printf("sql语句:%d \n", userDB.Id)
	if userDB.Id == 0 {
		//fmt.Printf("sql语句sss:%v \n", userDB)
		return false
	}
	return true
}
func AddUser(username string, pwd string) {
	// o := orm.NewOrm()
	//var userDB UserDB
	qb, _ := orm.NewQueryBuilder("mysql")
	//qb.Select("*").From("user_d_b").Where(fmt.Sprintf("name='%s'", username)).And(fmt.Sprintf("pass='%s'", pwd))
	qb.InsertInto("user_d_b", "name", "pass").Values(Mark(username), Mark(pwd))
	sql := qb.String()
	fmt.Printf("sql语句:%v \n", sql)
	//o.Raw(sql).QueryRow(&userDB)
	o.Raw(sql).Exec()
}
func UpdateUser(username string, pwd string) {
	// o := orm.NewOrm()
	//var userDB UserDB
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("user_d_b").Set(fmt.Sprintf("pass='%s'", pwd)).Where(fmt.Sprintf("name='%s'", username))
	sql := qb.String()
	fmt.Printf("sql语句:%v \n", sql)
	//o.Raw(sql).QueryRow(&userDB)
	o.Raw(sql).Exec()
}
func DeleteUser(username string) {
	//多数据插入
	// users := []User{
	//    {Name: "slene"},
	//    {Name: "astaxie"},
	//    {Name: "unknown"},
	//    ...
	// }
	// successNums, err := o.InsertMulti(100, users)

	// o := orm.NewOrm()
	//var userDB UserDB
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Delete().From("user_d_b").Where(fmt.Sprintf("name='%s'", username))
	sql := qb.String()
	fmt.Printf("sql语句:%v \n", sql)
	//o.Raw(sql).QueryRow(&userDB)
	o.Raw(sql).Exec()
}
func Mark(val string) string {
	return fmt.Sprintf("'%s'", val)
}
func PoolRedis() {

	pool := &redis.Pool{
		MaxIdle:   3, /*最大的空闲连接数*/
		MaxActive: 8, /*最大的激活连接数*/
		Dial: func() (redis.Conn, error) {
			// c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("密码"))
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	c := pool.Get()

	defer c.Close()
	//存值,
	// _, err := c.Do("SET", "key", "value")
	// if err != nil {
	// 	beego.Notice("redis set failed:", err)
	// } else {
	// 	beego.Notice("save success")
	// }
	//设置过期时间
	_, err := c.Do("SET", "key", "value", "EX", 360)
	if err != nil {
		beego.Notice("redis set failed:", err)
	} else {
		beego.Notice("save success")
	}
	// //存int
	// _, err := c.Do("SET", "key", 2)

	// //取值
	v, err1 := redis.String(c.Do("GET", "key"))
	if err1 != nil {
		beego.Notice("redis get failed:", err1)
	} else {
		beego.Notice("Get mykey: %v \n", v)
	}
	// bytes, err := redis.Bytes(c.Do("GET", "key"))
}
