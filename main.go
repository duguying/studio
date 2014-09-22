package main

import (
	"blog/config"
	_ "blog/routers"
	"blog/utils"
	"github.com/astaxie/beego"
)

func init() {
	config.InitSql()
	config.InitCache()
}

func main() {
	beego.AddFuncMap("tags", utils.TagSplit)
	beego.Run()
}
