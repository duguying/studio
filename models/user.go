package models

import (
	"blog/utils"
	// "fmt"
	"github.com/astaxie/beego/orm"
	// "log"
	"errors"
	"regexp"
)

type Users struct {
	Id       int
	Username string
	Password string
	Salt     string
	Email    string
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
	user.Salt = utils.RandString(10)
	user.Password = utils.Md5(password + user.Salt)
	return o.Insert(user)
}

func FindUser(username string) (Users, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username: username}
	err := o.Read(&user, "username")

	return user, err
}

func ChangeUsername(oldUsername string, newUsername string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("users").Filter("username", oldUsername).Update(orm.Params{
		"username": newUsername,
	})
	return err
}

/**
 * 修改邮箱
 */
func ChangeEmail(username string, email string) error {
	o := orm.NewOrm()
	o.Using("default")

	reg := regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`)
	result := reg.MatchString(email)
	if !result {
		return errors.New("not a email")
	}

	num, err := o.QueryTable("users").Filter("username", username).Update(orm.Params{"email": email})

	if nil != err {
		return err
	} else if 0 == num {
		return errors.New("not update")
	} else {
		return nil
	}
}
