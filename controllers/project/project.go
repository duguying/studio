package project

import (
	"fmt"
	"github.com/duguying/blog/controllers"
	"github.com/duguying/blog/models"
	"github.com/gogather/com/log"
	"strconv"
	"time"
)

type ProjectListController struct {
	controllers.BaseController
}

func (this *ProjectListController) PageProjects() {
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

	maps, nextPageFlag, totalPages, err := models.ListProjects(int(page), 10)

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

func (this *ProjectListController) AddProject() {
	name := this.GetString("name", "")
	icon := this.GetString("icon", "")
	author := this.GetString("author", "")
	description := this.GetString("description", "")
	createTime := time.Now()

	id, err := models.AddProject(name, icon, author, description, createTime)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "添加失败", "error": err}
	} else {
		if id <= 0 {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "添加失败"}
		}
	}

	this.Data["json"] = map[string]interface{}{"result": true, "msg": "添加成功", "id": id}
	this.ServeJson()
}

func (this *ProjectListController) DeleteProject() {
	id, err := this.GetInt64("id", 0)
	if err != nil {
		id = 0
	}

	err = models.DeleteProject(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "删除失败", "error": err}
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "删除成功"}
	}

	this.ServeJson()
}
