package controllers

import (
	// . "blog/models"
	// "fmt"
	"github.com/astaxie/beego"
	// "log"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.TplNames = "adminpannel.tpl"
}

func (this *AdminController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}
