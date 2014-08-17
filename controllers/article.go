package controllers

import (
	. "blog/models"
	"github.com/astaxie/beego"
)

type AddArticleController struct {
	beego.Controller
}

/**
 * 添加文章
 */
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
