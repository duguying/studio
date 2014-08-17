package controllers

import (
	"blog/utils"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("home page, random string: " + utils.RandString(10))
}

func (this *MainController) Post() {
	this.Ctx.WriteString("home page")
}
