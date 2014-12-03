package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
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
	beego.Router("/project/:page", &controllers.ProjectListController{})
	beego.Router("/resume/statistics", &controllers.StatisticsController{})
	beego.Router("/logo", &controllers.LogoController{})

	model := beego.AppConfig.String("runmode")
	if "dev" == model {
		beego.Router("/test", &controllers.TestController{})
	}

}
