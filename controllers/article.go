package controllers

import (
	. "blog/models"
	"fmt"
	"github.com/astaxie/beego"
	"log"
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
	}

	username := user.(string)

	id, err := AddArticle(title, content, keywords, username)
	if nil == err {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "success added, id " + fmt.Sprintf("[%d] ", id), "refer": nil}
	} else {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": nil}
	}
	this.ServeJson()
}

/**
 * 获取文章
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

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
		this.ServeJson()
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

/**
 * 修改文章
 */
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

	id, err := this.GetInt("id")
	title := this.Ctx.Input.Param(":title")

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
	} else if "" != title {
		art, err = GetArticleByTitle(title)
	} else {
		this.Ctx.WriteString("not found")
	}

	art.Title = newTitle
	art.Content = newContent
	art.Keywords = newKeywords

	err = UpdateArticle(id, title, art)

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "update failed", "refer": "/"}
		this.ServeJson()
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "update success", "refer": "/"}
		this.ServeJson()
	}

}

/**
 * 删除文章
 */
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

	id, err := this.GetInt("id")
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
