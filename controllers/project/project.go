package project

import (
	"fmt"
	"github.com/duguying/blog/controllers"
	. "github.com/duguying/blog/models"
	"github.com/gogather/com/log"
	"strconv"
)

type ProjectListController struct {
	controllers.BaseController
}

func (this *ProjectListController) Get() {
	page, err := this.GetInt64("page")
	if nil != err || page < 0 {
		page = 0
	}

	s := this.Ctx.Input.Param(":page")
	pageParm, err := strconv.Atoi(s)
	if nil != err || pageParm < 0 {
		pageParm = 0
	} else {
		page = int64(pageParm)
	}

	if 0 == page {
		page = 1
	}

	maps, err := CountByMonth()

	if nil == err {
		this.Data["count_by_month"] = maps
	}

	maps, nextPageFlag, totalPages, err := ListProjects(int(page), 10)

	if totalPages < int(page) {
		page = int64(totalPages)
	}

	log.Blueln("[page]", page)

	var prevPageFlag bool
	if 1 >= page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}

	if nil == err {
		this.Data["title"] = "项目"
		this.Data["keywords"] = "项目"
		this.Data["description"] = "独孤影的项目"
		this.Data["prev_page"] = fmt.Sprintf("/project/%d", page-1)
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = fmt.Sprintf("/project/%d", page+1)
		this.Data["next_page_flag"] = nextPageFlag
		this.Data["projects_in_page"] = maps
	}

	this.TplNames = "project.tpl"

}

func (this *ProjectListController) Post() {
	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	this.ServeJson()
}
