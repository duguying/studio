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

	"github.com/blevesearch/bleve/v2"
	"gorm.io/gorm"
)

func PageArticle(tx *gorm.DB, keyword string,
	page uint, pageSize uint, statusList []int) (total int64, list []*dbmodels.Article, err error) {
	total = 0
	query := "status in (?)"
	params := []interface{}{statusList}

	if keyword != "" {
		query = query + " and keywords like ?"
		params = append(params, fmt.Sprintf("%%%s%%", keyword))
	}

	err = tx.Table("articles").Where(query, params...).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	list = []*dbmodels.Article{}
	err = tx.Table("articles").Where(query, params...).Order("id desc").Offset(int((page - 1) * pageSize)).Limit(int(
		pageSize)).Find(&list).Error
	if err != nil {
		return 0, nil, err
	}

	return total, list, nil
}

// SearchArticle 搜索文章
func SearchArticle(tx *gorm.DB, keyword string, page, size uint) (total uint, result *bleve.SearchResult, articleMap map[uint]*dbmodels.Article, err error) {
	query := bleve.NewQueryStringQuery(keyword)
	from := size * (page - 1)
	searchRequest := bleve.NewSearchRequestOptions(query, int(size), int(from), false)
	searchRequest.SortBy([]string{"-_score"})
	searchRequest.Highlight = bleve.NewHighlight()
	result, err = g.Index.Search(searchRequest)
	if err != nil {
		return 0, nil, nil, err
	}

	total = uint(result.Total)

	// gather ids
	ids := []uint{}
	for _, hit := range result.Hits {
		id, err := strconv.ParseUint(hit.ID, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, uint(id))
	}

	// load article list as map
	list, err := LoadArticleByIds(tx, ids)
	if err != nil {
		return 0, nil, nil, err
	}
	articleMap = make(map[uint]*dbmodels.Article)
	for _, article := range list {
		articleMap[article.ID] = article
	}

	return total, result, articleMap, nil
}

func LoadArticleByIds(tx *gorm.DB, ids []uint) (list []*dbmodels.Article, err error) {
	list = []*dbmodels.Article{}
	err = tx.Table("articles").Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func PageArticleMonthly(tx *gorm.DB, year, month uint, page uint, pageSize uint) (total int64,
	list []*dbmodels.Article,
	err error) {
	total = 0
	err = tx.Table("articles").Where("status=? and year(created_at)=? and month(created_at)=?", 1, year,
		month).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	list = []*dbmodels.Article{}
	err = tx.Table("articles").Where("status=? and year(created_at)=? and month(created_at)=?", 1, year,
		month).Order("id desc").Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&list).Error
	if err != nil {
		return 0, nil, err
	}

	return total, list, nil
}

// ArticleToShowContent article转显示结构
func ArticleToShowContent(articles []*dbmodels.Article) (articleContent []*models.ArticleShowContent) {
	articleContent = []*models.ArticleShowContent{}
	for _, article := range articles {
		articleContent = append(articleContent, article.ToArticleShowContent())
	}
	return articleContent
}

// ArticleToTitle article转标题
func ArticleToTitle(articles []*dbmodels.Article) (articleTitle []*models.ArticleTitle) {
	articleTitle = []*models.ArticleTitle{}
	for _, article := range articles {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle
}

// HotArticleTitle 获取topN热文标题
func HotArticleTitle(tx *gorm.DB, num uint) (articleTitle []*models.ArticleTitle, err error) {
	var list []*dbmodels.Article
	err = tx.Model(dbmodels.Article{}).Order("count desc").Limit(int(num)).Find(&list).Error
	if err != nil {
		return nil, err
	}
	articleTitle = []*models.ArticleTitle{}
	for _, article := range list {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle, nil
}

// MonthArch 按月归档
func MonthArch(tx *gorm.DB) (archInfos []*dbmodels.ArchInfo, err error) {
	var list []*dbmodels.Article
	archInfos = []*dbmodels.ArchInfo{}
	err = tx.Model(dbmodels.Article{}).Select("created_at").Where("status=?", 1).Find(&list).Error
	if err != nil {
		return nil, err
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

// GetArticle 通过URI获取文章
func GetArticle(tx *gorm.DB, uri string) (art *dbmodels.Article, err error) {
	art = &dbmodels.Article{}
	err = tx.Table("articles").Where("uri=?", uri).First(art).Error
	if err != nil {
		return nil, err
	}
	return art, nil
}

// GetArticleByID 通过ID获取文章
func GetArticleByID(tx *gorm.DB, aid uint) (art *dbmodels.Article, err error) {
	art = &dbmodels.Article{}
	err = tx.Table("articles").Where("id=?", aid).First(art).Error
	if err != nil {
		return nil, err
	}
	return art, nil
}

// AddArticle 添加文章
func AddArticle(tx *gorm.DB, aar *models.Article, author string, authorID uint) (art *dbmodels.Article, err error) {
	now := time.Now()
	art = &dbmodels.Article{
		Title:     aar.Title,
		URI:       aar.URI,
		Keywords:  strings.Join(aar.Keywords, ","),
		Abstract:  aar.Abstract,
		Type:      aar.Type,
		Content:   aar.Content,
		Author:    author,
		AuthorID:  authorID,
		Status:    dbmodels.ArtStatusDraft,
		CreatedAt: now,
	}

	if !aar.Draft {
		art.Status = dbmodels.ArtStatusPublish
		art.PublishTime = &now
	}

	err = tx.Model(dbmodels.Article{}).Create(art).Error
	if err != nil {
		return nil, err
	}

	err = g.Index.Index(fmt.Sprintf("%d", art.ID), art.ToArticleIndex())
	if err != nil {
		return nil, err
	}

	return art, nil
}

// PublishArticle 发布文章
func PublishArticle(tx *gorm.DB, aid uint, publish bool, uid uint) (err error) {
	now := time.Now()
	status := dbmodels.ArtStatusPublish
	if !publish {
		status = dbmodels.ArtStatusDraft
	}

	err = tx.Model(dbmodels.Article{}).Where("id=?", aid).UpdateColumns(dbmodels.Article{
		Status:      status,
		PublishTime: &now,
		UpdatedBy:   uid,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteArticle 删除文章
func DeleteArticle(tx *gorm.DB, aid uint, uid uint) (err error) {
	art, err := GetArticleByID(tx, aid)
	if err != nil {
		return err
	}

	if art.UpdatedBy != uid {
		return fmt.Errorf("无权限删除")
	}

	err = tx.Model(dbmodels.Article{}).Where("id=?", aid).Delete(&dbmodels.Article{}).Error
	if err != nil {
		return err
	}

	err = g.Index.Delete(fmt.Sprintf("%d", aid))
	if err != nil {
		return err
	}

	return nil
}

// ListAllArticleURI 列举所有文章URI
func ListAllArticleURI(tx *gorm.DB) (list []*dbmodels.Article, err error) {
	list = []*dbmodels.Article{}
	err = tx.Table("articles").Select("uri").Where("status=?", 1).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ListAllArticle 列举所有已发布文章
func ListAllArticle(tx *gorm.DB) (list []*dbmodels.Article, err error) {
	list = []*dbmodels.Article{}
	err = tx.Model(dbmodels.Article{}).Where("status=?", 1).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ListAllTags 列举所有标签
func ListAllTags(tx *gorm.DB) (tags []string, counts []int64, err error) {
	list := []*dbmodels.Article{}
	err = tx.Model(dbmodels.Article{}).Select("keywords").Where("status=?",
		dbmodels.ArtStatusPublish).Find(&list).Error
	if err != nil {
		return nil, nil, err
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
	counts = []int64{}
	for _, tag := range tags {
		total := int64(0)
		err := tx.Model(dbmodels.Article{}).Where("status=? and keywords like ?", dbmodels.ArtStatusPublish,
			fmt.Sprintf("%%%s%%", tag)).Count(&total).Error
		if err != nil {
			log.Println("count keyword failed, err:", err.Error())
		}
		counts = append(counts, total)
	}
	return tags, counts, nil
}

// UpdateArticleViewCount 更新文章阅读计数
func UpdateArticleViewCount(tx *gorm.DB, uri string, cnt int) (err error) {
	art, err := GetArticle(tx, uri)
	if err != nil {
		return err
	}
	err = tx.Model(dbmodels.Article{}).Where("uri=?", uri).Updates(map[string]interface{}{
		"count": art.Count + uint(cnt),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateArticle 更新文章
func UpdateArticle(tx *gorm.DB, id uint, article *models.Article) (err error) {
	_, err = GetArticleByID(tx, id)
	if err != nil {
		return err
	}
	fields := map[string]interface{}{
		"content": article.Content,
	}
	if article.Title != "" {
		fields["title"] = article.Title
	}
	if article.URI != "" {
		fields["uri"] = article.URI
	}
	err = tx.Model(dbmodels.Article{}).Where("id=?", id).Updates(fields).Error
	if err != nil {
		return err
	}
	return nil
}
