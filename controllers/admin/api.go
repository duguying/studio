package admin

import (
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
	// . "github.com/duguying/blog/models"
	// "github.com/duguying/blog/utils"
	// "github.com/gogather/com"
	// "time"
)

// 注册
type AdminApiController struct {
	controllers.BaseController
}

func (this *AdminApiController) NavList() {

	this.Data["json"] = [...]interface{}{
		map[string]interface{}{
			"title": "添加文章",
			"uri":   "new_article",
		},
		map[string]interface{}{
			"title": "文章管理",
			"uri":   "manage_article",
		},
		map[string]interface{}{
			"title": "项目管理",
			"uri":   "project_manage",
		},
		map[string]interface{}{
			"title": "OSS管理",
			"uri":   "oss_manage",
		},
	}

	this.ServeJson()
}
