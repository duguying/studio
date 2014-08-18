package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Article struct {
	Id      int
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

func List() {
	o := orm.NewOrm()
	o.Using("default")
	var art []*Article
	_, err := o.QueryTable("article").All(&art)

	if nil == err {
		for i := 0; i < len(art); i++ {
			fmt.Printf("item[" + strconv.Itoa(art[i].Id) + "]: " + art[i].Title + "\n")
		}
	}

}
