package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/registor", &controllers.RegistorController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/rename", &controllers.ChangeUsernameController{})
	beego.Router("/email", &controllers.SetEmailController{})
	beego.Router("/add", &controllers.AddArticleController{})
	beego.Router("/draft/get", &controllers.GetDraftController{})
	beego.Router("/draft/save", &controllers.SaveDraftController{})
	beego.Router("/update/:uri", &controllers.UpdateArticleController{})
	beego.Router("/update", &controllers.UpdateArticleController{})
	beego.Router("/delete/:uri", &controllers.DeleteArticleController{})
	beego.Router("/delete", &controllers.DeleteArticleController{})
	beego.Router("/password/getback", &controllers.GetBackPasswordController{})
	beego.Router("/password/sendemail", &controllers.SendEmailToGetBackPasswordController{})
	beego.Router("/password/reset", &controllers.SetPasswordController{})
	beego.Router("/password/change", &controllers.ChangePasswordController{})
	beego.Router("/password/reset/:varify", &controllers.SetPasswordController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/article/page/:page", &controllers.AdminArticleListController{}) //TODO
}
