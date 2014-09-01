package controllers

import (
	// "blog/utils"
	. "blog/models"
	"fmt"
	"github.com/astaxie/beego"
	// "log"
	"blog/utils"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	page, err := this.GetInt("page")
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

	maps, nextPageFlag, totalPages, err := ListPage(int(page))

	if totalPages < int(page) {
		page = int64(totalPages)
	}

	var prevPageFlag bool
	if 1 == page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}

	if nil == err {
		this.Data["prev_page"] = page - 1
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = page + 1
		this.Data["next_page_flag"] = nextPageFlag
		this.Data["articles_in_page"] = maps
	}

	hottest, err := HottestArticleList()

	if nil == err {
		this.Data["hottest"] = hottest
	}

	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	this.Ctx.WriteString("home page")
}

/**
 * Upload
 */
type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {
	conf := utils.ReadFile("conf/ueditor.json")
	this.Ctx.WriteString(conf)
}

func (this *UploadController) Post() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJson()
		return
	}

	f, h, err := this.GetFile("upfile")
	if nil == err {
		f.Close()
	}

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "upload failed", "refer": nil}
		this.ServeJson()
		return
	}

	err = this.SaveToFile("upfile", "static/upload/"+h.Filename)

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "upload failed", "refer": nil}
		this.ServeJson()
		return
	}

	this.Data["json"] = map[string]interface{}{
		"url":      fmt.Sprintf("/static/upload/%s", h.Filename), //保存后的文件路径
		"title":    "",                                           //文件描述，对图片来说在前端会添加到title属性上
		"original": h.Filename,                                   //原始文件名
		"state":    "SUCCESS",
	}
	this.ServeJson()
}

/**
 * 按关键词列表
 */
type TagController struct {
	beego.Controller
}

func (this *TagController) Get() {
	tag := this.Ctx.Input.Param(":tag")

	page, err := this.GetInt("page")
	if nil != err || page <= 0 {
		page = 1
	}

	maps, nextPageFlag, totalPages, err := ListByKeyword(tag, int(page))

	if totalPages < int(page) {
		page = int64(totalPages)
	}

	var prevPageFlag bool
	if 1 == page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}

	if nil == err {
		this.Data["prev_page"] = page - 1
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = page + 1
		this.Data["next_page_flag"] = nextPageFlag
		this.Data["articles_in_page"] = maps
	}

	this.TplNames = "index.tpl"
}

func (this *TagController) Post() {
	this.Ctx.WriteString("home page")
}
