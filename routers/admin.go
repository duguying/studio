package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers/admin"
	"github.com/duguying/blog/controllers/article"
	"github.com/duguying/blog/controllers/index"
)

func init() {
	beego.Router("/registor", &admin.RegistorController{})
	beego.Router("/login", &admin.LoginController{})
	beego.Router("/logout", &admin.LogoutController{})
	beego.Router("/rename", &admin.ChangeUsernameController{})
	beego.Router("/email", &admin.SetEmailController{})
	beego.Router("/password/getback", &admin.GetBackPasswordController{})
	beego.Router("/password/sendemail", &admin.SendEmailToGetBackPasswordController{})
	beego.Router("/password/reset", &admin.SetPasswordController{})
	beego.Router("/password/change", &admin.ChangePasswordController{})
	beego.Router("/password/reset/:varify", &admin.SetPasswordController{})
	beego.Router("/admin", &admin.AdminController{})
	beego.Router("/admin/article/page/:page", &article.AdminArticleListController{}) //TODO

	beego.Router("/add", &article.AddArticleController{})
	beego.Router("/draft/get", &article.GetDraftController{})
	beego.Router("/draft/save", &article.SaveDraftController{})
	beego.Router("/update/:uri", &article.UpdateArticleController{})
	beego.Router("/update", &article.UpdateArticleController{})
	beego.Router("/delete/:uri", &article.DeleteArticleController{})
	beego.Router("/delete", &article.DeleteArticleController{})

	beego.Router("/upload", &index.UploadController{})

	beego.Router("/install", &index.InstallController{})

	// ng api
	beego.Router("/admin/api/navlist", &admin.AdminApiController{}, "*:NavList")
}
