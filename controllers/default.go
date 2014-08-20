package controllers

import (
	// "blog/utils"
	// "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	// u := this.GetSession("username")

	// uinfo := "nil"
	// if nil != u {
	// 	uinfo = u.(string)
	// }

	// this.Ctx.WriteString("home page, random string: " + utils.RandString(10) + "\nuser:" + uinfo)

	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	this.Ctx.WriteString("home page")
}
