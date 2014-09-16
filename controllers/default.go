package controllers

import (
	// "blog/utils"
	. "blog/models"
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego"
	// "log"
	"os"
	"strconv"
	"time"
)

// 首页
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

	maps, nextPageFlag, totalPages, err := ListPage(int(page), 6)

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
		this.Data["prev_page"] = fmt.Sprintf("/page/%d", page-1)
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = fmt.Sprintf("/page/%d", page+1)
		this.Data["next_page_flag"] = nextPageFlag
		this.Data["articles_in_page"] = maps
	}

	hottest, err := HottestArticleList()

	if nil == err {
		this.Data["hottest"] = hottest
	}

	this.Data["title"] = ""

	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	this.Ctx.WriteString("home page")
}

// 上传
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

	// 获取上传文件
	f, h, err := this.GetFile("upfile")
	if nil == err {
		f.Close()
	}

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "upload failed", "refer": nil}
		this.ServeJson()
		return
	}

	mime := h.Header.Get("Content-Type")

	// 文件保存到本地
	err = this.SaveToFile("upfile", "./static/upload/"+h.Filename)
	if nil != err {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"state":  "upload to local FAILED",
			"msg":    "upload failed",
			"refer":  nil,
		}
		this.ServeJson()
		return
	}

	// 文件保存到OSS
	t := time.Now()
	ossFilename := fmt.Sprintf("%d/%d/%d/%s", t.Year(), t.Month(), t.Day(), h.Filename)
	err = utils.OssStore(ossFilename, "static/upload/"+h.Filename)

	if nil != err {
		// 保存到oss失败
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"state":  "upload to oss FAILED, " + fmt.Sprint(err),
			"msg":    "upload failed",
			"refer":  nil,
		}
		this.ServeJson()
		return
	} else {
		// 保存到oss成功
		os.Remove("./static/upload/" + h.Filename)
		if _, err = AddFile(h.Filename, ossFilename, "oss", mime); nil != err {
			this.Data["json"] = map[string]interface{}{
				"result": false,
				"state":  "save info to database FAILED, " + fmt.Sprint(err),
				"msg":    "upload failed",
				"refer":  nil,
			}
			this.ServeJson()
		}
	}

	this.Data["json"] = map[string]interface{}{
		"url":      utils.OssGetURL(ossFilename), //保存后的文件路径
		"title":    "",                           //文件描述，对图片来说在前端会添加到title属性上
		"original": h.Filename,                   //原始文件名
		"state":    "SUCCESS",
	}
	this.ServeJson()
}

// 按关键词列表
type TagController struct {
	beego.Controller
}

func (this *TagController) Get() {
	tag := this.Ctx.Input.Param(":tag")

	s := this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(s)
	if nil != err || page <= 0 {
		page = 1
	}

	maps, nextPageFlag, totalPages, err := ListByKeyword(tag, int(page), 6)

	if totalPages < int(page) {
		page = totalPages
	}

	var prevPageFlag bool
	if 1 == page {
		prevPageFlag = false
	} else {
		prevPageFlag = true
	}

	if nil == err {
		this.Data["prev_page"] = fmt.Sprintf("/tag/%s/%d", tag, page-1)
		this.Data["prev_page_flag"] = prevPageFlag
		this.Data["next_page"] = fmt.Sprintf("/tag/%s/%d", tag, page+1)
		this.Data["next_page_flag"] = nextPageFlag
		this.Data["articles_in_page"] = maps
	}

	hottest, err := HottestArticleList()

	if nil == err {
		this.Data["hottest"] = hottest
	}
	monthMaps, err := CountByMonth()

	if nil == err {
		this.Data["count_by_month"] = monthMaps
	}

	this.Data["title"] = "- " + tag

	this.TplNames = "index.tpl"
}

func (this *TagController) Post() {
	this.Ctx.WriteString("home page")
}
