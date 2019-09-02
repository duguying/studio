// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package db

import (
	"duguying/studio/g"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func PageArticle(keyword string, page uint, pageSize uint) (total uint, list []*dbmodels.Article, err error) {
	total = 0
	query := "status=?"
	params := []interface{}{1}
	if keyword != "" {
		query = query + " and keywords like ?"
		params = append(params, fmt.Sprintf("%%%s%%", keyword))
	}

	errs := g.Db.Table("articles").Where(query, params...).Count(&total).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	list = []*dbmodels.Article{}
	errs = g.Db.Table("articles").Where(query, params...).Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).GetErrors()
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

func ArticleToShowContent(articles []*dbmodels.Article) (articleContent []*models.ArticleShowContent) {
	articleContent = []*models.ArticleShowContent{}
	for _, article := range articles {
		articleContent = append(articleContent, article.ToArticleShowContent())
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

func MonthArch() (archInfos []*dbmodels.ArchInfo, err error) {
	list := []*dbmodels.Article{}
	archInfos = []*dbmodels.ArchInfo{}
	errs := g.Db.Table("articles").Select("created_at").Where("status=?", 1).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}

	// assemble ArchInfo
	archMap := map[string]uint{}
	for _, item := range list {
		key := item.CreatedAt.Format("2006-01")
		val, ok := archMap[key]
		if ok {
			val++
		} else {
			val = 1
		}
		archMap[key] = val
	}
	for key, value := range archMap {
		segs := strings.Split(key, "-")
		year, _ := strconv.ParseInt(segs[0], 10, 64)
		month, _ := strconv.ParseInt(segs[1], 10, 64)
		archInfos = append(archInfos, &dbmodels.ArchInfo{
			Date:   fmt.Sprintf("%d年%d月", year, month),
			Year:   uint(year),
			Month:  uint(month),
			Number: value,
		})
	}

	sort.Sort(dbmodels.ArchInfoList(archInfos))
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

func PublishArticle(aid uint, publish bool, uid uint) (err error) {
	status := dbmodels.ArtStatus_Publish
	if !publish {
		status = dbmodels.ArtStatus_Draft
	}

	errs := g.Db.Model(dbmodels.Article{}).Where("id=?", aid).UpdateColumns(dbmodels.Article{
		Status:      status,
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

func ListAllTags() (tags []string, counts []uint, err error) {
	list := []*dbmodels.Article{}
	errs := g.Db.Model(dbmodels.Article{}).Select("keywords").Where("status=?", dbmodels.ArtStatus_Publish).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, nil, errs[0]
	}
	tags = []string{}
	for _, item := range list {
		tgs := strings.Split(item.Keywords, ",")
		for _, tg := range tgs {
			if tg != "" {
				if !utils.StrContain(tg, tags) {
					tags = append(tags, strings.TrimSpace(tg))
				}
			}
		}
	}
	counts = []uint{}
	for _, tag := range tags {
		total := uint(0)
		errs := g.Db.Model(dbmodels.Article{}).Where("status=? and keywords like ?", dbmodels.ArtStatus_Publish, fmt.Sprintf("%%%s%%", tag)).Count(&total).GetErrors()
		if len(errs) > 0 && errs[0] != nil {
			log.Println("count keyword failed, err:", errs[0].Error())
		}
		counts = append(counts, total)
	}
	return tags, counts, nil
}
