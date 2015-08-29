package initial

import (
	"github.com/astaxie/beego"
)

func InitEnv() {
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		beego.SetStaticPath("/static", "static")
	}
}
