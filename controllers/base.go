package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

// Controller基类继承封装
type BaseController struct {
	beego.Controller
}

func (this *BaseController) Forbbiden(mark string, condition string) {
	mark = strings.ToLower(mark)
	condition = strings.ToLower(condition)

	if mark == "not" {
		if this.Data["userIs"] != condition {
			this.Redirect("/", 302)
		}
	} else {
		if this.Data["userIs"] == condition {
			this.Redirect("/", 302)
		}
	}
}

// run before get
func (this *BaseController) Prepare() {
	// login status
	user := this.GetSession("username")
	if user == nil {
		this.Data["userIs"] = ""
	} else {
		this.Data["userIs"] = "admin"
	}
	this.Data["inDev"] = beego.AppConfig.String("runmode") == "dev"
	this.Data["articleImage"] = "/logo"
	this.Data["articleTitle"] = "独孤影的博客"
}

// run after finished
func (this *BaseController) Finish() {

}
