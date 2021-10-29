package cms

import "github.com/astaxie/beego"

type MyController struct {
	beego.Controller
}

func (i *MyController) Get() {
	i.TplName = "cms/my.html"
}
