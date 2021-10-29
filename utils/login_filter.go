package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func CmsLoginFilter(c *context.Context) {
	cms_user_name := c.Input.Session("cms_user_name")
	if cms_user_name == nil {
		c.Redirect(302, beego.URLFor("LoginController.Get"))
	}
}
