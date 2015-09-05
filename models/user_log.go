package models

import (
	//	"errors"
	"github.com/astaxie/beego/orm"
	//	"github.com/gogather/com"
	// "log"
	//	"regexp"
	"time"
)

type UserLog struct {
	Id         int64
	User       int64
	Ip         string
	Ua         string
	Location   string
	Action     int
	CreateTime time.Time
}

func (this *UserLog) TableName() string {
	return "user_log"
}

func init() {
	orm.RegisterModel(new(UserLog))
}

func (this *UserLog) AddUserlog(user int64, ip string, ua string, location string, action int) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	userLog := new(UserLog)
	userLog.User = user
	userLog.Ip = ip
	userLog.Ua = ua
	userLog.Location = location
	userLog.Action = action
	return o.Insert(userLog)
}

func (this *UserLog) GetUserLogByIp(ip string) (UserLog, error) {
	o := orm.NewOrm()
	o.Using("default")
	userLog := UserLog{Ip: ip}
	err := o.Read(&userLog, "ip")
	return userLog, err
}

func (this *UserLog) IsValidLocation(data map[string]interface{}) bool {
	cityName := data["cityName"].(string)
	countryName := data["countryName"].(string)
	regionName := data["regionName"].(string)
	if len(cityName) == 0 && len(countryName) == 0 && len(regionName) == 0 {
		return false
	} else {
		return true
	}
}
