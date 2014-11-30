package models

import (
	// "errors"
	// "fmt"
	"github.com/astaxie/beego/orm"
	// "strconv"
	"time"
)

type Tags struct {
	Id   int
	Name string
	Time time.Time
}

func (this *Tags) TableName() string {
	return "tags"
}

func init() {
	orm.RegisterModel(new(Tags))
}

func NewTag(tagName string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	tag := new(Tags)
	tag.Name = tagName
	return o.Insert(tag)
}
