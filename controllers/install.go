package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/duguying/blog/utils"
	"strings"
)

// 系统安装
type InstallController struct {
	beego.Controller
}

func (this *InstallController) Get() {
	o := orm.NewOrm()

	if utils.FileExist("install.lock") {
		this.Abort("404")
	} else {
		sqls := utils.ReadFile("blog.sql")
		sqlArr := strings.Split(sqls, ";")
		for index, element := range sqlArr {
			this.Ctx.WriteString(fmt.Sprintf("[%d] ", index) + element)
			_, err := o.Raw(element).Exec()
			if err != nil {
				this.Ctx.WriteString(" ~~ ERROR!\n")
			}
		}
		utils.WriteFile("install.lock", "")
	}
}

func (this *InstallController) Post() {
	this.Abort("404")
}
