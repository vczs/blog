package front

import (
	"blog/models"
	"blog/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type FrontLoginController struct {
	beego.Controller
}

func (l *FrontLoginController) Get() {
	l.TplName = "front/login.html"
}

func (l *FrontLoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")
	md5_pwd := utils.GetMD5(password)

	o := orm.NewOrm()
	exist := o.QueryTable(new(models.User)).Filter("user_name", username).Filter("password", md5_pwd).Exist()
	if exist {
		l.SetSession("cms_user_name", username)
		l.Redirect(beego.URLFor("IndexController.Get"), 302)
	} else {
		l.Redirect(beego.URLFor("FrontLoginController.Get"), 302)
	}
}

func (l *FrontLoginController) UnLogin() {
	l.DelSession("cms_user_name")
	l.Redirect(beego.URLFor("IndexController.Get"), 302)
}
