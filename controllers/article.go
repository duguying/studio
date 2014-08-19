package controllers

import (
	. "blog/models"
	"github.com/astaxie/beego"
	// "strconv"
)

/**
 * 添加文章
 */
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

	id, err := AddArticle(title, content, keywords)
	if nil == err {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "success added, id " + string(id), "refer": nil}
	} else {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": nil}
	}
	this.ServeJson()
}

/**
 * 文章
 */
type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	id, err := this.GetInt("id")
	title := this.Ctx.Input.Param(":title")

	var art Article
	if nil == err {
		art, err = GetArticle(int(id))
	} else if "" != title {
		art, err = GetArticleByTitle(title)
	} else {
		this.Ctx.WriteString("not found")
	}

	this.Data["id"] = art.Id
	this.Data["title"] = art.Title
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
