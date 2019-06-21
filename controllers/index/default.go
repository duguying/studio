package index

import (
	"fmt"
	"log"
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
	. "github.com/duguying/blog/models"
	"github.com/duguying/blog/utils"
	"github.com/gogather/com"
	"os"
	"strconv"
	"time"
)

// 首页
type MainController struct {
	controllers.BaseController
}

func (this *MainController) Get() {
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

	// login status
	user := this.GetSession("username")
	if user == nil {
		this.Data["is_admin"] = ""
	} else {
		this.Data["is_admin"] = "admin"
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

	this.Data["title"] = "首页"
	this.Data["articleTitle"] = "独孤影的博客"
	this.Data["keywords"] = "个人网站,IT,技术,编程"
	this.Data["description"] = "独孤影的博客，记录我的编程学习之路"
	this.Data["host"] = beego.AppConfig.String("host")

	this.TplName = "index.tpl"
}

func (this *MainController) Post() {
	this.Ctx.WriteString("home page")
}

// 上传
type UploadController struct {
	controllers.BaseController
}

func (this *UploadController) Get() {
	conf, _ := com.ReadFile("conf/ueditor.json")
	this.Ctx.WriteString(conf)
}

func (this *UploadController) Post() {
	// if not login, permission deny
	user := this.GetSession("username")
	if user == nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
		this.ServeJSON()
		return
	}

	// 获取上传文件
	f, h, err := this.GetFile("upfile")
	if nil == err {
		f.Close()
	}

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "upload failed", "refer": nil}
		this.ServeJSON()
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
		this.ServeJSON()
		return
	}

	// 文件保存到OSS
	t := time.Now()
	ossFilename := fmt.Sprintf("%d/%d/%d/%s", t.Year(), t.Month(), t.Day(), h.Filename)
	err = utils.OssStore(ossFilename, "static/upload/"+h.Filename)

	// 文件名过长
	if len(h.Filename) > 96 {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"state":  "upload to oss FAILED, " + fmt.Sprint(err),
			"msg":    "upload failed",
			"refer":  nil,
		}
		this.ServeJSON()
	}

	if nil != err {
		// 保存到oss失败
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"state":  "upload to oss FAILED, " + fmt.Sprint(err),
			"msg":    "upload failed",
			"refer":  nil,
		}
		this.ServeJSON()
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
			this.ServeJSON()
		}
	}

	this.Data["json"] = map[string]interface{}{
		"url":      utils.OssGetURL(ossFilename), //保存后的文件路径
		"title":    "",                           //文件描述，对图片来说在前端会添加到title属性上
		"original": h.Filename,                   //原始文件名
		"state":    "SUCCESS",
	}
	this.ServeJSON()
}

// 按关键词列表
type TagController struct {
	controllers.BaseController
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

	this.TplName = "index.tpl"
}

func (this *TagController) Post() {
	this.Ctx.WriteString("home page")
}

// 统计
type StatisticsController struct {
	controllers.BaseController
}

func (this *StatisticsController) Get() {
	this.Data["title"] = "数据统计"
	this.Data["keywords"] = "数据统计"
	this.Data["description"] = "独孤影的代码数据统计，数据来自于Github.com"
	this.TplName = "about/statistics.tpl"
}

// 关于博客
type AboutBlogController struct {
	controllers.BaseController
}

func (this *AboutBlogController) Get() {
	this.TplName = "about/blog.tpl"
}

func (this *AboutBlogController) Post() {
	this.TplName = "about/blog.tpl"
}

// 简历
// 关于博客
type ResumeController struct {
	controllers.BaseController
}

func (this *ResumeController) Get() {
	this.Data["title"] = "个人简历"
	this.Data["keywords"] = "个人简历"
	this.Data["description"] = "我正在求职，如果你对我感兴趣欢迎联系我，这儿可以获取我的个人简历"
	this.TplName = "about/resume.tpl"
}

func (this *ResumeController) Post() {
	this.TplName = "about/resume.tpl"
}

// Logo
type LogoController struct {
	controllers.BaseController
}

func (this *LogoController) Get() {
	localfile := "tmp/logo.png"

	if !com.PathExist(localfile) {
		utils.GetImage(beego.AppConfig.String("logo"), localfile)
	}

	data, _ := utils.ReadFileByte(localfile)

	this.Ctx.Output.Body(data)
	this.Ctx.Output.ContentType("image/png")

}

func (this *LogoController) Post() {
	this.Ctx.WriteString("get method only")
}

type SiteIconController struct {
	controllers.BaseController
}

func (this *SiteIconController) Get() {
	localfile := "fis/img/favicon.ico"

	if !com.PathExist(localfile) {
		utils.GetImage(beego.AppConfig.String("logo"), localfile)
	}

	data, _ := utils.ReadFileByte(localfile)

	this.Ctx.Output.Body(data)
	this.Ctx.Output.ContentType("image/x-icon")
}

type SiteAdsenceController struct {
	controllers.BaseController
}

func (this *SiteAdsenceController) Get(){
	localfile := "static/ads.txt"

	data, err := utils.ReadFileByte(localfile)
	if err!=nil {
		log.Println("read ads.txt failed, err:", err.Error)
	}

	this.Ctx.Output.Body(data)
	this.Ctx.Output.ContentType("text/plain")
}