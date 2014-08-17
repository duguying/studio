package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id      int64
	Title   string
	Content string
}

func (u *Article) TableName() string {
	return "article"
}

func init() {
	orm.RegisterModel(new(Article))
}

func AddArticle(title string, content string, keywords string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := new(Article)
	art.Title = title
	art.Content = content
	return o.Insert(art)
}
