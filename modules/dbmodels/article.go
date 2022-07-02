// Copyright 2017. All rights reserved.
// This file is part of blog project
// Created by duguying on 2017/11/2.

package dbmodels

import (
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gogather/json"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	ArtStatus_Draft   = 0
	ArtStatus_Publish = 1
)

const (
	ContentType_HTML     = 0
	ContentType_MarkDown = 1
)

type Article struct {
	Id          uint       `json:"id"`
	Title       string     `json:"title"`
	Uri         string     `json:"uri"`
	Keywords    string     `json:"keywords"`
	Abstract    string     `json:"abstract"`
	Type        int        `json:"type" gorm:"default:0"`
	Content     string     `json:"content" sql:"type:longtext"`
	Author      string     `json:"author"`
	AuthorId    uint       `json:"author_id"`
	Count       uint       `json:"count"`
	Status      int        `json:"status"`
	PublishTime time.Time  `json:"publish_time"`
	UpdatedBy   uint       `json:"updated_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ArticleIndex struct {
	Id          uint       `json:"id"`
	Title       string     `json:"title"`
	Keywords    string     `json:"keywords"`
	Abstract    string     `json:"abstract"`
	Type        int        `json:"type"`
	Content     string     `json:"content"`
	Author      string     `json:"author"`
	Status      int        `json:"status"`
	PublishTime time.Time  `json:"publish_time"`
	UpdatedBy   uint       `json:"updated_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (a *Article) ToArticleIndex() *ArticleIndex {
	return &ArticleIndex{
		Id:          a.Id,
		Title:       a.Title,
		Keywords:    a.Keywords,
		Abstract:    a.Abstract,
		Type:        a.Type,
		Content:     utils.TrimHtml(a.Content),
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

func (a *Article) ToArticleShowContent() *models.ArticleShowContent {
	content := []byte(a.Content)
	if a.Type == ContentType_MarkDown {
		content = blackfriday.Run([]byte(a.Content))
	}
	tags := []string{}
	segs := strings.Split(strings.Replace(a.Keywords, "，", ",", -1), ",")
	for _, seg := range segs {
		tags = append(tags, strings.TrimSpace(seg))
	}
	return &models.ArticleShowContent{
		Id:        a.Id,
		Title:     a.Title,
		Uri:       a.Uri,
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
		Id:        a.Id,
		Title:     a.Title,
		Uri:       a.Uri,
		Author:    a.Author,
		Tags:      tags,
		Type:      a.Type,
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
		Content:   a.Content,
	}
}

func (a *Article) ToArticleTitle() *models.ArticleTitle {
	return &models.ArticleTitle{
		Id:        a.Id,
		Title:     a.Title,
		Uri:       "/article/" + a.Uri,
		Author:    a.Author,
		CreatedAt: a.CreatedAt,
		ViewCount: a.Count,
	}
}

func (a *Article) ToArticleSearchAbstract(keyword string) *models.ArticleSearchAbstract {
	title := strings.ReplaceAll(bluemonday.UGCPolicy().Sanitize(a.Title), keyword, fmt.Sprintf("<b>%s</b>", keyword))
	keywords := strings.ReplaceAll(bluemonday.UGCPolicy().Sanitize(a.Keywords), keyword, fmt.Sprintf("<b>%s</b>", keyword))

	content := bluemonday.UGCPolicy().Sanitize(a.Content)
	idx := strings.Index(content, keyword)
	if idx > 20 {
		idx = idx - 20
	}
	last := idx + 100
	if last > len(content) {
		last = len(content) - 1
	}
	content = content[idx:last]
	content = strings.ReplaceAll(content, keyword, fmt.Sprintf("<b>%s</b>", keyword))

	asa := &models.ArticleSearchAbstract{
		Id:        a.Id,
		Title:     title,
		Keywords:  keywords,
		Content:   content,
		CreatedAt: &a.CreatedAt,
	}
	return asa
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
