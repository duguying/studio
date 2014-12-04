package initial

import (
	"github.com/astaxie/beego"
	"github.com/duguying/blog/utils"
)

func InitTplFunc() {
	beego.AddFuncMap("tags", utils.TagSplit)
	beego.AddFuncMap("asset", utils.Fis)
}
