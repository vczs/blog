package cms

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (i *MainController) Get() {
	i.TplName = "cms/index.html"
}

func (i *MainController) Welcome() {
	i.TplName = "cms/welcome.html"
}
