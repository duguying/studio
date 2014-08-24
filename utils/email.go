package utils

import (
	"github.com/astaxie/beego"
	"net/smtp"
	"strings"
)

/*
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 */

func SendMail(to string, subject string, body string) error {
	user := beego.AppConfig.String("adminemail")
	password := beego.AppConfig.String("adminemailpass")
	host := beego.AppConfig.String("adminemailhost")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	content_type = "Content-type:text/html;charset=utf-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
