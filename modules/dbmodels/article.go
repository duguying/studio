// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package dbmodels

import (
	"duguying/studio/g"
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"strings"
	"time"

	"github.com/gogather/blackfriday/v2"
	"github.com/gogather/com"
	"github.com/gogather/json"
)

const (
	ArtStatusDraft   = 0
	ArtStatusPublish = 1
)

var (
	ArtStatusMap = map[int]string{
		ArtStatusDraft:   "草稿",
		ArtStatusPublish: "已发布",
	}
)

const (
	ContentTypeHTML     = 0
	ContentTypeMarkDown = 1
)

type Article struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title" gorm:"index:,unique"`
	URI         string     `json:"uri" gorm:"index"`
	Keywords    string     `json:"keywords" gorm:"index:,class:FULLTEXT"`
	Abstract    string     `json:"abstract"`
	Type        int        `json:"type" gorm:"default:0;index"`
	Content     string     `json:"content" gorm:"type:longtext;index:,class:FULLTEXT"`
	Author      string     `json:"author" gorm:"index"`
	AuthorID    uint       `json:"author_id" gorm:"index"`
	Count       uint       `json:"count" gorm:"index:,sort:desc"`
	Status      int        `json:"status" gorm:"index"`
	PublishTime *time.Time `json:"publish_time" gorm:"index"`
	UpdatedBy   uint       `json:"updated_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at" gorm:"index:,sort:desc"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
}

type ArticleIndex struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Keywords    string     `json:"keywords"`
	Abstract    string     `json:"abstract"`
	Type        int        `json:"type"`
	Content     string     `json:"content"`
	Author      string     `json:"author"`
	Status      int        `json:"status"`
	PublishTime *time.Time `json:"publish_time"`
	UpdatedBy   uint       `json:"updated_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (a *Article) ToArticleIndex() *ArticleIndex {
	return &ArticleIndex{
		ID:          a.ID,
		Title:       a.Title,
		Keywords:    a.Keywords,
		Abstract:    a.Abstract,
		Type:        a.Type,
		Content:     utils.TrimHTML(a.Content),
		Author:      a.Author,
		Status:      a.Status,
		PublishTime: a.PublishTime,
		UpdatedBy:   a.UpdatedBy,
		UpdatedAt:   a.UpdatedAt,
		CreatedAt:   a.CreatedAt,
		DeletedAt:   a.DeletedAt,
	}
}

func (a *Article) String() string {
	c, _ := json.Marshal(a)
	return string(c)
}

// MarkdownFull markdown全量转html，带缓存
func (a *Article) MarkdownFull(input []byte) []byte {
	md5sign := com.Md5(string(input))
	key := "art:" + md5sign
	output, err := g.Cache.Get(key)
	if err != nil {
		htmlContent := a.markdownFull(input)
		g.Cache.SetTTL(key, string(htmlContent), time.Hour*24*30)
		return htmlContent
	}
	return []byte(output)
}

func (a *Article) markdownFull(input []byte) []byte {
	// set up the HTML renderer
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags:      blackfriday.CommonHTMLFlags,
		Extensions: blackfriday.CommonExtensions | blackfriday.LaTeXMath,
	})
	options := blackfriday.Options{
		Extensions: blackfriday.CommonExtensions | blackfriday.LaTeXMath,
	}
	return blackfriday.Markdown(input, renderer, options)
}

func (a *Article) ToArticleShowContent() *models.ArticleShowContent {
	content := []byte(a.Content)
	if a.Type == ContentTypeMarkDown {
		content = a.MarkdownFull([]byte(a.Content))
		// content = []byte(utils.ParseMath(string(content)))
	}
	tags := []string{}
	segs := strings.Split(strings.Replace(a.Keywords, "，", ",", -1), ",")
	for _, seg := range segs {
		tags = append(tags, strings.TrimSpace(seg))
	}
	return &models.ArticleShowContent{
		ID:        a.ID,
		Title:     a.Title,
		URI:       a.URI,
		Author:    a.Author,
		Tags:      tags,
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
		Content:   string(content),
	}
}

func (a *Article) ToArticleContent() *models.ArticleContent {
	tags := []string{}
	segs := strings.Split(strings.Replace(a.Keywords, "，", ",", -1), ",")
	for _, seg := range segs {
		tags = append(tags, strings.TrimSpace(seg))
	}
	return &models.ArticleContent{
		ID:        a.ID,
		Title:     a.Title,
		URI:       a.URI,
		Author:    a.Author,
		Tags:      tags,
		Type:      a.Type,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
		Content:   a.Content,
	}
}

func (a *Article) ToArticleTitle() *models.ArticleTitle {
	return &models.ArticleTitle{
		ID:        a.ID,
		Title:     a.Title,
		URI:       "/article/" + a.URI,
		Author:    a.Author,
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
	}
}

func (a *Article) ToArticleAdminTitle() *models.ArticleAdminTitle {
	return &models.ArticleAdminTitle{
		ID:         a.ID,
		Title:      a.Title,
		URI:        "/article/" + a.URI,
		Author:     a.Author,
		CreatedAt:  a.CreatedAt,
		ViewCount:  a.Count,
		Status:     a.Status,
		StatusName: ArtStatusMap[a.Status],
	}
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

func (ai *ArchInfo) ToModel() *models.ArchInfo {
	return &models.ArchInfo{
		Date:   ai.Date,
		Number: ai.Number,
		Year:   ai.Year,
		Month:  ai.Month,
	}
}

type ArchInfoList []*ArchInfo

func (al ArchInfoList) Len() int {
	return len(al)
}

func (al ArchInfoList) Less(i, j int) bool {
	return (al[i].Year*100 + al[i].Month) > (al[j].Year*100 + al[j].Month)
}

func (al ArchInfoList) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}
