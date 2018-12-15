package models

type UserDB struct {
	Id        int        `json:"id"`
	Name      string     `json:"name" orm:"unique"`
	Nickname  string     `json:"nickname" orm:"default('')"`
	Mobile    string     `json:"mobile" orm:"default('')"`
	Age       int        `json:"age" `
	Sex       bool       `json:"sex"`
	Email     string     `json:"email" orm:"default('')"`
	Address   string     `json:"address" orm:"default('')"`
	Pass      string     `json:"pass"`
	Addtime   int        `json:"addtime"`
	Lastlogin int        `json:"lastlogin"`
	Articles  []*Article `orm:"reverse(many)"`
}

type Article struct {
	Id       int          `json:"id"`
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	Addtime  int          `json:"addtime"`
	Uptime   int          `json:"uptime"`
	UserDB   *UserDB      `json:"user" orm:"rel(fk)"`
	Link     string       `json:"link"`
	Intro    string       `json:"intro"`
	Type     *Articletype `json:"type" orm:"rel(fk)"`
	Comments []*Comment   `orm:"reverse(many)"` //反向一对多关联
}

type Articletype struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Orderno  int        `json:"orderno"`
	Articles []*Article `orm:"reverse(many)"`
}

type Comment struct {
	Id      int      `json:"id"`
	Cname   string   `json:"cname"`
	Cemail  string   `json:"cemail"`
	Content string   `json:"content"`
	Addtime int      `json:"addtime"`
	Aid     *Article `json:"article" orm:"rel(fk)"`
}

type Users struct {
	Id       int
	Username string `orm:"size(15);column(aaassaname)"`
	Pwd      string
	Age      int
	Sex      string
	Sex11    string
}

//----------------------------------
type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

/*
//我们也可以使用Tag对属性进行详细的设置
type Users struct {
	Id  int   `pk:"auto;column(id)"`  //设置主键自增长 列名设为id
	Name string `orm:"size(15);column(name)"`  //设置varchar长度为15 列名为name
	Pwd  string  `orm:"size(15);column(pwd)"`
	Age  int    `orm:"column(age)"`
	Sex  string  `orm:"size(15);column(sex)"`
}
*/
// type Model_DB interface {
// 	LoginSubmit(username string, pwd string) bool
// }
