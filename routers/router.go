package routers

import (
	"blog/controllers/cms"
	"blog/controllers/front"
	"github.com/astaxie/beego"
)

func init() {
	//后台
	beego.Router("/cms", &cms.LoginController{})
	beego.Router("/cms/unlogin", &cms.LoginController{}, "get:CmsUnlogin")
	beego.Router("/cms/my", &cms.MyController{})
	beego.Router("/cms/index/index", &cms.MainController{})
	beego.Router("/cms/index/welcome", &cms.MainController{}, "get:Welcome")
	beego.Router("/cms/index/post_list", &cms.PostController{})
	beego.Router("/cms/index/post_toadd", &cms.PostController{}, "get:ToAdd")
	beego.Router("/cms/index/post_doadd", &cms.PostController{}, "post:DoAdd")
	beego.Router("/cms/index/post_delete", &cms.PostController{}, "get:PostDelete")
	beego.Router("/cms/index/post_to_edit", &cms.PostController{}, "get:ToEdit")
	beego.Router("/cms/index/post_do_edit", &cms.PostController{}, "post:DoEdit")
	//前端
	beego.Router("/", &front.IndexController{})
	beego.Router("/detail", &front.IndexController{}, "get:PostDetail")
	beego.Router("/register", &front.RegisterController{})
	beego.Router("/login", &front.FrontLoginController{})
	beego.Router("/unlogin", &front.FrontLoginController{}, "get:UnLogin")
	beego.Router("/comment", &front.CommentController{})
}
