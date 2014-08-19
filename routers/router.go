package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.AddArticleController{})
	beego.Router("/registor", &controllers.RegistorController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/article/:title", &controllers.ArticleController{})
	beego.Router("/article", &controllers.ArticleController{})
}
