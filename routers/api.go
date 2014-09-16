package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/get/user", &controllers.CurrentUserController{})
	beego.Router("/api/get/total_article_number", &controllers.TotalArticleNumberController{})
	beego.Router("/api/get/total_user_number", &controllers.TotalUserNumberController{})
	beego.Router("/api/get/server_time", &controllers.ServerTimeController{})

	beego.Router("/xmlrpc", &controllers.XmlrpcController{})
	beego.Router("/xmlrpc/test", &controllers.XmlTestController{})
}
