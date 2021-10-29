package front

import (
	"blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Post() {
	post_id, err := c.GetInt("post_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 401, "msg": "请先登录！"}
		c.ServeJSON()
	} else {
		o := orm.NewOrm()
		post := &models.Post{}
		o.QueryTable(new(models.Post)).Filter("id", post_id).One(post)
		content := c.GetString("content")
		user_name := c.GetSession("cms_user_name")
		if user_name == nil {
			c.Data["json"] = map[string]interface{}{"code": 401, "msg": "请先登录！"}
			c.ServeJSON()
		} else {
			user := &models.User{}
			o.QueryTable(new(models.User)).Filter("user_name", user_name).One(user)
			pid, err := c.GetInt("pid")
			if err != nil {
				pid = 0
			}
			comment := &models.Comment{
				Post:    post,
				Content: content,
				PId:     pid,
				Author:  user,
			}
			_, errInsert := o.Insert(comment)
			if errInsert != nil {
				c.Data["json"] = map[string]interface{}{"code": 500, "msg": "评论出错，请重试！"}
			} else {
				c.Data["json"] = map[string]interface{}{"code": 200, "msg": "评论成功！"}
			}
			c.ServeJSON()
		}
	}
}
