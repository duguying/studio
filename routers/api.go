package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers/api"
	"github.com/duguying/blog/controllers/xmlrpc"
)

func init() {
	beego.Router("/api/get/user", &api.CurrentUserController{})
	beego.Router("/api/get/total_article_number", &api.TotalArticleNumberController{})
	beego.Router("/api/get/total_user_number", &api.TotalUserNumberController{})
	beego.Router("/api/get/server_time", &api.ServerTimeController{})
	beego.Router("/map.json", &api.MapJsonController{})

	beego.Router("/xmlrpc", &xmlrpc.XmlrpcController{})

}
