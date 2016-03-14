package api

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
	. "github.com/duguying/blog/models"
	"github.com/gogather/com"
	"time"
)

// 获取当前用户名
type CurrentUserController struct {
	controllers.BaseController
}

func (this *CurrentUserController) Get() {
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "username": nil, "msg": "user not login"}
		this.ServeJSON()
	} else {
		username := user.(string)
		this.Data["json"] = map[string]interface{}{"result": true, "username": username, "msg": "user have login"}
		this.ServeJSON()
	}

}

func (this *CurrentUserController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "username": nil, "msg": "only get request is avalible"}
	this.ServeJSON()
}

// 获取当前文章总数
type TotalArticleNumberController struct {
	controllers.BaseController
}

func (this *TotalArticleNumberController) Get() {
	num, err := CountArticle()
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "number": nil, "msg": "get number failed"}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "number": num, "msg": "get number success"}
		this.ServeJSON()
	}

}

func (this *TotalArticleNumberController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "only get request avalible"}
	this.ServeJSON()
}

// 获取当前用户总数
type TotalUserNumberController struct {
	controllers.BaseController
}

func (this *TotalUserNumberController) Get() {
	num, err := CountUser()
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "number": nil, "msg": "get number failed"}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "number": num, "msg": "get number success"}
		this.ServeJSON()
	}
}

func (this *TotalUserNumberController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "only get request avalible"}
	this.ServeJSON()
}

// 获取服务器当前时间
type ServerTimeController struct {
	controllers.BaseController
}

func (this *ServerTimeController) Get() {
	now := time.Now()
	this.Data["json"] = map[string]interface{}{
		"result": true,
		"msg":    "user not login",
		"year":   now.Year(),
		"month":  now.Month(),
		"day":    now.Day(),
		"h":      now.Hour(),
		"m":      now.Minute(),
		"s":      now.Second(),
	}
	this.ServeJSON()
}

func (this *ServerTimeController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "only get request is avalible"}
	this.ServeJSON()
}

// map.json 接口
type MapJsonController struct {
	controllers.BaseController
}

func (this *MapJsonController) Get() {
	staticMap := beego.AppConfig.String("static_map")
	content, _ := com.ReadFile(staticMap)
	data, err := com.JsonDecode(content)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "can not get map.json"}
	} else {
		this.Data["json"] = data
	}

	this.ServeJSON()
}
