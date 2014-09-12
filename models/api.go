package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

// 统计文章数目
func CountArticle() (int, error) {
	sql := "select count(*) as number from article"
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw(sql).Values(&maps)

	return strconv.Atoi(maps[0]["number"].(string))
}

// 统计用户数目
func CountUser() (int, error) {
	sql := "select count(*) as number from users"
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw(sql).Values(&maps)

	return strconv.Atoi(maps[0]["number"].(string))
}
