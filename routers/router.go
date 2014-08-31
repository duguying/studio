package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/page/:page", &controllers.MainController{})
	beego.Router("/tag/:tag", &controllers.TagController{})
	beego.Router("/registor", &controllers.RegistorController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/rename", &controllers.ChangeUsernameController{})
	beego.Router("/email", &controllers.SetEmailController{})
	beego.Router("/add", &controllers.AddArticleController{})
	beego.Router("/article/:title", &controllers.ArticleController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/update/:title", &controllers.UpdateArticleController{})
	beego.Router("/update", &controllers.UpdateArticleController{})
	beego.Router("/delete/:title", &controllers.DeleteArticleController{})
	beego.Router("/delete", &controllers.DeleteArticleController{})
	beego.Router("/password/sendemail", &controllers.SendEmailToGetBackPasswordController{})
	beego.Router("/password/reset", &controllers.SetPasswordController{})
	beego.Router("/password/change", &controllers.ChangePasswordController{})
	beego.Router("/password/reset/:varify", &controllers.SetPasswordController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/project", &controllers.ProjectListController{})

	beego.Router("/test", &controllers.TestController{})
}
