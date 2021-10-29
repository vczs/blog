package front

import (
	"blog/models"
	"blog/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "front/register.html"
}

func (r *RegisterController) Post() {
	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")
	if password != repassword {
		r.Ctx.WriteString("两次密码不一致")
	} else {
		md5_pwd := utils.GetMD5(password)
		o := orm.NewOrm()
		user := &models.User{
			UserName: username,
			Password: md5_pwd,
			IsAdmin:  2,
			Cover:    "static/upload/bq3.png",
		}
		_, err := o.Insert(user)
		if err != nil {
			r.Ctx.WriteString("用户已存在！")
		} else {
			r.Redirect("/login", 302)
		}
	}
}
