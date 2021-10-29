package cms

import (
	"blog/models"
	"blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "cms/login.html"
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	password := l.GetString("password")
	md5_pwd := utils.GetMD5(password)

	o := orm.NewOrm()
	user := models.User{}
	err := o.QueryTable(new(models.User)).Filter("user_name", username).One(&user)
	if err != nil {
		l.SetSession("cms_user_name", username)
		l.Redirect(beego.URLFor("LoginController.Get"), 302)
	} else {
		ok := md5_pwd == user.Password && user.IsAdmin == 1
		if ok {
			l.SetSession("cms_user_name", username)
			l.Redirect(beego.URLFor("MainController.Get"), 302)
		} else {
			l.SetSession("cms_user_name", username)
			l.Redirect(beego.URLFor("LoginController.Get"), 302)
		}
	}
}

func (l *LoginController) CmsUnlogin() {
	l.DelSession("cms_user_name")
	l.Redirect(beego.URLFor("LoginController.Get"), 302)
}
