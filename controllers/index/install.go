package index

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/duguying/blog/controllers"
	"github.com/duguying/blog/utils"
	"github.com/gogather/com"
	"strings"
)

// 系统安装
type InstallController struct {
	controllers.BaseController
}

func (this *InstallController) Get() {
	o := orm.NewOrm()

	if com.FileExist("install.lock") {
		this.Abort("404")
	} else {
		sqls, _ := com.ReadFile("etc/blog.sql")
		sqlArr := strings.Split(sqls, ";")
		for index, element := range sqlArr {
			this.Ctx.WriteString(fmt.Sprintf("[%d] ", index) + element)
			_, err := o.Raw(element).Exec()
			if err != nil {
				this.Ctx.WriteString(" ~~ ERROR!\n")
			}
		}
		utils.WriteFile("install.lock", " ")
	}
}

func (this *InstallController) Post() {
	this.Abort("404")
}
