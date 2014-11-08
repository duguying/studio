package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/duguying/blog/models"
	"log"
	"strconv"
)

// 添加文章
type AddArticleController struct {
	beego.Controller
}

func (this *AddArticleController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}

func (this *AddArticleController) Post() {
	title := this.GetString("title")
	content := this.GetString("content")
	keywords := this.GetString("keywords")

	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJson()
		return
	}

	if "" == title {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "title can not be empty", "refer": "/"}
		this.ServeJson()
		return
	}

	username := user.(string)

	id, err := AddArticle(title, content, keywords, username)
	if nil == err {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "success added, id " + fmt.Sprintf("[%d] ", id), "refer": nil}
	} else {
		log.Println(err)
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "added failed", "refer": nil}
	}
	this.ServeJson()
}

// 获取文章
type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	id, err := this.GetInt("id")
	uri := this.Ctx.Input.Param(":uri")

	var art Article
	if nil == err {
		art, err = GetArticle(int(id))
	} else if "" != uri {
		art, err = GetArticleByUri(uri)
	} else {
		this.Abort("404")
		return
	}

	if 0 == art.Id {
		this.Abort("404")
		return
	}

	maps, err := CountByMonth()
	if nil == err {
		this.Data["count_by_month"] = maps
	}

	hottest, err := HottestArticleList()
	if nil == err {
		this.Data["hottest"] = hottest
	}

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
		this.ServeJson()
	}

	if 0 != art.Id {
		UpdateCount(art.Id)
	}

	this.Data["id"] = art.Id
	this.Data["title"] = art.Title
	this.Data["uri"] = art.Uri
	this.Data["content"] = art.Content
	this.Data["author"] = art.Author
	this.Data["time"] = art.Time
	this.Data["count"] = art.Count
	this.Data["keywords"] = art.Keywords
	this.TplNames = "article.tpl"
}

func (this *ArticleController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}

// 修改文章
type UpdateArticleController struct {
	beego.Controller
}

func (this *UpdateArticleController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}

func (this *UpdateArticleController) Post() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJson()
		return
	}

	id, err := this.GetInt64("id")
	uri := this.Ctx.Input.Param(":uri")

	newTitle := this.GetString("title")
	newContent := this.GetString("content")
	newKeywords := this.GetString("keywords")

	if "" == newTitle {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "title can not be empty", "refer": "/"}
		this.ServeJson()
	}

	var art Article

	if nil == err {
		art, err = GetArticle(int(id))
	} else if "" != uri {
		art, err = GetArticleByUri(uri)
	} else {
		this.Ctx.WriteString("not found")
	}

	art.Title = newTitle
	art.Content = newContent
	art.Keywords = newKeywords

	err = UpdateArticle(id, uri, art)

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "update failed", "refer": "/"}
		this.ServeJson()
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "update success", "refer": "/"}
		this.ServeJson()
	}

}

// 删除文章
type DeleteArticleController struct {
	beego.Controller
}

func (this *DeleteArticleController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}

func (this *DeleteArticleController) Post() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJson()
		return
	}

	id, err := this.GetInt64("id")
	title := this.Ctx.Input.Param(":title")

	if err != nil {
		id = 0
	}

	num, err := DeleteArticle(id, title)

	if nil != err {
		log.Fatal(err)
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "delete faild", "refer": nil}
		this.ServeJson()
	} else if 0 == num {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "articles dose not exist", "refer": nil}
		this.ServeJson()
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": fmt.Sprintf("[%d]", num) + " articles deleted", "refer": nil}
		this.ServeJson()
	}
}

// 文章列表页
type ArticleListPageController struct {
	beego.Controller
}

func (this *ArticleListPageController) Get() {
	s := this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(s)
	if nil != err || page < 0 {
		page = 1
	}

	maps, nextPageFlag, _, err := ListPage(int(page), 30)
	var prevPageFlag bool
	if 1 == page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJson()
	} else {
		this.Data["list"] = maps
		this.Data["prev_page"] = page - 1
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = page + 1
		this.Data["next_page_flag"] = nextPageFlag
		this.TplNames = "list.tpl"
	}
}

func (this *ArticleListPageController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request, only get is avalible", "refer": "/"}
	this.ServeJson()
}

// 管理- 获取文章列表
type AdminArticleListController struct {
	beego.Controller
}

func (this *AdminArticleListController) Get() {
	s := this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(s)
	if nil != err || page < 0 {
		page = 1
	}

	maps, nextPage, pages, err := ListPage(int(page), 10)
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJson()
	} else {
		this.Data["json"] = map[string]interface{}{
			"result":   true,
			"msg":      "get list success",
			"refer":    "/",
			"pages":    pages,
			"nextPage": nextPage,
			"data":     maps,
			"page":     page,
		}
		this.ServeJson()
	}

}

func (this *AdminArticleListController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request, only get is avalible", "refer": "/"}
	this.ServeJson()
}

// 按月归档-按月文章列表
type ArchiveController struct {
	beego.Controller
}

func (this *ArchiveController) Get() {
	s := this.Ctx.Input.Param(":year")
	year, err := strconv.Atoi(s)
	if nil != err || year < 0 {
		year = 1970
	}

	s = this.Ctx.Input.Param(":month")
	month, err := strconv.Atoi(s)
	if nil != err || month < 0 {
		month = 1
	}

	s = this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(s)
	if nil != err || page < 0 {
		page = 1
	}

	maps, nextPageFlag, pages, err := ListByMonth(year, month, page, 10)

	if pages < int(page) {
		page = pages
	}

	var prevPageFlag bool
	if 1 == page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}

	if nil == err {
		this.Data["prev_page"] = fmt.Sprintf("/archive/%d/%d/%d", year, month, page-1)
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = fmt.Sprintf("/archive/%d/%d/%d", year, month, page+1)
		this.Data["next_page_flag"] = nextPageFlag
		this.Data["articles_in_page"] = maps
	}

	hottest, err := HottestArticleList()

	if nil == err {
		this.Data["hottest"] = hottest
	}
	monthMaps, err := CountByMonth()

	if nil == err {
		this.Data["count_by_month"] = monthMaps
	}

	this.Data["title"] = fmt.Sprintf("- %d年%d月", year, month)

	this.TplNames = "index.tpl"
}

func (this *ArchiveController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request, only get is avalible", "refer": "/"}
	this.ServeJson()
}
