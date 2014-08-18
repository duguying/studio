package controllers

import (
	. "blog/models"
	"blog/utils"
	// "fmt"
	"github.com/astaxie/beego"
)

/**
 * 注册
 */
type RegistorController struct {
	beego.Controller
}

func (this *RegistorController) Get() {
	this.TplNames = "registor.tpl"
}

func (this *RegistorController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	if !utils.CheckUsername(username) {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "illegal username", "refer": "/"}
		this.ServeJson()
		return
	}

	id, err := AddUser(username, password)
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "registor failed", "refer": "/"}
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": string(id) + "registor success", "refer": "/"}
	}
	this.ServeJson()
}

/**
 * 登录
 */
type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	if username == "" || password == "" {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	}

	user, err := FindUser(username)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "user does not exist", "refer": "/"}
	} else {
		passwd := utils.Md5(password + user.Salt)
		if passwd == user.Password {
			// TODO Session
			this.Data["json"] = map[string]interface{}{"result": true, "msg": "user[" + user.Username + "] login success ", "refer": "/"}
		}
	}
	this.ServeJson()
}
