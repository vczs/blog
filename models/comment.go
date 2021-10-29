package models

import "time"

type CommentTree struct {
	Id         int
	Content    string
	Author     *User
	CreateTime time.Time
	Children   []*CommentTree
}

type Comment struct {
	Id         int       `orm:"pk;auto"`
	Content    string    `orm:"size(4000);description(评论内容)"`
	Post       *Post     `orm:"rel(fk);description(帖子外键)"`
	PId        int       `orm:"description(父级评论);default(0)"`
	Author     *User     `orm:"rel(fk);description(评论人)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
}

func (c *Comment) TableName() string {
	return "sys_post_comment"
}
