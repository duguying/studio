package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/duguying/blog/utils"
	"github.com/gogather/com"
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

func (this *Article) TableName() string {
	return "article"
}

func init() {
	orm.RegisterModel(new(Article))
}

// 添加文章
func AddArticle(title string, content string, keywords string, author string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")

	sql := "insert into article(title, uri, keywords, content, author) values(?, ?, ?, ?, ?)"
	res, err := o.Raw(sql, title, strings.Replace(title, "/", "-", -1), keywords, content, author).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

// 通过id获取文章-cached
func GetArticle(id int) (Article, error) {
	var err error
	var art Article

	cache := utils.GetCache("GetArticle.id." + fmt.Sprintf("%d", id))
	if cache != nil { // check cache
		json.Unmarshal([]byte(cache.(string)), &art)
		return art, nil
	} else {
		o := orm.NewOrm()
		o.Using("default")
		art = Article{Id: id}
		err = o.Read(&art, "id")

		data, _ := com.JsonEncode(art)
		utils.SetCache("GetArticle.id."+fmt.Sprintf("%d", id), data, 600)
	}

	return art, err
}

// 通过uri获取文章-cached
func GetArticleByUri(uri string) (Article, error) {
	var err error
	var art Article

	cache := utils.GetCache("GetArticleByUri.uri." + uri)
	if cache != nil {
		json.Unmarshal([]byte(cache.(string)), &art)
		// get view count
		count, err := GetArticleViewCount(art.Id)
		if err == nil {
			art.Count = int(count)
		}

		return art, nil
	} else {
		o := orm.NewOrm()
		o.Using("default")
		art = Article{Uri: uri}
		err = o.Read(&art, "uri")

		data, _ := com.JsonEncode(art)
		utils.SetCache("GetArticleByUri.uri."+uri, data, 600)
	}

	return art, err
}

// 通过文章标题获取文章-cached
func GetArticleByTitle(title string) (Article, error) {
	var err error
	var art Article

	cache := utils.GetCache("GetArticleByTitle.title." + title)
	if cache != nil {
		json.Unmarshal([]byte(cache.(string)), &art)
		// get view count
		count, err := GetArticleViewCount(art.Id)
		if err == nil {
			art.Count = int(count)
		}

		return art, nil
	} else {
		o := orm.NewOrm()
		o.Using("default")
		art = Article{Title: title}
		err = o.Read(&art, "title")

		data, _ := com.JsonEncode(art)
		utils.SetCache("GetArticleByTitle.title."+title, data, 600)
	}

	return art, err
}

// 获取文章浏览量
func GetArticleViewCount(id int) (int, error) {
	var maps []orm.Params

	sql := `select count from article where id=?`
	o := orm.NewOrm()
	num, err := o.Raw(sql, id).Values(&maps)
	if err == nil && num > 0 {
		count := maps[0]["count"].(string)

		return strconv.Atoi(count)
	} else {
		return 0, err
	}
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

// 按月份统计文章数-cached
// select DATE_FORMAT(time,'%Y年%m月') as date,count(*) as number ,year(time) as year, month(time) as month from article group by date order by year desc, month desc
func CountByMonth() ([]orm.Params, error) {
	var maps []orm.Params

	cache := utils.GetCache("CountByMonth")
	if nil != cache {
		json.Unmarshal([]byte(cache.(string)), &maps)
		return maps, nil
	} else {
		sql := "select DATE_FORMAT(time,'%Y年%m月') as date,count(*) as number ,year(time) as year, month(time) as month from article group by date order by year desc, month desc"
		o := orm.NewOrm()
		num, err := o.Raw(sql).Values(&maps)
		if err == nil && num > 0 {
			data, _ := com.JsonEncode(maps)
			utils.SetCache("CountByMonth", data, 3600)
			return maps, nil
		} else {
			return nil, err
		}
	}

}

// 获取某月的文章列表-cached
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

	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	var err error

	// get data - cached
	cache1 := utils.GetCache(fmt.Sprintf("ListByMonth.list.%d.%d.%d", year, month, page))
	if nil != cache1 {
		json.Unmarshal([]byte(cache1.(string)), &maps)
	} else {
		sql1 := "select * from article where year(time)=? and month(time)=? order by time desc limit ?,?"
		_, err = o.Raw(sql1, year, month, numPerPage*(page-1), numPerPage).Values(&maps)

		data1, _ := com.JsonEncode(maps)
		utils.SetCache(fmt.Sprintf("ListByMonth.list.%d.%d.%d", year, month, page), data1, 3600)
	}

	cache2 := utils.GetCache(fmt.Sprintf("ListByMonth.count.%d.%d", year, month))
	if nil != cache2 {
		json.Unmarshal([]byte(cache2.(string)), &maps2)
	} else {
		sql2 := "select count(*)as number from article where year(time)=? and month(time)=?"
		o.Raw(sql2, year, month).Values(&maps2)

		data2, _ := com.JsonEncode(maps2)
		utils.SetCache(fmt.Sprintf("ListByMonth.count.%d.%d", year, month), data2, 3600)
	}

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

	if err == nil {
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
	if err != nil {
		fmt.Println("execute sql1 error:")
		fmt.Println(err)
		return nil, false, 0, err
	}

	_, err = o.Raw(sql2).Values(&maps2)
	if err != nil {
		fmt.Println("execute sql2 error:")
		fmt.Println(err)
		return nil, false, 0, err
	}

	number, err := strconv.Atoi(maps2[0]["number"].(string))

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
	sql1 := "select * from article where keywords like ? order by time desc limit ?,?"
	sql2 := "select count(*) as number from article where keywords like ?"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, fmt.Sprintf("%%%s%%", keyword), numPerPage*(page-1), numPerPage).Values(&maps)
	o.Raw(sql2, fmt.Sprintf("%%%s%%", keyword)).Values(&maps2)

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
