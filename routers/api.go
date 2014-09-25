package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
)

func init() {
	beego.Router("/api/get/user", &controllers.CurrentUserController{})
	beego.Router("/api/get/total_article_number", &controllers.TotalArticleNumberController{})
	beego.Router("/api/get/total_user_number", &controllers.TotalUserNumberController{})
	beego.Router("/api/get/server_time", &controllers.ServerTimeController{})

	beego.Router("/xmlrpc", &controllers.XmlrpcController{})
}
