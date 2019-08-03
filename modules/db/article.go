// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"github.com/gogather/json"
	"strconv"
	"strings"
	"time"
)

func PageArticle(page uint, pageSize uint) (total uint, list []*dbmodels.Article, err error) {
	total = 0
	errs := g.Db.Table("articles").Where("status=?", 1).Count(&total).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	list = []*dbmodels.Article{}
	errs = g.Db.Table("articles").Where("status=?", 1).Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	return total, list, nil
}

func ArticleToContent(articles []*dbmodels.Article) (articleContent []*dbmodels.WrapperArticleContent) {
	articleContent = []*dbmodels.WrapperArticleContent{}
	for _, article := range articles {
		articleContent = append(articleContent, article.ToArticleContent())
	}
	return articleContent
}

func ArticleToTitle(articles []*dbmodels.Article) (articleTitle []*dbmodels.WrapperArticleTitle) {
	articleTitle = []*dbmodels.WrapperArticleTitle{}
	for _, article := range articles {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle
}

func HotArticleTitle(num uint) (articleTitle []*dbmodels.WrapperArticleTitle, err error) {
	list := []*dbmodels.Article{}
	errs := g.Db.Table("articles").Order("count desc").Limit(num).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	articleTitle = []*dbmodels.WrapperArticleTitle{}
	for _, article := range list {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle, nil
}

type ArchInfo struct {
	Date   string `json:"date"`
	Number uint   `json:"number"`
	Year   uint   `json:"year"`
	Month  uint   `json:"month"`
}

func (ai *ArchInfo) String() string {
	c, _ := json.Marshal(ai)
	return string(c)
}

func (ai *ArchInfo) Parse() {
	segs := strings.Split(ai.Date, "-")
	if len(segs) > 1 {
		month, _ := strconv.ParseInt(segs[1], 10, 32)
		ai.Month = uint(month)
	}
	if len(segs) > 0 {
		year, _ := strconv.ParseInt(segs[0], 10, 32)
		ai.Year = uint(year)
	}
}

func sortAndParse(arch []*ArchInfo) []*ArchInfo {
	sortArch := []*ArchInfo{}
	for _, item := range arch {
		item.Parse()
		sortArch = append([]*ArchInfo{item}, sortArch...)
	}
	return sortArch
}

func MonthArch() (archInfos []*ArchInfo, err error) {
	archInfos = []*ArchInfo{}
	errs := g.Db.Table("articles").Select("DATE_FORMAT(created_at,'%Y-%m') as date,count(*) as number").Where("status=?", 1).Group("date").Find(&archInfos).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	archInfos = sortAndParse(archInfos)
	for _, arch := range archInfos {
		arch.Date = strings.Replace(arch.Date, "-", "年", -1) + "月"
	}
	return archInfos, nil
}

func GetArticle(uri string) (art *dbmodels.Article, err error) {
	art = &dbmodels.Article{}
	errs := g.Db.Table("articles").Where("uri=?", uri).First(art).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return art, nil
}

func GetArticleById(aid uint) (art *dbmodels.Article, err error) {
	art = &dbmodels.Article{}
	errs := g.Db.Table("articles").Where("id=?", aid).First(art).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return art, nil
}

func AddArticle(title string, uri string, keywords []string, abstract string, content string, author string, authorId uint, status int) (art *dbmodels.Article, err error) {
	article := &dbmodels.Article{
		Title:       title,
		Uri:         uri,
		Keywords:    strings.Join(keywords, ","),
		Abstract:    abstract,
		Author:      author,
		AuthorId:    authorId,
		Status:      status,
		PublishTime: time.Now(),
		CreatedAt:   time.Now(),
	}
	errs := g.Db.Table("articles").Create(article).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return article, nil
}

func PublishArticle(aid uint, uid uint) (err error) {
	errs := g.Db.Table("articles").Where("id=?", aid).UpdateColumns(dbmodels.Article{
		Status:      dbmodels.ART_STATUS_PUBLISH,
		PublishTime: time.Now(),
	}).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}
	return nil
}

func DeleteArticle(aid uint, uid uint) (err error) {
	errs := g.Db.Table("articles").Where("id=?", aid).UpdateColumn(dbmodels.Article{
		Status: dbmodels.ART_STATUS_DELETE,
	}).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}
	return nil
}
