package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type File struct {
	Id       int
	Filename string
	Path     string
	Time     time.Time
	Store    string
	Mime     string
}

func (u *File) TableName() string {
	return "file"
}

func init() {
	orm.RegisterModel(new(File))
}

// 添加文件信息到数据库
// filename 文件名
// path 路径
// store 存储类型
// mime 文件类型信息
// error 返回错误
func AddFile(filename string, path string, store string, mime string) error {
	o := orm.NewOrm()
	var file File
	file.Filename = filename
	file.Path = path
	if "local" == store {
		file.Store = "local"
	} else {
		file.Store = "oss"
	}
	file.Mime = mime

	_, err := o.Insert(&file)
	if err == nil {
		return nil
	} else {
		return err
	}
}
