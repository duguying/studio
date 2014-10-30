package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
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
