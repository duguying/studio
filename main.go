package main

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/config"
	_ "github.com/duguying/blog/routers"
	"github.com/duguying/blog/utils"
)

func init() {
	config.InitSql()
	config.InitCache()
}

func main() {
	beego.AddFuncMap("tags", utils.TagSplit)
	beego.Run()
}
