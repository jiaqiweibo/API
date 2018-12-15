package myredis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	// "github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

// Conn return redis connection.
func Conn() redis.Conn {
	return pool.Get()
}

/*
func Close() {
	pool.Close()
}
MaxIdle：最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。

MaxActive：最大的激活连接数，表示同时最多有N个连接

IdleTimeout：最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
*/

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     16,
		MaxActive:   16,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	server := beego.AppConfig.String("cache::server")
	password := beego.AppConfig.String("cache::password")

	pool = newPool(server, password)
}

/*
---------------@redis缓存 写入单个key，value----------------
设置过期时间 :EX
*/
func SetToDB(key string, value string) (isSuccess bool) {
	conn := Conn()
	defer conn.Close()
	r, err := conn.Do("SET", key, value, "EX", "3600")
	if err != nil {
		isSuccess = false
		beego.Notice(fmt.Sprintf("redis set failed:%s", err))
	} else {
		// fmt.Println("save success")
		beego.Notice(fmt.Sprintf("SaveToDB,写入测试:%s", r))
		isSuccess = true
	}
	return
}

/*
@redis缓存测试 读取单个key
*/
func GetFromDB(key string) (value string, isSucess bool) {
	conn := Conn()
	defer conn.Close()

	r, err := redis.String(conn.Do("GET", key))
	if err != nil {
		isSucess = false
		beego.Notice(fmt.Sprintf("redis get failed:%s", err))
	} else {
		value = r
		isSucess = true
		beego.Notice(fmt.Sprintf("ReadFromDB,读取测试:%s", r))
	}
	return
}

/*
---------------@redis缓存 写入json单个key，value----------------
设置过期时间 :EX
*/
func SetJsonToDB(key string, value interface{}) (isSuccess bool) {
	conn := Conn()
	defer conn.Close()

	// var strMap map[string]string
	strMap := make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	strMap["France"] = "Paris"
	strMap["Italy"] = "罗马"
	strMap["Japan"] = "东京"
	strMap["India "] = "新德里"
	beego.Notice(fmt.Sprintf("Json:%s", strMap))
	data, _ := json.Marshal(strMap)
	r, err := conn.Do("SET", key, data, "EX", "3600")
	if err != nil {
		isSuccess = false
		beego.Notice(fmt.Sprintf("redis set failed:%s", err))
	} else {
		// fmt.Println("save success")
		beego.Notice(fmt.Sprintf("SaveToDB,写入测试:%s", r))
		isSuccess = true
	}
	return
}

/*
@redis缓存测试 读取单个key
*/
func GetJsonFromDB(key string) (value interface{}, isSucess bool) {
	conn := Conn()
	defer conn.Close()

	r, err := redis.String(conn.Do("GET", key))
	if err != nil {
		isSucess = false
		beego.Notice(fmt.Sprintf("redis get failed:%s", err))
	} else {
		value = r
		isSucess = true
		beego.Notice(fmt.Sprintf("ReadFromDB,读取测试:%s", r))
	}
	return
}

/*
-----------------@redis缓存 写入多个key，value（批量设置）-------------
*/

func MsetToDB() {
	conn := Conn()
	defer conn.Close()
	s, err := conn.Do("MSET", "name", "superWang", "SEX", "F", "EX", "50")
	if err != nil {
		beego.Notice(fmt.Println("redis set failed:", err))
	} else {
		// fmt.Println("save success")
		beego.Notice(fmt.Sprintf("MsetToDB,批量测试:%s", s))
	}

}

/*
@redis缓存 读取多个key（批量读取）
返回数组
*/
func MgetToDB() {
	conn := Conn()
	defer conn.Close()
	r, err := redis.Strings(conn.Do("MGET", "SEX", "name"))
	if err != nil {
		beego.Notice(fmt.Println("redis get failed:", err))
	} else {
		beego.Notice(fmt.Printf("MgetToDB,批量测试: %T \n", r))
	}
}
