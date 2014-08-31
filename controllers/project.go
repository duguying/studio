package controllers

import (
	// . "blog/models"
	// "fmt"
	"github.com/astaxie/beego"
	// "log"
)

type ProjectListController struct {
	beego.Controller
}

func (this *ProjectListController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}

func (this *ProjectListController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}
