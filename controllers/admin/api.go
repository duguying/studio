package admin

import (
	"github.com/duguying/blog/controllers"
)

type AdminApiController struct {
	controllers.AdminBaseController
}

func (this *AdminApiController) NavList() {
	this.Data["json"] = [...]interface{}{
		map[string]interface{}{
			"title": "管理首页",
			"uri":   "",
		},
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
			"uri":   "manage_project",
		},
		map[string]interface{}{
			"title": "OSS管理",
			"uri":   "manage_oss",
		},
	}

	this.ServeJSON()
}
