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

func AddArticle(title string, content string, keywords string, author string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := new(Article)
	art.Title = title
	art.Keywords = keywords
	art.Content = content
	art.Author = author
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

func UpdateArticle(id int64, title string, newArt Article) error {
	o := orm.NewOrm()
	o.Using("default")
	var art Article

	if 0 != id {
		art = Article{Id: int(id)}
	} else if "" != title {
		art = Article{Title: title}
	}

	art.Title = newArt.Title
	art.Keywords = newArt.Keywords
	art.Content = newArt.Content

	_, err := o.Update(&art, "title", "keywords", "content")
	return err
}

func DeleteArticle(id int64, title string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	var art Article

	if 0 != id {
		art.Id = int(id)
	} else if "" != title {
		art.Title = title
	}

	return o.Delete(&art)
}
