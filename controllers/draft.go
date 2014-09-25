package controllers

import (
	. "github.com/duguying/blog/models"
	// "fmt"
	"github.com/astaxie/beego"
)

// 保存草稿
type SaveDraftController struct {
	beego.Controller
}

func (this *SaveDraftController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "only post request avalible", "refer": nil}
	this.ServeJson()
}

func (this *SaveDraftController) Post() {
	id, err := this.GetInt("id")
	if nil != err || id < 1 {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "article id is required", "refer": nil}
		this.ServeJson()
		return
	}

	title := this.GetString("title")
	content := this.GetString("content")
	keywords := this.GetString("keywords")

	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJson()
		return
	}

	if _, err := SaveDraft(int(id), title, keywords, content); nil == err {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "draft saved", "refer": nil}
		this.ServeJson()
		return
	} else {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "save draft failed", "refer": nil}
		this.ServeJson()
		return
	}
}

// 获取草稿
type GetDraftController struct {
	beego.Controller
}

func (this *GetDraftController) Get() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "only post request avalible", "refer": nil}
	this.ServeJson()
}

func (this *GetDraftController) Post() {
	id, err := this.GetInt("id")
	if nil != err || id < 1 {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "article id is required", "refer": nil}
		this.ServeJson()
		return
	}

	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJson()
		return
	}

	draft, err := GetDraft(int(id))
	if nil == err {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "get draft success",
			"refer":  nil,
			"draft":  draft,
		}
		this.ServeJson()
		return
	} else {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get draft failed", "refer": nil}
		this.ServeJson()
		return
	}
}
