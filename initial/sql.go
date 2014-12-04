package initial

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogather/com"
)

func InitSql() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	if nil != err {
		port = 3306
	}

	if com.FileExist("install.lock") {
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
	} else {
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8", user, passwd, host, port))
	}

}
