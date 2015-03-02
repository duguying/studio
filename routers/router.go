package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
	"github.com/duguying/blog/controllers/admin"
	"github.com/duguying/blog/controllers/article"
	"github.com/duguying/blog/controllers/index"
	"github.com/duguying/blog/controllers/project"
)

func init() {
	beego.Router("/", &index.MainController{})
	beego.Router("/page/:page", &index.MainController{})
	beego.Router("/tag/:tag/:page", &index.TagController{})
	beego.Router("/article/:uri", &article.ArticleController{})
	beego.Router("/article", &article.ArticleController{})
	beego.Router("/archive/:year/:month/:page", &article.ArchiveController{})
	beego.Router("/list", &article.ArticleListPageController{})
	beego.Router("/list/:page", &article.ArticleListPageController{})
	beego.Router("/project", &project.ProjectListController{})
	beego.Router("/project/:page", &project.ProjectListController{})
	beego.Router("/about/statistics", &index.StatisticsController{})
	beego.Router("/about/blog", &index.AboutBlogController{})
	beego.Router("/about/resume", &index.ResumeController{})
	beego.Router("/logo", &index.LogoController{})
	beego.Router("/favicon", &index.SiteIconController{})

	beego.ErrorController(&controllers.ErrorController{})

	model := beego.AppConfig.String("runmode")
	if "dev" == model {
		beego.Router("/test", &admin.TestController{})
	}

}
