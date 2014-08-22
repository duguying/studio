package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	// "log"
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

func UpdateCount(id int) error {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Id: id}
	err := o.Read(&art)

	o.QueryTable("article").Filter("id", id).Update(orm.Params{
		"count": art.Count + 1,
	})

	return err
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

/**
 * 按月份统计文章数
 * select DATE_FORMAT(time,'%Y-%m') as month,count(*) as number from article group by month order by month
 */
func CountByMonth() ([]orm.Params, error) {
	sql := "select DATE_FORMAT(time,'%Y年%m月') as month,count(*) as number from article group by month order by month"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		return maps, nil
	} else {
		return nil, err
	}
}

/**
 * 文章分页列表
 * select * from article limit 0,6
 * 返回值:
 * []orm.Params 文章
 * bool 是否有下一页
 * error 错误
 */
func ListPage(page int) ([]orm.Params, bool, int, error) {
	pagePerNum := 6
	sql1 := "select * from article limit ?," + fmt.Sprintf("%d", pagePerNum)
	sql2 := "select count(*) as number from article"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, 6*(page-1)).Values(&maps)
	o.Raw(sql2).Values(&maps2)

	number, _ := strconv.Atoi(maps2[0]["number"].(string))

	var addFlag int
	if 0 == (number % pagePerNum) {
		addFlag = 0
	} else {
		addFlag = 1
	}

	pages := number/pagePerNum + addFlag

	var flagNextPage bool
	if pages == page {
		flagNextPage = false
	} else {
		flagNextPage = true
	}

	if err == nil && num > 0 {
		return maps, flagNextPage, pages, nil
	} else {
		return nil, false, pages, err
	}
}

/**
 * 同关键词文章列表
 * select * from article where keywords like '%keyword%'
 */
func ListByKeyword(keyword string) {

}
