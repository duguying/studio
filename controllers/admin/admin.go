package admin

import (
	"github.com/duguying/blog/controllers"
)

// 管理面板
type AdminController struct {
	controllers.BaseController
}

func (this *AdminController) Get() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Redirect("/login", 302)
	} else {
		this.TplNames = "adminpannel.tpl"
	}
}

func (this *AdminController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}
