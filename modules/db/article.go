// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/service/models"
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

func PageArticleMonthly(year, month uint, page uint, pageSize uint) (total uint, list []*dbmodels.Article, err error) {
	total = 0
	errs := g.Db.Table("articles").Where("status=? and year(created_at)=? and month(created_at)=?", 1, year, month).Count(&total).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	list = []*dbmodels.Article{}
	errs = g.Db.Table("articles").Where("status=? and year(created_at)=? and month(created_at)=?", 1, year, month).Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	return total, list, nil
}

func ArticleToContent(articles []*dbmodels.Article) (articleContent []*models.ArticleContent) {
	articleContent = []*models.ArticleContent{}
	for _, article := range articles {
		articleContent = append(articleContent, article.ToArticleContent())
	}
	return articleContent
}

func ArticleToTitle(articles []*dbmodels.Article) (articleTitle []*models.ArticleTitle) {
	articleTitle = []*models.ArticleTitle{}
	for _, article := range articles {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle
}

func HotArticleTitle(num uint) (articleTitle []*models.ArticleTitle, err error) {
	list := []*dbmodels.Article{}
	errs := g.Db.Table("articles").Order("count desc").Limit(num).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	articleTitle = []*models.ArticleTitle{}
	for _, article := range list {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle, nil
}

func sortAndParse(arch []*dbmodels.ArchInfo) []*dbmodels.ArchInfo {
	sortArch := []*dbmodels.ArchInfo{}
	for _, item := range arch {
		sortArch = append([]*dbmodels.ArchInfo{item}, sortArch...)
	}
	return sortArch
}

func MonthArch() (archInfos []*dbmodels.ArchInfo, err error) {
	archInfos = []*dbmodels.ArchInfo{}
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

func AddArticle(aar *models.Article, author string, authorId uint) (art *dbmodels.Article, err error) {
	art = &dbmodels.Article{
		Title:     aar.Title,
		Uri:       aar.Uri,
		Keywords:  strings.Join(aar.Keywords, ","),
		Abstract:  aar.Abstract,
		Type:      aar.Type,
		Content:   aar.Content,
		Author:    author,
		AuthorId:  authorId,
		Status:    dbmodels.ArtStatus_Draft,
		CreatedAt: time.Now(),
	}

	if !aar.Draft {
		art.Status = dbmodels.ArtStatus_Publish
		art.PublishTime = time.Now()
	}

	errs := g.Db.Model(dbmodels.Article{}).Create(art).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return art, nil
}

func PublishArticle(aid uint, uid uint) (err error) {
	errs := g.Db.Model(dbmodels.Article{}).Where("id=?", aid).UpdateColumns(dbmodels.Article{
		Status:      dbmodels.ArtStatus_Publish,
		PublishTime: time.Now(),
		UpdatedBy:   uid,
	}).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}
	return nil
}

func DeleteArticle(aid uint, uid uint) (err error) {
	errs := g.Db.Model(dbmodels.Article{}).Where("id=?", aid).UpdateColumn(dbmodels.Article{
		UpdatedBy: uid,
	}).Delete(&dbmodels.Article{}).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return errs[0]
	}
	return nil
}

func ListAllArticleUri() (list []*dbmodels.Article, err error) {
	list = []*dbmodels.Article{}
	errs := g.Db.Table("articles").Select("uri").Where("status=?", 1).Order("id desc").Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return list, nil
}
