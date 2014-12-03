package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
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

func (this *File) TableName() string {
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
func AddFile(filename string, path string, store string, mime string) (int64, error) {
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

	// select count(*) from file where path=path
	sql := "select count(*) as number from file where path=?"
	var maps []orm.Params
	o.Raw(sql, path).Values(&maps)
	num, _ := strconv.Atoi(maps[0]["number"].(string))

	var err error
	var id int64
	if 0 == num {
		id, err = o.Insert(&file)
	} else {
		id, err = o.Update(&file, "path")
	}

	if err == nil {
		return id, nil
	} else {
		return id, err
	}
}

// 移除文件
// id 文件id
func RemoveFile(id int) error {
	if id < 1 {
		return errors.New("id is illeage")
	}

	o := orm.NewOrm()
	_, err := o.Delete(&File{Id: id})

	return err
}

// 文件列表
// page 页码
// numPerPage 每页条数
// 返回值:
// []orm.Params 文章
// bool 是否有下一页
// int 总页数
// error 错误
func GetFileList(page int, numPerPage int) ([]orm.Params, bool, int, error) {
	// numPerPage := 6
	sql1 := "select * from file order by time desc limit ?," + fmt.Sprintf("%d", numPerPage)
	sql2 := "select count(*) as number from file"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql1, numPerPage*(page-1)).Values(&maps)
	o.Raw(sql2).Values(&maps2)

	number, _ := strconv.Atoi(maps2[0]["number"].(string))

	var addFlag int
	if 0 == (number % numPerPage) {
		addFlag = 0
	} else {
		addFlag = 1
	}

	pages := number/numPerPage + addFlag

	var flagNextPage bool
	if pages == page {
		flagNextPage = false
	} else {
		flagNextPage = true
	}

	if err == nil && num > 0 {
		return maps, flagNextPage, pages, nil
	} else {
		return nil, false, pages, err
	}
}
