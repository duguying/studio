package article

import (
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
	. "github.com/duguying/blog/models"
	// "github.com/gogather/com/log"
	"strconv"
)

// 管理- 获取文章列表
type AdminArticleController struct {
	controllers.BaseController
}

func (this *AdminArticleController) ListArticle() {
	s := this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(s)
	if nil != err || page < 0 {
		page = 1
	}

	maps, nextPage, pages, err := ArticleListForAdmin(int(page), 10)
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJson()
	} else {
		this.Data["json"] = map[string]interface{}{
			"result":   true,
			"msg":      "get list success",
			"refer":    "/",
			"pages":    pages,
			"nextPage": nextPage,
			"data":     maps,
			"page":     page,
		}
		this.ServeJson()
	}

}

func (this *AdminArticleController) GetArticle() {
	s := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(s)
	if nil != err || id < 0 {
		id = 1
	}
	art, err := GetArticle(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "get article success",
			"data":   art,
		}
	}
	this.ServeJson()
}
