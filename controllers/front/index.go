package front

import (
	"blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {
	o := orm.NewOrm()
	posts := []models.Post{}
	o.QueryTable(new(models.Post)).RelatedSel().All(&posts)
	front_user_name := i.GetSession("cms_user_name")
	if front_user_name == nil {
		front_user_name = ""
	}
	i.Data["username"] = front_user_name
	i.Data["posts"] = posts
	i.TplName = "front/index.html"
}

func (i *IndexController) PostDetail() {
	id, _ := i.GetInt("id")
	o := orm.NewOrm()
	post := models.Post{}
	qs := o.QueryTable(new(models.Post)).Filter("id", id)
	qs.RelatedSel().One(&post)

	// 阅读数+1
	qs.Update(orm.Params{"read_num": post.ReadNum + 1})

	front_user_name := i.GetSession("cms_user_name")
	if front_user_name == nil {
		front_user_name = ""
	}

	comments := []models.Comment{}
	o.QueryTable(new(models.Comment)).Filter("post_id", id).Filter("p_id", 0).RelatedSel().All(&comments)
	comment_trees := []models.CommentTree{}
	for _, v := range comments {
		pid := v.Id
		comment_tree := models.CommentTree{
			Id:         v.Id,
			Content:    v.Content,
			Author:     v.Author,
			CreateTime: v.CreateTime,
			Children:   []*models.CommentTree{},
		}
		GetChild(pid, &comment_tree)
		comment_trees = append(comment_trees, comment_tree)
	}

	i.Data["username"] = front_user_name
	i.Data["post"] = post
	i.Data["comment_trees"] = comment_trees
	i.TplName = "front/detail.html"
}

func GetChild(pid int, comment_tree *models.CommentTree) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Comment))
	comments := []models.Comment{}
	_, err := qs.Filter("p_id", pid).RelatedSel().All(&comments)
	if err != nil {
		return
	}
	//查看一级评论下面的楼层
	for i := 0; i < len(comments); i++ {
		pid := comments[i].Id
		child := models.CommentTree{
			Id:         comments[i].Id,
			Content:    comments[i].Content,
			Author:     comments[i].Author,
			CreateTime: comments[i].CreateTime,
			Children:   []*models.CommentTree{},
		}
		comment_tree.Children = append(comment_tree.Children, &child)
		GetChild(pid, &child)
	}
}
