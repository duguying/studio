package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/page/:page", &controllers.MainController{})
	beego.Router("/tag/:tag/:page", &controllers.TagController{})
	beego.Router("/article/:uri", &controllers.ArticleController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/archive/:year/:month/:page", &controllers.ArchiveController{})
	beego.Router("/list", &controllers.ArticleListPageController{})
	beego.Router("/list/:page", &controllers.ArticleListPageController{})
	beego.Router("/project", &controllers.ProjectListController{})

	model := beego.AppConfig.String("runmode")
	if "dev" == model {
		beego.Router("/test", &controllers.TestController{})
	}

}
