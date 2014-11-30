package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Draft struct {
	Id        int
	ArticleId int
	Title     string
	Keywords  string
	Content   string
	LastTime  time.Time
}

func (this *Draft) TableName() string {
	return "draft"
}

func init() {
	orm.RegisterModel(new(Draft))
}

// 通过文章ID获取其草稿
func GetDraft(articleId int) (Draft, error) {
	o := orm.NewOrm()
	o.Using("default")
	draft := Draft{ArticleId: articleId}
	_, _, err := o.ReadOrCreate(&draft, "article_id")
	return draft, err
}

// 通过文章ID保存其草稿
func SaveDraft(articleId int, title string, keywords string, content string) (int, error) {
	o := orm.NewOrm()
	draft := Draft{ArticleId: articleId}
	if _, _, err := o.ReadOrCreate(&draft, "ArticleId"); err == nil {
		draft.Title = title
		draft.Keywords = keywords
		draft.Content = content
		if num, err := o.Update(&draft); err == nil {
			return int(num), nil
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}
}
