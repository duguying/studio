package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Article struct {
	Id       int
	Title    string
	Keywords string
	Content  string
	Author   string
	Time     time.Time
	Count    int
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

func GetArticle(id int) (Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Id: id}
	err := o.Read(&art, "id")
	return art, err
}

func GetArticleByTitle(title string) (Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Title: title}
	err := o.Read(&art, "title")
	return art, err
}
