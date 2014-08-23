package main

import (
	_ "blog/routers"
	// "blog/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	if nil != err {
		port = 3306
	}

	// to := "xxxx@gmail.com;ssssss@gmail.com"
	// subject := "Test send email by golang"
	// body := `
	//    <html>
	//    <body>
	//    <h3>
	//    "Test send email by golang"
	//    </h3>
	//    </body>
	//    </html>
	// `
	// utils.SendMail(to, subject, body, "html")

	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
}

func main() {
	beego.Run()
}
