package controllers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/models"
	"github.com/duguying/blog/utils"
	"github.com/gogather/com/log"
	"strings"
)

// Controller基类继承封装
type AdminBaseController struct {
	beego.Controller
}

// run before get
func (this *AdminBaseController) Prepare() {
	// login status
	user := this.GetSession("username")
	if user == nil {
		this.Redirect("/login", 302)
	} else {
		// find user id
		username := user.(string)
		u, err := models.FindUser(username)
		if err != nil {
			log.Warnln(err)
		} else {
			userLog := &models.UserLog{}

			// get ip
			ipPort := this.Ctx.Request.RemoteAddr
			ipPortArr := strings.Split(ipPort, ":")
			ip := ipPortArr[0]

			// get location
			location := ""
			userLogIp, err := userLog.GetUserLogByIp(ip)
			if err == nil {
				if len(location) > 0 {
					location = userLogIp.Location
				} else {
					location = utils.GetLocation(ip)
				}
			} else {
				location = utils.GetLocation(ip)
			}

			// get user agent
			ua := this.Ctx.Request.UserAgent()

			// save data
			_, err = userLog.AddUserlog(int64(u.Id), ip, ua, location, 0)
			if err != nil {
				log.Warnln(err)
			}
		}
	}
}

// run after finished
func (this *AdminBaseController) Finish() {

}