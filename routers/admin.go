package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers/admin"
	"github.com/duguying/blog/controllers/article"
	"github.com/duguying/blog/controllers/index"
	"github.com/duguying/blog/controllers/project"
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
	beego.Router("/admin/*", &admin.AdminController{}) // 全匹配

	beego.Router("/add", &article.AddArticleController{})
	beego.Router("/update/:uri", &article.UpdateArticleController{})
	beego.Router("/update", &article.UpdateArticleController{})
	beego.Router("/delete/:uri", &article.DeleteArticleController{})
	beego.Router("/delete", &article.DeleteArticleController{})

	beego.Router("/upload", &index.UploadController{})

	beego.Router("/install", &index.InstallController{})

	// ng api
	beego.Router("/api/admin/navlist", &admin.AdminApiController{}, "*:NavList")
	beego.Router("/api/admin/article/page/:page", &article.AdminArticleController{}, "*:ListArticle")
	beego.Router("/api/admin/article/:id", &article.AdminArticleController{}, "*:GetArticle")
	beego.Router("/api/admin/add", &article.AdminArticleController{}, "*:AddArticle")
	beego.Router("/api/admin/save", &article.AdminArticleController{}, "*:SaveArticleAsDraft")
	beego.Router("/api/admin/delete", &article.AdminArticleController{}, "*:DelArticle")
	beego.Router("/api/admin/update", &article.AdminArticleController{}, "*:UpdateArticle")
	beego.Router("/api/admin/draft_publish", &article.AdminArticleController{}, "*:DraftPublish")
	beego.Router("/api/admin/project/:id", &article.AdminProjectController{}, "*:GetProject")
	beego.Router("/api/admin/project/list/:page", &article.AdminProjectController{}, "*:ListProject")
	beego.Router("/api/admin/project/delete", &project.ProjectListController{}, "*:DeleteProject")
	beego.Router("/api/admin/project/add", &project.ProjectListController{}, "*:AddProject")
	beego.Router("/api/admin/project/update", &project.ProjectListController{}, "*:UpdateProject")
}
