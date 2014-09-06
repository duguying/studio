package controllers

import (
	// . "blog/models"
	// "fmt"
	"github.com/astaxie/beego"
	// "log"
)

// 管理面板
type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Redirect("/login", 302)
	} else {
		this.TplNames = "adminpannel.tpl"
	}
}

func (this *AdminController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}

// 管理- 获取文章列表
type AdminArticleListController struct {
	beego.Controller
}

func (this *AdminArticleListController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}

func (this *AdminArticleListController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}
