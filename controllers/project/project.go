package project

import (
	"fmt"
	"github.com/duguying/blog/controllers"
	"github.com/duguying/blog/models"
	"github.com/gogather/com"
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

	this.TplName = "project.tpl"

}

func (this *ProjectListController) AddProject() {

	paramsBody := string(this.Ctx.Input.RequestBody)
	var params map[string]interface{}
	p, err := com.JsonDecode(paramsBody)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "parse params failed", "refer": "/"}
		this.ServeJSON()
		return
	} else {
		params = p.(map[string]interface{})["params"].(map[string]interface{})
	}

	name := params["name"].(string)
	icon := params["icon"].(string)
	description := params["description"].(string)
	createTime := time.Now()

	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJSON()
		return
	}

	author := user.(string)

	id, err := models.AddProject(name, icon, author, description, createTime)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "添加失败", "error": err}
	} else {
		if id <= 0 {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "添加失败"}
		}
	}

	this.Data["json"] = map[string]interface{}{"result": true, "msg": "添加成功", "id": id}
	this.ServeJSON()
}

func (this *ProjectListController) DeleteProject() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJSON()
		return
	}

	paramsBody := string(this.Ctx.Input.RequestBody)
	var params map[string]interface{}
	p, err := com.JsonDecode(paramsBody)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "parse params failed", "refer": "/"}
		this.ServeJSON()
		return
	} else {
		params = p.(map[string]interface{})["params"].(map[string]interface{})
	}

	id := int64(params["id"].(float64))

	err = models.DeleteProject(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "删除失败", "error": err}
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "删除成功"}
	}

	this.ServeJSON()
}

func (this *ProjectListController) UpdateProject() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJSON()
		return
	}

//	log.Pinkln(user)

	paramsBody := string(this.Ctx.Input.RequestBody)
	var params map[string]interface{}
	p, err := com.JsonDecode(paramsBody)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "parse params failed", "refer": "/"}
		this.ServeJSON()
		return
	} else {
		params = p.(map[string]interface{})["params"].(map[string]interface{})
	}

	log.Pinkln(params)

	id := int64(params["id"].(float64))
	name := params["name"].(string)
	icon := params["icon"].(string)
	description := params["description"].(string)

	log.Pinkln(id)

	err = models.UpdateProject(id, name, icon, description)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "修改失败", "error": err}
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "修改成功"}
	}

	this.ServeJSON()
}
