package fis

import (
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
	to := this.GetString("to")
	dir := getDir(to)
	createDirs(dir)
	this.SaveToFile("file", to)
	this.Ctx.WriteString("0")
}
