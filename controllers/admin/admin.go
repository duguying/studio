package admin

import (
	"github.com/duguying/blog/controllers"
)

// 管理面板
type AdminController struct {
	controllers.AdminBaseController
}

func (this *AdminController) Get() {
	this.TplName = "admin/index.tpl"
}

func (this *AdminController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJSON()
}
