package fis

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/controllers"
	"github.com/gogather/com"
	"regexp"
	"strings"
)

type FisController struct {
	controllers.BaseController
}

func createDirs(path string) {
	createPath := ""
	pathArr := strings.Split(path, "/")
	for i := 0; i < len(pathArr); i++ {
		createPath = createPath + pathArr[i] + "/"
		if !com.FileExist(createPath) {
			com.Mkdir(createPath)
		}
	}
}

func getDir(path string) string {
	reg := regexp.MustCompile(`/[\d\D][^/]+$`)
	return reg.ReplaceAllString(path, "")
}

func (this *FisController) Receiver() {
	key := beego.AppConfig.String("fis_receiver_key")
	upKey := this.GetString("key")
	if key != upKey {
		this.Ctx.WriteString("1")
		return
	}

	to := this.GetString("to")
	dir := getDir(to)
	createDirs(dir)
	err := this.SaveToFile("file", to)
	if err != nil {
		this.Ctx.WriteString("1")
	} else {
		this.Ctx.WriteString("0")
	}

}
