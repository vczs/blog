package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int       `orm:"pk;auto"`
	UserName   string    `orm:"description(用户名);index;unique"`
	Password   string    `orm:"description(密码)"`
	IsAdmin    int       `orm:"description(1是管理员 2是用户);default(2)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now_add"`
	Cover      string    `orm:"description(头像);default(static/upload/bq3.png)"`
	Posts      []*Post   `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User), new(Post), new(Comment))
}
