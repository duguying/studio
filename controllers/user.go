package controllers

import (
	. "blog/models"
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego"
	// "strconv"
	"log"
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

/**
 * 登出
 */
type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.DelSession("username")
	this.Ctx.WriteString("you have logout")
}

func (this *LogoutController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request ", "refer": "/"}
	this.ServeJson()
}

/**
 * 测试暂用页
 */
type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	this.Data["username"] = this.GetSession("username")
	this.TplNames = "test.tpl"
}

func (this *TestController) Post() {
	this.Data["username"] = this.GetSession("username")
	this.TplNames = "test.tpl"
}

/**
 * 修改用户名
 */
type ChangeUsernameController struct {
	beego.Controller
}

func (this *ChangeUsernameController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request ", "refer": "/"}
	this.ServeJson()
}

func (this *ChangeUsernameController) Post() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJson()
		return
	}

	oldUsername := user.(string)
	newUsername := this.GetString("username")

	err := ChangeUsername(oldUsername, newUsername)

	if nil != err {
		log.Println(err)
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "change username failed", "refer": "/"}
		this.ServeJson()
	} else {
		this.SetSession("username", newUsername)
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "change username success", "refer": "/"}
		this.ServeJson()
	}
}
