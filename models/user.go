package models

import (
	"blog/utils"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int64
	Username string
	Password string
	Salt     string
}

func (u *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

func AddUser(username string, password string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.Username = username
	user.Password = password
	user.Salt = utils.RandString(10)
	return o.Insert(user)
}
