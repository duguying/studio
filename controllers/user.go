package controllers

import (
	. "blog/models"
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego"
	// "strconv"
)

/**
 * 注册
 */
type RegistorController struct {
	beego.Controller
}

func (this *RegistorController) Get() {
	registorable, err := beego.AppConfig.Bool("registorable")
	if registorable || nil != err {
		this.TplNames = "registor.tpl"
	} else {
		this.Ctx.WriteString("registor closed")
	}
}

func (this *RegistorController) Post() {
	registorable, err := beego.AppConfig.Bool("registorable")
	if nil != err {
		// default registorable is true, do nothing
	} else if !registorable {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "registorable is false", "refer": "/"}
		this.ServeJson()
		return
	}

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
		this.Data["json"] = map[string]interface{}{"result": true, "msg": fmt.Sprintf("[%d] ", id) + "registor success", "refer": "/"}
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
			this.SetSession("username", username)
			this.Data["json"] = map[string]interface{}{"result": true, "msg": "user[" + user.Username + "] login success ", "refer": "/"}
		} else {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "login failed ", "refer": "/"}
		}
	}
	this.ServeJson()
}
