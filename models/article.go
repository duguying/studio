package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	// "log"
	"strconv"
	"strings"
	"time"
)

type Article struct {
	Id       int
	Title    string
	Uri      string
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

// 添加文章
func AddArticle(title string, content string, keywords string, author string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := new(Article)
	art.Title = title
	art.Uri = strings.Replace(title, "/", "-", -1)
	art.Keywords = keywords
	art.Content = content
	art.Author = author
	return o.Insert(art)
}

// 通过id获取文章
func GetArticle(id int) (Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Id: id}
	err := o.Read(&art, "id")
	return art, err
}

// 通过uri获取文章
func GetArticleByUri(uri string) (Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Uri: uri}
	err := o.Read(&art, "uri")
	return art, err
}

// 通过文章标题获取文章
func GetArticleByTitle(title string) (Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Title: title}
	err := o.Read(&art, "title")
	return art, err
}

// 更新阅览数统计
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

// 更新文章
func UpdateArticle(id int64, uri string, newArt Article) error {
	o := orm.NewOrm()
	o.Using("default")
	var art Article

	if 0 != id {
		art = Article{Id: int(id)}
	} else if "" != uri {
		art = Article{Uri: uri}
	}

	art.Title = newArt.Title
	art.Keywords = newArt.Keywords
	art.Content = newArt.Content

	_, err := o.Update(&art, "title", "keywords", "content")
	return err
}

// 通过uri删除文章
func DeleteArticle(id int64, uri string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	var art Article

	if 0 != id {
		art.Id = int(id)
	} else if "" != uri {
		art.Uri = uri
	}

	return o.Delete(&art)
}

// 按月份统计文章数
// select DATE_FORMAT(time,'%Y年%m月') as date,count(*) as number ,year(time) as year, month(time) as month from article group by month order by month
func CountByMonth() ([]orm.Params, error) {
	sql := "select DATE_FORMAT(time,'%Y年%m月') as date,count(*) as number ,year(time) as year, month(time) as month from article group by month order by month"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		return maps, nil
	} else {
		return nil, err
	}
}

// 获取某月的文章列表
// select * from article where year(time)=2014 and month(time)=8
// year 年
// month 月
// page 页码
// numPerPage 每页条数
// 返回值:
// []orm.Params 文章
// bool 是否有下一页
// int 总页数
// error 错误
func ListByMonth(year int, month int, page int, numPerPage int) ([]orm.Params, bool, int, error) {
	if year < 0 {
		year = 1970
	}

	if month < 0 || month > 12 {
		month = 1
	}

	if page < 1 {
		page = 1
	}

	if numPerPage < 1 {
		numPerPage = 10
	}

	sql1 := "select * from article where year(time)=? and month(time)=? limit ?,?"
	sql2 := "select count(*)as number from article where year(time)=? and month(time)=?"

	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, year, month, numPerPage*(page-1), numPerPage).Values(&maps)
	o.Raw(sql2, year, month).Values(&maps2)

	// calculate pages
	number, _ := strconv.Atoi(maps2[0]["number"].(string))
	var addFlag int
	if 0 == (number % numPerPage) {
		addFlag = 0
	} else {
		addFlag = 1
	}
	pages := number/numPerPage + addFlag

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

// 文章分页列表
// select * from article order by time desc limit 0,6
// page 页码
// numPerPage 每页条数
// 返回值:
// []orm.Params 文章
// bool 是否有下一页
// int 总页数
// error 错误
func ListPage(page int, numPerPage int) ([]orm.Params, bool, int, error) {
	// pagePerNum := 6
	sql1 := "select * from article order by time desc limit ?," + fmt.Sprintf("%d", numPerPage)
	sql2 := "select count(*) as number from article"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, numPerPage*(page-1)).Values(&maps)
	o.Raw(sql2).Values(&maps2)

	number, _ := strconv.Atoi(maps2[0]["number"].(string))

	var addFlag int
	if 0 == (number % numPerPage) {
		addFlag = 0
	} else {
		addFlag = 1
	}

	pages := number/numPerPage + addFlag

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

// 同关键词文章列表
// select * from article where keywords like '%keyword%'
// 返回值:
// []orm.Params 文章
// bool 是否有下一页
// error 错误
func ListByKeyword(keyword string, page int, numPerPage int) ([]orm.Params, bool, int, error) {
	// numPerPage := 6
	sql1 := "select * from article where keywords like '%" + keyword + "%' limit ?," + fmt.Sprintf("%d", numPerPage)
	sql2 := "select count(*) as number from article where keywords like '%" + keyword + "%'"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, numPerPage*(page-1)).Values(&maps)
	o.Raw(sql2).Values(&maps2)

	number, _ := strconv.Atoi(maps2[0]["number"].(string))

	var addFlag int
	if 0 == (number % numPerPage) {
		addFlag = 0
	} else {
		addFlag = 1
	}

	pages := number/numPerPage + addFlag

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

// 最热文章列表
// select * from article order by count desc limit 10
func HottestArticleList() ([]orm.Params, error) {
	sql := "select * from article order by count desc limit 20"
	var maps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)
	return maps, err
}
