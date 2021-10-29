package cms

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Get() {
	o := orm.NewOrm()
	posts := []models.Post{}
	qs := o.QueryTable(new(models.Post))
	qs.RelatedSel().All(&posts)

	//查询到总行数
	count, _ := qs.Count()
	current_page, err := p.GetInt("pages")
	if err != nil {
		current_page = 1
	}
	total_pages := utils.GetPageNum(count, 10)

	left_pages, right_pages, left_has_more, right_has_more := utils.Get_pagination_data(total_pages, current_page, 4)
	has_pre_page, has_next_page, pre_page, next_page := utils.HasNext(current_page, total_pages)

	p.Data["left_pages"] = left_pages
	p.Data["left_has_more"] = left_has_more
	p.Data["page"] = current_page

	p.Data["has_pre_page"] = has_pre_page
	p.Data["pre_page"] = pre_page
	p.Data["has_next_page"] = has_next_page
	p.Data["next_page"] = next_page

	p.Data["right_pages"] = right_pages
	p.Data["right_has_more"] = right_has_more

	p.Data["num_pages"] = total_pages //总页数
	p.Data["count"] = count           //总数量

	p.Data["posts"] = posts

	p.TplName = "cms/post-list.html"
}

func (p *PostController) ToAdd() {
	p.TplName = "cms/post-add.html"
}

func (p *PostController) DoAdd() {
	//获取前端数据
	title := p.GetString("title")
	desc := p.GetString("desc")
	content := p.GetString("content")
	f, h, err := p.GetFile("cover")
	if f != nil {
		defer f.Close()
	}

	//cover处理
	var cover string
	if err != nil {
		cover = "static/upload/no_pic.jpg"
	} else {
		timeUnix := time.Now().Unix()
		time_str := strconv.FormatInt(timeUnix, 10)
		path := "static/upload/" + time_str + h.Filename
		err = p.SaveToFile("cover", path)
		if err != nil {
			cover = "static/upload/no_pic.jpg"
		} else {
			cover = path
		}
	}

	//保存到数据库
	o := orm.NewOrm()
	author := p.GetSession("cms_user_name")
	user := &models.User{}
	o.QueryTable(new(models.User)).Filter("user_name", author).One(user)
	post := &models.Post{
		Title:   title,
		Desc:    desc,
		Content: content,
		Cover:   cover,
		Author:  user,
	}
	_, err_insert := o.Insert(post)
	if err_insert != nil {
		p.Data["json"] = map[string]interface{}{"code": 500, "msg": err}
		p.ServeJSON()
	} else {
		p.Data["json"] = map[string]interface{}{"code": 200, "msg": "添加成功"}
		p.ServeJSON()
	}
}

func (p *PostController) PostDelete() {
	id, err := p.GetInt("id")
	if err != nil {
		p.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()
	_, err_delete := o.QueryTable(new(models.Post)).Filter("id", id).Delete()

	if err_delete != nil {
		p.Ctx.WriteString("删除错误")
	}

	p.Redirect(beego.URLFor("PostController.Get"), 302)
}

func (p *PostController) ToEdit() {
	id, err := p.GetInt("id")
	if err != nil {
		p.Data["json"] = map[string]interface{}{"code": 500, "msg": "id参数错误"}
	}

	o := orm.NewOrm()
	post := models.Post{}
	o.QueryTable(new(models.Post)).Filter("id", id).One(&post)
	p.Data["post"] = post
	p.TplName = "cms/post-edit.html"
}

func (p *PostController) DoEdit() {
	id, err := p.GetInt("id")
	if err != nil {
		p.Data["json"] = map[string]interface{}{"code": 500, "msg": "id参数错误"}
	}
	title := p.GetString("title")
	desc := p.GetString("desc")
	content := p.GetString("content")
	f, h, err := p.GetFile("cover")

	var errUpdate error
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Post)).Filter("id", id)
	if err != nil {
		_, errUpdate = qs.Update(orm.Params{
			"title":   title,
			"desc":    desc,
			"content": content,
		})
	} else {
		defer f.Close()
		timeUnix := time.Now().Unix()
		time_str := strconv.FormatInt(timeUnix, 10)
		path := "static/upload/" + time_str + h.Filename
		errSave := p.SaveToFile("cover", path)
		if errSave != nil {
			errUpdate = fmt.Errorf("封面上传失败:%v", errSave)
		} else {
			_, errUpdate = qs.Update(orm.Params{
				"title":   title,
				"desc":    desc,
				"content": content,
				"cover":   path,
			})
		}
	}

	if errUpdate != nil {
		p.Data["json"] = map[string]interface{}{"code": 500, "msg": "更新失败"}
	} else {
		p.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
	}
	p.ServeJSON()
}
